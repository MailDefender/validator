package services

import "maildefender/validator/internal/configuration"

type TokenValidationOut struct {
	Error string `json:"error"`
}

func ValidateTokenWithEngine(token string) (TokenValidationOut, int, error) {
	var out TokenValidationOut
	httpCode, err := doReq(configuration.EngineBaseEndpoint(), request{
		reqType:  post,
		endpoint: "/v1/engine/token/validate/" + token,
	},
		&out)

	return out, httpCode, err
}
