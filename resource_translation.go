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

type ResourceTranslation struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Strings struct {
			Other string `json:"other"`
		} `json:"strings"`
		Reviewed           bool      `json:"reviewed"`
		Proofread          bool      `json:"proofread"`
		Finalized          bool      `json:"finalized"`
		Origin             string    `json:"origin"`
		DatetimeCreated    time.Time `json:"datetime_created"`
		DatetimeTranslated time.Time `json:"datetime_translated"`
		DatetimeReviewed   time.Time `json:"datetime_reviewed"`
		DatetimeProofread  time.Time `json:"datetime_proofread"`
	} `json:"attributes"`
	Relationships struct {
		Resource struct {
			Data struct {
				Type string `json:"type"`
				ID   string `json:"id"`
			} `json:"data"`
			Links struct {
				Related string `json:"related"`
			} `json:"links"`
		} `json:"resource"`
		Language struct {
			Data struct {
				Type string `json:"type"`
				ID   string `json:"id"`
			} `json:"data"`
			Links struct {
				Related string `json:"related"`
			} `json:"links"`
		} `json:"language"`
		Translator struct {
			Data struct {
				Type string `json:"type"`
				ID   string `json:"id"`
			} `json:"data"`
			Links struct {
				Related string `json:"related"`
			} `json:"links"`
		} `json:"translator"`
		Reviewer struct {
			Data struct {
				Type string `json:"type"`
				ID   string `json:"id"`
			} `json:"data"`
			Links struct {
				Related string `json:"related"`
			} `json:"links"`
		} `json:"reviewer"`
		Proofreader struct {
			Data struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
			Links struct {
				Related string `json:"related"`
			} `json:"links"`
		} `json:"proofreader"`
		ResourceString struct {
			Data struct {
				Type string `json:"type"`
				ID   string `json:"id"`
			} `json:"data"`
			Links struct {
				Related string `json:"related"`
			} `json:"links"`
		} `json:"resource_string"`
	} `json:"relationships"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
}

// Get a Resource Translations collection.
// https://developers.transifex.com/reference/get_resource-translations
func (t *TransifexApiClient) GetResourceTranslationsCollection(resourceID, language string) ([]ResourceTranslation, error) {

	// Define the variable to decode the service response
	var rtc struct {
		Data     []ResourceTranslation `json:"data"`
		Included []struct {
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
		} `json:"included"`
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
			"/resource_translations",
			fmt.Sprintf("?filter[resource]=%s", resourceID),
			fmt.Sprintf("&filter[language]=%s", language),
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
	err = json.NewDecoder(resp.Body).Decode(&rtc)
	if err != nil {
		t.l.Error(err)
		return nil, err
	}

	return rtc.Data, nil
}

// Get a Resource Translation details.
// https://developers.transifex.com/reference/get_resource-translations
func (t *TransifexApiClient) GetResourceTranslationDetails(resource_translation_id string) (ResourceTranslation, error) {

	// Define the variable to decode the service response
	var rt struct {
		Data ResourceTranslation `json:"data"`
	}

	// Create an API request
	req, err := http.NewRequest(
		"GET",
		strings.Join([]string{
			t.apiURL,
			"/resource_translations/",
			resource_translation_id,
		}, ""),
		bytes.NewBuffer(nil))
	if err != nil {
		t.l.Error(err)
		return ResourceTranslation{}, err
	}

	// Set authorization and Accept HTTP request headers
	req.Header.Set("Authorization", "Bearer "+t.token)
	req.Header.Add("Accept", "application/vnd.api+json")

	// Perform the request
	resp, err := t.client.Do(req)
	if err != nil {
		t.l.Error(err)
		return ResourceTranslation{}, err
	}

	// Decode the JSON response into the corresponding variable
	err = json.NewDecoder(resp.Body).Decode(&rt)
	if err != nil {
		t.l.Error(err)
		return ResourceTranslation{}, err
	}

	return rt.Data, nil
}

// The function prints the information about a resource translation
func (t *TransifexApiClient) PrintResourceTranslation(r ResourceTranslation, formatter string) {

	switch formatter {

	case "text":
		fmt.Printf("  ID: %v\n", r.ID)
		fmt.Printf("  Type: %v\n", r.Type)
		fmt.Printf("  Attributes:\n")
		fmt.Printf("    Strings:\n")
		fmt.Printf("      Other: %v\n", r.Attributes.Strings.Other)
		fmt.Printf("    Reviewed: %v\n", r.Attributes.Reviewed)
		fmt.Printf("    Proofread: %v\n", r.Attributes.Proofread)
		fmt.Printf("    Finalized: %v\n", r.Attributes.Finalized)
		fmt.Printf("    Origin: %v\n", r.Attributes.Origin)
		fmt.Printf("    DatetimeCreated: %v\n", r.Attributes.DatetimeCreated)
		fmt.Printf("    DatetimeTranslated: %v\n", r.Attributes.DatetimeTranslated)
		fmt.Printf("    Reviewed: %v\n", r.Attributes.Reviewed)
		fmt.Printf("    Proofread: %v\n", r.Attributes.Proofread)
		fmt.Printf("  Relationships:\n")
		fmt.Printf("    Resource:\n")
		fmt.Printf("      Data:\n")
		fmt.Printf("        Type: %v\n", r.Relationships.Resource.Data.Type)
		fmt.Printf("        ID: %v\n", r.Relationships.Resource.Data.ID)
		fmt.Printf("      Links:\n")
		fmt.Printf("        Related: %v\n", r.Relationships.Resource.Links.Related)
		fmt.Printf("    Language:\n")
		fmt.Printf("      Data:\n")
		fmt.Printf("        Type: %v\n", r.Relationships.Language.Data.Type)
		fmt.Printf("        ID: %v\n", r.Relationships.Language.Data.ID)
		fmt.Printf("      Links:\n")
		fmt.Printf("        Related: %v\n", r.Relationships.Language.Links.Related)
		fmt.Printf("    Translator:\n")
		fmt.Printf("      Data:\n")
		fmt.Printf("        Type: %v\n", r.Relationships.Translator.Data.Type)
		fmt.Printf("        ID: %v\n", r.Relationships.Translator.Data.ID)
		fmt.Printf("      Links:\n")
		fmt.Printf("        Related: %v\n", r.Relationships.Translator.Links.Related)
		fmt.Printf("    Reviewer:\n")
		fmt.Printf("      Data:\n")
		fmt.Printf("        Type: %v\n", r.Relationships.Reviewer.Data.Type)
		fmt.Printf("        ID: %v\n", r.Relationships.Reviewer.Data.ID)
		fmt.Printf("      Links:\n")
		fmt.Printf("        Related: %v\n", r.Relationships.Reviewer.Links.Related)
		fmt.Printf("    Proofreader: %v\n", r.Relationships.Proofreader)
		fmt.Printf("    ResourceString:\n")
		fmt.Printf("      Data:\n")
		fmt.Printf("        Type: %v\n", r.Relationships.ResourceString.Data.Type)
		fmt.Printf("        ID: %v\n", r.Relationships.ResourceString.Data.ID)
		fmt.Printf("      Links:\n")
		fmt.Printf("        Related: %v\n", r.Relationships.ResourceString.Links.Related)
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
