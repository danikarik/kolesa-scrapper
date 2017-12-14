package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// GetPhoneNumber получает номера
func GetPhoneNumber(id int) (string, error) {

	showURL := fmt.Sprintf("%s%d", SHOWURL, id)
	phoneURL := fmt.Sprintf("%s?id=%d", PHONEURL, id)

	log.Printf("parsing: %s\n", showURL)

	client := &http.Client{}
	req, err := http.NewRequest("GET", phoneURL, nil)
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	if err != nil {
		return "", err
	}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	ajaxModel := &AjaxModel{}
	err = json.Unmarshal(body, ajaxModel)
	if err != nil {
		return "", err
	}

	return ajaxModel.Data.Model.Phone, nil
}
