package types

type Pessoa struct {
	Name string `json:"name" valid:"required~Campo nome é obrigatório"`
	Age  int    `json:"age" valid:"range(1|2)~Você deve adicionar numeros entre 1 e 2"`
}
