package zoomAPI

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
   API Documentation
   https://marketplace.zoom.us/docs/api-reference/zoom-api/meetings/meetingstatus
   This API actually ends an active meeting if it has been started.
*/

type RecordAction string

var (
	Start  RecordAction = "recording.start"
	Stop   RecordAction = "recording.stop"
	Pause  RecordAction = "recording.pause"
	Resume RecordAction = "recording.resume"
)

func (client Client) ToggleMeetingRecording(meetingId string, action RecordAction) (err error) {

	toggleMeetingRecordingRequest := ToggleMeetingRecordingRequest{
		Method: string(action),
	}

	var reqBody []byte
	reqBody, err = json.Marshal(toggleMeetingRecordingRequest)
	if err != nil {
		return
	}

	endpoint := fmt.Sprintf("/live_meetings/%v/events", meetingId)
	httpMethod := http.MethodPatch

	_, err = client.executeRequestWithBody(endpoint, httpMethod, reqBody)
	if err != nil {
		return
	}

	return

}
