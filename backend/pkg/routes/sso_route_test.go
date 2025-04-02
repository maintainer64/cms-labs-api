package routes

import (
	"fmt"
	"strings"
	"testing"

	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/usecases/auth"
	"gitlab.com/a10869/api-modules/shared/cms_client"
)

func TestSSOAuthorize(t *testing.T) {
	description := "successful authorization code generation"
	f := NewFiberTestHTTP()

	_, clientID := f.AuthorizationServiceBasic()

	// Создаем тестового пользователя
	user := models.User{}
	user.Email = "test" + uuid.New().String() + "@example.com"
	user.UserRole = models.UsersRoleStudent
	user.Name = "Name " + uuid.New().String()
	f.DB.Create(&user)
	authHeader := f.AuthorizationUser(user.ID, 0, "")

	input := auth.SSOAuthorizeInputDTO{
		ClientID:     clientID,
		RedirectUri:  "https://localhost/callback",
		ResponseType: "code",
		Scope:        "openid",
		UserID:       user.ID,
	}

	expectedCode := 200
	statusCode, body := f.Request(&FiberTestHttpRequest{
		Method:        "POST",
		Route:         "/api/v1/sso/authorize",
		Body:          FiberRequestPayload(input),
		Authorization: authHeader,
	})

	response := auth.SSOAuthorizeResponse{}
	_ = json.Unmarshal([]byte(body), &response)

	assert.Equal(t, expectedCode, statusCode, description)
	assert.NotEmpty(t, response.Result.Code, description)
	assert.Equal(t, input.RedirectUri, response.Result.RedirectUri, description)
	assert.Equal(t, input.ClientID, response.Result.Application, description)
}

func TestSSOAuthorizeInvalidClient(t *testing.T) {
	description := "authorization with invalid client"
	f := NewFiberTestHTTP()

	// Создаем тестового пользователя
	user := models.User{}
	user.Email = "test" + uuid.New().String() + "@example.com"
	user.UserRole = models.UsersRoleStudent
	user.Name = "Name " + uuid.New().String()
	f.DB.Create(&user)
	authHeader := f.AuthorizationUser(user.ID, 0, "")

	input := auth.SSOAuthorizeInputDTO{
		ClientID:     "invalid-client",
		RedirectUri:  "https://invalid.example.com/callback",
		ResponseType: "code",
		Scope:        "openid",
		UserID:       user.ID,
	}

	expectedCode := 500
	statusCode, body := f.Request(&FiberTestHttpRequest{
		Method:        "POST",
		Route:         "/api/v1/sso/authorize",
		Body:          FiberRequestPayload(input),
		Authorization: authHeader,
	})

	assert.Equal(t, expectedCode, statusCode, description)
	assert.Contains(t, body, "PNETServer not found", description)
}

func TestSSOTokenByAuthCode(t *testing.T) {
	description := "get tokens by authorization code"
	f := NewFiberTestHTTP()

	authHeaderServer, clientID := f.AuthorizationServiceBasic()
	server := models.PNETServer{}
	f.DB.Where("client_id = ?", clientID).Find(&server)

	// Создаем тестового пользователя
	user := models.User{}
	user.Email = "test" + uuid.New().String() + "@example.com"
	user.UserRole = models.UsersRoleStudent
	user.Name = "Name " + uuid.New().String()
	f.DB.Create(&user)

	// Создаем authorization code
	attempt := models.TokenAttempt{}
	attempt.UserID = user.ID
	attempt.ServerID = server.ID
	attempt.State = uuid.New().String()
	attempt.AuthorizationCode = uuid.New().String()
	f.DB.Create(&attempt)

	input := map[string]string{
		"grant_type":   "authorization_code",
		"code":         attempt.AuthorizationCode,
		"redirect_uri": server.Url + "/callback",
	}

	expectedCode := 200
	statusCode, body := f.Request(&FiberTestHttpRequest{
		Method:        "POST",
		Route:         "/api/v1/sso/token",
		Body:          FiberRequestFormPayload(input),
		Authorization: authHeaderServer,
		ContentType:   "application/x-www-form-urlencoded",
	})

	response := auth.SwaggerSSOToken{}
	_ = json.Unmarshal([]byte(body), &response)

	assert.Equal(t, expectedCode, statusCode, description)
	assert.NotEmpty(t, response.AccessToken, description)
	assert.NotEmpty(t, response.IdToken, description)
	assert.NotEmpty(t, response.RefreshToken, description)
	assert.Equal(t, "bearer", response.TokenType, description)
	assert.Equal(t, fmt.Sprintf("%d", user.ID), response.UserId, description)
	assert.Greater(t, response.ExpiresIn, int64(0), description)
}

