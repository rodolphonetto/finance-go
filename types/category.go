package types

type Category struct {
	Name string `json:"name" valid:"required~Campo nome é obrigatório"`
}
