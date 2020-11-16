# jogadores-service

REST API in go-lang, using Fiber framework and MongoDB

## Run

Run app with Atlas MongoDB:

```sh
PORT=":3000" MONGODB_CONNECTION_STRING="mongodb+srv://<ATLAS_URI>" go run server.go
```

Run app with local MongoDB:

```sh
PORT=":3000" go run server.go
```
