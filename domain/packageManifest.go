package domain

type PackageManifest struct {
	Name  string   `json:"name"`
	Hash  string   `json:"hash"`
	Files []string `json:"files"`
}
