package apicall

import (
	"encoding/json"
	"net/url"
)

// SingleTiddler represents a single tiddler
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

type MinimalSingleTiddler struct {
	Title   string `json:"title"`
	Creator string `json:"creator"`
	Text    string `json:"text"`
}

// GetAllTiddlers returns the tiddlers from the API
func (apicall *APICall) GetAllTiddlers() string {
	uri := "/recipes/default/tiddlers.json"

	return apicall.Get(uri)
}

// GetTiddlerByName gets a tiddler from the api by name
func (apicall *APICall) GetTiddlerByName(tiddler string) SingleTiddler {
	path := url.PathEscape(tiddler)
	uri := "/recipes/default/tiddlers/" + path

	resp := apicall.Get(uri)

	var tiddlerJSON SingleTiddler
	json.Unmarshal([]byte(resp), &tiddlerJSON)
	return tiddlerJSON
}

// AddNewTiddler puts a new tiddler to the api
func (apicall *APICall) AddNewTiddler(title string, creator string, text string) bool {
	putpath := url.PathEscape(title)

	uri := "/recipes/default/tiddlers/" + putpath

	tiddler := MinimalSingleTiddler{
		Title:   title,
		Creator: creator,
		Text:    text,
	}
	return apicall.Put(uri, tiddler)
}

// UpdateTiddler PUTs a new version of the text to the server
func (apicall *APICall) UpdateTiddler(title string, text string) bool {

	putpath := url.PathEscape(title)

	uri := "/recipes/default/tiddlers/" + putpath

	tiddler := MinimalSingleTiddler{
		Text: text,
	}
	return apicall.Put(uri, tiddler)
}
