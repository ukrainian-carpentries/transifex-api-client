package transifex_api_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Project struct {
	Attributes struct {
		Archived                 bool     `json:"archived"`
		DatetimeCreated          string   `json:"datetime_created"`
		DatetimeModified         string   `json:"datetime_modified"`
		Description              string   `json:"description"`
		HomepageURL              string   `json:"homepage_url"`
		InstructionsURL          string   `json:"instructions_url"`
		License                  string   `json:"license"`
		LogoURL                  string   `json:"logo_url"`
		LongDescription          string   `json:"long_description"`
		MachineTranslationFillup bool     `json:"machine_translation_fillup"`
		Name                     string   `json:"name"`
		Private                  bool     `json:"private"`
		RepositoryURL            string   `json:"repository_url"`
		Slug                     string   `json:"slug"`
		Tags                     []string `json:"tags"`
		TranslationMemoryFillup  bool     `json:"translation_memory_fillup"`
		Type                     string   `json:"type"`
	} `json:"attributes"`
	ID    string `json:"id"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Relationships struct {
		Languages struct {
			Links struct {
				Related string `json:"related"`
				Self    string `json:"self"`
			} `json:"links"`
		} `json:"languages"`
		Maintainers struct {
			Links struct {
				Related string `json:"related"`
				Self    string `json:"self"`
			} `json:"links"`
		} `json:"maintainers"`
		Organization struct {
			Data struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
			Links struct {
				Related string `json:"related"`
			} `json:"links"`
		} `json:"organization"`
		Resources struct {
			Links struct {
				Related string `json:"related"`
			} `json:"links"`
		} `json:"resources"`
		SourceLanguage struct {
			Data struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
			Links struct {
				Related string `json:"related"`
			} `json:"links"`
		} `json:"source_language"`
		Team struct {
			Data struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
			Links struct {
				Related string `json:"related"`
				Self    string `json:"self"`
			} `json:"links"`
		} `json:"team"`
	} `json:"relationships"`
	Type string `json:"type"`
}

type ListProjectsParameters struct {
	Organization string
	Cursor       string
	Slug         string
	Name         string
}

type GetProjectMaintainersParameters struct {
	Project_id string
	Cursor     string
}

type ListLanguageRelationshipsParameters struct {
	Project_id string
	Cursor     string
}

type GetProjectMaintainerRelationshipsParameters struct {
	Project_id string
	Cursor     string
}

// Get the list of projects that belong to a single organization.
// https://developers.transifex.com/reference/get_projects
func (t *TransifexApiClient) ListProjects(params ListProjectsParameters) ([]Project, error) {

	paramStr, err := t.createListProjectsParametersString(params)
	if err != nil {
		return nil, err
	}

	// Define the variable to decode the service response
	var lpr struct {
		Data  []Project `json:"data"`
		Links struct {
			Next     string `json:"next"`
			Previous string `json:"previous"`
			Self     string `json:"self"`
		} `json:"links"`
	}

	// Create an API request
	req, err := http.NewRequest(
		"GET",
		strings.Join([]string{
			t.apiURL,
			"/projects",
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
	err = json.NewDecoder(resp.Body).Decode(&lpr)
	if err != nil {
		t.l.Error(err)
		return nil, err
	}

	return lpr.Data, nil
}

// Get the details of a specific project.
// https://developers.transifex.com/reference/get_projects-project-id
func (t *TransifexApiClient) GetProjectDetails(project_id string) (Project, error) {

	// Define the variable to decode the service response
	var pd struct {
		Data Project `json:"data"`
	}

	// Create an API request
	req, err := http.NewRequest(
		"GET",
		strings.Join([]string{
			t.apiURL,
			"/projects/",
			project_id,
		}, ""),
		bytes.NewBuffer(nil))
	if err != nil {
		t.l.Error(err)
		return Project{}, err
	}

	// Set authorization and Accept HTTP request headers
	req.Header.Set("Authorization", "Bearer "+t.token)
	req.Header.Add("Accept", "application/vnd.api+json")

	// Perform the request
	resp, err := t.client.Do(req)
	if err != nil {
		t.l.Error(err)
		return Project{}, err
	}

	// Decode the JSON response into the corresponding variable
	err = json.NewDecoder(resp.Body).Decode(&pd)
	if err != nil {
		t.l.Error(err)
		return Project{}, err
	}

	return pd.Data, nil
}

// Get a list of all target languages of a specific project.
// https://developers.transifex.com/reference/get_projects-project-id-languages
func (t *TransifexApiClient) ListProjectLanguages(project_id string) ([]Language, error) {

	// Define the variable to decode the service response
	var pl struct {
		Data []Language `json:"data"`
	}

	// Create an API request
	req, err := http.NewRequest(
		"GET",
		strings.Join([]string{
			t.apiURL,
			"/projects/",
			project_id,
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
	err = json.NewDecoder(resp.Body).Decode(&pl)
	if err != nil {
		t.l.Error(err)
		return nil, err
	}

	return pl.Data, nil
}

// Get project maintainers.
// https://developers.transifex.com/reference/get_projects-project-id-maintainers
func (t *TransifexApiClient) GetProjectMaintainers(params GetProjectMaintainersParameters) ([]Maintainer, error) {

	paramStr, err := t.createGetProjectMaintainersParametersString(params)
	if err != nil {
		return nil, err
	}

	// Define the variable to decode the service response
	var pl struct {
		Data []Maintainer `json:"data"`
	}

	// Create an API request
	req, err := http.NewRequest(
		"GET",
		strings.Join([]string{
			t.apiURL,
			"/projects/",
			params.Project_id,
			"/maintainers",
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
	err = json.NewDecoder(resp.Body).Decode(&pl)
	if err != nil {
		t.l.Error(err)
		return nil, err
	}

	return pl.Data, nil
}

// List language relationships.
// https://developers.transifex.com/reference/get_projects-project-id-relationships-languages
func (t *TransifexApiClient) GetLanguageRelationships(params ListLanguageRelationshipsParameters) ([]LanguageRelationship, error) {

	paramStr, err := t.createListLanguageRelationshipsParametersString(params)
	if err != nil {
		return nil, err
	}

	// Define the variable to decode the service response
	var lr struct {
		Data  []LanguageRelationship `json:"data"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	}

	// Create an API request
	req, err := http.NewRequest(
		"GET",
		strings.Join([]string{
			t.apiURL,
			"/projects/",
			params.Project_id,
			"/relationships/languages",
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
	err = json.NewDecoder(resp.Body).Decode(&lr)
	if err != nil {
		t.l.Error(err)
		return nil, err
	}

	return lr.Data, nil
}

// Get project maintainer relationships.
// https://developers.transifex.com/reference/get_projects-project-id-relationships-maintainers
func (t *TransifexApiClient) GetProjectMaintainerRelationships(params GetProjectMaintainerRelationshipsParameters) ([]MaintainerRelationship, error) {

	paramStr, err := t.createGetProjectMaintainerRelationshipsParametersString(params)
	if err != nil {
		return nil, err
	}

	// Define the variable to decode the service response
	var pmr struct {
		Data  []MaintainerRelationship `json:"data"`
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
			"/projects/",
			params.Project_id,
			"/relationships/maintainers",
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
	err = json.NewDecoder(resp.Body).Decode(&pmr)
	if err != nil {
		t.l.Error(err)
		return nil, err
	}

	return pmr.Data, nil
}

// Get team relationship.
// https://developers.transifex.com/reference/get_projects-project-id-relationships-team
func (t *TransifexApiClient) GetTeamRelationship(project_id string) (TeamRelationship, error) {

	// Define the variable to decode the service response
	var tr struct {
		Data  TeamRelationship `json:"data"`
		Links struct {
			Related string `json:"related"`
			Self    string `json:"self"`
		} `json:"links"`
	}

	// Create an API request
	req, err := http.NewRequest(
		"GET",
		strings.Join([]string{
			t.apiURL,
			"/projects/",
			project_id,
			"/relationships/team",
		}, ""),
		bytes.NewBuffer(nil))
	if err != nil {
		t.l.Error(err)
		return TeamRelationship{}, err
	}

	// Set authorization and Accept HTTP request headers
	req.Header.Set("Authorization", "Bearer "+t.token)
	req.Header.Add("Accept", "application/vnd.api+json")

	// Perform the request
	resp, err := t.client.Do(req)
	if err != nil {
		t.l.Error(err)
		return TeamRelationship{}, err
	}

	// Decode the JSON response into the corresponding variable
	err = json.NewDecoder(resp.Body).Decode(&tr)
	if err != nil {
		t.l.Error(err)
		return TeamRelationship{}, err
	}

	return tr.Data, nil
}

// The function prints the information about a project
func (t *TransifexApiClient) PrintProject(p Project, formatter string) {

	switch formatter {

	case "text":
		fmt.Printf("  ID: %v\n", p.ID)
		fmt.Printf("  Type: %v\n", p.Type)
		fmt.Printf("  Attributes:\n")
		fmt.Printf("    Slug: %v\n", p.Attributes.Slug)
		fmt.Printf("    Name: %v\n", p.Attributes.Name)
		fmt.Printf("    Type: %v\n", p.Attributes.Type)
		fmt.Printf("    DatetimeCreated: %v\n", p.Attributes.DatetimeCreated)
		fmt.Printf("    DatetimeModified: %v\n", p.Attributes.DatetimeModified)
		if len(p.Attributes.Tags) > 0 {
			fmt.Printf("    Tags:\n")
			for _, v := range p.Attributes.Tags {
				fmt.Printf("      - %v\n", v)
			}
		}
		fmt.Printf("    Description: %v\n", p.Attributes.Description)
		fmt.Printf("    LongDescription: %v\n", p.Attributes.LongDescription)
		fmt.Printf("    Private: %v\n", p.Attributes.Private)
		fmt.Printf("    Archived: %v\n", p.Attributes.Archived)
		fmt.Printf("    TranslationMemoryFillup: %v\n", p.Attributes.TranslationMemoryFillup)
		fmt.Printf("    MachineTranslationFillup: %v\n", p.Attributes.MachineTranslationFillup)
		fmt.Printf("    HomepageURL: %v\n", p.Attributes.HomepageURL)
		fmt.Printf("    RepositoryURL: %v\n", p.Attributes.RepositoryURL)
		fmt.Printf("    InstructionsURL: %v\n", p.Attributes.InstructionsURL)
		fmt.Printf("    License: %v\n", p.Attributes.License)
		fmt.Printf("    LogoURL: %v\n", p.Attributes.LogoURL)

		fmt.Printf("  Relationships:\n")
		fmt.Printf("    Organization:\n")
		fmt.Printf("      Links:\n")
		fmt.Printf("        Related: %v\n", p.Relationships.Organization.Links.Related)
		fmt.Printf("      Data:\n")
		fmt.Printf("        Type: %v\n", p.Relationships.Organization.Data.Type)
		fmt.Printf("        ID: %v\n", p.Relationships.Organization.Data.ID)

		fmt.Printf("    SourceLanguage:\n")
		fmt.Printf("      Links:\n")
		fmt.Printf("        Related: %v\n", p.Relationships.SourceLanguage.Links.Related)
		fmt.Printf("      Data:\n")
		fmt.Printf("        Type: %v\n", p.Relationships.SourceLanguage.Data.Type)
		fmt.Printf("        ID: %v\n", p.Relationships.SourceLanguage.Data.ID)
		fmt.Printf("    Languages:\n")
		fmt.Printf("      Links:\n")
		fmt.Printf("        Self: %v\n", p.Relationships.Languages.Links.Self)
		fmt.Printf("        Related: %v\n", p.Relationships.Languages.Links.Related)
		fmt.Printf("    Team:\n")
		fmt.Printf("      Data:\n")
		fmt.Printf("        Type: %v\n", p.Relationships.Team.Data.Type)
		fmt.Printf("        ID: %v\n", p.Relationships.Team.Data.ID)
		fmt.Printf("      Links:\n")
		fmt.Printf("        Related: %v\n", p.Relationships.Team.Links.Related)
		fmt.Printf("        Self: %v\n", p.Relationships.Team.Links.Self)
		fmt.Printf("    Maintainers:\n")
		fmt.Printf("      Links:\n")
		fmt.Printf("        Related: %v\n", p.Relationships.Maintainers.Links.Related)
		fmt.Printf("        Self: %v\n", p.Relationships.Maintainers.Links.Self)
		fmt.Printf("    Resources:\n")
		fmt.Printf("      Links:\n")
		fmt.Printf("        Related: %v\n", p.Relationships.Resources.Links.Related)
		fmt.Printf("  Links:\n")
		fmt.Printf("    Self: %v\n", p.Links.Self)

	case "json":
		text2print, err := json.Marshal(p)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(text2print))

	default:
	}
}

// The function checks the input set of parameters and converts it into a valid URL parameters string
func (t *TransifexApiClient) createListProjectsParametersString(params ListProjectsParameters) (string, error) {
	// Initialize the parameters string
	paramStr := ""

	// Add mandatory Organization ID option
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
func (t *TransifexApiClient) createGetProjectMaintainersParametersString(params GetProjectMaintainersParameters) (string, error) {
	// Initialize the parameters string
	paramStr := ""

	// Add optional Code value
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
func (t *TransifexApiClient) createListLanguageRelationshipsParametersString(params ListLanguageRelationshipsParameters) (string, error) {
	// Initialize the parameters string
	paramStr := ""

	// Add optional Code value
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
func (t *TransifexApiClient) createGetProjectMaintainerRelationshipsParametersString(params GetProjectMaintainerRelationshipsParameters) (string, error) {
	// Initialize the parameters string
	paramStr := ""

	// Add optional Code value
	if params.Cursor != "" {
		paramStr += "&page[cursor]=" + params.Cursor
	}

	// Replace the & with ? symbol if the string is not empty
	if len(paramStr) > 0 {
		paramStr = "?" + strings.TrimPrefix(paramStr, "&")
	}

	return paramStr, nil
}
