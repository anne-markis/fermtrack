# fermtrack

## What it is

IDK some sort of fancy CLI driven wine tool for home winemakers. I'm just riffing here.

## How to Run locally

This requires go >= 1.22

You will need a chat gpt key. Create a `.env` file at the root and specify the key there.

.env example
```
CHATGPT_KEY=your-key
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


### Migrations

Migrations are by goose.

```
% brew install goose # or follow instructions here: https://github.com/pressly/goose
% GOOSE_DRIVER=mysql GOOSE_MIGRATION_DIR=migrations goose create SOMENAME sql

```

### Future changes
* Add users, and users to own fermentations
* Add filtering for GET /v1/fermentations
* Edit/Create fermentation
* Auth
* Better test coverage
* Put on server
* Comments

### Mermaid class diagram
```
graph TD
    %% Bubble Tea CLI Client
    subgraph CLI Client

        B[Bubble Tea CLI]
        B --> |User Commands| C[FermTrack HTTP Client]
    end

    %% FermTrack Service
    subgraph FermTrack Service
        C --> |HTTP Requests| D[HTTP Handlers]
        D --> T[FermTrackService]
        T --> |Fetches Data| E[Fermentation Database]
        T --> |Sends Query| F[LLM Service]
        F --> |Returns Processed Data| T
        E --> |Returns Fermentation Data| T
        D --> |Sends Response| B
    end

    %% Relationships and Flows
    classDef cli fill:#f9f,stroke:#333,stroke-width:2px;
    classDef fermService fill:#bff,stroke:#333,stroke-width:2px;
    classDef db fill:#bbf,stroke:#333,stroke-width:2px;
    classDef llm fill:#bfb,stroke:#333,stroke-width:2px;

    class A,B cli;
    class D fermService;
    class E db;
    class F llm;
```

### Links I found useful

https://platform.openai.com/docs/guides/prompt-engineering/tactic-provide-examples

https://github.com/sashabaranov/go-openai

https://platform.openai.com/usage

https://charm.sh/blog/commands-in-bubbletea/


TODO
add lint