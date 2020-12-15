package uim

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/xopenapi/uim-api-go/internal/errorsx"
	"github.com/stretchr/testify/assert"
)

var dummyUimErr = errorsx.String("dummy_error_from_uim")

type viewsHandler struct {
	rawResponse string
}

func (h *viewsHandler) handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(h.rawResponse))
}

func TestUIM_OpenView(t *testing.T) {
	once.Do(startServer)
	api := New("testing-token", OptionAPIURL("http://"+serverAddr+"/"))

	cases := []struct {
		caseName     string
		triggerID    string
		rawResp      string
		expectedResp *ViewResponse
		expectedErr  error
	}{
		{
			caseName:     "pass empty trigger_id",
			triggerID:    "",
			rawResp:      "",
			expectedResp: nil,
			expectedErr:  ErrParametersMissing,
		},
		{
			caseName:  "raise an error from UIM API",
			triggerID: "dummy_trigger_id",
			rawResp: `{
				"ok": false,
				"error": "dummy_error_from_uim",
				"response_metadata": {
					"messages": [
						"dummy error response"
					]
				}
			}`,
			expectedResp: &ViewResponse{
				UimResponse{
					Ok:    false,
					Error: dummyUimErr.Error(),
				},
				View{},
			},
			expectedErr: dummyUimErr,
		},
		{
			caseName:  "success",
			triggerID: "dummy_trigger_id",
			rawResp: `{
				"ok": true,
				"view": {
					"id": "VMHU10V25",
					"team_id": "T8N4K1JN",
					"type": "modal",
					"title": {
						"type": "plain_text",
						"text": "Quite a plain modal"
					},
					"submit": {
						"type": "plain_text",
						"text": "Create"
					},
					"blocks": [
						{
							"type": "rich_text",
							"block_id": "a_block_id",
							"label": {
								"type": "plain_text",
								"text": "A simple label",
								"emoji": true
							},
							"optional": false,
							"element": {
								"type": "plain_text_input",
								"action_id": "an_action_id"
							}
						}
					],
					"private_metadata": "Shh it is a secret",
					"callback_id": "identify_your_modals",
					"external_id": "",
					"state": {
						"values": {}
					},
					"hash": "156772938.1827394",
					"clear_on_close": false,
					"notify_on_close": false,
					"root_view_id": "VMHU10V25",
					"app_id": "AA4928AQ",
					"bot_id": "BA13894H"
				}
			}`,
			expectedResp: &ViewResponse{
				UimResponse{
					Ok:    true,
					Error: "",
				},
				View{
					ID:   "VMHU10V25",
					Type: VTModal,
				},
			},
			expectedErr: nil,
		},
	}

	h := &viewsHandler{}
	http.HandleFunc("/views.open", h.handler)
	for _, c := range cases {
		t.Run(c.caseName, func(t *testing.T) {
			h.rawResponse = c.rawResp

			resp, err := api.OpenView(c.triggerID, ModalViewRequest{})
			if c.expectedErr == nil && err != nil {
				t.Errorf("unexpected error: %s\n", err)
				return
			}
			if c.expectedErr != nil && err == nil {
				t.Errorf("expected %s, but did not raise an error", c.expectedErr)
				return
			}
			if c.expectedErr != nil && err != nil && c.expectedErr.Error() != err.Error() {
				t.Errorf("expected %s as error but got %s\n", c.expectedErr, err)
				return
			}
			if resp == nil || c.expectedResp == nil {
				return
			}
			if c.expectedResp.ID != resp.ID || c.expectedResp.Type != resp.Type {
				t.Errorf("expected:\n\t%v\nas response but got:\n\t%v\n", c.expectedResp, resp)
			}
		})
	}
}

