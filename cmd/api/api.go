package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rogedev/expenses_api/service/user"
)

type ApiServer struct {
	addr string
	db   *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *ApiServer {
	return &ApiServer{
		addr: addr,
		db:   db,
	}
}

func (s *ApiServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("api/").Subrouter()
	userStore := user.NewStore(s.db)
	userService := user.NewHandler(userStore)
	userService.RegisterRoutes(subrouter)
	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, router)
}
