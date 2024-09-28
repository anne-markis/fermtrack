package server

import "net/http"

type FermtrackServer struct {
}

// TODO would be nice to shut server down good
func NewServer() http.Handler {
	// config
	// logger
	// mysql store
	mux := http.NewServeMux()
	// addRoutes(
	// 	mux,
	// 	Logger,
	// 	Config,
	// 	commentStore,
	// 	anotherStore,
	// )
	var handler http.Handler = mux
	// handler = logging.NewLoggingMiddleware(logger, handler)
	// handler = checkAuthHeaders(handler)

	return handler
}

func (fs *FermtrackServer) ListProjectsHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("GET /list route"))
}

func (fs *FermtrackServer) EditProjectHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("POST /edit route"))
}
