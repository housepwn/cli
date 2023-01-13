/*
Copyright © 2022 Doppler <support@doppler.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package http

import (
	"encoding/json"
	"fmt"

	"github.com/DopplerHQ/cli/pkg/utils"
)

func CaptureCommand(command string) ([]byte, Error) {
	postBody := map[string]interface{}{"command": command}
	body, err := json.Marshal(postBody)
	if err != nil {
		return nil, Error{Err: err, Message: "Unable to marshal command"}
	}

	utils.LogDebug(fmt.Sprintf("Sending anonymous analytics payload: '%s'", body))

	_, _, resp, err := PostRequest("https://cli.doppler.com", true, map[string]string{"Content-Type": "application/json"}, "/v1/analytics", nil, body)
	if err != nil {
		return nil, Error{Err: err, Message: "Unable to send anonymous analytics"}
	}
	return resp, Error{}
}

func CaptureEvent(event string) ([]byte, Error) {
	postBody := map[string]interface{}{"event": event}
	body, err := json.Marshal(postBody)
	if err != nil {
		return nil, Error{Err: err, Message: "Unable to marshal event"}
	}

	utils.LogDebug(fmt.Sprintf("Sending anonymous analytics payload: '%s'", body))

	_, _, resp, err := PostRequest("https://cli.doppler.com", true, map[string]string{"Content-Type": "application/json"}, "/v1/analytics", nil, body)
	if err != nil {
		return nil, Error{Err: err, Message: "Unable to send anonymous analytics"}
	}
	return resp, Error{}
}