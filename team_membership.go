package transifex_api_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type TeamMembership struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes struct {
		Role string `json:"role"`
	} `json:"attributes"`
	Relationships struct {
		Team struct {
			Data struct {
				Type string `json:"type"`
				ID   string `json:"id"`
			} `json:"data"`
			Links struct {
				Related string `json:"related"`
			} `json:"links"`
		} `json:"team"`
		Language struct {
			Data struct {
				Type string `json:"type"`
				ID   string `json:"id"`
			} `json:"data"`
			Links struct {
				Related string `json:"related"`
			} `json:"links"`
		} `json:"language"`
		User struct {
			Data struct {
				Type string `json:"type"`
				ID   string `json:"id"`
			} `json:"data"`
			Links struct {
				Related string `json:"related"`
			} `json:"links"`
		} `json:"user"`
	} `json:"relationships"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
}

// Get a list of all resources (in a specific project).
// https://developers.transifex.com/reference/get_team-memberships
func (t *TransifexApiClient) ListTeamMemberships(organizationID string) ([]TeamMembership, error) {

	// Define the variable to decode the service response
	var tms struct {
		Data  []TeamMembership `json:"data"`
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
			"/team_memberships",
			fmt.Sprintf("?filter[organization]=%s", organizationID),
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
	err = json.NewDecoder(resp.Body).Decode(&tms)
	if err != nil {
		t.l.Error(err)
		return nil, err
	}

	return tms.Data, nil
}

// Get single team membership.
// https://developers.transifex.com/reference/get_team-memberships-team-membership-id
func (t *TransifexApiClient) GetSingleTeamMembership(team_membership_id string) (TeamMembership, error) {

	// Define the variable to decode the service response
	var tms struct {
		Data TeamMembership `json:"data"`
	}

	// Create an API request
	req, err := http.NewRequest(
		"GET",
		strings.Join([]string{
			t.apiURL,
			"/team_memberships/",
			team_membership_id,
		}, ""),
		bytes.NewBuffer(nil))
	if err != nil {
		t.l.Error(err)
		return TeamMembership{}, err
	}

	// Set authorization and Accept HTTP request headers
	req.Header.Set("Authorization", "Bearer "+t.token)
	req.Header.Add("Accept", "application/vnd.api+json")

	// Perform the request
	resp, err := t.client.Do(req)
	if err != nil {
		t.l.Error(err)
		return TeamMembership{}, err
	}

	// Decode the JSON response into the corresponding variable
	err = json.NewDecoder(resp.Body).Decode(&tms)
	if err != nil {
		t.l.Error(err)
		return TeamMembership{}, err
	}

	return tms.Data, nil
}

// The function prints the information about an team membership
func (t *TransifexApiClient) PrintTeamMembership(tm TeamMembership, formatter string) {

	switch formatter {
		
	case "text":
		fmt.Printf("Team membership information:\n")
		fmt.Printf("  Type: %v\n", tm.Type)
		fmt.Printf("  ID: %v\n", tm.ID)
		fmt.Printf("  Attributes:\n")
		fmt.Printf("    Role: %v\n", tm.Attributes.Role)
		fmt.Printf("  Relationships:\n")
		fmt.Printf("    Team:\n")
		fmt.Printf("      Data:\n")
		fmt.Printf("        Type: %v\n", tm.Relationships.Team.Data.Type)
		fmt.Printf("        ID: %v\n", tm.Relationships.Team.Data.ID)
		fmt.Printf("      Links:\n")
		fmt.Printf("        Related: %v\n", tm.Relationships.Team.Links.Related)
		fmt.Printf("    Language:\n")
		fmt.Printf("      Data:\n")
		fmt.Printf("        Type: %v\n", tm.Relationships.Language.Data.Type)
		fmt.Printf("      Links:\n")
		fmt.Printf("        Related: %v\n", tm.Relationships.Language.Links.Related)
		fmt.Printf("    User:\n")
		fmt.Printf("      Data:\n")
		fmt.Printf("        Type: %v\n", tm.Relationships.User.Data.Type)
		fmt.Printf("        ID: %v\n", tm.Relationships.User.Data.ID)
		fmt.Printf("      Links:\n")
		fmt.Printf("        Related: %v\n", tm.Relationships.User.Links.Related)
		fmt.Printf("  Links:\n")
		fmt.Printf("    Self: %v\n", tm.Links.Self)

	case "json":
		text2print, err := json.Marshal(tm)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(text2print))

	default:
	}
}
