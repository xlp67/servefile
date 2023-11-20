package internal

import (
	"fileserve/configs"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"path/filepath"
	"strings"
)
func HandleDir(w http.ResponseWriter, r *http.Request) {
	config, err := configs.Config(".")
	if err != nil {panic(err)}
	url_path := strings.TrimPrefix(r.URL.Path, config.PathDirectory)
	filePath := filepath.Join(config.Directory, url_path)
	var check string
	for _, blockdir := range config.BlockDirectories {
		if strings.Contains(filePath, blockdir) {
			check = blockdir
			break
		}
	}
	trim_handle := strings.TrimPrefix(r.URL.Path, config.PathDirectory)
	all_files := fmt.Sprintf("%s/%s", config.Directory, trim_handle)
	_, err = exec.Command("ls", all_files).CombinedOutput()
	if check == "" && err == nil {
		log.Printf("request succeeded on path: %s\n", trim_handle)
		http.ServeFile(w, r, filePath)
	} else {
			if err != nil {
				log.Println("path not found")
				w.Write([]byte("path not found"))
				w.WriteHeader(http.StatusNotFound)
			} else {
			log.Printf("request blocked in path: %s", check)	
			w.Write([]byte("request blocked"))
			w.WriteHeader(http.StatusBadRequest)}
	}
}