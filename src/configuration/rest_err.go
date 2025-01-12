package rest_err //este pacote serve para representar e gerenciar os erros, podendo ser utilizado na aplicação inteira
import (
	"net/http"
)

type RestErr struct {
	Message string   `json:"message"`
	Err     string   `json:"errors"`
	Code    int64    `json:"code"`
	Causes  []Causes `json:"causes,omitempty"` //slice de structs (lista dinamica) usada para representar erros com multiplas causas
}

type Causes struct { //struct para detalhar melhor os erros
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *RestErr) Error() string { //receber e retornar erros por parametro
	return r.Message
}

func NewRestErr(message, err string, code int64, causes []Causes) *RestErr { //construtor apontando para o metodo RestErr
	return &RestErr{
		Message: message, //Mensagem que descreve o erro
		Err:     err,     //Tipo de erro (not found)
		Code:    code,    //Código HTTP
		Causes:  causes,  //Causas do erro
	}
}

func NewBadRequestError(message string) *RestErr { //struct usada para retornar quando acontece um bad request simples
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest, //400 Bad Request: O servidor não pode processar a requisição devido a erro do cliente (dados inválidos, formato incorreto, etc.).
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *RestErr { //struct usada para especificar um bad request, por ex: "campo email obrigatório"
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest, //400 Bad Request: O servidor não pode processar a requisição devido a erro do cliente (dados inválidos, formato incorreto, etc.).
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *RestErr { //retorna erros internos do servidor
	return &RestErr{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError, //500 Internal Server Error: O servidor encontrou uma falha inesperada e não conseguiu processar a solicitação.
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound, //404 Not Found: O recurso solicitado não foi encontrado no servidor. Ex: página não encontrada.
	}
}

func NewForbiddenError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden, //403 Forbidden: O cliente está autenticado, mas não tem permissão para acessar o recurso ou realizar a ação.
	}
}
