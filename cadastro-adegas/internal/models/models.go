package models

import (
	"github.com/Israel-Ferreira/canabrava-shop/cadastro-adegas/pkg/requests"
	"gorm.io/gorm"
)

type Adega struct {
	gorm.Model
	NomeFantasia string `gorm:"column:nome_fantasia"`
	RazaoSocial  string `gorm:"column:razao_social"`
	Cnpj         string `gorm:"unique"`
	Email        string `gorm:"unique"`
	Telefone     string `gorm:"column:telefone"`
	Aproved      bool   `gorm:"column:aproved"`
}

type EnderecoAdega struct {
	gorm.Model
	AdegaId int
	Adega   Adega 

	Logradouro string `gorm:"column:logradouro"`
	Numero     uint   `gorm:"column:numero"`
	Cep        string `gorm:"column:cep"`
	Bairro     string `gorm:"column:bairro"`
	Cidade     string `gorm:"column:cidade"`
	Uf         string `gorm:"column:uf"`
}

func NewAdega(adegaReq *requests.AdegaReq) *Adega {
	return &Adega{
		NomeFantasia: adegaReq.NomeFantasia,
		RazaoSocial:  adegaReq.RazaoSocial,
		Cnpj:         adegaReq.Cnpj,
		Email:        adegaReq.Email,
		Telefone:     adegaReq.Telefone,
	}
}
