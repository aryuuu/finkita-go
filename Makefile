PG_DSN = "postgres://$(PG_USERNAME):$(PG_PASSWORD)@$(PG_HOST):$(PG_PORT)/$(PG_DATABASE)?sslmode=disable"

run:
	export $$(xargs < .env) && \
	go run cmd/rest/main.go
 
dev-air:
	export $$(xargs < .env) && \
	air

test:
	export $$(xargs < .env) && \	
	go test -v ./...

build: 
	go build -o bin/rest cmd/rest/main.go 

run-scraper:
	export $$(xargs < .env) && \
	go run cmd/scraper/main.go

build-scraper:
	go build -o bin/scraper cmd/scraper/main.go 

migration-generate: 
	# get file name from argument
	# touch migrations/$$(date +%s)-$(name).sql
	# create file in  
	migrate create -ext sql -dir migrations/ $(name)

migration-up:
	migrate -database=$(PG_DSN) -path=migrations up

migration-down: 
	migrate -database=$(PG_DSN) -path=migrations down
