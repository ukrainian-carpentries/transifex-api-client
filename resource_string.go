package transifex_api_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type ResourceString struct {
	Attributes struct {
		AppearanceOrder          int    `json:"appearance_order"`
		CharacterLimit           int    `json:"character_limit"`
		Context                  string `json:"context"`
		DatetimeCreated          string `json:"datetime_created"`
		DeveloperComment         string `json:"developer_comment"`
		Instructions             string `json:"instructions"`
		Key                      string `json:"key"`
		MetadataDatetimeModified string `json:"metadata_datetime_modified"`
		Occurrences              string `json:"occurrences"`
		Pluralized               bool   `json:"pluralized"`
		StringHash               string `json:"string_hash"`
		Strings                  struct {
			One   string `json:"one"`
			Other string `json:"other"`
		} `json:"strings"`
		StringsDatetimeModified string   `json:"strings_datetime_modified"`
		Tags                    []string `json:"tags"`
	} `json:"attributes"`
	ID    string `json:"id"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Relationships struct {
		Committer struct {
			Data struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
			Links struct {
				Related string `json:"related"`
			} `json:"links"`
		} `json:"committer"`
		Language struct {
			Data struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
			Links struct {
				Related string `json:"related"`
			} `json:"links"`
		} `json:"language"`
		Resource struct {
			Data struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
			Links struct {
				Related string `json:"related"`
			} `json:"links"`
		} `json:"resource"`
	} `json:"relationships"`
	Type string `json:"type"`
}

type ResourceStringRevision interface{}

// Get resource strings collection.
// https://developers.transifex.com/reference/get_resource-strings
func (t *TransifexApiClient) GetResourceStringsCollection(resourceID string) ([]ResourceString, error) {

	// Define the variable to decode the service response
	var rsc struct {
		Data  []ResourceString `json:"data"`
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
			"/resource_strings",
			fmt.Sprintf("?filter[resource]=%s", resourceID),
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
	err = json.NewDecoder(resp.Body).Decode(&rsc)
	if err != nil {
		t.l.Error(err)
		return nil, err
	}

	return rsc.Data, nil
}

// Get the details of a specific resource string.
// https://developers.transifex.com/reference/get_resource-strings-resource-string-id
func (t *TransifexApiClient) GetResourceStringDetails(resourceStringID string) (ResourceString, error) {

	// Define the variable to decode the service response
	var rsd struct {
		Data ResourceString `json:"data"`
	}

	// Create an API request
	req, err := http.NewRequest(
		"GET",
		strings.Join([]string{
			t.apiURL,
			"/resource_strings/",
			resourceStringID,
		}, ""),
		bytes.NewBuffer(nil))
	if err != nil {
		t.l.Error(err)
		return ResourceString{}, err
	}

	// Set authorization and Accept HTTP request headers
	req.Header.Set("Authorization", "Bearer "+t.token)
	req.Header.Add("Accept", "application/vnd.api+json")

	// Perform the request
	resp, err := t.client.Do(req)
	if err != nil {
		t.l.Error(err)
		return ResourceString{}, err
	}

	// Decode the JSON response into the corresponding variable
	err = json.NewDecoder(resp.Body).Decode(&rsd)
	if err != nil {
		t.l.Error(err)
		return ResourceString{}, err
	}

	return rsd.Data, nil
}

// Get revisions of resource strings.
// https://developers.transifex.com/reference/get_resource-strings-revisions
func (t *TransifexApiClient) GetRevisionsOfResourceStrings(resourceStringID string) ([]ResourceStringRevision, error) {

	// Define the variable to decode the service response
	var rors struct {
		Data  []ResourceStringRevision `json:"data"`
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
			"/resource_strings_revisions",
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
	err = json.NewDecoder(resp.Body).Decode(&rors)
	if err != nil {
		t.l.Error(err)
		return nil, err
	}

	return rors.Data, nil
}

// The function prints the information about a resource string
func (t *TransifexApiClient) PrintResourseString(s ResourceString, formatter string) {

	switch formatter {
	case "text":
		fmt.Printf("  Type: %v\n", s.Type)
		fmt.Printf("  ID: %v\n", s.ID)
		fmt.Printf("  Attributes:\n")
		fmt.Printf("    AppearanceOrder: %v\n", s.Attributes.AppearanceOrder)
		fmt.Printf("    Key: %v\n", s.Attributes.Key)
		fmt.Printf("    Context: %v\n", s.Attributes.Context)
		fmt.Printf("    Strings:\n")
		fmt.Printf("      Other: %v\n", s.Attributes.Strings.Other)

		// !ToDo: Check the Tags type
		if len(s.Attributes.Tags) > 0 {
			fmt.Printf("    Tags:\n")
			for _, v := range s.Attributes.Tags {
				fmt.Printf("      - %v\n", v)
			}
		}

		fmt.Printf("    Occurrences: %v\n", s.Attributes.Occurrences)
		fmt.Printf("    DeveloperComment: %v\n", s.Attributes.DeveloperComment)
		fmt.Printf("    Instructions: %v\n", s.Attributes.Instructions)
		fmt.Printf("    CharacterLimit: %v\n", s.Attributes.CharacterLimit)
		fmt.Printf("    Pluralized: %v\n", s.Attributes.Pluralized)
		fmt.Printf("    StringHash: %v\n", s.Attributes.StringHash)
		fmt.Printf("    DatetimeCreated: %v\n", s.Attributes.DatetimeCreated)
		fmt.Printf("    MetadataDatetimeModified: %v\n", s.Attributes.MetadataDatetimeModified)
		fmt.Printf("    StringsDatetimeModified: %v\n", s.Attributes.StringsDatetimeModified)

		fmt.Printf("  Relationships:\n")
		fmt.Printf("    Resource:\n")
		fmt.Printf("      Data:\n")
		fmt.Printf("        Type: %v\n", s.Relationships.Resource.Data.Type)
		fmt.Printf("        ID: %v\n", s.Relationships.Resource.Data.ID)
		fmt.Printf("      Links:\n")
		fmt.Printf("        Related: %v\n", s.Relationships.Resource.Links.Related)
		fmt.Printf("    Language:\n")
		fmt.Printf("      Data:\n")
		fmt.Printf("        Type: %v\n", s.Relationships.Language.Data.Type)
		fmt.Printf("        ID: %v\n", s.Relationships.Language.Data.ID)
		fmt.Printf("      Links:\n")
		fmt.Printf("        Related: %v\n", s.Relationships.Language.Links.Related)
		fmt.Printf("    Committer:\n")
		fmt.Printf("      Data:\n")
		fmt.Printf("        Type: %v\n", s.Relationships.Committer.Data.Type)
		fmt.Printf("        ID: %v\n", s.Relationships.Committer.Data.ID)
		fmt.Printf("      Links:\n")
		fmt.Printf("        Related: %v\n", s.Relationships.Committer.Links.Related)
		fmt.Printf("  Links:\n")
		fmt.Printf("    Self: %v\n", s.Links.Self)
	case "json":
		text2print, err := json.Marshal(s)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(text2print))

	default:
	}
}
