package auth

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	oauth2ns "github.com/loomnetwork/oauth2-noserver"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/linkedin"
)

type LoginAuth struct {
	Email  string `json:"email,omitempty" form:"id"`
	ApiKey string `json:"apikey,omitempty" form:"id"`
}

func Login(network string, loomnetworkHost string) string {
	n := strings.ToLower(network)
	var c *http.Client
	if strings.Index(n, "linkedin") >= 0 {
		c = loginLinkedIn()
	} else if strings.Index(n, "github") >= 0 {
		c = loginGithub()
	} else {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Please enter which network you want (linkedin/github): \n")
		text, _ := reader.ReadString('\n')
		return Login(text, loomnetworkHost)
	}
	if c == nil {
		fmt.Printf("Failed logging into %s\n", network)
	}
	return validateLoomNetwork(loomnetworkHost, c, n)
	//return extractGithubEmail(c, "")
}

type GithubEmail struct {
	Email    string `json:"email,omitempty"`
	Verified bool   `json:"verified,omitempty""`
	Primary  bool   `json:"primary,omitempty""`
}

func extractGithubEmail(c *http.Client, auth string) string {
	githubEmailURL := "https://api.github.com/user/emails"
	req, err := http.NewRequest("GET", githubEmailURL, nil)
	if auth != "" {
		req.Header.Add("Authorization", auth)
	}
	resp, err := c.Do(req)
	if err != nil {
		fmt.Println("error:", err)
		return ""
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 { // OK
		fmt.Printf("bad response code %d\n", resp.StatusCode)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var gemails []GithubEmail
	err = json.Unmarshal(bodyBytes, &gemails)
	if err != nil {
		fmt.Println("error:", err)
	}
	if len(gemails) > 0 {
		for _, email := range gemails {
			if email.Verified == true && email.Primary == true {
				return email.Email
			}
		}
	}

	return ""
}

func validateLoomNetwork(loomnetworkHost string, c *http.Client, network string) string {
	u, err := url.Parse(loomnetworkHost)
	u.Path = path.Join(u.Path, fmt.Sprintf("/login_oauth"))

	req, err := http.NewRequest("POST", u.String(), nil)
	req.Header.Add("Loom-Oauth-Provider", network)
	resp, err := c.Do(req)

	defer resp.Body.Close()

	if resp.StatusCode != 200 { // OK
		fmt.Printf("bad response code %d\n", resp.StatusCode)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//TODO remove this
	bodyString := string(bodyBytes)
	fmt.Printf(bodyString)

	var lauth LoginAuth
	err = json.Unmarshal(bodyBytes, &lauth)
	if err != nil {
		fmt.Println("error:", err)
		return ""
	}

	return lauth.ApiKey
}

func loginLinkedIn() *http.Client {
	clientID := "86zs2w1g2j8hfu"
	clientSecret := "mBd0gDHQdEwSRgt8"

	scopes := []string{"r_emailaddress", "r_basicprofile"}

	conf := &oauth2.Config{
		ClientID:     clientID,     // also known as slient key sometimes
		ClientSecret: clientSecret, // also known as secret key
		Scopes:       scopes,
		Endpoint:     linkedin.Endpoint,
	}
	r := oauth2ns.Authorize(conf)
	// use client.Get / client.Post for further requests, the token will automatically be there

	if len(r.Token.AccessToken) > 0 {
		return r.Client
	}
	return nil
}

func loginGithub() *http.Client {
	fmt.Printf("Attempting to login to Github\n")
	clientID := "57bc10263596c4739845"
	clientSecret := "952e625011098759a2f7ebc02e12c01cf1ef2e80"
	scopes := []string{"user:email"}

	conf := &oauth2.Config{
		ClientID:     clientID,     // also known as slient key sometimes
		ClientSecret: clientSecret, // also known as secret key
		Scopes:       scopes,
		Endpoint:     github.Endpoint,
	}
	r := oauth2ns.Authorize(conf)
	// use client.Get / client.Post for further requests, the token will automatically be there

	if len(r.Token.AccessToken) > 0 {
		return r.Client
	}

	return nil
}
