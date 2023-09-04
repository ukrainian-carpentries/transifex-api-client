package transifex_api_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type I18nFormat struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Name           string   `json:"name"`
		MediaType      string   `json:"media_type"`
		FileExtensions []string `json:"file_extensions"`
		Description    string   `json:"description"`
	} `json:"attributes"`
}

type ListI18nFormatsParameters struct {
	OrganizationID string
	Name           string
}

// Get information for all the supported i18n formats.
// https://developers.transifex.com/reference/get_i18n-formats
//
// For more information check
// https://help.transifex.com/en/articles/6219670-introduction-to-file-formats
func (t *TransifexApiClient) ListI18nFormats(params ListI18nFormatsParameters) ([]I18nFormat, error) {

	paramStr, err := t.createListI18nParametersString(params)
	if err != nil {
		return nil, err
	}

	// Define the variable to decode the service response
	var i18nfs struct {
		Data []I18nFormat `json:"data"`
	}

	// Create an API request
	req, err := http.NewRequest(
		"GET",
		strings.Join([]string{
			t.apiURL,
			"/i18n_formats",
			paramStr,
		}, ""),
		bytes.NewBuffer(nil))
	if err != nil {
		t.l.Error(err)
		return nil, err
	}

	// Set authorization and Accept HTTP request headers
	req.Header.Set("Authorization", "Bearer "+t.token)
	req.Header.Add("Accept", "application/vnd.api+json")

	// Perform the request
	resp, err := t.client.Do(req)
	if err != nil {
		t.l.Error(err)
		return nil, err
	}

	// Decode the JSON response into the corresponding variable
	err = json.NewDecoder(resp.Body).Decode(&i18nfs)
	if err != nil {
		t.l.Error(err)
		return nil, err
	}

	return i18nfs.Data, nil
}

// The function prints the information about an i18nFormat
func (t *TransifexApiClient) PrintI18nFormat(i I18nFormat, formatter string) {

	switch formatter {

	case "text":
		fmt.Printf("  ID: %v\n", i.ID)
		fmt.Printf("  Type: %v\n", i.Type)
		fmt.Printf("  Attributes:\n")
		fmt.Printf("    Name: %v\n", i.Attributes.Name)
		fmt.Printf("    MediaType: %v\n", i.Attributes.MediaType)
		fmt.Printf("    FileExtensions: %v\n", i.Attributes.FileExtensions)
		for _, v := range i.Attributes.FileExtensions {
			fmt.Printf("      - %s\n", v)
		}
		fmt.Printf("    Description: %v\n", i.Attributes.Description)

	case "json":
		text2print, err := json.Marshal(i)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(text2print))

	default:
	}
}

func (t *TransifexApiClient) createListI18nParametersString(params ListI18nFormatsParameters) (string, error) {
	// Initialize the parameters string
	paramStr := ""

	// Add mandatory Organization ID option
	if params.OrganizationID == "" {
		return "", fmt.Errorf("mandatory parameter 'OrganizationID' is missed")
	}
	paramStr += "&filter[organization]=" + params.OrganizationID

	// Add I18n format name option
	if params.Name != "" {
		paramStr += "&filter[name]=" + params.Name
	}

	// Replace the & with ? symbol if the string is not empty
	if len(paramStr) > 0 {
		paramStr = "?" + strings.TrimPrefix(paramStr, "&")
	}

	return paramStr, nil
}
