package transifex_api_client

import (
	"encoding/json"
	"fmt"
	"log"
)

type Maintainer struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Username string `json:"username"`
	} `json:"attributes"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
}

type MaintainerRelationship struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// The function prints the information about a maintainer
func (t *TransifexApiClient) PrintMaintainer(m Maintainer, formatter string) {

	switch formatter {

	case "text":
		fmt.Printf("Maintainer information:\n")
		fmt.Printf("  ID: %v\n", m.ID)
		fmt.Printf("  Type: %v\n", m.Type)
		fmt.Printf("  Attributes:\n")
		fmt.Printf("    Username: %v\n", m.Attributes.Username)
		fmt.Printf("  Links:\n")
		fmt.Printf("    Self: %v\n", m.Links.Self)

	case "json":
		text2print, err := json.Marshal(m)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(text2print))

	default:
	}
}
