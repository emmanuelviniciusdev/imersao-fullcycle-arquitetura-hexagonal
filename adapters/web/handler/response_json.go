package handler

import "encoding/json"

func ResponseJson(msg string) []byte {
	response := struct {
		Message string `json:"message"`
	}{
		msg,
	}

	responseJsonMarshal, err := json.Marshal(response)

	if err != nil {
		return []byte(err.Error())
	}

	return responseJsonMarshal
}
