package manager

type Manager interface {
	EncodeURL(url string) string
	DecodeURL(url string) string
	DeleteURL(url string) error
}
