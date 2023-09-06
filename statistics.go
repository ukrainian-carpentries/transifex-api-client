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

type ResourseLanguageStat struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		LastProofreadUpdate   time.Time `json:"last_proofread_update"`
		LastReviewUpdate      time.Time `json:"last_review_update"`
		LastTranslationUpdate time.Time `json:"last_translation_update"`
		LastUpdate            time.Time `json:"last_update"`
		ProofreadStrings      int       `json:"proofread_strings"`
		ProofreadWords        int       `json:"proofread_words"`
		ReviewedStrings       int       `json:"reviewed_strings"`
		ReviewedWords         int       `json:"reviewed_words"`
		TotalStrings          int       `json:"total_strings"`
		TotalWords            int       `json:"total_words"`
		TranslatedStrings     int       `json:"translated_strings"`
		TranslatedWords       int       `json:"translated_words"`
		UntranslatedStrings   int       `json:"untranslated_strings"`
		UntranslatedWords     int       `json:"untranslated_words"`
	} `json:"attributes"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Relationships struct {
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
}

type GetResourceLanguageStatsCollectionParameters struct {
	Project  string
	Resource string
	Language string
	Cursor   string
}

// Get the statistics for a set of resources.
// You must specify at least a project and optionally a language/resource to filter against.
// https://developers.transifex.com/reference/get_resource-language-stats
func (t *TransifexApiClient) GetResourceLanguageStatsCollection(params GetResourceLanguageStatsCollectionParameters) ([]ResourseLanguageStat, error) {

	paramStr, err := t.createGetResourceLanguageStatsCollectionParametersString(params)
	if err != nil {
		return nil, err
	}

	// Define the variable to decode the service response
	var rlsc struct {
		Data  []ResourseLanguageStat `json:"data"`
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
			"/resource_language_stats",
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
	err = json.NewDecoder(resp.Body).Decode(&rlsc)
	if err != nil {
		t.l.Error(err)
		return nil, err
	}

	return rlsc.Data, nil
}

// Get information for a specific supported language.
// https://developers.transifex.com/reference/get_resource-language-stats-resource-language-stats-id
func (t *TransifexApiClient) GetResourceLanguageStats(resource_language_stats_id string) (ResourseLanguageStat, error) {

	// Define the variable to decode the service response
	var rls struct {
		Data ResourseLanguageStat `json:"data"`
	}

	// Create an API request
	req, err := http.NewRequest(
		"GET",
		strings.Join([]string{
			t.apiURL,
			"/resource_language_stats/",
			resource_language_stats_id,
		}, ""),
		bytes.NewBuffer(nil))
	if err != nil {
		t.l.Error(err)
		return ResourseLanguageStat{}, err
	}

	// Set authorization and Accept HTTP request headers
	req.Header.Set("Authorization", "Bearer "+t.token)
	req.Header.Add("Accept", "application/vnd.api+json")

	// Perform the request
	resp, err := t.client.Do(req)
	if err != nil {
		t.l.Error(err)
		return ResourseLanguageStat{}, err
	}

	// Decode the JSON response into the corresponding variable
	err = json.NewDecoder(resp.Body).Decode(&rls)
	if err != nil {
		t.l.Error(err)
		return ResourseLanguageStat{}, err
	}

	return rls.Data, nil
}

// The function prints the information about a resource
func (t *TransifexApiClient) PrintResourseLanguageStat(r ResourseLanguageStat, formatter string) {

	switch formatter {

	case "text":
		fmt.Printf("  ID: %v\n", r.ID)
		fmt.Printf("  Type: %v\n", r.Type)
		fmt.Printf("  Attributes:\n")
		fmt.Printf("    LastProofreadUpdate: %v\n", r.Attributes.LastProofreadUpdate)
		fmt.Printf("    LastReviewUpdate: %v\n", r.Attributes.LastReviewUpdate)
		fmt.Printf("    LastTranslationUpdate: %v\n", r.Attributes.LastTranslationUpdate)
		fmt.Printf("    LastUpdate: %v\n", r.Attributes.LastUpdate)
		fmt.Printf("    ProofreadStrings: %v\n", r.Attributes.ProofreadStrings)
		fmt.Printf("    ProofreadWords: %v\n", r.Attributes.ProofreadWords)
		fmt.Printf("    ReviewedStrings: %v\n", r.Attributes.ReviewedStrings)
		fmt.Printf("    ReviewedWords: %v\n", r.Attributes.ReviewedWords)
		fmt.Printf("    TotalStrings: %v\n", r.Attributes.TotalStrings)
		fmt.Printf("    TotalWords: %v\n", r.Attributes.TotalWords)
		fmt.Printf("    TranslatedStrings: %v\n", r.Attributes.TranslatedStrings)
		fmt.Printf("    TranslatedWords: %v\n", r.Attributes.TranslatedWords)
		fmt.Printf("    UntranslatedStrings: %v\n", r.Attributes.UntranslatedStrings)
		fmt.Printf("    UntranslatedWords: %v\n", r.Attributes.UntranslatedWords)
		fmt.Printf("  Links:\n")
		fmt.Printf("    Self: %v\n", r.Links.Self)
		fmt.Printf("  Relationships:\n")
		fmt.Printf("    Language:\n")
		fmt.Printf("      Data:\n")
		fmt.Printf("        ID: %v\n", r.Relationships.Language.Data.ID)
		fmt.Printf("        Type: %v\n", r.Relationships.Language.Data.Type)
		fmt.Printf("      Links:\n")
		fmt.Printf("        Related: %v\n", r.Relationships.Language.Links.Related)
		fmt.Printf("    Resource:\n")
		fmt.Printf("      Data: %v\n", r.Relationships.Resource.Data)
		fmt.Printf("        ID: %v\n", r.Relationships.Resource.Data.ID)
		fmt.Printf("        Type: %v\n", r.Relationships.Resource.Data.Type)
		fmt.Printf("      Links:\n")
		fmt.Printf("        Related: %v\n", r.Relationships.Resource.Links.Related)

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
func (t *TransifexApiClient) createGetResourceLanguageStatsCollectionParametersString(params GetResourceLanguageStatsCollectionParameters) (string, error) {
	// Initialize the parameters string
	paramStr := ""

	// Add mandatory Project option
	if params.Project == "" {
		return "", fmt.Errorf("mandatory parameter 'Project' is missed")
	}
	paramStr += "&filter[project]=" + params.Project

	// Add Resource option
	if params.Resource != "" {
		paramStr += "&filter[resource]=" + params.Resource
	}

	// Add Language option
	if params.Language != "" {
		paramStr += "&filter[language]=" + params.Language
	}

	// Add Cursor option
	if params.Cursor != "" {
		paramStr += "&page[cursor]=" + params.Cursor
	}

	// Replace the & with ? symbol if the string is not empty
	if len(paramStr) > 0 {
		paramStr = "?" + strings.TrimPrefix(paramStr, "&")
	}

	return paramStr, nil
}
