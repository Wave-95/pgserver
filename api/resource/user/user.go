package user

import (
	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v5/pgxpool"
)

type API struct {
	//Dependencies stored here
	Repo Repository
}

func NewUserApi(db *pgxpool.Pool) *API {
	userRepository := NewUserRepository(db)
	return &API{Repo: &userRepository}
}

func (api *API) SetupRoutes(r chi.Router) {
	r.Get("/users/{userID}", api.handleGetUser)
}
