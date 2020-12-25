package uim

import (
	"context"
)

type IMAccount struct {
	Id          string      `json:"id"`
	AccountId   string      `json:"accountId"`
	CustomId    string      `json:"customId"`
	Nickname    string      `json:"nickname"`
	Avatar      string      `json:"avatar"`
	Gender      int32       `json:"gender"`
	Mobile      string      `json:"mobile"`
	Country     string      `json:"country"`
	State       string      `json:"state"`
	City        string      `json:"city"`
	Signature   string      `json:"signature"`
	Alias       string      `json:"alias"`
	Qrcode      string      `json:"qrcode"`
	ExtendProps interface{} `json:"extendProps"`
	IsDeleted   bool        `json:"isDeleted"`
	CreateAt    int64       `json:"createAt"`
	UpdateAt    int64       `json:"updateAt"`
}

type IMAccountUpateParameters struct {
	AccountId    string      `json:"accountId"`
	CustomId     string      `json:"customeId"`
	Nickname     string      `json:"nickname"`
	Avatar       string      `json:"avatar"`
	Gender       int32       `json:"gender"`
	Mobile       string      `json:"mobile"`
	Country      string      `json:"country"`
	State        string      `json:"state"`
	City         string      `json:"city"`
	Alias        string      `json:"alias"`
	Qrcode       string      `json:"qrcode"`
	Signature    string      `json:"signature"`
	TenantId     string      `json:"tenantId"`
	TenantUserId string      `json:"tenantUserId"`
	ExtendProps  interface{} `json:"extendProps"`
	IsDeleted    bool        `json:"isDeleted"`
}

type IMAccountAddParameters struct {
	AccountId    string                 `json:"accountId"`
	CustomId     string                 `json:"customId"`
	Nickname     string                 `json:"nickname"`
	Avatar       string                 `json:"avatar"`
	Gender       int32                  `json:"gender"`
	Mobile       string                 `json:"mobile"`
	Country      string                 `json:"country"`
	State        string                 `json:"state"`
	City         string                 `json:"city"`
	Alias        string                 `json:"alias"`
	Qrcode       string                 `json:"qrcode"`
	Signature    string                 `json:"signature"`
	TenantId     string                 `json:"tenantId"`
	TenantUserId string                 `json:"tenantUserId"`
	ExtendProps  map[string]interface{} `json:"extendProps"`
}

func (api *Client) AccountAdd(req *IMAccountAddParameters) error {
	return api.AccountAddContext(context.Background(), req)
}

func (api *Client) AccountAddContext(ctx context.Context, req *IMAccountAddParameters) error {
	response := UimResponse{}
	err := api.postJSONMethod(ctx, "account.add", req, &response)
	if err != nil {
		return err
	}
	return response.Err()
}

func (api *Client) AccountUpdate(req *IMAccountUpateParameters) error {
	return api.AccountUpdateContext(context.Background(), req)
}

func (api *Client) AccountUpdateContext(ctx context.Context, req *IMAccountUpateParameters) error {
	response := UimResponse{}
	err := api.postJSONMethod(ctx, "account.update", req, &response)
	if err != nil {
		return err
	}
	return response.Err()
}
