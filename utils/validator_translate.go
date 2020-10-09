package utils

import (
	"fmt"
	"github.com/go-playground/locales"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"log"
)

type MessageValidator struct {
	Message interface{}
}

type CustomValidator struct {
	CustomV *validator.Validate
	Trans   locales.Translator
	Dto     interface{}
}

func (u CustomValidator) TranslateValidator() (ms []string, er error) {
	v := u.CustomV
	translator := u.Trans
	uni := ut.New(translator, translator)

	trans, found := uni.GetTranslator("en")
	if !found {
		log.Fatal("translator not found")
	}

	_ = v.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is a required field", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	_ = v.RegisterTranslation("email", trans, func(ut ut.Translator) error {
		return ut.Add("email", "{0} must be a valid email", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("email", fe.Field())
		return t
	})

	_ = v.RegisterTranslation("min", trans, func(ut ut.Translator) error {
		return ut.Add("min", "{0} minimum 8 characters", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("min", fe.Field())
		return t
	})

	err := v.Struct(u.Dto)
	var messageValidator []string
	if err != nil {
		//fmt.Println(validator.ValidationErrors)
		for _, e := range err.(validator.ValidationErrors) {
			fmt.Println(e.Translate(trans))
			//messageValidator = append(messageValidator, MessageValidator{
			//	Message: e.Translate(trans),
			//})
			messageValidator = append(messageValidator, fmt.Sprintf("%v", e.Translate(trans)))
		}
	}
	return messageValidator, err
}

func (u CustomValidator) Translate() {

}
