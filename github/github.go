package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Eldius/github-cli/config"
)

func GenerateDeviceCode2() {
	c := http.Client{
		Timeout: 2 * time.Second,
	}

	//http.NewRequest(http.MethodPost, config.GetDeviceVerificationCodeUri(), )
	var payload = []byte(
		fmt.Sprintf(
			`{"client_id":"%s", "scope": "%s"}`,
			config.GetClientID(),
			strings.Join(config.GetScopes(), ", ")),
	)
	res, err := c.Post(config.GetDeviceVerificationCodeUri(), "application/json", bytes.NewBuffer(payload))
	if err != nil {
		log.Fatalln("Failed to request values:", err.Error())
	}
	defer res.Body.Close()

	bBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("Failed to read response:", err.Error())
	}

	fmt.Printf("Response:\n%s\n", string(bBody))

}

func GenerateDeviceCode() {
	c := http.Client{
		Timeout: 2 * time.Second,
	}

	req, err := http.NewRequest(http.MethodPost, config.GetDeviceVerificationCodeUri(), nil)
	if err != nil {
		log.Fatalln("Failed to create request:", err.Error())
	}
	req.Header.Set("accept", "application/json")
	q := req.URL.Query()
	q.Add("client_id", config.GetClientID())
	q.Add("scope", strings.Join(config.GetScopes(), ", "))
	req.URL.RawQuery = q.Encode()

	res, err := c.Do(req)
	if err != nil {
		log.Fatalln("Failed to request values:", err.Error())
	}
	defer res.Body.Close()

	var response *DeviceCodeResponse
	json.NewDecoder(res.Body).Decode(&response)

	fmt.Printf("Response:\n%v\n", response)

	fmt.Printf(
		`Open the url above in your browser and put the code '%s'
		%s

`,
		response.UserCode,
		response.VerificationURI,
	)

	fmt.Println("Press the Enter Key after confirmed your user ID.")
	fmt.Scanln()

	req1, err := http.NewRequest(http.MethodPost, config.GetAccessCodeUri(), nil)
	if err != nil {
		log.Fatalln("Failed to create request:", err.Error())
	}
	req1.Header.Set("accept", "application/json")
	q1 := req.URL.Query()
	q1.Add("client_id", config.GetClientID())
	q1.Add("device_code", response.DeviceCode)
	q1.Add("grant_type", config.GetGrantType())
	req1.URL.RawQuery = q1.Encode()

	res1, err := c.Do(req1)
	if err != nil {
		log.Fatalln("Failed to request values:", err.Error())
	}
	defer res1.Body.Close()

	/* 	bBody, err := ioutil.ReadAll(res1.Body)
	   	if err != nil {
	   		log.Fatalln("Failed to read response body:", err.Error())
	   	}

	   	fmt.Println("response body:", string(bBody))
	*/

	var response1 AccessCodeResponse
	json.NewDecoder(res1.Body).Decode(&response1)
	if err != nil {
		log.Fatalln("Failed to read response body:", err.Error())
	}

	fmt.Printf("response body: %v\n", response1)

}
