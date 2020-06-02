package apicall

import (
	"encoding/json"
	"net/url"
)

type SingleTiddler struct {
	Title    string `json:"title"`
	Created  string `json:"created"`
	Creator  string `json:"creator"`
	Modified string `json:"modified"`
	Modifier string `json:"modifier"`
	Tags     string `json:"tags"`
	Type     string `json:"type"`
	Text     string `json:"text"`
	Revision string `json:"revision"`
}

// GetAllTiddlers returns the tiddlers from the API
func (apicall *APICall) GetAllTiddlers() string {
	uri := "/recipes/default/tiddlers.json"

	return apicall.Get(uri)
}

func (apicall *APICall) GetTiddlerByName(tiddler string) SingleTiddler {
	apicall.log.Info("GetTiddlerByName", tiddler)
	path := url.PathEscape(tiddler)
	uri := "/recipes/default/tiddlers/" + path

	resp := apicall.Get(uri)

	var tiddlerJSON SingleTiddler
	json.Unmarshal([]byte(resp), &tiddlerJSON)
	return tiddlerJSON
}
