package uim

type IMProviderGroupPostMessageParameters struct {
	UserId    string      `json:"userId"`
	ToGroupId string      `json:"toGroupId"`
	Content   interface{} `json:"content"`
}

type IMProviderGroupPostMessageResponse struct {
	UimResponse
	Message IMMessage `json:"message"`
}
