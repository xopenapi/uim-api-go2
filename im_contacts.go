package uim

import (
	"context"
)

type IMContact struct {
	AccountId     string      `json:"accountId"`
	UserId        string      `json:"userId"`
	CustomId      string      `json:"customId"`
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
	CreateAt      int64       `json:"createAt"`
	UpdateAt      int64       `json:"updateAt"`
}

type IMContactUpateParameters struct {
	AccountId     string                 `json:"accountId"`
	UserId        string                 `json:"userId"`
	Nickname      string                 `json:"nickname"`
	Avatar        string                 `json:"avatar"`
	Gender        int32                  `json:"gender"`
	Mobile        string                 `json:"mobile"`
	Country       string                 `json:"country"`
	State         string                 `json:"state"`
	City          string                 `json:"city"`
	Signature     string                 `json:"signature"`
	Alias         string                 `json:"alias"`
	Tags          []string               `json:"tags"`
	Source        string                 `json:"source"`
	IsBlackListed bool                   `json:"isBlackListed"`
	ExtendProps   map[string]interface{} `json:"extendProps"`
}

type IMContactAddParameters struct {
	AccountId     string                 `json:"accountId"`
	UserId        string                 `json:"userId"`
	CustomId      string                 `json:"customId"`
	Nickname      string                 `json:"nickname"`
	Avatar        string                 `json:"avatar"`
	Gender        int32                  `json:"gender"`
	Mobile        string                 `json:"mobile"`
	Country       string                 `json:"country"`
	State         string                 `json:"state"`
	City          string                 `json:"city"`
	Signature     string                 `json:"signature"`
	Alias         string                 `json:"alias"`
	Tags          []string               `json:"tags"`
	Source        string                 `json:"source"`
	IsBlackListed bool                   `json:"isBlackListed"`
	ExtendProps   map[string]interface{} `json:"extendProps"`
}

type IMContactRequestParameters struct {
	RequestId   string                 `json:"requestId"`
	AccountId   string                 `json:"accountId"`
	UserId      string                 `json:"userId"`
	CustomId    string                 `json:"customId"`
	Nickname    string                 `json:"nickname"`
	Avatar      string                 `json:"avatar"`
	Gender      int32                  `json:"gender"`
	Mobile      string                 `json:"mobile"`
	Country     string                 `json:"country"`
	State       string                 `json:"state"`
	City        string                 `json:"city"`
	Signature   string                 `json:"signature"`
	Alias       string                 `json:"alias"`
	Tags        []string               `json:"tags"`
	Source      string                 `json:"source"`
	ExtendProps map[string]interface{} `json:"extendProps"`
	Content     string                 `json:"content"`
}

func (api *Client) ContactAdd(req *IMContactAddParameters) error {
	return api.ContactAddContext(context.Background(), req)
}

func (api *Client) ContactAddContext(ctx context.Context, req *IMContactAddParameters) error {
	response := UimResponse{}
	err := api.postJSONMethod(ctx, "contact.add", req, &response)
	if err != nil {
		return err
	}
	return response.Err()
}

func (api *Client) ContactRequest(req *IMContactRequestParameters) error {
	return api.ContactRequestContext(context.Background(), req)
}

func (api *Client) ContactRequestContext(ctx context.Context, req *IMContactRequestParameters) error {
	response := UimResponse{}
	err := api.postJSONMethod(ctx, "contact.request", req, &response)
	if err != nil {
		return err
	}
	return response.Err()
}

func (api *Client) ContactUpdate(req *IMContactUpateParameters) error {
	return api.ContactUpdateContext(context.Background(), req)
}

func (api *Client) ContactUpdateContext(ctx context.Context, req *IMContactUpateParameters) error {
	response := UimResponse{}
	err := api.postJSONMethod(ctx, "contact.update", req, &response)
	if err != nil {
		return err
	}
	return response.Err()
}
