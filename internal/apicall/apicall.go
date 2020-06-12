package apicall

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"tikli/internal/config"
	"tikli/internal/logger"
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
	return a.config.Get(config.CKURL) + uri
}

// IsValidConnection checks the server status
func (a *APICall) IsValidConnection() bool {
	uri := "/status"
	url := a.getFullURL(uri)

	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.SetBasicAuth(a.config.Get(config.CKUsername), a.config.Get(config.CKPassword))

	resp, err := a.client.Do(req)
	if err != nil {
		a.log.Fatal(err)
	}

	if resp.StatusCode == 200 {
		return true
	}
	return false
}

// Get a URI
func (a *APICall) Get(uri string) string {
	url := a.getFullURL(uri)

	req, _ := http.NewRequest(http.MethodGet, url, nil)

	a.log.Info("GET " + url)
	req.SetBasicAuth(a.config.Get(config.CKUsername), a.config.Get(config.CKPassword))

	resp, err := a.client.Do(req)
	if err != nil {
		a.log.Fatal(err)
	}

	bodyText, _ := ioutil.ReadAll(resp.Body)
	return string(bodyText)
}

// Put a Single Tiddler
func (a *APICall) Put(uri string, tiddler MinimalSingleTiddler) bool {
	url := a.getFullURL(uri)

	// Convert the tiddler to json
	json, jerr := json.Marshal(tiddler)
	if jerr != nil {
		a.log.Fatal(jerr)
	}

	resp := a.makeRequest(http.MethodPut, url, bytes.NewBuffer(json))

	//a.log.Info(bytes.NewBuffer(json))

	//a.log.Info(resp)

	if resp.StatusCode == 204 {
		return true
	}
	return false
}

func (a *APICall) makeRequest(method string, url string, body io.Reader) *http.Response {

	req, _ := http.NewRequest(method, url, body)

	//a.log.Info("makeRequest() :: " + method + " " + url)
	req.SetBasicAuth(a.config.Get(config.CKUsername), a.config.Get(config.CKPassword))
	req.Header.Set("Content-Type", "application/json")
	// This is used for "authentication" by tiddlywiki (Major pain to figure out)
	req.Header.Add("X-Requested-With", "TiddlyWiki")
	//a.log.Info("makeRequest() :: req ", req)
	resp, err := a.client.Do(req)
	if err != nil {
		a.log.Fatal(err)
	}
	return resp
}
