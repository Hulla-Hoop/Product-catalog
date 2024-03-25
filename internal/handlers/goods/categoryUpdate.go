package handlers

import "net/http"

func (h *marketHandlers) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	reqID := r.Context().Value("reqID").(string)
	if reqID == "" {
		reqID = ""
	}

	id := r.URL.Query().Get("id")
	name := r.URL.Query().Get("name")

	category, err := h.service.UpdateCategory(reqID, id, name)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Write(category)
	w.WriteHeader(http.StatusOK)
}
