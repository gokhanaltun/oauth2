package oauth2

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"
)

type OAuth struct {
	Provider Provider
	Config   Config
}

type StateFunc func() (string, error)

type Config struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
	Scopes       []string
	ExtraParams  map[string]string
}

type ExchangeResponse struct {
	PostFormResponse
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    uint64 `json:"expires_in"`
	IdToken      string `json:"id_token,omitempty"`
}

func (o *OAuth) AuthCodeURL(state string) (string, error) {
	if err := o.ValidateConfig(); err != nil {
		return "", err
	}

	params := url.Values{}
	params.Set("response_type", "code")
	params.Set("client_id", o.Config.ClientID)
	params.Set("redirect_uri", o.Config.RedirectURL)
	params.Set("scope", strings.Join(o.Config.Scopes, " "))

	if len(o.Config.ExtraParams) > 0 {
		for key, value := range o.Config.ExtraParams {
			if value != "" && key != "state" {
				params.Set(key, value)
			}
		}
	}

	if state != "" {
		params.Set("state", state)
	}

	return fmt.Sprintf("%s?%s", o.Provider.AuthURL, params.Encode()), nil
}

func (o *OAuth) Exchange(code string) (*ExchangeResponse, error) {
	if err := o.ValidateConfig(); err != nil {
		return nil, err
	}

	data := url.Values{}
	data.Set("client_id", o.Config.ClientID)
	data.Set("client_secret", o.Config.ClientSecret)
	data.Set("code", code)
	data.Set("redirect_uri", o.Config.RedirectURL)
	data.Set("grant_type", "authorization_code")

	postFormResponse, err := PostForm(o.Provider.TokenURL, data)
	if err != nil {
		return &ExchangeResponse{PostFormResponse: *postFormResponse}, err
	}

	exchangeResponse := &ExchangeResponse{
		PostFormResponse: *postFormResponse,
	}

	jsonParseErr := json.Unmarshal(postFormResponse.RawResponseBody, &exchangeResponse)
	if jsonParseErr != nil {
		return exchangeResponse, errors.New("failed to parse json response")
	}

	return exchangeResponse, nil
}

func (o *OAuth) ValidateConfig() error {
	// Check if the provider is set
	if o.Provider.AuthURL == "" {
		return errors.New("provider AuthURL is not set")
	}
	if o.Provider.TokenURL == "" {
		return errors.New("provider TokenURL is not set")
	}

	// Check if the config is set properly
	if o.Config.ClientID == "" {
		return errors.New("client ID is not set")
	}
	if o.Config.ClientSecret == "" {
		return errors.New("client secret is not set")
	}
	if o.Config.RedirectURL == "" {
		return errors.New("redirect URL is not set")
	}
	if len(o.Config.Scopes) == 0 {
		return errors.New("scopes are not set")
	}
	return nil
}
