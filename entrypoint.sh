# mysql起動を待つ
wait-for "${DATABASE_HOST}:${DATABSE_PORT}" --"$@"

# watch for .go file changes
CompileDaemon --build="go build -o main main.go" --command=./main