package handler

import "github.com/gorilla/mux"

func NewRouter(r *mux.Router) {
	// Routes
	// Fallback / error route
	registerErrorHandler(r);

	registerZipperHanler(r)
	registerCombineHandler(r)
}

// Register Error Handler
func registerErrorHandler(r *mux.Router) {
	// Error handler (Default)
	errorH := NewDefaultErrorHandler()

	r.MethodNotAllowedHandler = errorH.MethodNotAllowedHandler()
	r.NotFoundHandler = errorH.NotFoundHandler()
}

// Create new ZIP
func registerZipperHanler(r *mux.Router) {
	// Creat enew handler
	zipperH := NewZipperHandler()

	zipperGroup := r.PathPrefix("/zipper")

	zipperGroup.Methods("POST").HandlerFunc(zipperH.ZipperByPathHandler)
}

// Create new Merged PDF
func registerCombineHandler(r *mux.Router) {
	// Creat enew handler
	mergerH := NewMergeHandler()

	mergerGroup := r.PathPrefix("/merger")

	mergerGroup.Methods("POST").HandlerFunc(mergerH.MergerByPathHandler)
}