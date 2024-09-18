# fermtrack

## What it is

IDK some sort of fancy CLI driven wine tool for home winemakers. I'm just riffing here.

## How to Run

You will need a chat gpt key. Create a `.env` file at the root and specify the key there.

.env example
```
CHATGPT3_KEY=your-key
```

To run, simply run main directly and build and run.
```
% go mod vendor
% go run main.go
```
OR
```
% go build
% ./fermtrack
```

If you want to test out some functionality but don't have an openAI token or don't want to use credits, start the process with the `cheap` arg. This will return a hard coded dummy answer for all questions.

```
% ./fermtrack cheap
```

### More docs

https://platform.openai.com/docs/guides/prompt-engineering/tactic-provide-examples

https://github.com/sashabaranov/go-openai

https://platform.openai.com/usage