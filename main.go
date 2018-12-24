package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
	"github.com/kataras/iris/websocket"
	"myiris/web/routes"
)

func main() {
	app :=iris.New()
	temp :=iris.HTML("web/views/",".html")
	app.RegisterView(temp)
	ws :=websocket.New(websocket.Config{})
	ws.OnConnection(routes.HandleConnection)
	app.Handle("GET","/", func(context iris.Context) {
		context.HTML("<h1>1</h1>")
	})
	app.Get("/ping", func(context iris.Context) {
		context.WriteString("pong")
	})
	app.PartyFunc("my_endpoint", func(p iris.Party) {
		p.Get("/",ws.Handler())
	})
	app.PartyFunc("hellorout",func(r iris.Party){
		r.Get("/",hero.Handler(routes.Hellos))
	})
	app.Get("/h", func(context iris.Context) {
		context.JSON(iris.Map{"message":"hello iris"})
	})
	app.Get("/temp", func(context iris.Context) {
		context.ViewData("Title","hel1")
		context.ViewData("Content","123445")
		context.View("index.html")
	})
	app.Run(iris.Addr(":8086"))
}
