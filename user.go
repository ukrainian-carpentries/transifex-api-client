package transifex_api_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type User struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Username string `json:"username"`
	} `json:"attributes"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
}

// Get the details of the user specified by the required path parameter user_id.
// https://developers.transifex.com/reference/get_users-user-id
func (t *TransifexApiClient) GetUserDetails(user_id string) (User, error) {

	// Define the variable to decode the service response
	var u struct {
		Data User `json:"data"`
	}

	// Create an API request
	req, err := http.NewRequest(
		"GET",
		strings.Join([]string{
			t.apiURL,
			"/users/",
			user_id,
		}, ""),
		bytes.NewBuffer(nil))
	if err != nil {
		t.l.Error(err)
		return User{}, err
	}

	// Set authorization and Accept HTTP request headers
	req.Header.Set("Authorization", "Bearer "+t.token)
	req.Header.Add("Accept", "application/vnd.api+json")

	// Perform the request
	resp, err := t.client.Do(req)
	if err != nil {
		t.l.Error(err)
		return User{}, err
	}

	// Decode the JSON response into the corresponding variable
	err = json.NewDecoder(resp.Body).Decode(&u)
	if err != nil {
		t.l.Error(err)
		return User{}, err
	}

	return u.Data, nil
}

// The function prints the information about a user
func (t *TransifexApiClient) PrintUser(u User, formatter string) {
	switch formatter {

	case "text":
		fmt.Printf("  ID: %v\n", u.ID)
		fmt.Printf("  Type: %v\n", u.Type)
		fmt.Printf("  Attributes:\n")
		fmt.Printf("    Username: %v\n", u.Attributes.Username)
		fmt.Printf("  Links:\n")
		fmt.Printf("    Self: %v\n", u.Links.Self)

	case "json":
		text2print, err := json.Marshal(u)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(text2print))

	default:
	}
}
