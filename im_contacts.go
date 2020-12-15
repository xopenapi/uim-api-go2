package uim

import (
	"context"
)

type IMContact struct {
	Id       string `json:"id"`
	CreateAt int64  `json:"createAt"`
	UpdateAt int64  `json:"updateAt"`
	IMContactUpateParameters
}

type IMContactUpateParameters struct {
	AccountId     string      `json:"accountId"`
	UserId        string      `json:"userId"`
	Nickname      string      `json:"nickname"`
	Avatar        string      `json:"avatar"`
	Gender        int32       `json:"gender"`
	Mobile        string      `json:"mobile"`
	Country       string      `json:"country"`
	State         string      `json:"state"`
	City          string      `json:"city"`
	Signature     string      `json:"signature"`
	Alias         string      `json:"alias"`
	Tags          []string    `json:"tags"`
	Source        string      `json:"source"`
	IsBlackListed bool        `json:"isBlackListed"`
	ExtendProps   interface{} `json:"extendProps"`
}

type IMContactAddParamemters struct {
	IMContactUpateParameters
}

func (api *Client) ContactAdd(req *IMContactAddParamemters, channelID string) error {
	return api.ContactAddContext(context.Background(), req, channelID)
}

func (api *Client) ContactAddContext(ctx context.Context, req *IMContactAddParamemters, channelID string) error {
	response := UimResponse{}
	err := api.postJSONMethod(ctx, "contact.add", req, &response)
	if err != nil {
		return err
	}
	return response.Err()
}

func (api *Client) ContactUpdate(req *IMContactUpateParameters, channelID string) error {
	return api.ContactUpdateContext(context.Background(), req, channelID)
}

func (api *Client) ContactUpdateContext(ctx context.Context, req *IMContactUpateParameters, channelID string) error {
	response := UimResponse{}
	err := api.postJSONMethod(ctx, "contact.update", req, &response)
	if err != nil {
		return err
	}
	return response.Err()
}
