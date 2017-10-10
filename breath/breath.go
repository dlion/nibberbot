package breath

import (
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	baseURL = "http://gorthaur.biancalana.me/vaporfont"
)

type BreatRequest struct {
	Message string `json:"message"`
}

func Breath(str string) (string, error) {
	payload := url.Values{}
	payload.Set("payload", str)
	resp, err := http.PostForm(baseURL, payload)
	if err != nil {
		return "", err
	}

	var br BreatRequest
	jb := json.NewDecoder(resp.Body)
	err = jb.Decode(&br)
	if err != nil {
		return "", err
	}
	return br.Message, nil
}
