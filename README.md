# GO Fiber RESTful API

My GO Fiber RESTful API Application

I will be using the following versions:

- Go: 1.21.6
- Fiber: v3

# 1 Hello World Guide

### 1.1 Install GO

https://go.dev/dl/

### 1.2 Create the go.mod file

Inside your project folder run this command:
```
go mod init <module-name>
```
### 1.3 Create your project structure

```
/your-module-name
├── go.mod
├── src
│   └── main.go
```


### 1.4 Installing Fiber

Inside the project folder run this to install fiber
```
go get -u github.com/gofiber/fiber/v3
```

### 1.5 main.go

```GO
package main

import (
	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	app.Get("/hello", func(c fiber.Ctx) error {
		return c.SendString("Hello, World from Fiber!")
	})

	app.Listen(":3000")
}
```