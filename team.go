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

type ListTeamsParameters struct {
	Organization string
	Cursor       string
	Slug         string
	Name         string
}

type GetTeamManagersParameters struct {
	Team   string
	Cursor string
}

type GetTeamManagerRelationshipsParameters struct {
	Team   string
	Cursor string
}

// Get the list of teams that belong to a single organization.
// https://developers.transifex.com/reference/get_teams
func (t *TransifexApiClient) ListTeams(params ListTeamsParameters) ([]Team, error) {

	paramStr, err := t.createListTeamsParametersString(params)
	if err != nil {
		return nil, err
	}

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
	err = json.NewDecoder(resp.Body).Decode(&ts)
	if err != nil {
		t.l.Error(err)
		return nil, err
	}

	return ts.Data, nil
}

// Get the details of a single team.
// https://developers.transifex.com/reference/get_teams-team-id
func (t *TransifexApiClient) GetTeamDetail(team_id string) (Team, error) {

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
			team_id,
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
func (t *TransifexApiClient) GetTeamManagers(params GetTeamManagersParameters) ([]TeamManager, error) {

	paramStr, err := t.createGetTeamManagersParametersString(params)
	if err != nil {
		return nil, err
	}

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
			params.Team,
			"/managers",
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

// Get team manager relationships.
// https://developers.transifex.com/reference/get_teams-team-id-relationships-managers
func (t *TransifexApiClient) GetTeamManagerRelationships(params GetTeamManagerRelationshipsParameters) ([]TeamManagerRelationship, error) {

	paramStr, err := t.createGetTeamManagerRelationshipsParametersString(params)
	if err != nil {
		return nil, err
	}
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
			params.Team,
			"/relationships/managers",
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

// The function checks the input set of parameters and converts it into a valid URL parameters string
func (t *TransifexApiClient) createListTeamsParametersString(params ListTeamsParameters) (string, error) {
	// Initialize the parameters string
	paramStr := ""

	// Add mandatory Organization option
	if params.Organization == "" {
		return "", fmt.Errorf("mandatory parameter 'Organization' is missed")
	}
	paramStr += "&filter[organization]=" + params.Organization

	// Add optional Cursor value (from the previous response!)
	// The cursor used for pagination.
	// The value of the cursor must be retrieved from pagination links included in previous responses;
	// you should not attempt to write them on your own.
	if params.Cursor != "" {
		paramStr += "&page[cursor]=" + params.Cursor
	}

	// Add optional Slug value
	if params.Slug != "" {
		paramStr += "&filter[slug]=" + params.Slug
	}

	// Add optional Name value
	if params.Name != "" {
		paramStr += "&filter[name]=" + params.Name
	}

	// Replace the & with ? symbol if the string is not empty
	if len(paramStr) > 0 {
		paramStr = "?" + strings.TrimPrefix(paramStr, "&")
	}

	return paramStr, nil
}

// The function checks the input set of parameters and converts it into a valid URL parameters string
func (t *TransifexApiClient) createGetTeamManagersParametersString(params GetTeamManagersParameters) (string, error) {
	// Initialize the parameters string
	paramStr := ""

	// Add mandatory Team option
	if params.Team == "" {
		return "", fmt.Errorf("mandatory parameter 'Team' is missed")
	}

	// Add optional Cursor value (from the previous response!)
	// The cursor used for pagination.
	// The value of the cursor must be retrieved from pagination links included in previous responses;
	// you should not attempt to write them on your own.
	if params.Cursor != "" {
		paramStr += "&page[cursor]=" + params.Cursor
	}

	// Replace the & with ? symbol if the string is not empty
	if len(paramStr) > 0 {
		paramStr = "?" + strings.TrimPrefix(paramStr, "&")
	}

	return paramStr, nil
}

// The function checks the input set of parameters and converts it into a valid URL parameters string
func (t *TransifexApiClient) createGetTeamManagerRelationshipsParametersString(params GetTeamManagerRelationshipsParameters) (string, error) {
	// Initialize the parameters string
	paramStr := ""

	// Add mandatory Team option
	if params.Team == "" {
		return "", fmt.Errorf("mandatory parameter 'Team' is missed")
	}

	// Add optional Cursor value (from the previous response!)
	// The cursor used for pagination.
	// The value of the cursor must be retrieved from pagination links included in previous responses;
	// you should not attempt to write them on your own.
	if params.Cursor != "" {
		paramStr += "&page[cursor]=" + params.Cursor
	}

	// Replace the & with ? symbol if the string is not empty
	if len(paramStr) > 0 {
		paramStr = "?" + strings.TrimPrefix(paramStr, "&")
	}

	return paramStr, nil
}
