package types

type Domain interface {
	Me() string
	Validate() error
}
