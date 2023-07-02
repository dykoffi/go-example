all: dev
dev: db-start
	ENV=development \
\
	KEYCLOAK_HOST=http://localhost:8180 \
	KEYCLOAK_REALM=test-go \
	KEYCLOAK_CLIENT_ID=go \
	KEYCLOAK_CLIENT_SECRET=8LfYm6ShIdKsfTg7UYBtl0uRPQbNlJ08 \
\
	POSTGRES_DB=db \
	POSTGRES_HOST=localhost \
	POSTGRES_PORT=21822 \
	POSTGRES_USER=user \
	POSTGRES_PASSWORD=4134 \
\
	air

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags netgo -a -v -ldflags='-s' -o app .
doc:
	swag init -o public -ot json
test:
	go test -v ./...

docker-build: build
	docker build -t exp1 .
	
# database docker config

db-start: 
	docker compose up -d
db-stop: 
	docker compose down

clean:
	rm -r app