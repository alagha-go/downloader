package tvshows

import (
	"encoding/json"
)


func UnmarshalLinkResponse(data []byte) (LinkResponse, error) {
	var r LinkResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *LinkResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}