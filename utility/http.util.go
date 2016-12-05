package utility

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if err = json.NewDecoder(r.Body).Decode(target); err != nil {
		log.Println(err)
	}
	return err
}
