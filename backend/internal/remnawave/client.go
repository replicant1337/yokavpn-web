package remnawave

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Client struct {
	BaseURL string
	Token   string
}

func NewClient() *Client {
	return &Client{
		BaseURL: os.Getenv("REMNAWAVE_API_URL"),
		Token:   os.Getenv("REMNAWAVE_API_TOKEN"),
	}
}

func (c *Client) CreateUser(username, email string) (*RemnaUser, error) {
	reqBody, _ := json.Marshal(CreateUserRequest{Username: username, Email: email})
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/api/users", c.BaseURL), bytes.NewBuffer(reqBody))
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var response UserResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}
	return &response.Data, nil
}

func (c *Client) CreateSubscription(userID string) (*RemnaSubscription, error) {
	reqBody, _ := json.Marshal(CreateSubscriptionRequest{UserID: userID})
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/api/subscriptions", c.BaseURL), bytes.NewBuffer(reqBody))
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var response SubscriptionResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}
	return &response.Data, nil
}

func (c *Client) GetSubscriptionByShortID(shortID string) (*RemnaSubscription, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/api/subscriptions/%s", c.BaseURL, shortID), nil)
	req.Header.Set("Authorization", "Bearer "+c.Token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var response SubscriptionResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}
	return &response.Data, nil
}
