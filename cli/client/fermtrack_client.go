package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
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

type Fermentation struct {
	ID           int        `json:"id"`
	UUID         string     `json:"uuid"`
	Nickname     string     `json:"nickname"`
	StartAt      time.Time  `json:"start_at"`
	BottledAt    *time.Time `json:"bottled_at"`
	RecipeNotes  string     `json:"recipe_notes"`
	TastingNotes *string    `json:"tasting_notes"`
}

func (f Fermentation) ToString() string {
	recipeNotes := "[No Receipe Notes]"
	tastingNotes := "[No Tasting Notes]"
	if f.TastingNotes != nil {
		tastingNotes = *f.TastingNotes
	}
	if f.RecipeNotes != "" {
		recipeNotes = f.RecipeNotes
	}
	return fmt.Sprintf(`%s

Recipe Notes: 
	%s

Tasting Notes: 
	%s
	`, f.Nickname, recipeNotes, tastingNotes)
}

type Fermtracker interface {
	AskQuestion(ctx context.Context, question *FermentationQuestion) (*FermentationAdvice, error)
	ListFermentations(ctx context.Context) ([]Fermentation, error)
	GetFermentation(ctx context.Context, uuid string) (*Fermentation, error)
	Login(ctx context.Context, username, password string) error
}

type FermentationQuestion struct {
	Question string `json:"question"`
}
type FermentationAdvice struct {
	Answer string `json:"answer"`
}

func (fc *FermentationClient) Login(ctx context.Context, username, password string) error {
	return fmt.Errorf("unimplemented") // TODO
}

func (fc *FermentationClient) AskQuestion(ctx context.Context, question *FermentationQuestion) (*FermentationAdvice, error) {
	url := fmt.Sprintf("%s/v1/fermentations/advice", fc.baseURL)
	body, err := json.Marshal(question)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal fermentation: %v", err)
	}

	resp, err := fc.client.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create fermentation: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var answer FermentationAdvice

	if err := json.NewDecoder(resp.Body).Decode(&answer); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &answer, nil
}

func (fc *FermentationClient) ListFermentations(ctx context.Context) ([]Fermentation, error) {
	url := fmt.Sprintf("%s/v1/fermentations", fc.baseURL)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("error response: %s", body)
	}

	var fermentations []Fermentation
	if err := json.NewDecoder(resp.Body).Decode(&fermentations); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return fermentations, nil
}

func (fc *FermentationClient) GetFermentation(ctx context.Context, uuid string) (*Fermentation, error) {
	url := fmt.Sprintf("%s/v1/fermentations/%s", fc.baseURL, uuid)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("error response: %s", body)
	}

	var fermentation Fermentation
	if err := json.NewDecoder(resp.Body).Decode(&fermentation); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return &fermentation, nil
}
