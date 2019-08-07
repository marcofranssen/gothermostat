package nest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

type nest struct {
	config *viper.Viper
}

// Nest interact with the Nest API
type Nest interface {
	Authenticate() error
	All(combined *Combined) error
	Devices(devices *Devices) error
}

// New creates a new instance of Nest using the given config
func New(config *viper.Viper) Nest {
	return &nest{config: config}
}

// Authenticate Authenticate with the nest API
func (n *nest) Authenticate() error {
	if len(n.config.GetString("authCode")) <= 0 {
		fmt.Printf("Go to %s and get a authCode and put it in your config file.\n", n.config.GetString("authURL"))
	}

	if len(n.config.GetString("accessToken")) <= 0 {
		tokenResp, err := getAccessToken(n.config)
		if err != nil {
			return err
		}

		n.config.Set("accessToken", tokenResp.AccessToken)

		fmt.Println(tokenResp)
	}
	return nil
}

// Devices Get all devices from the Nest API
func (n *nest) Devices(devices *Devices) error {
	return n.get("/devices", devices)
}

// All Get all data from the nest API
func (n *nest) All(combined *Combined) error {
	return n.get("/", combined)
}

func (n *nest) get(path string, response interface{}) error {
	resp, err := n.authorizedRequest(http.MethodGet, path, "application/json")
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("Api request failed: %s\n%s", resp.Status, body)
	}

	err = json.Unmarshal(body, response)
	return err
}

func (n *nest) authorizedRequest(method string, path string, contentType string) (*http.Response, error) {
	client := http.Client{
		CheckRedirect: checkRedirect,
	}
	url := fmt.Sprintf("https://developer-api.nest.com%s", path)
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Authorization", fmt.Sprint("Bearer", n.config.GetString("accessToken")))
	resp, err := client.Do(req)
	return resp, err
}

func checkRedirect(redirRequest *http.Request, via []*http.Request) error {
	// Go's http.DefaultClient does not forward headers when a redirect 3xx
	// response is received. Thus, the header (which in this case contains the
	// Authorization token) needs to be passed forward to the redirect
	// destinations.
	redirRequest.Header = via[0].Header

	// Go's http.DefaultClient allows 10 redirects before returning an
	// an error. We have mimicked this default behavior.s
	if len(via) >= 10 {
		return errors.New("stopped after 10 redirects")
	}
	return nil
}

func getAccessToken(cfg *viper.Viper) (token, error) {
	var tokenResp token
	tokenURL := cfg.GetString("tokenUrl")
	clientID := cfg.GetString("clientId")
	clientSecret := cfg.GetString("clientSecret")
	authCode := cfg.GetString("authCode")
	authURL := fmt.Sprintf("%s?client_id=%s&client_secret=%s&code=%s&grant_type=authorization_code", tokenURL, clientID, clientSecret, authCode)
	resp, err := http.Post(authURL, "x-www-form-urlencoded", nil)
	if err != nil {
		return tokenResp, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return tokenResp, err
	}

	if resp.StatusCode != 200 {
		return tokenResp, fmt.Errorf("accesstoken failed: %s\n%s", resp.Status, body)
	}

	err = json.Unmarshal(body, &tokenResp)
	return tokenResp, err
}
