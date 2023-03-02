# Group

O Group é uma funcionalidade do Framework Quick que permite agrupar rotas e aplicar middleware a elas.

Por exemplo, se você tiver um conjunto de rotas que precisam de autenticação antes de serem acessadas, em vez de adicionar o middleware de autenticação individualmente para cada rota, pode agrupá-las usando a funcionalidade Group e aplicar o middleware a todas as rotas do grupo de uma só vez. Isso pode tornar o código mais legível e organizado, além de evitar a repetição de código.

#### Group01

```go
package main

import "github.com/gojeffotoni/quick"

func main() {
	app := quick.New(quick.Config{
		MaxBodySize: 5 * 1024 * 1024,
	})

	group := app.Group("/v1")
	group.Get("/user", func(c *quick.Ctx) {
		c.Status(200).SendString("[GET] [GROUP] /v1/user ok!!!")
		return
	})
	group.Post("/user", func(c *quick.Ctx) {
		c.Status(200).SendString("[POST] [GROUP] /v1/user ok!!!")
		return
	})

 app.Listen("0.0.0.0:8080")
}
```

#### Group02

```go
package main

import "github.com/gojeffotoni/quick"

func main() {
	app := quick.New(quick.Config{
		MaxBodySize: 5 * 1024 * 1024,
	})

	group2 := app.Group("/v2")

	group2.Get("/user", func(c *quick.Ctx) {
		c.Set("Content-Type", "application/json")
		c.Status(200).SendString("Quick em ação com [GET] /v2/user ❤️!")
		return
	})

	group2.Post("/user", func(c *quick.Ctx) {
		c.Set("Content-Type", "application/json")
		c.Status(200).SendString("Quick em ação com [POST] /v2/user ❤️!")
		return
	})


 app.Listen("0.0.0.0:8080")
}
```


