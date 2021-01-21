package zoomAPI

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
   API Documentation
   https://marketplace.zoom.us/docs/api-reference/zoom-api/cloud-recording/recordingget
   This API actually ends an active meeting if it has been started.
*/

func (client Client) GetMeetingRecording(meetingId string) (GetMeetingRecordingResponse, error) {

	resp := GetMeetingRecordingResponse{}

	endpoint := fmt.Sprintf("/meetings/%v/recordings", meetingId)
	httpMethod := http.MethodGet

	var b []byte
	b, err := client.executeRequest(endpoint, httpMethod)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(b, &resp)
	if err != nil {
		return resp, err
	}

	return resp, nil

}
