package transports

type Server interface {
	Start() error
	Stop() error
}
