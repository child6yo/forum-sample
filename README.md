## Forum sample

👋
That's my simple pet-project - REST forum service sample.
I don't mind that it's pretty safety code for using in real projects.

The sercice is write using:
- Golang
- [gin](https://github.com/gin-gonic/gin)
- [sqlx](https://github.com/jmoiron/sqlx)
- [golang-jwt/jwt](https://github.com/golang-jwt/jwt/) for jwt bearer generation
- [migrate](https://github.com/golang-migrate/migrate) for migrations
- [swaggo](https://github.com/swaggo/swag) for documentation
- [uber-go/mock](https://github.com/uber-go/mock) for service mocks
- [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) for repository mocks

Service documentation is available after server startup at:
```
/docs/index.html
```
WARNING: for authentication in swagger (lock icon, after sign in) you have to use:
```
Bearer YOUR_JWT_TOKEN
```

Forum Features:
- User
  - Sign Up
  - Sign In (JWT Bearer)
- Posts
  - Create post (using title and post text)
  - Read someone else
  - Update your post
  - Delete your post
- Thread (discussions/post comments)
  - Create thread under the post (you can answer at someone's thread as well)
  - Read thread
  - Read threads tree (all threads under post)
  - Update your thread

## How to install
1. Clone repository
```
git clone https://github.com/child6yo/forum-sample
cd forum-sample
```

2. Configurate app:
   - in docker-compose
   - in config/config.yml (using in code)
     
3. Run Makefile
```
make migrate
make build && make run
```
Or use:
```
migrate -path ./schemas -database 'db path' up
docker-compose up
```

## Project Structure

```
forum-sample
├── cmd
|   └── main.go
├── config/ # contains config files
├── docs/ # contains swagger documentation
├── pkg/validation # contains validation library
|   └── validation.go
├── internal
|   ├── handler
|   |   ├── auth.go
|   |   ├── handler.go # contains routes initialization
|   |   ├── middleware.go
|   |   ├── posts.go
|   |   ├── response.go # contains response structure & funcs for different response generation
|   |   └── threads.go
|   ├── repository # database logic
|   |   ├── auth.go
|   |   ├── postgres_db.go # contains postgres initialization
|   |   ├── posts.go
|   |   ├── repository.go
|   |   └── threads.go
|   ├── service # buisness logic
|   |   ├── mocks/
|   |   ├── auth.go
|   |   ├── posts.go
|   |   ├── service.go
|   |   └── threads.go
├── posts.go # post model
├── server.go # server interface
├── threads.go # thread model
└── user.go # user model
```
