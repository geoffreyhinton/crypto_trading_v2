export $(cat .env.dev | grep -v '^#' | xargs) && go run cmd/server/main.go

