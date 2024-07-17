package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type apiWhatsapp struct {
	Token  string
	client *http.Client
}

func (ws *apiWhatsapp) ResquestWhatsapp(data DataSendingToken) (result []byte, status int, err error) {
	if ws.client == nil {
		ws.client = &http.Client{}
		ws.Token = "EAAUvdOTJkE4BOzxW0CPt4kWfgChJN8tUGdGY1dpAQ5cmXND1ZAPcnZBl34DnxQG7YrMBjRzV8c3St3V2j9YdY6Dfrf9gdhvj7VyySbCyk1WnV6SSdGTqhD9BCdJ51BYwj40AaZA6gxZAxt6eH2nZAU25BPuZBmxnUyPRChJLtCsJ6oSXtMfpwoGisKvPXzws7zvf8p4E9kXoylfshxsqxqs8zSbEQZD"
	}

	dataJson, err := json.Marshal(data)
	log.Printf(string(dataJson))
	if err != nil {
		log.Fatal("error marshalling data")
		return nil, http.StatusInternalServerError, err
	}

	req, err := http.NewRequest("POST", "https://graph.facebook.com/v19.0/351317348068042/messages", bytes.NewBuffer(dataJson))
	if err != nil {
		return
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", ws.Token))

	resp, err := ws.client.Do(req)
	if resp.Body != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return
	}
	result, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	log.Printf("%v", string(result))
	status = resp.StatusCode
	return
}
