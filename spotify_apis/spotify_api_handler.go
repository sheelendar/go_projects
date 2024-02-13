package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

var CurrentAccessToken = ""
var dbPool *pgxpool.Pool

func main() {
	dbPool = getDBIsntance()
	InitHttpClient()

}

// searchTracks search track from spotify api and return
func searchTracks(query string) *SearchTrackResponse {
	if CurrentAccessToken == "" {
		fmt.Println("access token not found in response")
		connectWithSpotify()
	}
	queryURL := fmt.Sprintf("%s?q=%s&type=track&limit=5&offset=0", SearchURL, query)
	respnse := &SearchTrackResponse{}
	executGettRequest(queryURL, CurrentAccessToken, respnse)
	return respnse
}

// connectWithSpotify check currentAccesstoken and update it.
func connectWithSpotify() {
	CurrentAccessToken = getAccessToken()
	fmt.Println(CurrentAccessToken)
}

// executGettRequest execute Get request and request URL.
func executGettRequest(getURL, accessToken string, response interface{}) interface{} {

	req, err := http.NewRequest(http.MethodGet, getURL, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error while requesting url: ", getURL, " Error: ", err)
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(response)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	if res.StatusCode != http.StatusOK {
		fmt.Println(res.StatusCode)
	}
	return response
}

// getAccessToken hit spotify auth api and return access token.
func getAccessToken() string {
	authString := base64.StdEncoding.EncodeToString([]byte(ClientID + ":" + ClientSecret))

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", ClientID) // Add the scopes your application needs
	data.Set("client_secret", ClientSecret)

	req, err := http.NewRequest(http.MethodPost, TokenURL, ioutil.NopCloser(strings.NewReader(data.Encode())))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Authorization", "Basic "+authString)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error while requesting url: ", TokenURL, " Error: ", err)
	}

	defer res.Body.Close()
	resposne := &AccessTokenResponse{}
	err = json.NewDecoder(res.Body).Decode(resposne)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	if res.StatusCode != http.StatusOK {
		fmt.Println(res.StatusCode)
	}
	fmt.Println(resposne)
	return resposne.AccessToken
}

//  http://localhost:8080/errors
//  client ID :  3807b039fed240b5a7dcfab3cad4717d
//  client secret :  f3ec4d5cc2d847d5bdd269a7200d9ecd

/*
	curl -X POST "https://accounts.spotify.com/api/token" \
	     -H "Content-Type: application/x-www-form-urlencoded" \
	     -d "grant_type=client_credentials&client_id=3807b039fed240b5a7dcfab3cad4717d&client_secret=f3ec4d5cc2d847d5bdd269a7200d9ecd"

//{"access_token":"BQATNGpuooH2ZEPL-iYZuiYmC4uiW4s0RDUqg8VxtU0hea5EBAdiEQNsdgbzS4ITe3z_Tw784dNxiy92VeCNX49-crVYpPZrWlQnubAlURd7GTwOPcw","token_type":"Bearer","expires_in":3600}%

	curl --request GET \
	  --url 'https://api.spotify.com/v1/search?q=remaster&type=track&limit=5&offset=0' \
	  --header 'Authorization: Bearer BQAv4ixJvt7O1ldFj15PCr1D9Y8DoQ_WRng0LuBJaIFzjuQEw8rE8QR7wUA7Fy3QVkYwCAu-v52CdSo3-0s6iJ-NjKdcGuE0JtYWTuYPn4T-QgWthJ4'
*/
