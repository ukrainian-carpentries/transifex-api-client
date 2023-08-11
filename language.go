package transifex_api_client

import (
	"encoding/json"
	"fmt"
	"log"
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

// The function prints the information about a language
func (t *TransifexApiClient) PrintLanguage(l Language, formatter string) {

	switch formatter {
	case "text":
		fmt.Printf("Language information:\n")
		fmt.Printf("  ID: %+v\n", l.ID)
		fmt.Printf("  Type: %+v\n", l.Type)
		fmt.Printf("  Attributes:\n")
		fmt.Printf("    Code: %+v\n", l.Attributes.Code)
		fmt.Printf("    Name: %+v\n", l.Attributes.Name)
		fmt.Printf("    Rtl: %+v\n", l.Attributes.Rtl)
		fmt.Printf("    PluralEquation: %+v\n", l.Attributes.PluralEquation)
		fmt.Printf("    PluralRules:\n")
		fmt.Printf("      One: %+v\n", l.Attributes.PluralRules.One)
		fmt.Printf("      Other: %+v\n", l.Attributes.PluralRules.Other)
		fmt.Printf("  Links:\n")
		fmt.Printf("    Self: %+v\n", l.Links.Self)
	case "json":
		text2print, err := json.Marshal(l)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(text2print))

	default:
	}
}
