package uim

type IMMessage struct {
	Id        string      `json:"id"`
	UserId    string      `json:"userId"`
	ToUserId  string      `json:"toUserId"`
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

type IMText struct {
	Message string `json:"message"`
}

type IMImage struct {
	Format string `json:"format"`
	Width  int32  `json:"width"`
	Height int32  `json:"height"`
	Size   int32  `json:"size"`
	Thumb  string `json:"thumb"`
	URL    string `json:"url"`
}

type IMVoice struct {
	Format   string `json:"format"`
	Duration int64  `json:"duration"`
	Size     int64  `json"size"`
	URL      string `json:"url"`
}

type IMVideo struct {
	Format   string       `json:"format"`
	Duration int64        `json:"duration"`
	Size     int64        `json:"size"`
	URL      string       `json:"url"`
	Width    int64        `json:"width"`
	Height   int64        `json:"height"`
	Thumb    IMVideoThumb `json:"thumb"`
}

type IMVideoThumb struct {
	Format string `json:"format"`
	Width  int32  `json:"width"`
	Height int32  `json:"height"`
	Size   int32  `json:"size"`
	URL    string `json:"url"`
}

type IMLocation struct {
	Name string  `json:"name"`
	Lng  float64 `json:"lng"`
	Lat  float64 `json:"lat"`
}

type IMFile struct {
	Format string `json:"format"`
	Name   string `json:"name"`
	Size   int64  `json:"size"`
	URL    string `json:"url"`
}
