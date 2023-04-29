all: dev
dev: doc
	air

build:
	go build .
doc:
	swag init -o public -ot json
test:
	go test -v ./...

# database docker config

db-start: 
	docker compose up -d
db-stop: 
	docker compose down

clean:
	rm -r exp1