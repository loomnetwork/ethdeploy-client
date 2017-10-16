package auth

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
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

func Login(network string, loomnetworkHost string) string {
	n := strings.ToLower(network)
	if strings.Index(n, "linkedin") >= 0 {
		authToken := loginLinkedIn()
		return validateLoomNetwork(authToken, loomnetworkHost)
	} else if strings.Index(n, "github") >= 0 {
		authToken := loginGithub()
		return validateLoomNetwork(authToken, loomnetworkHost)
	} else {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Please enter which network you want (linkedin/github): \n")
		text, _ := reader.ReadString('\n')
		return Login(text, loomnetworkHost)
	}
	return ""
}

func validateLoomNetwork(authToken, loomnetworkHost string) string {
	return ""
}

func loginLinkedIn() string {
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

	peopleURL := "https://api.linkedin.com/v1/people/~:(email-address)?format=json"
	//authValidate(r.Code, r.RedirectURL)
	resp, err := r.Client.Get(peopleURL)

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

	fmt.Printf("%s -\n", r.Token)

	return ""
}

func loginGithub() string {
	fmt.Printf("Attempting to login to Github\n")

	return ""
}
