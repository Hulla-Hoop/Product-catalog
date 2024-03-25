package handlers

import "net/http"

func (h *marketHandlers) AllCategories(w http.ResponseWriter, r *http.Request) {
	reqID, ok := r.Context().Value("reqID").(string)
	if !ok {
		reqID = ""
	}

	category, err := h.service.AllCategories(reqID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Write(category)
	w.WriteHeader(http.StatusOK)
}
