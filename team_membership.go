package transifex_api_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type ListTeamMembershipsParameters struct {
	Organization string
	Team         string
	Language     string
	User         string
	Role         string
	Cursor       string
	Include      string
}

type GetSingleTeamMembershipParameters struct {
	TeamMembership string
	Include        string
}

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

// List team memberships.
// https://developers.transifex.com/reference/get_team-memberships
func (t *TransifexApiClient) ListTeamMemberships(params ListTeamMembershipsParameters) ([]TeamMembership, error) {

	paramStr, err := t.createListTeamMembershipsParametersString(params)
	if err != nil {
		return nil, err
	}

	// Define the variable to decode the service response
	var tms struct {
		Data     []TeamMembership `json:"data"`
		Included []struct {
			ID         string `json:"id"`
			Type       string `json:"type"`
			Attributes struct {
				Username string `json:"username"`
			} `json:"attributes"`
			Links struct {
				Self string `json:"self"`
			} `json:"links"`
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
			"/team_memberships",
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
	err = json.NewDecoder(resp.Body).Decode(&tms)
	if err != nil {
		t.l.Error(err)
		return nil, err
	}

	return tms.Data, nil
}

// Get single team membership.
// https://developers.transifex.com/reference/get_team-memberships-team-membership-id
func (t *TransifexApiClient) GetSingleTeamMembership(params GetSingleTeamMembershipParameters) (TeamMembership, error) {

	paramStr, err := t.createGetSingleTeamMembershipParametersString(params)
	if err != nil {
		return TeamMembership{}, err
	}

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
			params.TeamMembership,
			paramStr,
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

// The function checks the input set of parameters and converts it into a valid URL parameters string
func (t *TransifexApiClient) createListTeamMembershipsParametersString(params ListTeamMembershipsParameters) (string, error) {
	// Initialize the parameters string
	paramStr := ""

	// Add mandatory Organization option
	if params.Organization == "" {
		return "", fmt.Errorf("mandatory parameter 'Organization' is missed")
	}
	paramStr += "&filter[organization]=" + params.Organization

	// Add optional Team value
	if params.Team != "" {
		paramStr += "&filter[team]=" + params.Team
	}

	// Add optional Language value
	if params.Language != "" {
		paramStr += "&filter[language]=" + params.Language
	}

	// Add optional User value
	if params.User != "" {
		paramStr += "&filter[user]=" + params.User
	}

	// Add optional Role value
	switch strings.ToLower(params.Role) {
	case "coordinator":
		fallthrough
	case "translator":
		fallthrough
	case "reviewer":
		paramStr += "&filter[role]=" + strings.ToLower(params.Role)
	case "":
	default:
		return "", fmt.Errorf("unknown 'Origin' value")
	}

	// Add optional Cursor value (from the previous response!)
	// The cursor used for pagination.
	// The value of the cursor must be retrieved from pagination links included in previous responses;
	// you should not attempt to write them on your own.
	if params.Cursor != "" {
		paramStr += "&page[cursor]=" + params.Cursor
	}

	// Add optional Include value
	if params.Include != "" {
		if params.Include != "user" {
			return "", fmt.Errorf("unknown 'Include' value")
		}
		paramStr += "&include=user"
	}

	// Replace the & with ? symbol if the string is not empty
	if len(paramStr) > 0 {
		paramStr = "?" + strings.TrimPrefix(paramStr, "&")
	}

	return paramStr, nil
}

// The function checks the input set of parameters and converts it into a valid URL parameters string
func (t *TransifexApiClient) createGetSingleTeamMembershipParametersString(params GetSingleTeamMembershipParameters) (string, error) {
	// Initialize the parameters string
	paramStr := ""

	// Add mandatory TeamMembership option
	if params.TeamMembership == "" {
		return "", fmt.Errorf("mandatory parameter 'TeamMembership' is missed")
	}

	// Add optional Include value
	if params.Include != "" {
		if params.Include != "user" {
			return "", fmt.Errorf("unknown 'Include' value")
		}
		paramStr += "&include=user"
	}

	// Replace the & with ? symbol if the string is not empty
	if len(paramStr) > 0 {
		paramStr = "?" + strings.TrimPrefix(paramStr, "&")
	}

	return paramStr, nil
}
