package main

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func GoQuery() {

	Load("")

	// create a new http transport to ignore self-signed certificates
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	// create a new http client with the transport
	client := &http.Client{Transport: transport}

	// Request the HTML page.
	res, err := client.Get("https://ibank.ubagroup.com")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// get the input for user id
	userIdElement := doc.Find("type_UserPrincipal ui-autocomplete-input")

	// add the user id as value to the input element
	userIdElement.SetAttr("value", Get("USER_ID", ""))
}
