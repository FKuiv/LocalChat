# LocalChat

Setup a messaging app locally on your network

## Development environment

Type `go run main.go` in the root folder of the project to run the Go backend.

### Database

Docker compose file will setup the DB for you. Just run `docker compose up -d` from the root of the directory. This will create a new folder `db` in this project where the database is stored.

#### MinIO for file storage

Using MinIO to store all the files in this application. Golang docs: https://min.io/docs/minio/linux/developers/go/API.html

## Environment variables

Check the `.env.example` file for creating a .env
