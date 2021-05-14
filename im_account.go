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
	Status      int32       `json:"status"`
	Online      int32       `json:"online"`
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
	Status       int32       `json:"status"`
	Online       int32       `json:"online"`
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
