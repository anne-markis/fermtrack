package answer

import (
	"context"
	"strings"
	"time"
)

type CheapClient struct {
}

func (o CheapClient) AskQuestion(ctx context.Context, question string) (string, error) {
	question = strings.Join(strings.Fields(question), "")
	if question == "" {
		return "Ask me, the wine wizard, anything you like.", nil
	}
	time.Sleep(1 * time.Second) // simulate slow AI answer
	return `My apologies, but at the moment, my thoughts seem to be a bit hazy, much like a foggy morning.

It's as if my mind is a glass of wine, swirling with ideas, but not quite focused enough to give you a clear answer.

Perhaps we could revisit this later when my mind is a bit more sober and my thoughts are clearer.

Enjoy this poem instead:

Fermenting mind swirls,
Grape juice turned to liquid fire,
Drunkard dreams of wine.
	`, nil
}
