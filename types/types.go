package types

type Software struct {
	Name    string `json:"name"`
	Version string `json:"vers,omitempty"`
}
