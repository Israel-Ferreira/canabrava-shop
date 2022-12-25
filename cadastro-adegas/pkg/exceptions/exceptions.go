package exceptions

import "errors"

var (
	ErrCnpjBlank         = errors.New("error: o cnpj da adega não pode estar em branco")
	ErrNomeFantasiaBlank = errors.New("error: O nome fantasia da adega não pode estar em branco")
	ErrEmailBlank        = errors.New("error: O email não deve estar em branco")
	ErrEmailValidationFailed = errors.New("error: o email informado não é valido")
)
