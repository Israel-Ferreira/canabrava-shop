package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Israel-Ferreira/canabrava-shop/cadastro-adegas/internal/services"
	"github.com/Israel-Ferreira/canabrava-shop/cadastro-adegas/pkg/exceptions"
	"github.com/Israel-Ferreira/canabrava-shop/cadastro-adegas/pkg/requests"
)

func CadastrarAdega(rw http.ResponseWriter, r *http.Request) {
	var adegaReq requests.AdegaReq

	if err := json.NewDecoder(r.Body).Decode(&adegaReq); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	errService := services.SaveAdegaAndSendToSQS(adegaReq)

	if errService != nil {
		checkError(rw, errService)
		return
	}


	
	
	rw.WriteHeader(http.StatusCreated)
	
}

func checkError(rw http.ResponseWriter, err error) {

	jsonErrorMsg := map[string]string{
		"msg": err.Error(),
	}

	resp, _ := json.Marshal(jsonErrorMsg)

	switch err {
	case exceptions.ErrCnpjBlank:
	case exceptions.ErrEmailBlank:
	case exceptions.ErrNomeFantasiaBlank:
	case exceptions.ErrEmailValidationFailed:
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write(resp)
		return
	default:
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write(resp)
		return
	}

}
