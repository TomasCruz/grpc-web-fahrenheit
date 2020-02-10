package model

var (
	client Client
)

// SetClient sets client interface
func SetClient(cli Client) {
	client = cli
}
