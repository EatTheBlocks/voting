package ipfs

type Service interface {
	PinJSON(name string, body interface{}) (string, error)
}
