package uim

import "context"

type TimelineType int32

type TimelineMediaType int32

const (
	TextTimeline  TimelineType = 1
	ImageTimeline TimelineType = 2
	VoiceTimeline TimelineType = 3
	VideoTimeline TimelineType = 4
	LinkTimeline  TimelineType = 7

	ImageTimelineMedia TimelineMediaType = 2
	VoiceTimelineMedia TimelineMediaType = 3
	VideoTimelineMedia TimelineMediaType = 4
	LinkTimelineMedia  TimelineMediaType = 7
)

type IMTimelineAddParameters struct {
	AccountId  string              `json:"accountId"`
	UserId     string              `json:"userId"`
	Type       TimelineType        `json:"type"`
	TimelineId string              `json:"timelineId"`
	Content    string              `json:"content"`
	Location   *TimelineLocation   `json:"location"`
	LikeUsers  []*TimelineLikeUser `json:"likeUsers"`
	Comments   []*TimelineComment  `json:"comments"`
	Medias     []*TimelineMedia    `json:"medias"`
	CreateAt   int64               `json:"createAt"`
}

type TimelineLocation struct {
	Name string  `json:"name"`
	Lng  float64 `json:"lng"`
	Lat  float64 `json:"lat"`
}

type TimelineLikeUser struct {
	UserId   string `json:"userId"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	CreateAt int64  `json:"likeAt"`
}

type TimelineComment struct {
	CommentId       string `json:"commentId"`
	UserId          string `json:"userId"`
	Nickname        string `json:"nickname"`
	Avatar          string `json:"avatar"`
	Content         string `json:"content"`
	CreateAt        int64  `json:"createAt"`
	ReplayCommentId string `json:"replayId"`
}

type TimelineMedia struct {
	Type    TimelineMediaType `json:"type"`
	Content interface{}       `json:"content"`
}

type IMTimelineAddResponse struct {
	UimResponse
}

type TimelineImage struct {
	Format string `json:"format"`
	Width  int32  `json:"width"`
	Height int32  `json:"height"`
	Size   int32  `json:"size"`
	Thumb  string `json:"thumb"`
	URL    string `json:"url"`
}

type TimelineVideo struct {
	Format   string              `json:"format"`
	Duration int64               `json:"duration"`
	Size     int64               `json:"size"`
	URL      string              `json:"url"`
	Width    int64               `json:"width"`
	Height   int64               `json:"height"`
	Thumb    *TimelineVideoThumb `json:"thumb"`
}

type TimelineVideoThumb struct {
	Format string `json:"format"`
	Width  int32  `json:"width"`
	Height int32  `json:"height"`
	Size   int32  `json:"size"`
	URL    string `json:"url"`
}

type TimelineLink struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	PicURL      string `json:"picurl"`
	URL         string `json:"url"`
}

func (api *Client) TimelineAdd(req *IMTimelineAddParameters) error {
	return api.TimelineAddContext(context.Background(), req)
}

func (api *Client) TimelineAddContext(ctx context.Context, req *IMTimelineAddParameters) error {
	response := IMTimelineAddResponse{}
	err := api.postJSONMethod(ctx, "timeline.add", req, &response)
	if err != nil {
		return err
	}
	return response.Err()
}

//func (api *Client) TimelineRemove(req *IMGroupPostMessageParameters) error {
//	return api.GroupPostMessageContext(context.Background(), req)
//}
//
//func (api *Client) TimelineRemoveContext(ctx context.Context, req *IMGroupPostMessageParameters) error {
//	response := IMChatPostMessageResponse{}
//	err := api.postJSONMethod(ctx, "group.postMessage", req, &response)
//	if err != nil {
//		return err
//	}
//	return response.Err()
//}
//
//func (api *Client) TimelineLikeAdd(req *IMGroupPostMessageParameters) error {
//	return api.GroupPostMessageContext(context.Background(), req)
//}
//
//func (api *Client) TimelineLikeAddContext(ctx context.Context, req *IMGroupPostMessageParameters) error {
//	response := IMChatPostMessageResponse{}
//	err := api.postJSONMethod(ctx, "group.postMessage", req, &response)
//	if err != nil {
//		return err
//	}
//	return response.Err()
//}
//
//func (api *Client) TimelineLikeRemove(req *IMGroupPostMessageParameters) error {
//	return api.GroupPostMessageContext(context.Background(), req)
//}
//
//func (api *Client) TimelineLikeRemoveContext(ctx context.Context, req *IMGroupPostMessageParameters) error {
//	response := IMChatPostMessageResponse{}
//	err := api.postJSONMethod(ctx, "group.postMessage", req, &response)
//	if err != nil {
//		return err
//	}
//	return response.Err()
//}
//
//func (api *Client) TimelineMessageAdd(req *IMGroupPostMessageParameters) error {
//	return api.GroupPostMessageContext(context.Background(), req)
//}
//
//func (api *Client) TimelineMessageAddContext(ctx context.Context, req *IMGroupPostMessageParameters) error {
//	response := IMChatPostMessageResponse{}
//	err := api.postJSONMethod(ctx, "group.postMessage", req, &response)
//	if err != nil {
//		return err
//	}
//	return response.Err()
//}
//
//func (api *Client) TimelineMessageRemove(req *IMGroupPostMessageParameters) error {
//	return api.GroupPostMessageContext(context.Background(), req)
//}
//
//func (api *Client) TimelineMessageRemoveContext(ctx context.Context, req *IMGroupPostMessageParameters) error {
//	response := IMChatPostMessageResponse{}
//	err := api.postJSONMethod(ctx, "group.postMessage", req, &response)
//	if err != nil {
//		return err
//	}
//	return response.Err()
//}