func TestUIM_View_PublishView(t *testing.T) {
	once.Do(startServer)
	api := New("testing-token", OptionAPIURL("http://"+serverAddr+"/"))

	cases := []struct {
		caseName     string
		userID       string
		rawResp      string
		expectedResp *ViewResponse
		expectedErr  error
	}{
		{
			caseName:     "pass empty user_id",
			userID:       "",
			rawResp:      "",
			expectedResp: nil,
			expectedErr:  ErrParametersMissing,
		},
		{
			caseName: "raise an error from UIM API",
			userID:   "dummy_user_id",
			rawResp: `{
				"ok": false,
				"error": "dummy_error_from_uim",
				"response_metadata": {
					"messages": [
						"dummy error response"
					]
				}
			}`,
			expectedResp: &ViewResponse{
				UimResponse{
					Ok:    false,
					Error: dummyUimErr.Error(),
				},
				View{},
			},
			expectedErr: dummyUimErr,
		},
		{
			caseName: "success",
			userID:   "dummy_user_id",
			rawResp: `{
				"ok": true,
				"view": {
					"id": "VMHU10V25",
					"team_id": "T8N4K1JN",
					"type": "home",
					"close": null,
					"submit": null,
					"blocks": [
						{
							"type": "section",
							"block_id": "2WGp9",
							"text": {
								"type": "mrkdwn",
								"text": "A simple section with some sample sentence.",
								"verbatim": false
							}
						}
					],
					"private_metadata": "Shh it is a secret",
					"callback_id": "identify_your_home_tab",
					"state": {
						"values": {}
					},
					"hash": "156772938.1827394",
					"clear_on_close": false,
					"notify_on_close": false,
					"root_view_id": "VMHU10V25",
					"previous_view_id": null,
					"app_id": "AA4928AQ",
					"external_id": "",
					"bot_id": "BA13894H"
				}
			}`,
			expectedResp: &ViewResponse{
				UimResponse{
					Ok:    true,
					Error: "",
				},
				View{
					ID:   "VMHU10V25",
					Type: VTHomeTab,
				},
			},
			expectedErr: nil,
		},
	}

	h := &viewsHandler{}
	http.HandleFunc("/views.publish", h.handler)
	for _, c := range cases {
		t.Run(c.caseName, func(t *testing.T) {
			h.rawResponse = c.rawResp

			resp, err := api.PublishView(c.userID, HomeTabViewRequest{}, "dummy_hash")
			if c.expectedErr == nil && err != nil {
				t.Errorf("unexpected error: %s\n", err)
				return
			}
			if c.expectedErr != nil && err == nil {
				t.Errorf("expected %s, but did not raise an error", c.expectedErr)
				return
			}
			if c.expectedErr != nil && err != nil && c.expectedErr.Error() != err.Error() {
				t.Errorf("expected %s as error but got %s\n", c.expectedErr, err)
				return
			}
			if resp == nil || c.expectedResp == nil {
				return
			}
			if c.expectedResp.ID != resp.ID || c.expectedResp.Type != resp.Type {
				t.Errorf("expected:\n\t%v\nas response but got:\n\t%v\n", c.expectedResp, resp)
			}
		})
	}
}

