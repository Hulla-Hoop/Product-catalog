package handlers

import "net/http"

func (h *marketHandlers) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	reqID := r.Context().Value("reqID").(string)
	if reqID == "" {
		reqID = ""
	}

	id := r.URL.Query().Get("id")

	category, err := h.service.DeleteCategory(reqID, id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Write(category)
	w.WriteHeader(http.StatusOK)
}
