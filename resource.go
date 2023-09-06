package transifex_api_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Resource struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		AcceptTranslations bool     `json:"accept_translations"`
		Categories         []string `json:"categories"`
		DatetimeCreated    string   `json:"datetime_created"`
		DatetimeModified   string   `json:"datetime_modified"`
		I18NOptions        struct {
			AllowDuplicateStrings bool `json:"allow_duplicate_strings"`
		} `json:"i18n_options"`
		I18NVersion int    `json:"i18n_version"`
		Mp4URL      string `json:"mp4_url"`
		Name        string `json:"name"`
		OggURL      string `json:"ogg_url"`
		Priority    string `json:"priority"`
		Slug        string `json:"slug"`
		StringCount int    `json:"string_count"`
		WebmURL     string `json:"webm_url"`
		WordCount   int    `json:"word_count"`
		YoutubeURL  string `json:"youtube_url"`
	} `json:"attributes"`
	Relationships struct {
		I18NFormat struct {
			Data struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
		} `json:"i18n_format"`
		Project struct {
			Data struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
			Links struct {
				Related string `json:"related"`
			} `json:"links"`
		} `json:"project"`
	} `json:"relationships"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
}

type ListResourcesParameters struct {
	Project string
	Cursor  string
	Slug    string
	Name    string
}

// Get a list of all resources (in a specific project).
// https://developers.transifex.com/reference/get_resources
func (t *TransifexApiClient) ListResources(params ListResourcesParameters) ([]Resource, error) {

	paramStr, err := t.createListResourcesParametersString(params)
	if err != nil {
		return nil, err
	}

	// Define the variable to decode the service response
	var r struct {
		Data  []Resource `json:"data"`
		Links struct {
			Self     string `json:"self"`
			Next     string `json:"next"`
			Previous string `json:"previous"`
		} `json:"links"`
	}

	// Create an API request
	req, err := http.NewRequest(
		"GET",
		strings.Join([]string{
			t.apiURL,
			"/resources",
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
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		t.l.Error(err)
		return nil, err
	}

	return r.Data, nil
}

// Get details of a specific resource.
// https://developers.transifex.com/reference/get_resources-resource-id
func (t *TransifexApiClient) GetResourceDetails(resource_id string) (Resource, error) {

	// Define the variable to decode the service response
	var r struct {
		Data Resource `json:"data"`
	}

	// Create an API request
	req, err := http.NewRequest(
		"GET",
		strings.Join([]string{
			t.apiURL,
			"/resources/",
			resource_id,
		}, ""),
		bytes.NewBuffer(nil))
	if err != nil {
		t.l.Error(err)
		return Resource{}, err
	}

	// Set authorization and Accept HTTP request headers
	req.Header.Set("Authorization", "Bearer "+t.token)
	req.Header.Add("Accept", "application/vnd.api+json")

	// Perform the request
	resp, err := t.client.Do(req)
	if err != nil {
		t.l.Error(err)
		return Resource{}, err
	}

	// Decode the JSON response into the corresponding variable
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		t.l.Error(err)
		return Resource{}, err
	}

	return r.Data, nil
}

// The function prints the information about a resource
func (t *TransifexApiClient) PrintResource(r Resource, formatter string) {

	switch formatter {

	case "text":
		fmt.Printf("  ID: %v\n", r.ID)
		fmt.Printf("  Type: %v\n", r.Type)
		fmt.Printf("  Attributes:\n")
		fmt.Printf("    Slug: %v\n", r.Attributes.Slug)
		fmt.Printf("    Name: %v\n", r.Attributes.Name)
		fmt.Printf("    Priority: %v\n", r.Attributes.Priority)
		fmt.Printf("    I18NVersion: %v\n", r.Attributes.I18NVersion)
		fmt.Printf("    AcceptTranslations: %v\n", r.Attributes.AcceptTranslations)
		fmt.Printf("    StringCount: %v\n", r.Attributes.StringCount)
		fmt.Printf("    WordCount: %v\n", r.Attributes.WordCount)
		fmt.Printf("    DatetimeCreated: %v\n", r.Attributes.DatetimeCreated)
		fmt.Printf("    DatetimeModified: %v\n", r.Attributes.DatetimeModified)
		if len(r.Attributes.Categories) > 0 {
			for _, v := range r.Attributes.Categories {
				fmt.Printf("%v\n", v)
			}
		}
		fmt.Printf("    I18NOptions: %v\n", r.Attributes.I18NOptions)
		fmt.Printf("    Mp4URL: %v\n", r.Attributes.Mp4URL)
		fmt.Printf("    OggURL: %v\n", r.Attributes.OggURL)
		fmt.Printf("    YoutubeURL: %v\n", r.Attributes.YoutubeURL)
		fmt.Printf("    WebmURL: %v\n", r.Attributes.WebmURL)
		fmt.Printf("  Relationships:\n")
		fmt.Printf("    Project:\n")
		fmt.Printf("      Links:\n")
		fmt.Printf("        Related: %v\n", r.Relationships.Project.Links.Related)
		fmt.Printf("      Data:\n")
		fmt.Printf("        Type: %v\n", r.Relationships.Project.Data.Type)
		fmt.Printf("        ID: %v\n", r.Relationships.Project.Data.ID)
		fmt.Printf("    I18NFormat:\n")
		fmt.Printf("      Data:\n")
		fmt.Printf("        Type: %v\n", r.Relationships.I18NFormat.Data.Type)
		fmt.Printf("        ID: %v\n", r.Relationships.I18NFormat.Data.ID)
		fmt.Printf("  Links:\n")
		fmt.Printf("    Self: %v\n", r.Links.Self)

	case "json":
		text2print, err := json.Marshal(r)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(text2print))

	default:
	}
}

// The function checks the input set of parameters and converts it into a valid URL parameters string
func (t *TransifexApiClient) createListResourcesParametersString(params ListResourcesParameters) (string, error) {
	// Initialize the parameters string
	paramStr := ""

	// Add mandatory Organization ID option
	if params.Project == "" {
		return "", fmt.Errorf("mandatory parameter 'Project' is missed")
	}
	paramStr += "&filter[project]=" + params.Project

	// Add Cursor option
	if params.Cursor != "" {
		paramStr += "&page[cursor]=" + params.Cursor
	}

	// Add Slug option
	if params.Slug != "" {
		paramStr += "&filter[slug]=" + params.Slug
	}

	// Add Name option
	if params.Name != "" {
		paramStr += "&filter[name]=" + params.Name
	}

	// Replace the & with ? symbol if the string is not empty
	if len(paramStr) > 0 {
		paramStr = "?" + strings.TrimPrefix(paramStr, "&")
	}

	return paramStr, nil
}
