package handlers

import "net/http"

func (h *marketHandlers) GoodsOnCateory(w http.ResponseWriter, r *http.Request) {
	reqID, ok := r.Context().Value("reqID").(string)
	if !ok {
		reqID = ""
	}

	categor := r.URL.Query().Get("name")

	category, err := h.service.GoodsOnCateory(reqID, categor)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Write(category)

}
