package services

import (
	"net/mail"

	"github.com/Israel-Ferreira/canabrava-shop/cadastro-adegas/internal/config"
	"github.com/Israel-Ferreira/canabrava-shop/cadastro-adegas/internal/models"
	"github.com/Israel-Ferreira/canabrava-shop/cadastro-adegas/pkg/exceptions"
	"github.com/Israel-Ferreira/canabrava-shop/cadastro-adegas/pkg/requests"
)

func SaveAdegaAndSendToSQS(adegaReqBody requests.AdegaReq) error {
	validationError := validateRequestBody(adegaReqBody)

	if validationError != nil {
		return validationError
	}

	adega := models.NewAdega(&adegaReqBody)

	

	if txn := config.Db.Save(&adega); txn.Error != nil {
		return txn.Error
	}

	return nil
}

func EmailIsValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func validateRequestBody(adegaReqBody requests.AdegaReq) error {
	if adegaReqBody.Cnpj == "" {
		return exceptions.ErrCnpjBlank
	}

	if adegaReqBody.NomeFantasia == "" {
		return exceptions.ErrNomeFantasiaBlank
	}

	if adegaReqBody.Email == "" {
		return exceptions.ErrEmailBlank
	}

	if !EmailIsValid(adegaReqBody.Email) {
		return exceptions.ErrEmailValidationFailed
	}

	return nil
}
