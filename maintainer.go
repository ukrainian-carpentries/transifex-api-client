package transifex_api_client

import "fmt"

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
func (t *TransifexApiClient) PrintMaintainer(m Maintainer) {

	fmt.Printf("Maintainer information:\n")
	fmt.Printf("  ID: %v\n", m.ID)
	fmt.Printf("  Type: %v\n", m.Type)
	fmt.Printf("  Attributes:\n")
	fmt.Printf("    Username: %v\n", m.Attributes.Username)
	fmt.Printf("  Links:\n")
	fmt.Printf("    Self: %v\n", m.Links.Self)
}
