package uim

import (
	"context"
)

type TokenRefreshParameters struct {
	AppId     string `json:"appId"`
	AppSecret string `json:"appSecret"`
}

type TokenRefreshResponse struct {
	UimResponse
	Token   string `json:"token"`
	Expires int    `json:"expires"`
}

func (api *Client) TokenRefresh(req *TokenRefreshParameters) (*TokenRefreshResponse, error) {
	return api.TokenRefreshContext(context.Background(), req)
}

func (api *Client) TokenRefreshContext(ctx context.Context, req *TokenRefreshParameters) (*TokenRefreshResponse, error) {
	response := TokenRefreshResponse{}
	err := api.postJSONMethod(ctx, "token.refresh", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
