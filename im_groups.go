package uim

import "context"

type IMGroupPostMessageParameters struct {
	Id        string      `json:"id"`
	UserId    string      `json:"userId"`
	ToGroupId string      `json:"toGroupId"`
	Content   interface{} `json:"content"`
	IsDeleted bool        `json:"isDeleted"`
	IsRevoked bool        `json:"isRevoked"`
	SendAt    int64       `json:"sendAt"`
	RevokeAt  int64       `json:"revokeAt"`
	DeleteAt  int64       `json"deleteAt"`
	CreateAt  int64       `json:"createAt"`
	UpdateAt  int64       `json:"updateAt"`
}

type IMGroupPostMessageResponse struct {
	UimResponse
}

func (api *Client) GroupPostMessage(req *IMGroupPostMessageParameters, channelID string) error {
	return api.GroupPostMessageContext(context.Background(), req, channelID)
}

func (api *Client) GroupPostMessageContext(ctx context.Context, req *IMGroupPostMessageParameters, channelID string) error {
	response := IMChatPostMessageResponse{}
	err := api.postJSONMethod(ctx, "group.postMessage", req, &response)
	if err != nil {
		return err
	}
	return response.Err()
}