func TestSSOIntrospectValidToken(t *testing.T) {
	description := "introspect valid access token"
	f := NewFiberTestHTTP()

	authHeaderClient, clientID := f.AuthorizationServiceBasic()
	server := models.PNETServer{}
	f.DB.Where("client_id = ?", clientID).Find(&server)
	authHeader := f.AuthorizationUser(0, server.ID, "")

	input := map[string]string{
		"token": strings.Replace(authHeader, "Bearer ", "", 1),
	}

	expectedCode := 200
	statusCode, body := f.Request(&FiberTestHttpRequest{
		Method:        "POST",
		Route:         "/api/v1/sso/introspect",
		Body:          FiberRequestFormPayload(input),
		Authorization: authHeaderClient,
		ContentType:   "application/x-www-form-urlencoded",
	})

	response := auth.SSOTokenIntrospect{}
	_ = json.Unmarshal([]byte(body), &response)

	assert.Equal(t, expectedCode, statusCode, description)
	assert.True(t, response.Active, description)
	assert.Equal(t, server.ClientID, response.ClientID, description)
	assert.Equal(t, server.Name, response.Username, description)
	assert.Equal(t, "access", response.TokenType, description)
}

func TestSSOUserInfo(t *testing.T) {
	description := "get user info with valid token"
	f := NewFiberTestHTTP()

	_, clientID := f.AuthorizationServiceBasic()
	server := models.PNETServer{}
	f.DB.Where("client_id = ?", clientID).Find(&server)
	user := models.User{}
	user.Email = uuid.New().String() + "@admin.com"
	user.UserRole = models.UsersRoleAdmin
	f.DB.Create(&user)
	authHeader := f.AuthorizationUser(user.ID, server.ID, "")

	expectedCode := 200
	statusCode, body := f.Request(
		&FiberTestHttpRequest{
			Method:        "GET",
			Route:         "/api/v1/sso/userinfo",
			Authorization: authHeader,
		},
	)

	response := cms_client.SSOTokenPublicData{}
	err := json.Unmarshal([]byte(body), &response)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, expectedCode, statusCode, description)
	assert.Equal(t, user.Email, response.Email, description)
	assert.Equal(t, user.Name, response.Name, description)
	assert.Equal(t, user.UserRole, response.UserRoleMain(), description)
	assert.Equal(t, fmt.Sprintf("%d", user.ID), response.Sub, description)
	assert.Equal(t, server.ClientID, response.Aud, description)
	assert.Equal(t, server.ID, response.ServerID, description)
}

func TestSSOOpenIdConfiguration(t *testing.T) {
	description := "get OpenID configuration"
	f := NewFiberTestHTTP()

	expectedCode := 200
	statusCode, body := f.Request(&FiberTestHttpRequest{
		Method: "GET",
		Route:  "/api/v1/sso/.well-known/openid-configuration",
	})

	response := auth.SSOOpenidConfigurationOutputDTO{}
	_ = json.Unmarshal([]byte(body), &response)

	assert.Equal(t, expectedCode, statusCode, description)
	assert.Contains(t, response.Issuer, "/api/v1/sso", description)
	assert.NotEmpty(t, response.AuthorizationEndpoint, description)
	assert.NotEmpty(t, response.TokenEndpoint, description)
	assert.NotEmpty(t, response.UserInfoEndpoint, description)
	assert.NotEmpty(t, response.JwksUri, description)
}

func TestSSOJwks(t *testing.T) {
	description := "get JWKS"
	f := NewFiberTestHTTP()

	expectedCode := 200
	statusCode, body := f.Request(&FiberTestHttpRequest{
		Method: "GET",
		Route:  "/api/v1/sso/jwks",
	})

	response := auth.SSOJWKSOutputDTO{}
	_ = json.Unmarshal([]byte(body), &response)

	assert.Equal(t, expectedCode, statusCode, description)
	assert.Len(t, response.Keys, 2, description) // access и refresh ключи
	for _, key := range response.Keys {
		assert.Equal(t, "RSA", key.Kty, description)
		assert.Equal(t, "sig", key.Use, description)
		assert.Contains(t, []string{"access", "refresh"}, key.Kid, description)
		assert.Equal(t, "RS256", key.Alg, description)
		assert.NotEmpty(t, key.N, description)
		assert.NotEmpty(t, key.E, description)
	}
}
