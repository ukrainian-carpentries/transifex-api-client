package transifex_api_client

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type TransifexApiClient struct {
	apiURL string         // URL of the Transifex service
	l      *logrus.Logger // An instance of the logrus logger
	token  string         // An auth token for the API client
	client *http.Client   // HTTP client to send the requests to the service API
}

// The function returns a new instance of the transifex API client
// with the configured logger
func New(config *Config) (*TransifexApiClient, error) {

	// Create a transifex API client instance
	tr := &TransifexApiClient{
		apiURL: config.ApiURL,  // save the service URL (as string)
		l:      logrus.New(),   // create a logger instance
		token:  config.Token,   // save the service API token
		client: &http.Client{}, // create an HTTP client to send API requests
	}

	// Configure the logger
	tr.configureLogger(config)
	return tr, nil
}
