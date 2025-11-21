package main

import (
	"errors"
	"net/http"
)

func getClientID() error {

	var url = "https://accounts.spotify.com/api/token"

	//fazer um body aqui e passar na request o client id e secret

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return errors.New("requisition error")
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return nil
}

func makeAction() {

}
