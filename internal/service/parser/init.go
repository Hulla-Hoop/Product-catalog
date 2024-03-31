package parser

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testinhousead/internal/model"
	"testinhousead/internal/service"
)

func Parse(p service.Parser, reqID string) error {

	var v []model.Product

	r, err := http.Get("https://emojihub.yurace.pro/api/all")
	if err != nil {
		return err
	}

	sl, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(sl, &v)
	if err != nil {
		fmt.Println(err)
	}
	for _, a := range v {
		p.CreateGoods(reqID, a.Name, a.Category)
	}

	return nil
}
