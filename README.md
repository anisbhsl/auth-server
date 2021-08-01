auth-server
------------------

A simple JWT based auth server that lets user to register, login and fetch user profile.
This project uses `JWT` for authentication and `sqlite` for storing user info.

Example Run:
```
export APP_SECRET=<app_secret_key> ; go run main.go --host=127.0.0.1 --port=5000
```
API server will run at: `127.0.0.1:5000`

Run using docker:
```
docker build .
docker run -p 5000:5000 -e APP_SECRET=<app_secret_key> <imageTag>
```


Run tests:
```
go test ./...
```

