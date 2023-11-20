package internal

import (
	"fileserve/configs"
	"fileserve/internal/entity"
	"log"
	"net/http"
	"strings"
)

func AddBlockHandle(w http.ResponseWriter, r *http.Request) {
	config, err := configs.Config(".")
	if err != nil {log.Fatal(err)}
	if r.URL.Path == "/favicon.ico/" {
	} else {
		block := strings.TrimPrefix(r.URL.Path, config.PathAddBlock)
		internal.Insert(block)
	}
		
}
func RemoveBlockHandle(w http.ResponseWriter, r *http.Request) {
	config, err := configs.Config(".")
	if err != nil {log.Fatal(err)}
	block := strings.TrimPrefix(r.URL.Path, config.PathRemoveBlock)
	internal.Remove(block)
}