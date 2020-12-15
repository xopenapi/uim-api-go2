package uim

import "context"

type IMChatPostMessageParameters struct {
	Id        string      `json:"id"`
	UserId    string      `json:"userId"`
	ToUserId  string      `json:"toUserId"`
	Content   interface{} `json:"content"`
	IsDeleted bool        `json:"isDeleted"`
	IsRevoked bool        `json:"isRevoked"`
	SendAt    int64       `json:"sendAt"`
	RevokeAt  int64       `json:"revokeAt"`
	DeleteAt  int64       `json"deleteAt"`
	CreateAt  int64       `json:"createAt"`
	UpdateAt  int64       `json:"updateAt"`
}

type IMChatPostMessageResponse struct {
	UimResponse
}

func (api *Client) ChatPostMessage(req *IMChatPostMessageParameters, channelID string) error {
	return api.ChatPostMessageContext(context.Background(), req, channelID)
}

func (api *Client) ChatPostMessageContext(ctx context.Context, req *IMChatPostMessageParameters, channelID string) error {
	response := IMChatPostMessageResponse{}
	err := api.postJSONMethod(ctx, "chat.postMessage", req, &response)
	if err != nil {
		return err
	}
	return response.Err()
}
