package transifex_api_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

type Resource struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Slug               string        `json:"slug"`
		Name               string        `json:"name"`
		Priority           string        `json:"priority"`
		I18NType           string        `json:"i18n_type"`
		I18NVersion        int           `json:"i18n_version"`
		AcceptTranslations bool          `json:"accept_translations"`
		StringCount        int           `json:"string_count"`
		WordCount          int           `json:"word_count"`
		DatetimeCreated    time.Time     `json:"datetime_created"`
		DatetimeModified   time.Time     `json:"datetime_modified"`
		Categories         []interface{} `json:"categories"`
		I18NOptions        struct {
		} `json:"i18n_options"`
		Mp4URL     interface{} `json:"mp4_url"`
		OggURL     interface{} `json:"ogg_url"`
		YoutubeURL interface{} `json:"youtube_url"`
		WebmURL    interface{} `json:"webm_url"`
	} `json:"attributes"`
	Relationships struct {
		Project struct {
			Links struct {
				Related string `json:"related"`
			} `json:"links"`
			Data struct {
				Type string `json:"type"`
				ID   string `json:"id"`
			} `json:"data"`
		} `json:"project"`
		I18NFormat struct {
			Data struct {
				Type string `json:"type"`
				ID   string `json:"id"`
			} `json:"data"`
		} `json:"i18n_format"`
		Base interface{} `json:"base"`
	} `json:"relationships"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
}

// Get a list of all resources (in a specific project).
// https://developers.transifex.com/reference/get_resources
func (t *TransifexApiClient) ListResources(projectID string) ([]Resource, error) {

	// Define the variable to decode the service response
	var r struct {
		Data  []Resource `json:"data"`
		Links struct {
			Self     string      `json:"self"`
			Next     interface{} `json:"next"`
			Previous interface{} `json:"previous"`
		} `json:"links"`
	}

	// Create an API request
	req, err := http.NewRequest(
		"GET",
		strings.Join([]string{
			t.apiURL,
			"/resources",
			fmt.Sprintf("?filter[project]=%s", projectID),
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
func (t *TransifexApiClient) GetResourceDetails(resourceID string) (Resource, error) {

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
			resourceID,
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

// The function prints the information about an resource
func (t *TransifexApiClient) PrintResource(r Resource, formatter string) {

	switch formatter {
	case "text":

		fmt.Printf("  ID: %v\n", r.ID)
		fmt.Printf("  Type: %v\n", r.Type)
		fmt.Printf("  Attributes:\n")
		fmt.Printf("    Slug: %v\n", r.Attributes.Slug)
		fmt.Printf("    Name: %v\n", r.Attributes.Name)
		fmt.Printf("    Priority: %v\n", r.Attributes.Priority)
		fmt.Printf("    I18NType: %v\n", r.Attributes.I18NType)
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
		fmt.Printf("    Base: %v\n", r.Relationships.Base)
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
