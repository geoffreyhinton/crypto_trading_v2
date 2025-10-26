lsof -ti:8080 | xargs kill -9
pkill -f "go run cmd/server/main.go"