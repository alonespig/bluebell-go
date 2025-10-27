package global

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"go.uber.org/zap"
)

var Trans ut.Translator

func InitTranslation() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zhT := zh.New()
		uni := ut.New(zhT, zhT)
		var ok bool
		Trans, ok = uni.GetTranslator("zh")
		if !ok {
			zap.S().Fatalln("get zh translator failed")
		}
		err := zh_translations.RegisterDefaultTranslations(v, Trans)
		if err != nil {
			zap.S().Fatalln("init translator failed", err)
		}
	}
}
