package validator

import (
	"fmt"
	"gin-project/models"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var Trans ut.Translator

func InitTrans(lang string) (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := field.Tag.Get("json")
			if name == "-" {
				return ""
			}
			return name
		})

		v.RegisterStructValidation(SignUpParamStructLevelValidation, models.ParamsRegister{})

		zhT := zh.New()
		enT := en.New()
		uni := ut.New(enT, zhT, enT)

		var ok bool
		Trans, ok = uni.GetTranslator(lang)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", lang)
		}
		switch lang {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, Trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, Trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, Trans)
		}
		return
	}
	return
}

func RemoveTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.LastIndex(field, ".")+1:]] = err
	}
	return res
}

func SignUpParamStructLevelValidation(sl validator.StructLevel) {
	su := sl.Current().Interface().(models.ParamsRegister)
	if su.Password != su.RePassword {
		sl.ReportError(su.RePassword, "re_password", "rePassword", "eqfield", "password")
	}
}
