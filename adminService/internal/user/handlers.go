package user

import (
	"adminService/internal/context"
	"net/http"
)

type Handler struct {
}

func (h *Handler) HomePage(ctx *context.Context) {
	http.ServeFile(ctx.Response, ctx.Request, "./internal/static/html/index.html")
	return
}

func (h *Handler) LoginPage(ctx *context.Context) {
	http.ServeFile(ctx.Response, ctx.Request, "./internal/static/html/loginPage.html")
	return
}
