package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
)

const (
	CONCURRENT_USER   = 10000
	TARGET_PRODUCT_ID = 1
)

type TokenType struct {
	Token    string `json:"token"`
	ExpireIn int    `json:"expire_in"`
}

type AuthData struct {
	AccessToken  TokenType `json:"access_token"`
	RefreshToken TokenType `json:"refresh_token"`
}

type AuthResponse struct {
	Data AuthData `json:"data"`
}

func workerPool() {
	var wg sync.WaitGroup

	wg.Add(CONCURRENT_USER)

	for i := 0; i < CONCURRENT_USER; i++ {
		go worker(i, &wg)
	}

	wg.Wait()
}

func worker(workerID int, wg *sync.WaitGroup) {
	defer wg.Done()

	username := fmt.Sprintf("user-%d", rand.Intn(1001))
	sendRequest(workerID, username)
}

func sendRequest(workerID int, username string) error {
	password := "super_safe_password"
	deviceId := uuid.New().String()

	reqBody := []byte(fmt.Sprintf(`{
  "username": "%s",
  "password": "%s",
  "device_id": "%s"
  }`, username, password, deviceId))

	req, err := http.NewRequest("POST", "http://localhost:8080/v1/auth/login", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Printf("[Worker %d]: error creating login request: %v\n", workerID, err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := sendRequestWithRetry(client, req)
	if err != nil || resp.StatusCode != 200 {
		fmt.Printf("[Worker %d]: error sending login request: %v\n", workerID, err)
		return errors.New("error")
	}
	defer resp.Body.Close()

	var authResp AuthResponse

	err = json.NewDecoder(resp.Body).Decode(&authResp)
	if err != nil {
		fmt.Printf("[Worker %d]: error reading login response body: %v\n", workerID, err)
		return err
	}

	// reqBody = []byte(fmt.Sprintf(`{
	//    "items": [{
	//      "product_id": %d,
	//      "quantity": 1
	//    }]
	//  }`, TARGET_PRODUCT_ID))
	//
	// req, err = http.NewRequest("POST", "http://localhost:8080/v1/orders", bytes.NewBuffer(reqBody))
	// if err != nil {
	// 	fmt.Printf("[Worker %d]: error creating order request: %v\n", workerID, err)
	// 	return err
	// }
	//
	// req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", authResp.Data.AccessToken.Token))
	// req.Header.Set("Content-Type", "application/json")
	//
	// resp, err = sendRequestWithRetry(client, req)
	// if err != nil {
	// 	fmt.Printf("[Worker %d]: error sending order request: %v\n", workerID, err)
	// 	return err
	// }

	if status := resp.StatusCode; status >= 200 || status < 300 {
		fmt.Printf("[Worker %d]: Response received. Status: %d, Data: %v\n", workerID, resp.StatusCode, authResp)
		return nil
	} else {
		fmt.Printf("[Worker %d]: Error. Status: %d\n", workerID, resp.StatusCode)
	}
	return errors.New("error")
}

func sendRequestWithRetry(client *http.Client, req *http.Request) (*http.Response, error) {
	var e error

	for i := 0; i < 5; i++ {
		resp, err := client.Do(req)
		if err == nil && (resp.StatusCode >= 200 || resp.StatusCode < 300) {
			return resp, nil
		}

		if err != nil {
			e = err
		} else {
			e = errors.New("time out response")
		}
		time.Sleep(2 * time.Second)
	}

	return nil, e
}

func main() {
	rand.Seed(time.Now().Unix())

	workerPool()
}
