package manager

type Manager interface {
	CreateURL(message string) (string, error)
}

type manager struct {
	// database Database
}
