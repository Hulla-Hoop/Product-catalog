package handlers

import "net/http"

func (h *marketHandlers) AllCategories(w http.ResponseWriter, r *http.Request) {
	reqID := r.Context().Value("reqID").(string)
	if reqID == "" {
		reqID = ""
	}

	category, err := h.service.AllCategories(reqID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Write(category)
	w.WriteHeader(http.StatusOK)
}
