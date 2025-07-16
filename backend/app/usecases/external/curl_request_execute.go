package external

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	resty "github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	"gitlab.com/a10869/api-modules/backend/app/queries"
)

type CurlRequestExecuteUC struct {
	CurlRequestQueries *queries.CurlRequestQueries
	Client             *resty.Client
	xServiceId         string
}

type CurlRequestExecuteInputDTO struct {
	CurlRequestID uint              `json:"curl_request_id" validate:"required"`
	Overrides     map[string]string `json:"override"`
}

func (u *CurlRequestExecuteUC) SetContext(xServiceId string) *CurlRequestExecuteUC {
	u.xServiceId = xServiceId
	return u
}

func ApplyOverrides(original string, overrides map[string]string) string {
	result := original
	for key, value := range overrides {
		result = strings.ReplaceAll(result, "$"+key, value)
	}
	return result
}

func (u *CurlRequestExecuteUC) Execute(dto CurlRequestExecuteInputDTO) (*http.Response, error) {
	if u.xServiceId == "" {
		return nil, errors.New("not authorized service")
	}
	log.Info().Msg("CurlRequestExecuteUC: Execute curl request")
	curlRequestEntity, err := u.CurlRequestQueries.Get(dto.CurlRequestID)
	if err != nil {
		return nil, err
	}
	var timeout int64
	if curlRequestEntity.Timeout > 0 && curlRequestEntity.Timeout <= 300 { // 5 minutes max
		timeout = curlRequestEntity.Timeout
	} else {
		timeout = 300 // 3 minutes
	}
	log.Info().Msg(fmt.Sprintf("CurlRequestExecuteUC: Execute curl request set timeout=%d", timeout))
	u.Client.SetTimeout(time.Duration(timeout) * time.Second)
	u.Client.SetRedirectPolicy(resty.FlexibleRedirectPolicy(10))

	finalURL := ApplyOverrides(curlRequestEntity.URL, dto.Overrides)
	finalBody := ApplyOverrides(curlRequestEntity.Body, dto.Overrides)
	finalHeaders := make(map[string]string)
	for k, v := range curlRequestEntity.Headers {
		finalHeaders[k] = ApplyOverrides(v, dto.Overrides)
	}
	finalMethod := ApplyOverrides(curlRequestEntity.Method, dto.Overrides)
	req := u.Client.R().
		SetDoNotParseResponse(true). // Важно! Не парсить ответ автоматически
		SetBody(finalBody).
		SetHeaders(finalHeaders)
	resp, err := req.Execute(finalMethod, finalURL)
	if err != nil {
		return nil, err
	}
	httpResp := resp.RawResponse

	// Для корректной работы с телом ответа нужно "перемотать" его
	if httpResp.Body != nil {
		bodyBytes, err := io.ReadAll(httpResp.Body)
		if err != nil {
			return nil, err
		}
		_ = httpResp.Body.Close()
		httpResp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}

	return httpResp, nil
}
