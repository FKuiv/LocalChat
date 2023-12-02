# LocalChat

Setup a messaging app locally on your network

## Development environment

Type `go run main.go` in the root folder of the project to run the Go backend.

### Database

Docker compose file will setup the DB for you. Just run `docker compose up -d` from the root of the directory. This will create a new folder `db` in this project where the database is stored.

## Environment variables

Check the `.env.example` file for creating a .env

## Errors I encountered

When building Go docker image I got an error that said 'terminal prompts disabled'.
To tackle it I used this command: `export GOPRIVATE=github.com/FKuiv/*`
