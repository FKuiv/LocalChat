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

## Some explanations

### Backend

Backend is built with Golang and structured with a repository pattern. The whole project is split into three main pieces: user, group, and message. Using the repo pattern I can easily distribute the interactions with a database from a single place. Check the `utils/constants.go` for some extra constants used in the backend.

### Frontend

This branch contains frontend created with SvelteKit. Check the frontend folder's README.md for more info.

The `/api` folder contains all the endpoints defined in the backend but just translated into Typescript. Yes it is a little bit tedious/stupid to just copy the endpoints over but I can't think of a better solution at the moment. The `api/endpoints.ts` file includes the backend REST API endpoint URL definitions. Other files in the `/api` folder include the actual Axios requests that can be easily used throughout the frontend.
