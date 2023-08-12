package transifex_api_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Language struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Code           string `json:"code"`
		Name           string `json:"name"`
		Rtl            bool   `json:"rtl"`
		PluralEquation string `json:"plural_equation"`
		PluralRules    struct {
			One   string `json:"one"`
			Many  string `json:"many"`
			Few   string `json:"few"`
			Other string `json:"other"`
		} `json:"plural_rules"`
	} `json:"attributes"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
}

type LanguageRelationship struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

// Get information for all the supported languages.
// https://developers.transifex.com/reference/get_languages
func (t *TransifexApiClient) ListLanguages() ([]Language, error) {

	// Define the variable to decode the service response
	var ls struct {
		Data  []Language `json:"data"`
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
			"/languages",
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
	err = json.NewDecoder(resp.Body).Decode(&ls)
	if err != nil {
		t.l.Error(err)
		return nil, err
	}

	return ls.Data, nil
}

// Get information for a specific supported language.
// https://developers.transifex.com/reference/get_languages-language-id
func (t *TransifexApiClient) GetLanguageDetails(language_id string) (Language, error) {

	// Define the variable to decode the service response
	var ld struct {
		Data Language `json:"data"`
	}

	// Create an API request
	req, err := http.NewRequest(
		"GET",
		strings.Join([]string{
			t.apiURL,
			"/languages/",
			language_id,
		}, ""),
		bytes.NewBuffer(nil))
	if err != nil {
		t.l.Error(err)
		return Language{}, err
	}

	// Set authorization and Accept HTTP request headers
	req.Header.Set("Authorization", "Bearer "+t.token)
	req.Header.Add("Accept", "application/vnd.api+json")

	// Perform the request
	resp, err := t.client.Do(req)
	if err != nil {
		t.l.Error(err)
		return Language{}, err
	}

	// Decode the JSON response into the corresponding variable
	err = json.NewDecoder(resp.Body).Decode(&ld)
	if err != nil {
		t.l.Error(err)
		return Language{}, err
	}

	return ld.Data, nil
}

// The function prints the information about a language
func (t *TransifexApiClient) PrintLanguage(l Language, formatter string) {

	switch formatter {

	case "text":
		fmt.Printf("Language information:\n")
		fmt.Printf("  ID: %v\n", l.ID)
		fmt.Printf("  Type: %v\n", l.Type)
		fmt.Printf("  Attributes:\n")
		fmt.Printf("    Code: %v\n", l.Attributes.Code)
		fmt.Printf("    Name: %v\n", l.Attributes.Name)
		fmt.Printf("    Rtl: %v\n", l.Attributes.Rtl)
		fmt.Printf("    PluralEquation: %v\n", l.Attributes.PluralEquation)
		fmt.Printf("    PluralRules:\n")
		fmt.Printf("      One: %v\n", l.Attributes.PluralRules.One)
		fmt.Printf("      Many: %v\n", l.Attributes.PluralRules.Many)
		fmt.Printf("      Few: %v\n", l.Attributes.PluralRules.Few)
		fmt.Printf("      Other: %v\n", l.Attributes.PluralRules.Other)
		fmt.Printf("  Links:\n")
		fmt.Printf("    Self: %v\n", l.Links.Self)

	case "json":
		text2print, err := json.Marshal(l)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(text2print))

	default:
	}
}
