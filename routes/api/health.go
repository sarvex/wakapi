package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
)

type HealthApiHandler struct {
	db *gorm.DB
}

func NewHealthApiHandler(db *gorm.DB) *HealthApiHandler {
	return &HealthApiHandler{db: db}
}

func (h *HealthApiHandler) RegisterRoutes(router *mux.Router) {
	r := router.PathPrefix("/health").Subrouter()
	r.Methods(http.MethodGet).HandlerFunc(h.Get)
}

func (h *HealthApiHandler) Get(w http.ResponseWriter, r *http.Request) {
	var dbStatus int
	if sqlDb, err := h.db.DB(); err == nil {
		if err := sqlDb.Ping(); err == nil {
			dbStatus = 1
		}
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(fmt.Sprintf("app=1\ndb=%d", dbStatus)))
}