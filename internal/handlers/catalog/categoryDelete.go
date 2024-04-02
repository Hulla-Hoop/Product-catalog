package handlers

import "net/http"

// example /category/delete?id=1

func (h *marketHandlers) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	reqID, ok := r.Context().Value("reqID").(string)
	if !ok {
		reqID = ""
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "пустой параметр id", http.StatusBadRequest)
		return
	}

	category, err := h.service.DeleteCategory(reqID, id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Write(category)

}
