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
		client: &http.Client{
			Timeout: time.Second * 15,
		},
	}
}

// Setting jwt value in context
type jwtKey struct{}

var ContextKeyJWT = jwtKey{}

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
	Login(ctx context.Context, username, password string) (*LoginResponse, error)
}

type FermentationQuestion struct {
	Question string `json:"question"`
}
type FermentationAdvice struct {
	Answer string `json:"answer"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type LoginResponse struct {
	Token    string `json:"token"`
	Username string `json:"username"`
	UUID     string `json:"uuid"`
}

func (fc *FermentationClient) Login(ctx context.Context, username, password string) (*LoginResponse, error) {
	url := fmt.Sprintf("%s/v1/login", fc.baseURL)

	req := LoginRequest{
		Username: username,
		Password: password,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal user: %v", err)
	}

	resp, err := fc.client.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed login: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var loginResp LoginResponse

	if err := json.NewDecoder(resp.Body).Decode(&loginResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &loginResp, nil
}

func (fc *FermentationClient) AskQuestion(ctx context.Context, question *FermentationQuestion) (*FermentationAdvice, error) {
	url := fmt.Sprintf("%s/v1/fermentations/advice", fc.baseURL)

	body, err := json.Marshal(question)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal fermentation: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", ctx.Value(ContextKeyJWT)))

	resp, err := fc.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to ask question: %v", err)
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

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", ctx.Value(ContextKeyJWT)))

	resp, err := fc.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to ask question: %v", err)
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
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", ctx.Value(ContextKeyJWT)))

	resp, err := fc.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to ask question: %v", err)
	}
	defer resp.Body.Close()

	var fermentation Fermentation
	if err := json.NewDecoder(resp.Body).Decode(&fermentation); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return &fermentation, nil
}
