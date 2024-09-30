package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type FermentationClient struct {
	baseURL string
	client  *http.Client
}

func NewFermentationClient(baseURL string) *FermentationClient {
	return &FermentationClient{
		baseURL: baseURL,
		client:  &http.Client{},
	}
}

// type Fermentation struct {
// 	ID           int        `json:"id"`
// 	UUID         string     `json:"uuid"`
// 	Nickname     string     `json:"nickname"`
// 	StartAt      time.Time  `json:"start_at"`
// 	BottledAt    *time.Time `json:"bottled_at"`
// 	RecipeNotes  string     `json:"recipe_notes"`
// 	TastingNotes *string    `json:"tasting_notes"`
// 	DeletedAt    *time.Time `json:"deleted_at"`
// }

type Fermtracker interface {
	AskQuestion(ctx context.Context, question string) (string, error)
}

type FermentationQuestion struct {
	Question string `json:"question"`
}
type FermentationAdvice struct {
	Answer string `json:"answer"`
}

// TODO the input and output is weird here
func (fc *FermentationClient) AskQuestion(ctx context.Context, question string) (string, error) {
	fermQuestion := FermentationQuestion{
		Question: question,
	}
	url := fmt.Sprintf("%s/fermentations/advice", fc.baseURL)
	body, err := json.Marshal(fermQuestion)
	if err != nil {
		return "", fmt.Errorf("failed to marshal fermentation: %v", err)
	}

	resp, err := fc.client.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("failed to create fermentation: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var answer FermentationAdvice

	if err := json.NewDecoder(resp.Body).Decode(&answer); err != nil {
		return "", fmt.Errorf("failed to decode response: %v", err)
	}

	return answer.Answer, nil
}
