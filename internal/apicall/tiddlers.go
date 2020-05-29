package apicall

// GetAllTiddlers returns the tiddlers from the API
func (apicall *APICall) GetAllTiddlers() string {
	uri := "/recipes/default/tiddlers.json"

	return apicall.Get(uri)
}
