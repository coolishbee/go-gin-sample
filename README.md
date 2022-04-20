# go-gin-sample

[![rcard](https://goreportcard.com/badge/github.com/coolishbee/go-gin-sample)](https://goreportcard.com/report/github.com/coolishbee/go-gin-sample) [![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/coolishbee/go-gin-sample/master/LICENSE)

an simple example for gin + mariadb + swagger + google login idtoken validate.

## Installation
```
$ go get github.com/coolishbee/go-gin-sample
```

## How to run

### Required

- MariaDB

### Ready

Create a **game database** and import [SQL](https://github.com/coolishbee/go-gin-sample/blob/main/docs/sql/game_2022-03-25.sql)

### Conf

You should modify `conf/app.ini`

```
[database]
Type = mysql
User = root
Password =
Host = 127.0.0.1:3306
Name = game

...
```

### Run
```
$ cd $GOPATH/src/go-gin-sample
$ go run main.go 
```

[Swagger doc](http://localhost:8000/swagger/index.html)

## Features

- RESTful API
- Gorm
- Swagger
- Gin
- App configurable
- Logging

## TODO

- [Facebook Login verify access token](https://stackoverflow.com/questions/8605703/how-to-verify-facebook-access-token).
- Redis
- LevelDB

## Reference

* https://github.com/eddycjy/go-gin-example