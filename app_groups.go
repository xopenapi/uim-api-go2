package uim

type AppGroupPostMessageParameters struct {
	UserId    string      `json:"userId"`
	ToGroupId string      `json:"toGroupId"`
	Content   interface{} `json:"content"`
}

type AppGroupPostMessageResponse struct {
	UimResponse
	Message IMMessage `json:"message"`
}
