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


