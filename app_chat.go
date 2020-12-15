package uim

type AppChatPostMessageParameters struct {
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

type AppChatPostMessageResponse struct {
	UimResponse
	Message IMMessage `json:"message"`
}
