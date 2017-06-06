# gosvctest
Test for go based services based largely on: https://github.com/komand/gosea

generate TLS credentials:
  openssl genrsa -out server.key 2048

  openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650

setup mysql:
  sql/create.sql

sample curl:

  curl -k https://localhost:8080/ping

  curl -k https://localhost:8080/token

  curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6ZmFsc2UsInVzZXIiOnsiSWQiOjEsIkZpcnN0TmFtZSI6IkFkbWluIiwiTGFzdE5hbWUiOiJVc2VyIiwiUm9sZXMiOlsiYWRtaW5pc3RyYXRvciJdfSwiZXhwIjoxNDk2ODA3MTM3LCJqdGkiOiIxIiwiaXNzIjoidG9rZW4tc2VydmljZSJ9.BuQV7JiZNRPlNrkFLjwNcebT4AN4dcLmyQqe25x0Whw" -k https://localhost:8080/test
