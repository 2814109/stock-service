gen-gql:
	docker compose exec app sh -c 'go run github.com/99designs/gqlgen generate'

gen-db:
	docker compose exec app sh -c 'cd ./cmd go run main.go'