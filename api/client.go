package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/yardstick/terraform-provider-mailtrap/log"
)

type Client struct {
	token string
}

const BASE_URL = "https://mailtrap.io/api/v1"

func NewClient(token string) *Client {

	return &Client{token: token}
}

func (c *Client) Get(url string) (map[string]interface{}, error) {
	uri := BASE_URL + url + "?api_token=" + c.token
	resp, err := http.Get(uri)

	err = handleError(resp, nil, err)
	if err != nil {
		return nil, err
	}

	return responseToJson(resp)
}

func (c *Client) GetArray(url string) ([]map[string]interface{}, error) {
	uri := BASE_URL + url + "?api_token=" + c.token
	resp, err := http.Get(uri)

	err = handleError(resp, nil, err)
	if err != nil {
		return nil, err
	}

	return responseToJsonArray(resp)
}

func (c *Client) Post(url string, body map[string]interface{}) (map[string]interface{}, error) {
	json, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	uri := BASE_URL + url + "?api_token=" + c.token
	resp, err := http.Post(uri, "application/json", bytes.NewBuffer(json))

	err = handleError(resp, json, err)
	if err != nil {
		return nil, err
	}

	return responseToJson(resp)
}

func (c *Client) Patch(url string, body map[string]interface{}) (map[string]interface{}, error) {
	json, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", BASE_URL+url+"?api_token="+c.token, bytes.NewBuffer(json))
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)

	err = handleError(resp, json, err)
	if err != nil {
		return nil, err
	}

	return responseToJson(resp)
}

func (c *Client) Delete(url string) (map[string]interface{}, error) {
	req, err := http.NewRequest("DELETE", BASE_URL+url+"?api_token="+c.token, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)

	err = handleError(resp, nil, err)
	if err != nil {
		return nil, err
	}

	return responseToJson(resp)
}

func handleError(resp *http.Response, json []byte, err error) error {
	if err != nil {
		log.Error(resp.Request.Method + " Not Successful")
		log.Error("URL: " + resp.Request.RequestURI)
		log.Error("REQUEST BODY: " + string(json))
		log.Error("STATUS:" + resp.Status)
		return err
	}

	if resp.StatusCode != http.StatusOK {
		defer resp.Request.Body.Close()
		reqBody, _ := ioutil.ReadAll(resp.Request.Body)
		log.Error("NON-OK Status Code for Request")
		log.Error("URL: " + resp.Request.URL.String())
		log.Error("BODY: " + string(reqBody))
		log.Error("STATUS: " + resp.Status)
		return errors.New("NON-OK Status Code Returned From Mailtrap")
	}

	return nil
}

func responseToJson(resp *http.Response) (map[string]interface{}, error) {
	defer resp.Body.Close()
	var data map[string]interface{}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("Could not read body: " + err.Error())
		return nil, err
	}

	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		log.Error("Could not transform to Map: " + err.Error())
		return nil, err
	}

	return data, nil
}

func responseToJsonArray(resp *http.Response) ([]map[string]interface{}, error) {
	defer resp.Body.Close()
	var data []map[string]interface{}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("Could not read body: " + err.Error())
		return nil, err
	}

	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		log.Error("Could not transform to Map: " + err.Error())
		return nil, err
	}

	return data, nil
}
