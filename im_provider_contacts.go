package uim

type IMProviderContactListParameters struct {
	Cursor    string `json:"cursor"`
	Limit     int64  `json:"limit"`
	AccountId string `json:"accountId"`
}

type IMProviderContactListResponse struct {
	UimResponse
	ResponseMetaData ResponseMetadata `json:"response_metadata"`
	Contacts         []IMContact      `json:"contacts"`
}

type IMProviderContactInfoParameters struct {
	AccountId string `json:"accountId"`
	UserId    string `json:"userId"`
}

type IMProviderContactInfoResponse struct {
	UimResponse
	Contact IMContact `json:"contact"`
}

type IMProviderContactUpdateParamemters struct {
	AccountId   string      `json:"accountId"`
	UserId      string      `json:"userId"`
	Alias       string      `json:"alias"`
	Tags        string      `json:"tags"`
	Source      string      `json:"source"`
	ExtendProps interface{} `json:"extendsProps"`
}

type IMProviderContactUpdateResponse struct {
	UimResponse
	Contact IMContact `json:"contact"`
}