func TestUIM_PushView(t *testing.T) {
	once.Do(startServer)
	api := New("testing-token", OptionAPIURL("http://"+serverAddr+"/"))

	cases := []struct {
		caseName     string
		triggerID    string
		rawResp      string
		expectedResp *ViewResponse
		expectedErr  error
	}{
		{
			caseName:     "pass empty trigger_id",
			triggerID:    "",
			rawResp:      "",
			expectedResp: nil,
			expectedErr:  ErrParametersMissing,
		},
		{
			caseName:  "raise an error from UIM API",
			triggerID: "dummy_trigger_id",
			rawResp: `{
				"ok": false,
				"error": "dummy_error_from_uim",
				"response_metadata": {
					"messages": [
						"dummy error response"
					]
				}
			}`,
			expectedResp: &ViewResponse{
				UimResponse{
					Ok:    false,
					Error: dummyUimErr.Error(),
				},
				View{},
			},
			expectedErr: dummyUimErr,
		},
		{
			caseName:  "success",
			triggerID: "dummy_trigger_id",
			rawResp: `{
				"ok": true,
				"view": {
					"id": "VMHU10V25",
					"team_id": "T8N4K1JN",
					"type": "modal",
					"title": {
						"type": "plain_text",
						"text": "Quite a plain modal"
					},
					"submit": {
						"type": "plain_text",
						"text": "Create"
					},
					"blocks": [
						{
							"type": "rich_text",
							"block_id": "a_block_id",
							"label": {
								"type": "plain_text",
								"text": "A simple label",
								"emoji": true
							},
							"optional": false,
							"element": {
								"type": "plain_text_input",
								"action_id": "an_action_id"
							}
						}
					],
					"private_metadata": "Shh it is a secret",
					"callback_id": "identify_your_modals",
					"external_id": "",
					"state": {
						"values": {}
					},
					"hash": "156772938.1827394",
					"clear_on_close": false,
					"notify_on_close": false,
					"root_view_id": "VMHU10V25",
					"app_id": "AA4928AQ",
					"bot_id": "BA13894H"
				}
			}`,
			expectedResp: &ViewResponse{
				UimResponse{
					Ok:    true,
					Error: "",
				},
				View{
					ID:   "VMHU10V25",
					Type: VTModal,
				},
			},
			expectedErr: nil,
		},
	}

	h := &viewsHandler{}
	http.HandleFunc("/views.push", h.handler)
	for _, c := range cases {
		t.Run(c.caseName, func(t *testing.T) {
			h.rawResponse = c.rawResp

			resp, err := api.PushView(c.triggerID, ModalViewRequest{})
			if c.expectedErr == nil && err != nil {
				t.Errorf("unexpected error: %s\n", err)
				return
			}
			if c.expectedErr != nil && err == nil {
				t.Errorf("expected %s, but did not raise an error", c.expectedErr)
				return
			}
			if c.expectedErr != nil && err != nil && c.expectedErr.Error() != err.Error() {
				t.Errorf("expected %s as error but got %s\n", c.expectedErr, err)
				return
			}
			if resp == nil || c.expectedResp == nil {
				return
			}
			if c.expectedResp.ID != resp.ID || c.expectedResp.Type != resp.Type {
				t.Errorf("expected:\n\t%v\nas response but got:\n\t%v\n", c.expectedResp, resp)
			}
		})
	}
}

func TestUIM_UpdateView(t *testing.T) {
	once.Do(startServer)
	api := New("testing-token", OptionAPIURL("http://"+serverAddr+"/"))

	cases := []struct {
		caseName     string
		externalID   string
		viewID       string
		rawResp      string
		expectedResp *ViewResponse
		expectedErr  error
	}{
		{
			caseName:     "pass empty external_id and empty view_id",
			externalID:   "",
			viewID:       "",
			rawResp:      "",
			expectedResp: nil,
			expectedErr:  ErrParametersMissing,
		},
		{
			caseName:   "raise an error from UIM API",
			externalID: "dummy_external_id",
			viewID:     "",
			rawResp: `{
				"ok": false,
				"error": "dummy_error_from_uim",
				"response_metadata": {
					"messages": [
						"dummy error response"
					]
				}
			}`,
			expectedResp: &ViewResponse{
				UimResponse{
					Ok:    false,
					Error: dummyUimErr.Error(),
				},
				View{},
			},
			expectedErr: dummyUimErr,
		},
		{
			caseName:   "success",
			externalID: "",
			viewID:     "dummy_view_id",
			rawResp: `{
				"ok": true,
				"view": {
					"id": "VMHU10V25",
					"team_id": "T8N4K1JN",
					"type": "modal",
					"title": {
						"type": "plain_text",
						"text": "Quite a plain modal"
					},
					"submit": {
						"type": "plain_text",
						"text": "Create"
					},
					"blocks": [
						{
							"type": "rich_text",
							"block_id": "a_block_id",
							"label": {
								"type": "plain_text",
								"text": "A simple label",
								"emoji": true
							},
							"optional": false,
							"element": {
								"type": "plain_text_input",
								"action_id": "an_action_id"
							}
						}
					],
					"private_metadata": "Shh it is a secret",
					"callback_id": "identify_your_modals",
					"external_id": "",
					"state": {
						"values": {}
					},
					"hash": "156772938.1827394",
					"clear_on_close": false,
					"notify_on_close": false,
					"root_view_id": "VMHU10V25",
					"app_id": "AA4928AQ",
					"bot_id": "BA13894H"
				}
			}`,
			expectedResp: &ViewResponse{
				UimResponse{
					Ok:    true,
					Error: "",
				},
				View{
					ID:   "VMHU10V25",
					Type: VTModal,
				},
			},
			expectedErr: nil,
		},
	}

	h := &viewsHandler{}
	http.HandleFunc("/views.update", h.handler)
	for _, c := range cases {
		t.Run(c.caseName, func(t *testing.T) {
			h.rawResponse = c.rawResp

			resp, err := api.UpdateView(ModalViewRequest{}, c.externalID, "dummy_hash", c.viewID)
			if c.expectedErr == nil && err != nil {
				t.Errorf("unexpected error: %s\n", err)
				return
			}
			if c.expectedErr != nil && err == nil {
				t.Errorf("expected %s, but did not raise an error", c.expectedErr)
				return
			}
			if c.expectedErr != nil && err != nil && c.expectedErr.Error() != err.Error() {
				t.Errorf("expected %s as error but got %s\n", c.expectedErr, err)
				return
			}
			if resp == nil || c.expectedResp == nil {
				return
			}
			if c.expectedResp.ID != resp.ID || c.expectedResp.Type != resp.Type {
				t.Errorf("expected:\n\t%v\nas response but got:\n\t%v\n", c.expectedResp, resp)
			}
		})
	}
}

