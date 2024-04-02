package handlers

import (
	"net/http"
)

// example /goods/create?name=Apple&category=Fresh

func (h *marketHandlers) CreateGoods(w http.ResponseWriter, r *http.Request) {
	reqID, ok := r.Context().Value("reqID").(string)
	if !ok {
		reqID = ""
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "пустой параметр name", http.StatusBadRequest)
		return
	}
	categorys := r.URL.Query().Get("category")
	if categorys == "" {
		http.Error(w, "пустой параметр categorys", http.StatusBadRequest)
		return
	}
	h.logger.L.WithField("MARKETHANDLERS.CREATEGOODS", reqID).Debug("Name-", name, "Category--", categorys)

	product, err := h.service.CreateGoods(reqID, name, categorys)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Write(product)
}
