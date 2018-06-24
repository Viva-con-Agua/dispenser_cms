package main

         import "github.com/kataras/iris"
         

         func main() {
           app := iris.Default()

           // Method:   GET
           // Resource: http://localhost:8080/
           app.RegisterView(iris.HTML("./public", ".html"))
           app.Handle("GET", "/", func(ctx iris.Context) {
             ctx.View("index.html")
           })

           // same as app.Handle("GET", "/ping", [...])
           // Method:   GET
           // Resource: http://localhost:8080/ping
           app.Get("/ping", func(ctx iris.Context) {
             ctx.WriteString("pong")
           })
					 
	   app.StaticWeb("/static/js", "./public/static/js")
           app.StaticWeb("/static/css", "./public/static/css")

           // Method:   GET
           // Resource: http://localhost:8080/hello
           app.Get("/hello", func(ctx iris.Context) {
             ctx.JSON(iris.Map{"message": "Hello iris web framework."})
           })

           // http://localhost:8080
           // http://localhost:8080/ping
           // http://localhost:8080/hello
           app.Run(iris.Addr(":8080"))
         }
