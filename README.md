# go-rest-api-gorilla-mux

## Setup

#### 1. brew install:
   - ```go-task``` (https://formulae.brew.sh/formula/go-task)
   - ```golangci-lint``` (optional)
#### 2. Install VS Code:
   - https://code.visualstudio.com/download


## Run (Taskfile Options)
(after installing ```go-task```)(https://taskfile.dev)
```
task build
  (go build -o app cmd/server/main.go)

task test:
  (go test -v ./...)

task lint:
  (golangci-lint run)

task run:
  (docker-compose up --build)
      
integration-tests:
  (docker-compose up -d db
  go test -tags=integration -v ./...
    env:
      DB_USERNAME: postgres
      DB_PASSWORD: postgres
      DB_TABLE: postgres
      DB_HOST: localhost
      DB_PORT: 5432
      DB_DB: postgres
      SSL_MODE: disable)

task acceptance-tests:
  (docker-compose up -d --build
  go test -tags=e2e -v ./...)
```

## Endpoints

#### 1. Create comment
- /api/v1/comment

- curl -d "@testdata/createComment.json" http://localhost:8080/api/v1/comment -v --header 'Content-Type: application/json' --header 'Authorization: bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.2yPNTTY3Y5jUwYBPAJAfUc84Ybv2qPbZY_OHI7tzuug'

#### 2. Get comments
- /api/v1/comments
- curl http://localhost:8080/api/v1/comments -v --header 'Content-Type: application/json'

#### 3. Get comment by id
- /api/v1/comment/{uuid}
- curl http://localhost:8080/api/v1/comment/{id} -v --header 'Content-Type: application/json'

#### 4. Update comment by id
- /api/v1/comment/{uuid}
- curl -d "@testdata/updateComment.json" -X PUT http://localhost:8080/api/v1/comment/{id} -v --header 'Content-Type: application/json'

#### 4. Delete comment by id
- /api/v1/comment/{uuid}
- curl -X DELETE http://localhost:8080/api/v1/comment/{id} -v --header 'Content-Type: application/json'


## External Dependencies
   - github.com/jmoiron/sqlx
   - github.com/lib/pq

   - github.com/golang-migrate/migrate/v4
   - github.com/golang-migrate/migrate/v4/database/postgres
   - github.com/golang-migrate/migrate/v4/source/file

   - github.com/satori/go.uuid

   - github.com/gorilla/mux

   - github.com/dgrijalva/jwt-go
   - github.com/go-playground/validator/v10

   - github.com/stretchr/testify
   - github.com/go-resty/resty/v2
