package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/syamilAbdillah/post-book/internal"
	"github.com/syamilAbdillah/post-book/pkg/webrpc"
)


func Handler(w http.ResponseWriter, r *http.Request) {
	server().ServeHTTP(w, r)
}

func server() http.Handler {
	r := chi.NewRouter()
	
	r.Use(middleware.Logger)

	// mount services
	r.Handle(
		"/rpc/UserService/*", 
		webrpc.NewUserServiceServer(internal.NewUserService()),
	)
	r.Handle(
		"/rpc/PostService/*", 
		webrpc.NewPostServiceServer(internal.NewPostService()),
	)


	return r
}