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

type Team struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes struct {
		Name            string    `json:"name"`
		Slug            string    `json:"slug"`
		AutoJoin        bool      `json:"auto_join"`
		ClaRequired     bool      `json:"cla_required"`
		Cla             string    `json:"cla"`
		DatetimeCreated time.Time `json:"datetime_created"`
	} `json:"attributes"`
	Relationships struct {
		Organization struct {
			Links struct {
				Related string `json:"related"`
			} `json:"links"`
			Data struct {
				Type string `json:"type"`
				ID   string `json:"id"`
			} `json:"data"`
		} `json:"organization"`
		Managers struct {
			Links struct {
				Related string `json:"related"`
				Self    string `json:"self"`
			} `json:"links"`
		} `json:"managers"`
	} `json:"relationships"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
}

type TeamRelationship struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type TeamManager interface{}

type TeamManagerRelationship struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes struct {
		Name            string    `json:"name"`
		Slug            string    `json:"slug"`
		AutoJoin        bool      `json:"auto_join"`
		ClaRequired     bool      `json:"cla_required"`
		Cla             string    `json:"cla"`
		DatetimeCreated time.Time `json:"datetime_created"`
	} `json:"attributes"`
	Relationships struct {
		Organization struct {
			Links struct {
				Related string `json:"related"`
			} `json:"links"`
			Data struct {
				Type string `json:"type"`
				ID   string `json:"id"`
			} `json:"data"`
		} `json:"organization"`
		Managers struct {
			Links struct {
				Related string `json:"related"`
				Self    string `json:"self"`
			} `json:"links"`
		} `json:"managers"`
	} `json:"relationships"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
}

// Get the list of teams that belong to a single organization.
// https://developers.transifex.com/reference/get_teams
func (t *TransifexApiClient) ListTeams(organizationID string) ([]Team, error) {

	// Define the variable to decode the service response
	var ts struct {
		Data  []Team `json:"data"`
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
			"/teams",
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
	err = json.NewDecoder(resp.Body).Decode(&ts)
	if err != nil {
		t.l.Error(err)
		return nil, err
	}

	return ts.Data, nil
}

// Get the details of a single team.
// https://developers.transifex.com/reference/get_teams-team-id
func (t *TransifexApiClient) GetTeamDetail(teamID string) (Team, error) {

	// Define the variable to decode the service response
	var td struct {
		Data Team `json:"data"`
	}

	// Create an API request
	req, err := http.NewRequest(
		"GET",
		strings.Join([]string{
			t.apiURL,
			"/teams/",
			teamID,
		}, ""),
		bytes.NewBuffer(nil))
	if err != nil {
		t.l.Error(err)
		return Team{}, err
	}

	// Set authorization and Accept HTTP request headers
	req.Header.Set("Authorization", "Bearer "+t.token)
	req.Header.Add("Accept", "application/vnd.api+json")

	// Perform the request
	resp, err := t.client.Do(req)
	if err != nil {
		t.l.Error(err)
		return Team{}, err
	}

	// Decode the JSON response into the corresponding variable
	err = json.NewDecoder(resp.Body).Decode(&td)
	if err != nil {
		t.l.Error(err)
		return Team{}, err
	}

	return td.Data, nil
}

// Get the managers of a team.
// https://developers.transifex.com/reference/get_teams-team-id-managers
func (t *TransifexApiClient) GetTeamManagers(teamID string) ([]TeamManager, error) {

	// Define the variable to decode the service response
	var tms struct {
		Data  []TeamManager `json:"data"`
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
			"/teams/",
			teamID,
			"/managers",
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

// Get team manager relationships.
// https://developers.transifex.com/reference/get_teams-team-id-relationships-managers
func (t *TransifexApiClient) GetTeamManagerRelationships(teamID string) ([]TeamManagerRelationship, error) {

	// Define the variable to decode the service response
	var tmrs struct {
		Data  []TeamManagerRelationship `json:"data"`
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
			"/teams/",
			teamID,
			"/relationships/managers",
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
	err = json.NewDecoder(resp.Body).Decode(&tmrs)
	if err != nil {
		t.l.Error(err)
		return nil, err
	}

	return tmrs.Data, nil
}

// The function prints the information about an organization
func (t *TransifexApiClient) PrintTeam(tt Team, formatter string) {

	switch formatter {
	case "text":
		fmt.Printf("Team information:\n")
		fmt.Printf("  Type: %v\n", tt.Type)
		fmt.Printf("  ID: %v\n", tt.ID)
		fmt.Printf("  Attributes:\n")
		fmt.Printf("    Name: %v\n", tt.Attributes.Name)
		fmt.Printf("    Slug: %v\n", tt.Attributes.Slug)
		fmt.Printf("    AutoJoin: %v\n", tt.Attributes.AutoJoin)
		fmt.Printf("    ClaRequired: %v\n", tt.Attributes.ClaRequired)
		fmt.Printf("    Cla: %v\n", tt.Attributes.Cla)
		fmt.Printf("    DatetimeCreated: %v\n", tt.Attributes.DatetimeCreated)

		fmt.Printf("  Relationships:\n")
		fmt.Printf("    Organization:\n")
		fmt.Printf("      Links:\n")
		fmt.Printf("        Related: %v\n", tt.Relationships.Organization.Links.Related)
		fmt.Printf("      Data:\n")
		fmt.Printf("        Type: %v\n", tt.Relationships.Organization.Data.Type)
		fmt.Printf("        ID: %v\n", tt.Relationships.Organization.Data.ID)
		fmt.Printf("    Managers:\n")
		fmt.Printf("      Links:\n")
		fmt.Printf("        Related: %v\n", tt.Relationships.Managers.Links.Related)
		fmt.Printf("        Self: %v\n", tt.Relationships.Managers.Links.Self)
		fmt.Printf("  Links:\n")
		fmt.Printf("    Self: %v\n", tt.Links.Self)
	case "json":
		text2print, err := json.Marshal(tt)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(text2print))

	default:
	}
}
