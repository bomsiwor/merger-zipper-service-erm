package entity

type DocumentSource struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type Documents []DocumentSource

type DocumentRequest struct {
	Source     Documents `json:"source"`
	OutputName string    `json:"outputFileName"`
	OutputDir string `json:"targetDir"`
}
