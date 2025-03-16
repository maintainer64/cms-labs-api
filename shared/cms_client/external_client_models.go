package cms_client

type PnetServerPingAttemptDTO struct {
	AttemptID string `json:"attempt_id" validate:"required"`
	UserEmail string `json:"user_email"`
	UserID    uint   `json:"user_id"`
}

type PNETServerPingInputDTO struct {
	Attempts []PnetServerPingAttemptDTO `json:"attempts"`
}

type PNETServerPingOutputDTO struct {
	Count int `json:"count"`
}

type PNETServerPingResponse struct {
	Error  bool                    `json:"error" validate:"required"`
	Msg    string                  `json:"msg" validate:"required"`
	Result PNETServerPingOutputDTO `json:"result"`
}
