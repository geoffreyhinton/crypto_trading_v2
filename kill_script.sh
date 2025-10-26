lsof -ti:8080 | xargs kill -9
pkill -f "go run cmd/server/main.go"

lsof -ti:8080 | xargs kill -9 2>/dev/null || true