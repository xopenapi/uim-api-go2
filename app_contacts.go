package uim

type AppContactListParameters struct {
	Cursor    string `json:"cursor"`
	Limit     int64  `json:"limit"`
	AccountId string `json:"accountId"`
}

type AppContactListResponse struct {
	UimResponse
	ResponseMetaData ResponseMetadata `json:"response_metadata"`
	Contacts         []IMContact      `json:"contacts"`
}

type AppContactInfoParameters struct {
	AccountId string `json:"accountId"`
	UserId    string `json:"userId"`
}

type AppContactInfoResponse struct {
	UimResponse
	Contact IMContact `json:"contact"`
}

type AppContactUpdateParamemters struct {
	AccountId   string      `json:"accountId"`
	UserId      string      `json:"userId"`
	Alias       string      `json:"alias"`
	Tags        string      `json:"tags"`
	Source      string      `json:"source"`
	ExtendProps interface{} `json:"extendsProps"`
}

type AppContactUpdateResponse struct {
	UimResponse
	Contact IMContact `json:"contact"`
}
