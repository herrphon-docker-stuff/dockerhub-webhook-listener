package api

type Config struct {
	ListenAddr string
	Tls        struct {
		Key  string
		Cert string
	}
	Apikeys struct {
		Key []string
	}
}
