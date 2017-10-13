package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
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

func main() {

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

}

func authValidate(token, redirectUri string) {
	tokenURL := "https://www.linkedin.com/oauth/v2/accessToken"
	//
	//token := "AQR6v7_yb2i00wgW_diPb49Wu4-oBCdZyUCoTzqXm-FSPV3iBiM7oWNpfTyDCu58BTmBhsHap3m3BuoFTzrY_tGeKptGE-oxVRKmXQFbtfwLHPE28yV13WZiQQAYi97daGD5lg8NBIo7C_E-6yE"
	//redirect_uri := "http%3A%2F%2F127.0.0.1%3A14565%2Foauth%2Fcallback&response_type=code&scope=r_emailaddress+r_basicprofile&state=state"
	//"http://127.0.0.1:14565/oauth/callback"

	hc := http.Client{}

	fmt.Print(redirectUri)

	form := url.Values{}
	form.Add("grant_type", "authorization_code")
	form.Add("redirect_uri", redirectUri)
	form.Add("client_id", clientID)
	form.Add("client_secret", clientSecret)
	form.Add("code", token)

	req, err := http.NewRequest("POST", tokenURL, strings.NewReader(form.Encode()))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	//req.ParseForm()
	//"application/x-www-form-urlencoded"

	//req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// Save a copy of this request for debugging.
	requestDump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestDump))

	resp, err := hc.Do(req)
	if err != nil {
		log.Fatal(err)
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

}
