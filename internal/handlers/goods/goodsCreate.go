package handlers

import (
	"fmt"
	"net/http"
)

func (h *marketHandlers) CreateGoods(w http.ResponseWriter, r *http.Request) {
	reqID, ok := r.Context().Value("reqID").(string)
	if !ok {
		reqID = ""
	}

	name := r.URL.Query().Get("name")
	fmt.Println("-----------------------", name)
	category, err := h.service.CreateGoods(reqID, name)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Write(category)
}
