package transifex_api_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
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

type GetResourceTranslationsCollectionParameters struct {
	Resource         string
	Language         string
	Cursor           string
	TranslatedAfter  time.Time
	TranslatedBefore time.Time
	Key              string
	ModifiedAfter    time.Time
	ModifiedBefore   time.Time
	IsTranslated     string
	IsReviewed       string
	IsProofreaded    string
	IsFinalized      string
	TranslatedBy     string
	Origin           string
	Include          string
	Tags             []string
	Limit            string
}

type GetResourceTranslationDetailsParameters struct {
	ResourceTranslation string
	Include             string
}

// Get a Resource Translations collection.
// https://developers.transifex.com/reference/get_resource-translations
func (t *TransifexApiClient) GetResourceTranslationsCollection(params GetResourceTranslationsCollectionParameters) ([]ResourceTranslation, error) {

	paramStr, err := t.createGetResourceTranslationsCollectionParametersString(params)
	if err != nil {
		return nil, err
	}

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
	err = json.NewDecoder(resp.Body).Decode(&rtc)
	if err != nil {
		t.l.Error(err)
		return nil, err
	}

	return rtc.Data, nil
}

// Get a Resource Translation details.
// https://developers.transifex.com/reference/get_resource-translations
func (t *TransifexApiClient) GetResourceTranslationDetails(params GetResourceTranslationDetailsParameters) (ResourceTranslation, error) {

	paramStr, err := t.createGetResourceTranslationDetailsParametersString(params)
	if err != nil {
		return ResourceTranslation{}, err
	}

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
			params.ResourceTranslation,
			paramStr,
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

// The function checks the input set of parameters and converts it into a valid URL parameters string
func (t *TransifexApiClient) createGetResourceTranslationsCollectionParametersString(params GetResourceTranslationsCollectionParameters) (string, error) {
	// Initialize the parameters string
	paramStr := ""

	// Add mandatory Resource option
	if params.Resource == "" {
		return "", fmt.Errorf("mandatory parameter 'Resource' is missed")
	}
	paramStr += "&filter[resource]=" + params.Resource

	// Add mandatory Language option
	if params.Language == "" {
		return "", fmt.Errorf("mandatory parameter 'Language' is missed")
	}
	paramStr += "&filter[language]=" + params.Language

	// Add optional Cursor value (from the previous response!)
	// The cursor used for pagination.
	// The value of the cursor must be retrieved from pagination links included in previous responses;
	// you should not attempt to write them on your own.
	if params.Cursor != "" {
		paramStr += "&page[cursor]=" + params.Cursor
	}

	// Add optional date_translated->gte value
	if (params.TranslatedAfter != time.Time{}) {
		paramStr += "&filter[date_translated][gt]=" + params.TranslatedAfter.Format("2006-01-02T15:04:05Z")
	}

	// Add optional date_translated->lt value
	if (params.TranslatedBefore != time.Time{}) {
		paramStr += "&filter[date_translated][lt]=" + params.TranslatedBefore.Format("2006-01-02T15:04:05Z")
	}

	// Exact match for the key of the resource string.
	//! This filter is case sensitive.
	if params.Key != "" {
		paramStr += "&filter[resource_string][key]=" + params.Key
	}

	// Add optional date_translated->gte value
	if (params.ModifiedAfter != time.Time{}) {
		paramStr += "&filter[resource_string][date_modified][gte]=" + params.ModifiedAfter.Format("2006-01-02T15:04:05Z")
	}

	// Add optional date_translated->lt value
	if (params.ModifiedBefore != time.Time{}) {
		paramStr += "&filter[resource_string][date_modified][lte]=" + params.ModifiedBefore.Format("2006-01-02T15:04:05Z")
	}

	// Add allowed IsTranslated value
	switch strings.ToLower(params.IsTranslated) {
	case "true":
		fallthrough
	case "false":
		paramStr += "&filter[translated]=" + strings.ToLower(params.IsTranslated)
	case "":
	default:
		return "", fmt.Errorf("unknown 'IsTranslated' value")
	}

	// Add allowed IsReviewed value
	switch strings.ToLower(params.IsReviewed) {
	case "true":
		fallthrough
	case "false":
		paramStr += "&filter[reviewed]=" + strings.ToLower(params.IsReviewed)
	case "":
	default:
		return "", fmt.Errorf("unknown 'IsReviewed' value")
	}

	// Add allowed IsProofreaded value
	switch strings.ToLower(params.IsProofreaded) {
	case "true":
		fallthrough
	case "false":
		paramStr += "&filter[proofread]=" + strings.ToLower(params.IsProofreaded)
	case "":
	default:
		return "", fmt.Errorf("unknown 'IsProofreaded' value")
	}

	// Add allowed IsFinalized value
	switch strings.ToLower(params.IsFinalized) {
	case "true":
		fallthrough
	case "false":
		paramStr += "&filter[finalized]=" + strings.ToLower(params.IsFinalized)
	case "":
	default:
		return "", fmt.Errorf("unknown 'IsFinalized' value")
	}

	// Add optional valid Origin value
	switch strings.ToUpper(params.Origin) {
	case "API":
		fallthrough
	case "EDITOR":
		fallthrough
	case "UPLOAD":
		fallthrough
	case "TM":
		fallthrough
	case "VENDORS:GENGO":
		fallthrough
	case "VENDORS:TEXTMASTER":
		fallthrough
	case "VENDORS:E2F":
		fallthrough
	case "MT:GOOGLE":
		fallthrough
	case "MT:MICROSOFT":
		fallthrough
	case "MT:AMAZON":
		fallthrough
	case "MT:DEEPL":
		fallthrough
	case "AUTOFETCH":
		fallthrough
	case "TX:AUTOMATED":
		fallthrough
	case "TX:NATIVE_MIGRATION":
		fallthrough
	case "TX:PROPAGATED":
		fallthrough
	case "TX:MERGED":
		paramStr += "&filter[origin]=" + strings.ToUpper(params.Origin)
	case "":
	default:
		return "", fmt.Errorf("unknown 'Origin' value")
	}

	// Add optional Include value
	if params.Include != "" {
		if params.Include != "resource_string" {
			return "", fmt.Errorf("unknown 'Include' value")
		}
		paramStr += "&include=resource_string"
	}

	// Add Tags option
	if len(params.Tags) != 0 {
		paramStr += "&filter[resource_string][tags][all]=" + strings.Join(params.Tags, ",")
	}

	// The page size limit. If not set, the default value is 150.
	// If set, the minimum value it can take is 150 and the maximum 1000.
	if params.Limit != "" {
		num, err := strconv.Atoi(params.Limit)
		if err != nil {
			return "", fmt.Errorf("unable to convert 'Limit' value to int")
		}

		if num < 150 || num > 1000 {
			return "", fmt.Errorf("value of 'Limit' parameter should be in the range [150..1000]")
		}

		paramStr += "&limit=" + params.Limit
	} else {
		paramStr += "&limit=150"
	}

	// Replace the & with ? symbol if the string is not empty
	if len(paramStr) > 0 {
		paramStr = "?" + strings.TrimPrefix(paramStr, "&")
	}

	return paramStr, nil
}

// The function checks the input set of parameters and converts it into a valid URL parameters string
func (t *TransifexApiClient) createGetResourceTranslationDetailsParametersString(params GetResourceTranslationDetailsParameters) (string, error) {
	// Initialize the parameters string
	paramStr := ""

	// Add mandatory ResourceTranslation option
	if params.ResourceTranslation == "" {
		return "", fmt.Errorf("mandatory parameter 'ResourceTranslation' is missed")
	}

	// Add optional Include value
	if params.Include != "" {
		if params.Include != "resource_string" {
			return "", fmt.Errorf("unknown 'Include' value")
		}
		paramStr += "&include=resource_string"
	}

	// Replace the & with ? symbol if the string is not empty
	if len(paramStr) > 0 {
		paramStr = "?" + strings.TrimPrefix(paramStr, "&")
	}

	return paramStr, nil
}
