package lambroll

// Option represents common option.
type Option struct {
	Region   *string
	Profile  *string
	TFState  *string
	Endpoint *string
	Envfile  *[]string
	ExtStr   *map[string]string
	ExtCode  *map[string]string
}