func assertViewSubmissionResponse(t *testing.T, resp *ViewSubmissionResponse, encoded string) {
	var decoded *ViewSubmissionResponse
	assert.Nil(t, json.Unmarshal([]byte(encoded), &decoded))
	assert.Equal(t, decoded, resp)
}

func TestUIM_ClearViewSubmissionResponse(t *testing.T) {
	resp := NewClearViewSubmissionResponse()
	rawResp := `{
		"response_action": "clear"
	}`

	assertViewSubmissionResponse(t, resp, rawResp)
}

func TestUIM_UpdateViewSubmissionResponse(t *testing.T) {
	resp := NewUpdateViewSubmissionResponse(&ModalViewRequest{
		Type:   VTModal,
		Title:  NewTextBlockObject("plain_text", "Test update view submission response", false, false),
		Blocks: Blocks{BlockSet: []Block{NewFileBlock("file_block_id", "external_string", "source_string")}},
	})
	rawResp := `{
		"response_action": "update",
		"view": {
			"type": "modal",
			"title": {
				"type": "plain_text",
				"text": "Test update view submission response"
			},
			"blocks": [
				{
					"type": "file",
					"block_id": "file_block_id",
					"external_id": "external_string",
					"source": "source_string"
				}
			]
		}
	}`

	assertViewSubmissionResponse(t, resp, rawResp)
}

func TestUIM_PushViewSubmissionResponse(t *testing.T) {
	resp := NewPushViewSubmissionResponse(&ModalViewRequest{
		Type:  VTModal,
		Title: NewTextBlockObject("plain_text", "Test update view submission response", false, false),
		Blocks: Blocks{
			BlockSet: []Block{
				NewContextBlock(
					"context_block_id",
					NewTextBlockObject("plain_text", "Context text", false, false),
					NewImageBlockElement("image_url", "alt_text"),
				),
			},
		},
	})
	rawResp := `{
		"response_action": "push",
		"view": {
			"type": "modal",
			"title": {
				"type": "plain_text",
				"text": "Test update view submission response"
			},
			"blocks": [
				{
					"type": "context",
					"block_id": "context_block_id",
					"elements": [
						{
							"type": "plain_text",
							"text": "Context text"
						},
						{
							"type": "image",
							"image_url": "image_url",
							"alt_text": "alt_text"
						}
					]
				}
			]
		}
	}`

	assertViewSubmissionResponse(t, resp, rawResp)
}

func TestUIM_ErrorsViewSubmissionResponse(t *testing.T) {
	resp := NewErrorsViewSubmissionResponse(map[string]string{
		"input_text_action_id": "Please input a name that's at least 6 characters long",
		"file_action_id":       "File exceeded size limit of 5 KB",
	})
	rawResp := `{
		"response_action": "errors",
		"errors": {
			"input_text_action_id": "Please input a name that's at least 6 characters long",
			"file_action_id": "File exceeded size limit of 5 KB"
		}
	}`

	assertViewSubmissionResponse(t, resp, rawResp)
}