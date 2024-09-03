package handlers

import (
	"adminService/internal/configs"
	"adminService/internal/context"
	"adminService/internal/user"
	"log"
	"net/http"
	"reflect"
	"strings"
)

var types map[string]bool
var hdl *user.Handler
var urlMap map[string]map[string]reflect.Value

func init() {
	cfg := configs.Get()
	urlMap = make(map[string]map[string]reflect.Value)
	urlMap["POST"] = make(map[string]reflect.Value)
	urlMap["PUT"] = make(map[string]reflect.Value)
	urlMap["DELETE"] = make(map[string]reflect.Value)
	urlMap["GET"] = make(map[string]reflect.Value)

	maps := cfg.Handlers
	types = make(map[string]bool)
	types[".css"] = true
	types[".js"] = true
	types[".ico"] = true
	types[".jpg"] = true
	types[".png"] = true

	hdl = &user.Handler{}
	_struct := reflect.TypeOf(hdl)

	for methodNum := 0; methodNum < _struct.NumMethod(); methodNum++ {
		method := _struct.Method(methodNum)
		val, ok := maps[method.Name]
		if !ok {
			continue
		}

		urlMap[val.Method][val.Url] = reflect.ValueOf(hdl).MethodByName(method.Name)
	}
	log.Println("urlMap has been read")
}

func static(path string) bool {
	splitPath := strings.Split(path, "/")
	fileName := splitPath[len(splitPath)-1]
	splitName := strings.Split(fileName, ".")
	fileExt := "." + splitName[len(splitName)-1]
	if types[fileExt] {
		return true
	}
	return false
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	ctx := &context.Context{
		Response: w,
		Request:  r,
	}

	methodMap, ok := urlMap[r.Method]
	if !ok {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	log.Println("Page:", r.URL.Path)
	path := r.URL.Path[1:]
	if ok = static(path); ok {
		http.ServeFile(ctx.Response, ctx.Request, "./internal/static/"+path)
		return
	}

	method, ok := methodMap[path]
	if !ok {
		http.Error(w, "Path not found", http.StatusNotFound)
		return
	}

	log.Println("method: ", method)
	method.Call([]reflect.Value{reflect.ValueOf(ctx)})
}
