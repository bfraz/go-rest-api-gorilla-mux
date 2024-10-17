# go-rest-api-course-v2

- brew install
    - go-task
    - golangci-lint


- go get 
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


- curl -d "@testdata/createComment.json" http://localhost:8080/api/v1/comment -v --header 'Content-Type: application/json' --header 'Authorization: bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.2yPNTTY3Y5jUwYBPAJAfUc84Ybv2qPbZY_OHI7tzuug'

- curl http://localhost:8080/api/v1/comments -v --header 'Content-Type: application/json'

- curl http://localhost:8080/api/v1/comment/{id} -v --header 'Content-Type: application/json'

- curl -d "@testdata/updateComment.json" -X PUT http://localhost:8080/api/v1/comment/{id} -v --header 'Content-Type: application/json'

- curl -X DELETE http://localhost:8080/api/v1/comment/{id} -v --header 'Content-Type: application/json'


- go test -v ./...

- task integration-test
