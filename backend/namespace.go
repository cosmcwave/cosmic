package backend

import "net/http"

type Namespace struct {
	Services     []*Service `json:"namespace"`
	Hash         string     `json:"hash"`
	LatestChange string     `json:"latestChange"`
}

func namespaceHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("namespace handler"))
}
