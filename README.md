auth-server
------------------

A simple JWT based auth server that lets user to register, login and fetch user profile.
This project uses `JWT` for authentication (RSA signing method) and `sqlite` for storing user info.

**How to Run?**
```
APP_SECRET=<your_app_secret> make run-app
```
*API server will run at: `127.0.0.1:5000`*


**Run tests:**
```
make test
```

**API Info**

| Endpoint | Body Fields | Method| Remarks |
|----------|------|--------|--------|
| `api/v1/register-user` | `name`, `email`, `location`, `about`, `password`| POST| Password is saved in encrypted format|
| `/api/v1/auth/login`| `email`, `password` | POST | |
| `/api/v1/auth/refresh-token`| `refresh-token` | POST||
| `api/v1/me`| | GET| Use token in Authorization Header for User Profile|