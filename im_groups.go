package uim

import "context"

type IMGroupPostMessageParameters struct {
	Id        string        `json:"id"`
	UserId    string        `json:"userId"`
	ToUserId  string        `json:"toUserId"`
	ToGroupId string        `json:"toGroupId"`
	Content   *IMMsgContent `json:"content"`
	IsDeleted bool          `json:"isDeleted"`
	IsRevoked bool          `json:"isRevoked"`
	SendAt    int64         `json:"sendAt"`
	RevokeAt  int64         `json:"revokeAt"`
	DeleteAt  int64         `json"deleteAt"`
	CreateAt  int64         `json:"createAt"`
	UpdateAt  int64         `json:"updateAt"`
}

type IMGroupPostMessageResponse struct {
	UimResponse
}

type IMGroupAddParameters struct {
	AccountId string `json:"accountId"`
	Name      string `json:"name"`
	GroupId   string `json:"groupId"`
	Signature string `json:"signature"`
	Avatar    string `json:"avatar"`
	OwnerId   string `json:"ownerId"`
	Qrcode    string `json:"qrcode"`
}

type IMGroupAddResponse struct {
	UimResponse
}

type IMGroupRemoveParameters struct {
	AccountId string `json:"accountId"`
	GroupId   string `json:"groupId"`
}

type IMGroupRemoveResponse struct {
	UimResponse
}

type IMGroupMemberAddParameters struct {
	AccountId string         `json:"accountId"`
	GroupId   string         `json:"groupId"`
	Members   []*GroupMember `json:"members"`
}

type GroupMember struct {
	MemberId string `json:"memberId"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Alias    string `json:"alias"`
}

type IMGroupMemberAddResponse struct {
	UimResponse
}

type IMGroupMemberRemoveParameters struct {
	AccountId string `json:"accountId"`
	GroupId   string `json:"groupId"`
	MemberId  string `json:"memberId"`
}

type IMGroupMemberRemoveResponse struct {
	UimResponse
}

func (api *Client) GroupPostMessage(req *IMGroupPostMessageParameters) error {
	return api.GroupPostMessageContext(context.Background(), req)
}

func (api *Client) GroupPostMessageContext(ctx context.Context, req *IMGroupPostMessageParameters) error {
	response := IMChatPostMessageResponse{}
	err := api.postJSONMethod(ctx, "group.postMessage", req, &response)
	if err != nil {
		return err
	}
	return response.Err()
}

func (api *Client) GroupAdd(req *IMGroupAddParameters) (*IMGroupAddResponse, error) {
	return api.GroupAddContext(context.Background(), req)
}

func (api *Client) GroupAddContext(ctx context.Context, req *IMGroupAddParameters) (*IMGroupAddResponse, error) {
	response := IMGroupAddResponse{}
	err := api.postJSONMethod(ctx, "group.add", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (api *Client) GroupRemove(req *IMGroupRemoveParameters) (*IMGroupRemoveResponse, error) {
	return api.GroupRemoveContext(context.Background(), req)
}

func (api *Client) GroupRemoveContext(ctx context.Context, req *IMGroupRemoveParameters) (*IMGroupRemoveResponse, error) {
	response := IMGroupRemoveResponse{}
	err := api.postJSONMethod(ctx, "group.remove", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (api *Client) GroupMemberAdd(req *IMGroupMemberAddParameters) (*IMGroupMemberAddResponse, error) {
	return api.GroupMemberAddContext(context.Background(), req)
}

func (api *Client) GroupMemberAddContext(ctx context.Context, req *IMGroupMemberAddParameters) (*IMGroupMemberAddResponse, error) {
	response := IMGroupMemberAddResponse{}
	err := api.postJSONMethod(ctx, "group.member.add", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (api *Client) GroupMemberRemove(req *IMGroupMemberRemoveParameters) (*IMGroupMemberRemoveResponse, error) {
	return api.GroupMemberRemoveContext(context.Background(), req)
}

func (api *Client) GroupMemberRemoveContext(ctx context.Context, req *IMGroupMemberRemoveParameters) (*IMGroupMemberRemoveResponse, error) {
	response := IMGroupMemberRemoveResponse{}
	err := api.postJSONMethod(ctx, "group.member.remove", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
