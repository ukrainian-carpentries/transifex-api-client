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

type ResourceStringComment struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Category         string `json:"category"`
		Message          string `json:"message"`
		Priority         string `json:"priority"`
		Status           string `json:"status"`
		Type             string `json:"type"`
		DatetimeCreated  string `json:"datetime_created"`
		DatetimeModified string `json:"datetime_modified"`
		DatetimeResolved string `json:"datetime_resolved"`
	} `json:"attributes"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Relationships struct {
		Author struct {
			Data struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
			Links struct {
				Related string `json:"related"`
			} `json:"links"`
		} `json:"author"`
		Language struct {
			Data struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
			Links struct {
				Related string `json:"related"`
			} `json:"links"`
		} `json:"language"`
		Resolver struct {
			Data struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
			Links struct {
				Related string `json:"related"`
			} `json:"links"`
		} `json:"resolver"`
		Resource struct {
			Data struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
			Links struct {
				Related string `json:"related"`
			} `json:"links"`
		} `json:"resource"`
		ResourceString struct {
			Data struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
			Links struct {
				Related string `json:"related"`
			} `json:"links"`
		} `json:"resource_string"`
	} `json:"relationships"`
}

type ListResourceStringCommentsParameters struct {
	Organization          string
	Project               string
	Cursor                string
	Category              string
	Author                string
	DatetimeCreatedAfter  time.Time
	DatetimeCreatedBefore time.Time
	Priority              string
	Resource              string
	ResourceString        string
	Status                string
	Type                  string
}

// Get resource strings collection.
// Get a list of all resource string comments for an organization. You can further narrow down the list using the available filters.
// https://developers.transifex.com/reference/get_resource-string-comments
func (t *TransifexApiClient) ListResourceStringComments(params ListResourceStringCommentsParameters) ([]ResourceStringComment, error) {

	paramStr, err := t.createListResourceStringCommentsParametersString(params)
	if err != nil {
		return nil, err
	}

	// Define the variable to decode the service response
	var rscomm struct {
		Data  []ResourceStringComment `json:"data"`
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
			"/resource_string_comments",
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
	err = json.NewDecoder(resp.Body).Decode(&rscomm)
	if err != nil {
		t.l.Error(err)
		return nil, err
	}

	return rscomm.Data, nil
}

// Get resource strings collection.
// Get a list of all resource string comments for an organization. You can further narrow down the list using the available filters.
// https://developers.transifex.com/reference/get_resource-string-comments
func (t *TransifexApiClient) GetResourceStringComment(comment_id string) (ResourceStringComment, error) {

	// Define the variable to decode the service response
	var rscomm struct {
		Data ResourceStringComment `json:"data"`
	}

	// Create an API request
	req, err := http.NewRequest(
		"GET",
		strings.Join([]string{
			t.apiURL,
			"/resource_string_comments/",
			comment_id,
		}, ""),
		bytes.NewBuffer(nil))
	if err != nil {
		t.l.Error(err)
		return ResourceStringComment{}, err
	}

	// Set authorization and Accept HTTP request headers
	req.Header.Set("Authorization", "Bearer "+t.token)
	req.Header.Add("Accept", "application/vnd.api+json")

	// Perform the request
	resp, err := t.client.Do(req)
	if err != nil {
		t.l.Error(err)
		return ResourceStringComment{}, err
	}

	// Decode the JSON response into the corresponding variable
	err = json.NewDecoder(resp.Body).Decode(&rscomm)
	if err != nil {
		t.l.Error(err)
		return ResourceStringComment{}, err
	}

	return rscomm.Data, nil
}

// The function prints the information about a resource string comment
func (t *TransifexApiClient) PrintResourceStringComment(c ResourceStringComment, formatter string) {

	switch formatter {

	case "text":
		fmt.Printf("  ID: %v\n", c.ID)
		fmt.Printf("  Type: %v\n", c.Type)
		fmt.Printf("  Attributes:\n")
		fmt.Printf("    Category: %v\n", c.Attributes.Category)
		fmt.Printf("    Message: %v\n", c.Attributes.Message)
		fmt.Printf("    Priority: %v\n", c.Attributes.Priority)
		fmt.Printf("    Status: %v\n", c.Attributes.Status)
		fmt.Printf("    Type: %v\n", c.Attributes.Type)
		fmt.Printf("    DatetimeCreated: %v\n", c.Attributes.DatetimeCreated)
		fmt.Printf("    DatetimeModified: %v\n", c.Attributes.DatetimeModified)
		fmt.Printf("    DatetimeResolved: %v\n", c.Attributes.DatetimeResolved)
		fmt.Printf("  Links:\n")
		fmt.Printf("    Self: %v\n", c.Links.Self)
		fmt.Printf("  Relationships:\n")
		fmt.Printf("    Author:\n")
		fmt.Printf("      Data:\n")
		fmt.Printf("        ID: %v\n", c.Relationships.Author.Data.ID)
		fmt.Printf("        Type: %v\n", c.Relationships.Author.Data.Type)
		fmt.Printf("      Links:\n")
		fmt.Printf("        Related: %v\n", c.Relationships.Author.Links.Related)
		fmt.Printf("    Language:\n")
		fmt.Printf("      Data:\n")
		fmt.Printf("        ID: %v\n", c.Relationships.Language.Data.ID)
		fmt.Printf("        Type: %v\n", c.Relationships.Language.Data.Type)
		fmt.Printf("      Links:\n")
		fmt.Printf("        Related: %v\n", c.Relationships.Language.Links.Related)
		fmt.Printf("    Resolver:\n")
		fmt.Printf("      Data:\n")
		fmt.Printf("        ID: %v\n", c.Relationships.Resolver.Data.ID)
		fmt.Printf("        Type: %v\n", c.Relationships.Resolver.Data.Type)
		fmt.Printf("      Links:\n")
		fmt.Printf("        Related: %v\n", c.Relationships.Resolver.Links.Related)
		fmt.Printf("    Resource:\n")
		fmt.Printf("      Data:\n")
		fmt.Printf("        ID: %v\n", c.Relationships.Resource.Data.ID)
		fmt.Printf("        Type: %v\n", c.Relationships.Resource.Data.Type)
		fmt.Printf("      Links:\n")
		fmt.Printf("        Related: %v\n", c.Relationships.Resource.Links.Related)
		fmt.Printf("    ResourceString:\n")
		fmt.Printf("      Data:\n")
		fmt.Printf("        ID: %v\n", c.Relationships.ResourceString.Data.ID)
		fmt.Printf("        Type: %v\n", c.Relationships.ResourceString.Data.Type)
		fmt.Printf("      Links:\n")
		fmt.Printf("        Related: %v\n", c.Relationships.ResourceString.Links.Related)

	case "json":
		text2print, err := json.Marshal(c)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(text2print))

	default:
	}
}

// The function checks the input set of parameters and converts it into a valid URL parameters string
func (t *TransifexApiClient) createListResourceStringCommentsParametersString(params ListResourceStringCommentsParameters) (string, error) {
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

	// Add optional Project value
	if params.Project != "" {
		paramStr += "&filter[project]=" + params.Project
	}

	// Add optional Category value
	if params.Category != "" {
		paramStr += "&filter[category]=" + params.Category
	}

	// Add optional Author value
	if params.Author != "" {
		paramStr += "&filter[author]=" + params.Author
	}

	// Add optional datetime_created->gte value
	if (params.DatetimeCreatedAfter != time.Time{}) {
		paramStr += "&filter[datetime_created][gte]=" + params.DatetimeCreatedAfter.Format("2006-01-02T15:04:05Z")
	}

	// Add optional datetime_created->lt value
	if (params.DatetimeCreatedBefore != time.Time{}) {
		paramStr += "&filter[datetime_created][lt]=" + params.DatetimeCreatedBefore.Format("2006-01-02T15:04:05Z")
	}

	// Add optional datetime_created->lt value
	if (params.DatetimeCreatedBefore != time.Time{}) {
		paramStr += "&filter[datetime_created][lt]=" + params.DatetimeCreatedBefore.Format("2006-01-02T15:04:05Z")
	}

	// Add allowed Priority option
	switch strings.ToLower(params.Priority) {
	case "low":
		fallthrough
	case "normal":
		fallthrough
	case "high":
		fallthrough
	case "critical":
		fallthrough
	case "blocker":
		paramStr += "&filter[priority]=" + strings.ToLower(params.Priority)
	case "":
	default:
		return "", fmt.Errorf("unknown 'Priority' value")
	}

	// Add Resource option
	if params.Resource != "" {
		paramStr += "&filter[resource]=" + params.Resource
	}

	// Add Resource String option
	if params.ResourceString != "" {
		paramStr += "&filter[resource_string]=" + params.ResourceString
	}

	// Add allowed Status option
	switch strings.ToLower(params.Status) {
	case "open":
		fallthrough
	case "resolved":
		paramStr += "&filter[status]=" + strings.ToLower(params.Status)
	case "":
	default:
		return "", fmt.Errorf("unknown 'Status' value")
	}

	// Add allowed Type option
	switch strings.ToLower(params.Type) {
	case "issue":
		fallthrough
	case "comment":
		paramStr += "&filter[type]=" + strings.ToLower(params.Type)
	case "":
	default:
		return "", fmt.Errorf("unknown 'Type' value")
	}

	// Replace the & with ? symbol if the string is not empty
	if len(paramStr) > 0 {
		paramStr = "?" + strings.TrimPrefix(paramStr, "&")
	}

	return paramStr, nil
}
