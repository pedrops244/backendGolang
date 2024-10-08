package products

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sikozonpc/ecom/types"
	"github.com/sikozonpc/ecom/utils"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore, userStore types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.handleCreateProduct).Methods(http.MethodGet)
	router.HandleFunc("/products", h.handleCreateProduct).Methods(http.MethodPost)
}

func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	ps, err := h.store.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, ps)
}
