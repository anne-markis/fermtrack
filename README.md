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
% make start
```
 
To stop
```
% make stop
```

To view logs
```
% make logs
```


### Creating a user
Create a user locally by starting the server and running the following (substituting the username/pass as you like):

```
curl -X POST http://localhost:8080/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "username": "cooluser",
    "password": "cool123"
  }'
```

This is required to use the bubbletea cli in `cli` dir.


### Migrations

Migrations are by goose and will run upon server startup automatically.

```
% brew install goose # or follow instructions here: https://github.com/pressly/goose
% GOOSE_DRIVER=mysql GOOSE_MIGRATION_DIR=migrations goose create SOMENAME sql

```

### Future changes
* Add users, and users to own fermentations
* Add filtering for GET /v1/fermentations
* Edit/Create fermentation
* Auth
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