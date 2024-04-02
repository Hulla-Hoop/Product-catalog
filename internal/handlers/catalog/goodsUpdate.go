package handlers

import "net/http"

// example /goods/update?id=3&name=Fresh

func (h *marketHandlers) UpdateGoods(w http.ResponseWriter, r *http.Request) {
	reqID, ok := r.Context().Value("reqID").(string)
	if !ok {
		reqID = ""
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "пустой параметр id", http.StatusBadRequest)
		return
	}
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "пустой параметр name", http.StatusBadRequest)
		return
	}
	category, err := h.service.UpdateGoods(reqID, id, name)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Write(category)
	w.WriteHeader(http.StatusOK)
}
