# mysql起動を待つ
wait-for "${DATABSE_HOST}:${DATABSE_PORT}" --"$@"

# watch for .go file changes
go-basic-nakagome2 --build="go build -o main main.go" --command=./main