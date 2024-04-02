package handlers

import "net/http"

// example /allcategories

func (h *marketHandlers) AllCategories(w http.ResponseWriter, r *http.Request) {
	reqID, ok := r.Context().Value("reqID").(string)
	if !ok {
		reqID = ""
	}

	category, err := h.service.AllCategories(reqID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(category)
}
