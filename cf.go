package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"bytes"
)

func (dns *Dns) Handle() {
	httpClient := http.Client{}

	getDnsUrl := ApiBaseUrl + "zones/" + Config.Zone + "/dns_records?type=A&name=" + dns.Name
	putDnsUrl := ApiBaseUrl + "zones/" + Config.Zone + "/dns_records/"

	req, err := http.NewRequest("GET", getDnsUrl, nil)

	if err != nil {
		Quit("could not create request on " + dns.Name)
	}

	req.Header = http.Header{
		"Authorization": []string{"Bearer " + Config.Key},
		"Content-Type":  []string{"application/json"},
	}

	resp, err := httpClient.Do(req)

	if err != nil {
		Quit("could not fetch dns " + dns.Name)
	}

	rawBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		Quit("could not parse body on dns " + dns.Name)
	}

	var cfDns CloudflareDnsResult

	err = json.Unmarshal(rawBody, &cfDns)

	if err != nil {
		Quit("could not parse cloudflare dns entry")
	}

	if len(cfDns.Result) != 1 {
		Quit("error con retrieving dns list")
	}

	postBody, _ := json.Marshal(map[string]interface{}{
		"content": PublicIp,
		"type":    "A",
		"name":    dns.Name,
		"ttl":     strconv.Itoa(cfDns.Result[0].Ttl),
		"proxied": dns.Proxy,
	})

	req, err = http.NewRequest("PUT", putDnsUrl+cfDns.Result[0].Id, bytes.NewBuffer(postBody))

	if err != nil {
		Quit("could not create request to update on " + dns.Name)
	}

	req.Header = http.Header{
		"Authorization": []string{"Bearer " + Config.Key},
		"Content-Type":  []string{"application/json"},
	}

	resp, err = httpClient.Do(req)

	if err != nil {
		Quit("could not update dns " + dns.Name)
	}
}