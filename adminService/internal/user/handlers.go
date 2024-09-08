package user

import (
	"adminService/internal/context"
	"fmt"
	"net/http"
)

type Handler struct {
}

func (h *Handler) HomePage(ctx *context.Context) {
	token := ctx.Response.Header().Get("Authorization")
	fmt.Println(token)
	http.ServeFile(ctx.Response, ctx.Request, "./internal/static/html/index.html")
	return
}

func (h *Handler) LoginPage(ctx *context.Context) {
	http.ServeFile(ctx.Response, ctx.Request, "./internal/static/html/login.html")
	return
}
