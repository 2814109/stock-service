gen-gql:
	docker compose exec app sh -c 'go run github.com/99designs/gqlgen generate'