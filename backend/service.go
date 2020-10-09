package backend

import "net/http"

type Service struct {
	Name         string `json:"name"`
	Version      string `json:"version"`
	Hash         string `json:"name"`
	Location     string `json:"location"`
	Repository   string `json:"repository"`
	Alive        bool   `json:"alive"`
	LatestChange string `json:"latestChange"`
	CreatedAt    string `json:"created"`
}

func serviceHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("service handler"))
}
