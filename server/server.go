package server

import "net/http"

// TODO this naming bothers me
type FermtrackServer struct {
}

func (fs *FermtrackServer) GetProjectHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("GET /list route"))
}

func (fs *FermtrackServer) ListProjectsHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("GET /list route"))
}

func (fs *FermtrackServer) EditProjectHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("POST /edit route"))
}
