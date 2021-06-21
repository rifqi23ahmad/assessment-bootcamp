package main

func main() {
	r.gin.Default()

	r.Get("/users", UserGetAll)
	r.run()
}
