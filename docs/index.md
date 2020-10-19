## REST API using GO language and fiber framework

This project is using  [GoLang](https://golang.org/) and [Fiber](https://docs.gofiber.io/) for creation of [RestAPI](https://en.wikipedia.org/wiki/Representational_state_transfer)

### Rest Api's in go
Rest api's in go can be build using the following frameworks:
1. Gorilla
2. Gin 
3. Revel
4. Martini
5. web.go
6. Goji
7. Beego
8. echo (quite similar to fibre)
9. fibre

Every framework have its pro's and cons.
we will be using fibre here as this is inspired from [Express](https://expressjs.com/) which is battle tested and use [fasthttp](https://pkg.go.dev/gopkg.in/labstack/echo.v2/engine/fasthttp) under the hood.
### Landscape

we will be breaking down the development of the rest api in different phases:
1. Simple rest api with mocked data
2. Adding cassandra as backed using gocql package
3. Adding ES for search extensibility
4. custom Loggers for logging.
 