package handlers

import "net/http"

func (h *marketHandlers) CreateCategory(w http.ResponseWriter, r *http.Request) {
	reqID := r.Context().Value("reqID").(string)
	if reqID == "" {
		reqID = ""
	}

	name := r.URL.Query().Get("name")
	category, err := h.service.CreateCategory(reqID, name)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Write(category)
	w.WriteHeader(http.StatusOK)
}
