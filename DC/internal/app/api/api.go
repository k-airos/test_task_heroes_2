package api

//Base API server instance description
type API struct {
	Config *APIConfig
}

var Server *API

//API constructor: build base API instance
func New(config *APIConfig) *API {
	return &API{
		Config: config,
	}
}
