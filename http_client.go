package oauth2

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type PostFormResponse struct {
	RawResponseBody       []byte
	FormattedResponseBody map[string]interface{}
}

func PostForm(url string, data url.Values) (*PostFormResponse, error) {

	response, err := http.PostForm(url, data)
	if err != nil {
		return nil, fmt.Errorf("POST request failed: %v", err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	postFormResponse := &PostFormResponse{
		RawResponseBody: body,
	}

	var formattedResponseBody map[string]interface{}
	if err := json.Unmarshal(body, &formattedResponseBody); err != nil {
		return postFormResponse, fmt.Errorf("failed to parse JSON: %v", err)
	}

	postFormResponse.FormattedResponseBody = formattedResponseBody

	return postFormResponse, nil
}
