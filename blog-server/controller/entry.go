package controller

type BlogCo struct {
	Admin   AdminCo   // 生产环境的
	Example ExampleCo // 测试环境的
}

var CO = new(BlogCo)
