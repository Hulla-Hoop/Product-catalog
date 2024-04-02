package handlers

import "net/http"

// example /category/create?name=Fresh

func (h *marketHandlers) CreateCategory(w http.ResponseWriter, r *http.Request) {
	reqID, ok := r.Context().Value("reqID").(string)
	if !ok {
		reqID = ""
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "пустой параметр name", http.StatusBadRequest)
		return
	}

	category, err := h.service.CreateCategory(reqID, name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(category)
}
