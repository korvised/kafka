start-kafka:
	cd server && docker compose up -d
start-db:
	cd consumer && docker compose up -d
consumer:
	cd consumer && go run .
producer:
	cd producer && go run .

.PHONY: start-kafka start-db consumer producer