package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"time"

	"golang.org/x/exp/slices"
)

type Webhook struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	URL         string                 `json:"url"`
	Events      []string              `json:"events"`
	Headers     map[string]string     `json:"headers"`
	Secret      string                `json:"secret"`
	Enabled     bool                  `json:"enabled"`
	Conditions  map[string]any        `json:"conditions"`
	RetryConfig RetryConfig           `json:"retryConfig"`
	CreatedAt   time.Time             `json:"created_at"`
	UpdatedAt   time.Time             `json:"updated_at"`
}

type RetryConfig struct {
	MaxRetries  int           `json:"maxRetries"`
	RetryDelay  time.Duration `json:"retryDelay"`
	MaxDelay    time.Duration `json:"maxDelay"`
	Multiplier  float64       `json:"multiplier"`
}

func (w *Webhook) ValidatePayload(payload []byte, signature string) bool {
	mac := hmac.New(sha256.New, []byte(w.Secret))
	mac.Write(payload)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal([]byte(signature), expectedMAC)
}

func (w *Webhook) ShouldTrigger(event string, data map[string]any) bool {
	if !w.Enabled {
		return false
	}
	if !slices.Contains(w.Events, event) {
		return false
	}
	return w.matchConditions(data)
}

func (w *Webhook) matchConditions(data map[string]any) bool {
	// 实现条件匹配逻辑
	// TODO: 根据w.Conditions实现更复杂的条件判断
	return true
}

func main() {
	//Example Usage
	webhook := Webhook{
		ID:          "123",
		Name:        "My Webhook",
		URL:         "http://example.com/webhook",
		Events:      []string{"event1", "event2"},
		Headers:     map[string]string{"Content-Type": "application/json"},
		Secret:      "mysecret",
		Enabled:     true,
		Conditions:  map[string]any{"key1": "value1"},
		RetryConfig: RetryConfig{MaxRetries: 3, RetryDelay: 1 * time.Second},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	payload := []byte(`{"message": "Hello, world!"}`)
	signature := "some_signature" //Replace with actual signature generated using webhook.Secret

	isValid := webhook.ValidatePayload(payload, signature)
	shouldTrigger := webhook.ShouldTrigger("event1", map[string]any{"key1": "value1"})

	println("Payload is valid:", isValid)
	println("Webhook should trigger:", shouldTrigger)

}