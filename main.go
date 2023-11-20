package main

import (
	"fileserve/configs"
	"fileserve/internal/handlers"
	"net/http"
	"log"
)

func main() {
	config, err := configs.Config(".")
	if err != nil {log.Fatal(err)}
	mux := http.NewServeMux()
	mux.HandleFunc(config.PathDirectory, internal.HandleDir)
	mux.HandleFunc(config.PathAddBlock, internal.AddBlockHandle)
	mux.HandleFunc(config.PathRemoveBlock, internal.RemoveBlockHandle)
	err = http.ListenAndServe(config.AppPort, mux)
	if err != nil {log.Fatal(err)}
}
