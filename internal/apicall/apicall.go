package apicall

import (
	"tiddly-cli/internal/config"
	"tiddly-cli/internal/logger"
	"io/ioutil"
	"net/http"
)

// APICall struct
type APICall struct {
	client *http.Client
	log    logger.Logger
	config *config.Config
}

// New APICall
func New(log logger.Logger, config *config.Config) *APICall {
	a := new(APICall)
	a.client = &http.Client{}
	a.log = log
	a.config = config
	return a
}

func (a *APICall) getFullURL(uri string) string {
	return a.config.BaseURL + uri
}

// Get a URI
func (a *APICall) Get(uri string) string {
	url := a.getFullURL(uri)

	req, _ := http.NewRequest("GET", url, nil)

	a.log.Info("Getting " + url)
	req.SetBasicAuth(a.config.Auth.Username, a.config.Auth.Password)

	resp, err := a.client.Do(req)
	if err != nil {
		a.log.Fatal(err)
	}

	bodyText, _ := ioutil.ReadAll(resp.Body)
	return string(bodyText)
}
