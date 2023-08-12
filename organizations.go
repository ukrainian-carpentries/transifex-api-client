package transifex_api_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Organization struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Name    string `json:"name"`
		Slug    string `json:"slug"`
		LogoURL string `json:"logo_url"`
		Private bool   `json:"private"`
	} `json:"attributes"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Relationships struct {
		Projects struct {
			Links struct {
				Related string `json:"related"`
			} `json:"links"`
		} `json:"projects"`
		Teams struct {
			Links struct {
				Related string `json:"related"`
			} `json:"links"`
		} `json:"teams"`
	} `json:"relationships"`
}

// Get a list of all the Organizations the user belongs to.
// https://developers.transifex.com/reference/get_organizations
func (t *TransifexApiClient) ListOrganizations() ([]Organization, error) {

	// Define the variable to decode the service response
	var lor struct {
		Data  []Organization `json:"data"`
		Links struct {
			Next     string `json:"next"`
			Previous string `json:"previous"`
			Self     string `json:"self"`
		} `json:"links"`
	}

	// Create an API request
	req, err := http.NewRequest(
		"GET",
		strings.Join([]string{
			t.apiURL,
			"/organizations",
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
	err = json.NewDecoder(resp.Body).Decode(&lor)
	if err != nil {
		t.l.Error(err)
		return nil, err
	}

	return lor.Data, nil
}

// Get the details of an Organization.
// https://developers.transifex.com/reference/get_organizations-organization-id
func (t *TransifexApiClient) GetOrganizationDetails(id string) (Organization, error) {

	// Define the variable to decode the service response
	var od struct {
		Data Organization `json:"data"`
	}

	// Create an API request
	req, err := http.NewRequest(
		"GET",
		strings.Join([]string{
			t.apiURL,
			"/organizations/",
			id,
		}, ""),
		bytes.NewBuffer(nil),
	)
	if err != nil {
		t.l.Error(err)
		return Organization{}, err
	}

	// Set authorization and Accept HTTP request headers
	req.Header.Set("Authorization", "Bearer "+t.token)
	req.Header.Add("Accept", "application/vnd.api+json")

	// Perform the request
	resp, err := t.client.Do(req)
	if err != nil {
		t.l.Error(err)
		return Organization{}, err
	}

	// Decode the JSON response into the corresponding variable
	err = json.NewDecoder(resp.Body).Decode(&od)
	if err != nil {
		t.l.Error(err)
		return Organization{}, err
	}

	return od.Data, nil
}

// The function prints the information about an organization
func (t *TransifexApiClient) PrintOrganization(o Organization, formatter string) {

	switch formatter {

	case "text":
		fmt.Printf("  ID: %v\n", o.ID)
		fmt.Printf("  Type: %v\n", o.Type)
		fmt.Printf("  Attributes:\n")
		fmt.Printf("    Name: %v\n", o.Attributes.Name)
		fmt.Printf("    Slug: %v\n", o.Attributes.Slug)
		fmt.Printf("    LogoURL: %v\n", o.Attributes.LogoURL)
		fmt.Printf("    Private: %v\n", o.Attributes.Private)
		fmt.Printf("  Links:\n")
		fmt.Printf("    Self: %v\n", o.Links.Self)
		fmt.Printf("  Relationships:\n")
		fmt.Printf("    Projects:\n")
		fmt.Printf("      Links:\n")
		fmt.Printf("        Related: %v\n", o.Relationships.Projects.Links.Related)
		fmt.Printf("    Teams:\n")
		fmt.Printf("      Links:\n")
		fmt.Printf("        Related: %v\n", o.Relationships.Teams.Links.Related)

	case "json":
		text2print, err := json.Marshal(o)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(text2print))

	default:
	}
}
