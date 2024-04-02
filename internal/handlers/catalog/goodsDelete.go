package handlers

import "net/http"

// example /goods/delete?id=3

func (h *marketHandlers) DeleteGoods(w http.ResponseWriter, r *http.Request) {
	reqID, ok := r.Context().Value("reqID").(string)
	if !ok {
		reqID = ""
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "пустой параметр id", http.StatusBadRequest)
		return
	}
	category, err := h.service.DeleteGoods(reqID, id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Write(category)
}
