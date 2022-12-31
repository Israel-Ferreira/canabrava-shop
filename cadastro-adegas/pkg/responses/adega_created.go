package responses

type AdegaCreatedMsg struct {
	NomeFantasia string `json:"nome_fantasia"`
	RazaoSocial  string `json:"razao_social"`
	Cnpj         string `json:"cnpj"`
	Email        string `json:"email"`
	Telefone     string `json:"telefone"`
	Status       string `json:"status"`
}
