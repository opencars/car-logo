package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
)

// LogoHandler is responsible for returning URL of logo by vehicle manufacturer name.
type LogoHandler struct {
	images map[string]string
}

// Logo is object, that returns by the server for representing information about the logo URL.
type Logo struct {
	URL string `json:"url"`
}

// ServeHTTP returns information information about the car logo in JSON by its manufacturer.
// Implements http.Handler interface.
func (h *LogoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	make := strings.ToLower(mux.Vars(r)["make"])

	if _, ok := h.images[make]; !ok {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	fmt.Println(r.URL.String())

	w.Header().Set("Content-Type", "application/json")
	url := r.Host + "/emblems/" + make + ".png"
	if err := json.NewEncoder(w).Encode(Logo{URL: url}); err != nil {
		http.Error(w, "Broken", http.StatusInternalServerError)
		return
	}
}

func NewHandler() *LogoHandler {
	emblems := make(map[string]string)

	err := filepath.Walk("./emblems", func(path string, info os.FileInfo, err error) error {
		fixed := strings.ReplaceAll(filepath.Base(path), ".png", "")
		emblems[fixed] = path
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return &LogoHandler{emblems}
}
