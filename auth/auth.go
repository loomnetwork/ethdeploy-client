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
)

/*

github --- Cli tool

Client ID
57bc10263596c4739845
Client Secret
952e625011098759a2f7ebc02e12c01cf1ef2e80


github --- Website

Client ID
a6abecccefa53842aba4
Client Secret
836d8c9c4dc0c06bfcfdd9089fc59265fdc67a8a

*/
var (
	clientID     = "86zs2w1g2j8hfu"
	clientSecret = "mBd0gDHQdEwSRgt8"
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
}

func validateLoomNetwork(loomnetworkHost string, c *http.Client, network string) string {
	u, err := url.Parse(loomnetworkHost)
	u.Path = path.Join(u.Path, fmt.Sprintf("/login_oauth"))

	resp, err := c.Post(u.String(), "application/loom", nil) // we pass auth info in headers
	if err != nil {
		fmt.Printf("Failed validing againist Loom Network -%v\n", err)
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
	authURL := "https://www.linkedin.com/uas/oauth2/authorization"
	tokenURL := "https://www.linkedin.com/uas/oauth2/accessToken"
	scopes := []string{"r_emailaddress", "r_basicprofile"} // []string{"account"},

	conf := &oauth2.Config{
		ClientID:     clientID,     // also known as slient key sometimes
		ClientSecret: clientSecret, // also known as secret key
		Scopes:       scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  authURL,
			TokenURL: tokenURL,
		},
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

	return nil
}
