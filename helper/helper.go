package helper

import "github.com/go-playground/validator/v10"

type response struct { // maping untuk response json
	Meta meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type meta struct { // maping pesan, code, status
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}) response { //fungsi dan parsing parameter
	meta := meta{
		Message: message,
		Code:    code,
		Status:  status,
	}
	jsonResponse := response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse //return
}

func FormatValidationError(err error) []string {
	var errors []string // merubah pesan error menjadi array of string
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}
	return errors
}
