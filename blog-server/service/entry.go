package service

type Logic struct {
	AdminRL // Admin的注册和登录
	Article // 文章的crud
}

var LOC = new(Logic)
