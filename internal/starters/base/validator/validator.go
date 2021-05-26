package validator

import (
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	vtzh "gopkg.in/go-playground/validator.v9/translations/zh"
	"github.com/blackbinbinbinbin/Bingo-gin/internal/starters"
	"github.com/blackbinbinbinbin/Bingo-gin/internal/log"
)

var validate *validator.Validate
var translator ut.Translator

func Validate() *validator.Validate {
	starters.Check(validate)
	return validate
}

func Transtate() ut.Translator {
	starters.Check(translator)
	return translator
}

type ValidatorStarter struct {
	starters.BaseStarter
}

func (v *ValidatorStarter) Init(ctx starters.StarterContext) {
	validate = validator.New()
	//创建消息国际化通用翻译器
	cn := zh.New()
	uni := ut.New(cn, cn)
	var found bool
	translator, found = uni.GetTranslator("zh")
	if found {
		err := vtzh.RegisterDefaultTranslations(validate, translator)
		if err != nil {
			l := log.DefaultLogger
			log.Error(l).Log("msg", err)
		}
	} else {
		l := log.DefaultLogger
		log.Error(l).Log("msg", "Not found translator: zh")
	}

}

func ValidateStruct(s interface{}) (err error) {
	err = Validate().Struct(s)
	if err != nil {
		logHelp := LoggerHelp()
		_, ok := err.(*validator.InvalidValidationError)
		if ok {
			logHelp.Error("msg", "验证错误")
		}
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			for _, e := range errs {
				logHelp.Error("msg", e.Translate(Transtate()))
			}
		}
		return err
	}
	return nil
}