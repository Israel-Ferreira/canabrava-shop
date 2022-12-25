package requests

type EnderecoReq struct {
	Logradouro  string `json:"logradouro"`
	Numero      uint   `json:"numero"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Cidade      string `json:"cidade"`
	Uf          string `json:"uf"`
}

type InfoProprietario struct {
	Nome     string `json:"nome"`
	Email    string `json:"email"`
	Cpf      string `json:"cpf"`
	Telefone string `json:"telefone,omitempty"`
}

type AdegaReq struct {
	NomeFantasia string           `json:"nome_fantasia"`
	RazaoSocial  string           `json:"razao_social"`
	Cnpj         string           `json:"cnpj"`
	Email        string           `json:"email"`
	Telefone     string           `json:"telefone"`
	InfoProp     InfoProprietario `json:"info_proprietario"`
	Endereco     EnderecoReq      `json:"endereco"`
}
