[![License](https://img.shields.io/github/license/Massad/gin-boilerplate)](https://github.com/Massad/gin-boilerplate/blob/master/LICENSE) 
[![DB Version](https://img.shields.io/badge/DB-PostgreSQL--latest-blue)](https://github.com/Massad/gin-boilerplate/blob/master/go.mod) 


Welcome to **Golang Gin Gorm boilerplate** v1

The fastest way to deploy a restful api's with [Gin Framework](https://gin-gonic.github.io/gin/) with a structured project that defaults to **PostgreSQL** database

## Configured with

- [go-gorm](https://github.com/jinzhu/gorm): Go ORM for relational databases
- Go Modules
- Built-in **CORS Middleware**
- Feature **PostgreSQL** with JSON/JSONB queries & trigger functions
- Enviroment support
- And few other important utilties
### Installation

```
$ go mod init
```

```
$ go install
```


## Running Your Application

Rename .env.example to .env and place your credentials

```
$ mv .env.example .env
```



## Building Your Application

```
$ go build .
```

```
$ ./src.exe
```

## License

(The MIT License)

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
'Software'), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.