# GoBoard
A message board built in Golang

This project is for Go practicing and guided by [Building Go Web Applications and Microservices Using Gin](https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin) and [Test-driven Development of Go Web Applications with Gin](https://semaphoreci.com/community/tutorials/test-driven-development-of-go-web-applications-with-gin)

Install Dependency  
```sh
$ go get -u github.com/gin-gonic/gin
```  

Run
```sh
$ go run main.go
```

Test
```sh
$ go test -v tests/
```

TODO:  
- [x] Display the list of all articles on the home page
- [x] Display a single article on its own page
- [x] Users register with a username and a password
- [ ] Login
- [ ] Log out
- [ ] Create new articles (logged in users only)
- [ ] Use PostgreSQL as database
- [ ] Deploy on cloud