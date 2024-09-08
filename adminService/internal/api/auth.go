package api

import (
	"adminService/internal/context"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func (h *Handler) Auth(ctx *context.Context) {
	var data map[string]string
	err := json.NewDecoder(ctx.Request.Body).Decode(&data)
	if err != nil {
		log.Println(err)
	}

	dataJson, err := json.Marshal(data)

	res, err := http.Post("http://localhost:8081/auth/admin", "application/json", bytes.NewReader(dataJson))
	err = json.NewDecoder(ctx.Request.Body).Decode(res)
	token := res.Header.Get("Authorization")
	if token != "" {
		cookie := &http.Cookie{
			Name:  "token",
			Value: token,
			Path:  "/",
		}
		http.SetCookie(ctx.Response, cookie)

		response := struct {
			Ok bool `json:"ok"`
		}{
			Ok: true,
		}
		err = json.NewEncoder(ctx.Response).Encode(response)
		return
	}

	ctx.Response.WriteHeader(401)
	return
}
