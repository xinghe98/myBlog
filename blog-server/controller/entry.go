package controller

type BlogCo struct {
	Admin   AdminCo
	Example ExampleCo
}

var CO = new(BlogCo)
