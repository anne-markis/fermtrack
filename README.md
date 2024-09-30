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

### Mermaid class diagram
```
graph TD
    %% Main HTTP Service
    subgraph FermTrack Service
        B[HTTP Server]
        B --> |API Calls| D[MySQL DB]
        B --> |User questions| E[AI Client]
    end

    %% Bubbletea CLI Service
    subgraph CLI Interface
        C[User Commands, Questions]
        C --> |CLI Commands| F[Fermtrack Client]
        F --> |HTTP Requests| B
    end

    %% Components and Flows
    D --> |Data Store| B
    E --> |Fermentation Advice| B

    %% Labels for clarity
    classDef httpService fill:#f9f,stroke:#333,stroke-width:2px;
    classDef db fill:#bbf,stroke:#333,stroke-width:2px;
    classDef ai fill:#bfb,stroke:#333,stroke-width:2px;
    classDef cliService fill:#ffc,stroke:#333,stroke-width:2px;

    %% Style assignments
    class A,B httpService;
    class D db;
    class E ai;
    class F,C cliService;
```

### Links I found useful

https://platform.openai.com/docs/guides/prompt-engineering/tactic-provide-examples

https://github.com/sashabaranov/go-openai

https://platform.openai.com/usage

https://charm.sh/blog/commands-in-bubbletea/


TODO
add lint