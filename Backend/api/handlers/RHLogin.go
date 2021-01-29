package handlers

import (
	"errors"
	"github.com/karimabedrabbo/eyo/api/constants"
	"github.com/karimabedrabbo/eyo/api/models/response"
)

func (h *RhEnv) RHLoginUser() (int, *response.PayloadTokenResponse, error) {
	h.C.Set(k.GinJwtLoginTypeKey, k.GinJwtLoginTypeIdentity)

	h.Auth.LoginHandler(h.C)
	code := h.C.GetInt(k.GinJwtStatusCode)

	if message := h.C.GetString(k.GinJwtErrorMessage); message != "" {
		return code, nil, errors.New(message)
	}

	return code, &response.PayloadTokenResponse{
		Token: h.C.GetString(k.GinJwtToken),
		ExpiresAt: int64(h.C.GetInt(k.GinJwtExpireAt)),
	}, nil

}

func (h *RhEnv) RHLoginAnonymousUser() (int, *response.PayloadTokenResponse, error) {

	h.C.Set(k.GinJwtLoginTypeKey, k.GinJwtLoginTypeNoIdentity)

	h.Auth.LoginHandler(h.C)
	code := h.C.GetInt(k.GinJwtStatusCode)

	if message := h.C.GetString(k.GinJwtErrorMessage); message != "" {
		return code, nil, errors.New(message)
	}

	return code, &response.PayloadTokenResponse{
		Token: h.C.GetString(k.GinJwtToken),
		ExpiresAt: int64(h.C.GetInt(k.GinJwtExpireAt)),
	}, nil

}

func (h *RhEnv) RHLogoutUser() (int, error) {
	h.Auth.LogoutHandler(h.C)
	code := h.C.GetInt(k.GinJwtStatusCode)

	if message := h.C.GetString(k.GinJwtErrorMessage); message != "" {
		return code, errors.New(message)
	}

	return code, nil
}
