// code to unmarshal the json that we recieve in the request

package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err == nil {
		marshalError := json.Unmarshal([]byte(body), x)
		if marshalError != nil {
			return marshalError
		}
	}
	return nil
}
