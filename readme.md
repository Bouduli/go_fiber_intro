# Test av Go-Fiber
Detta är ett litet test av [Go-Fiber](https://docs.gofiber.io/) ramverket. 

## Requirements:
- [Golang](https://go.dev/) (språket Go)
- [Go-Fiber](https://docs.gofiber.io/) (ramverket i fråga)
#### Optional:
- [Air](https://github.com/cosmtrek/air) - Ett "Nodemon liknande" program fast gjort för Go.

## Starta ett Go-Fiber projekt: 
### 1. Initiera go-projektet
Ett go projekt startas genom att köra `go mod init example.com/m/v2`
```bash
$ go mod init example.com/m/v2
go: creating new go.mod: module example.com/m/v2
```
En fil `go.mod` skapas, och 'liknar' `en package.json` fil i NodeJS.
### 2. Installera tredjeparts-bibliotek
Paket installeras genom `go get` enligt
```bash
$ go get github.com/gofiber/fiber/v2
go: downloading github.com/gofiber/fiber/v2 v2.52.0
go: downloading github.com/gofiber/fiber v1.14.6
go: downloading github.com/google/uuid v1.5.0
go: downloading github.com/mattn/go-runewidth v0.0.15
/.../
```

### 3. Skapa en `.go` fil
#### 3.1 Hello world
Följande exempel är en "Hello World" applikation.
```go
//hello.go
package main

import "fmt"

func main() {
    fmt.Println("hello world")
}
```
Applikationen körs genom att skriva `go run`:
```bash
go run hello.go
> Hello World
```
#### 3.1 Hello Server (go-fiber)
Följande är ett exmpel av en Go-fiber server-applikation: 

```go
//server.go
package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/products", func(c *fiber.Ctx) error {
		return c.SendString("producs")
	})

	app.Get("/cars", func(c *fiber.Ctx) error {
		return c.SendString("cars")
	})

	app.Listen(":3000")
}
```

Raden `import "github.com/gofiber/fiber/v2"` ger tillgång till fiber *modulen (?)* likt en `require()` i NodeJS.

**Annars är det rätt bog-standard express.**


Servern startas likt `hello.go` genom `go run`:
```bash
$ go run server.go
 ┌───────────────────────────────────────────────────┐ 
 │                   Fiber v2.52.0                   │
 │               http://127.0.0.1:3000               │
 │       (bound on host 0.0.0.0 and port 3000)       │
 │                                                   │
 │ Handlers ............. 2  Processes ........... 1 │
 │ Prefork ....... Disabled  PID ............. 11196 │
 └───────────────────────────────────────────────────┘ 
```
**En server har nu startas, och här slutar introt.**

### 4. Quality of Life saker
#### 4.1 Air
Air är ett Nodemon liknande program för att starta-om program när filer ändras. Dock är det mer avancerat. 

**Air är ett extern paket, som laddas ned precis som go-fiber:**
```bash
$ go install github.com/cosmtrek/air@latest
go: downloading github.com/cosmtrek/air v1.49.0
go: downloading dario.cat/mergo v1.0.0
go: downloading github.com/fatih/color v1.14.1
go: downloading github.com/fsnotify/fsnotify v1.6.0
go: downloading github.com/gohugoio/hugo v0.111.3

/.../
```
Air körs genom att skriva (get this) `air` i terminalen, men måste först konfigureras (görs automatiskt) med `air init`:
```bash
$ air init

  __    _   ___
 / /\  | | | |_)
/_/--\ |_| |_| \_ v1.49.0, built with Go go1.21.6

.air.toml file created to the current directory with the default settings

```
Därefter kan din server, och andra applikationer startas direkt med `air`. 
```
$ air

  __    _   ___
 / /\  | | | |_)
/_/--\ |_| |_| \_ v1.49.0, built with Go go1.21.6

mkdir C:\Users\Oliver\Webb\go_fiber_test_02_03\tmp
watching .
!exclude tmp
building...
!exclude .git
!exclude .git
running...

 ┌───────────────────────────────────────────────────┐ 
 │                   Fiber v2.52.0                   │
 │               http://127.0.0.1:3000               │
 │       (bound on host 0.0.0.0 and port 3000)       │
 │                                                   │
 │ Handlers ............. 4  Processes ........... 1 │
 │ Prefork ....... Disabled  PID ............. 20188 │
 └───────────────────────────────────────────────────┘
```