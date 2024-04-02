package handlers

import "net/http"

func (h *marketHandlers) GoodsOnCateory(w http.ResponseWriter, r *http.Request) {
	reqID, ok := r.Context().Value("reqID").(string)
	if !ok {
		reqID = ""
	}

	categor := r.URL.Query().Get("name")
	if categor == "" {
		http.Error(w, "пустой параметр name", http.StatusBadRequest)
		return
	}

	category, err := h.service.GoodsOnCateory(reqID, categor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

	w.Write(category)

}
