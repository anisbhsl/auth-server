auth-server
------------------

A simple JWT based auth server that lets user to register, login and fetch user profile.
This project uses `JWT` for authentication and `sqlite` for storing user info.

Example Run:
```
export APP_SECRET=<app_secret_key> ; go run main.go --host=127.0.0.1 --port=8080
```
API server will run at: `127.0.0.1:8080`


Run tests:
```
go test ./...
```

