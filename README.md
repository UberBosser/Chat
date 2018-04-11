# GoTemplate
Gin + Infernojs = 💙

[Gin](https://github.com/gin-gonic/gin) and [Infernojs](https://github.com/infernojs/inferno) create a blazing fast server/client.

## Installation
```
$ git clone https://github.com/UberBosser/GoTemplate.git [<dir>]
$ go get github.com/gin-gonic/gin github.com/gin-contrib/gzip
$ cd [<dir>]
$ make init
```
## Commands
`$ make serve` builds the go files and starts the server, outputs to `log.txt`.  
`$ make stop` stops the server.  
`$ make build` alias to `$ yarn build`.  
`$ make clean` cleans static `*.bundle.css` & `*.bundle.js`.  

## Template structure
* `main.go` main Gin server.
* `src` `.jsx`, `.sass`, `.png`...
* `src/components` contain components of the UI.
* `src/containers` different web pages.
* `static` webpacked `.css` and `.js`.
* `templates/includes` Golang template "defines".
* `templates/layouts` Golang container templates.
