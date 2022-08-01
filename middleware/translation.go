package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"github.com/noovertime7/kubemanage/public"
	"gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
	"reflect"
)

//设置Translation
func TranslationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//参照：https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go

		//设置支持语言
		enLan := en.New()
		zhLan := zh.New()

		//设置国际化翻译器
		uni := ut.New(zhLan, zhLan, enLan)
		val := validator.New()

		//根据参数取翻译器实例
		locale := c.DefaultQuery("locale", "zh")
		trans, _ := uni.GetTranslator(locale)

		//翻译器注册到validator
		switch locale {
		case "en":
			err := en_translations.RegisterDefaultTranslations(val, trans)
			if err != nil {
				return
			}
			val.RegisterTagNameFunc(func(fld reflect.StructField) string {
				return fld.Tag.Get("en_comment")
			})
			break
		default:
			err := zh_translations.RegisterDefaultTranslations(val, trans)
			if err != nil {
				return
			}
			val.RegisterTagNameFunc(func(fld reflect.StructField) string {
				return fld.Tag.Get("comment")
			})

			////自定义验证方法
			////https://github.com/go-playground/validator/blob/v9/_examples/custom-validation/main.go
			//val.RegisterValidation("is-validuser", func(fl validator.FieldLevel) bool {
			//	return fl.Field().String() == "admin"
			//})
			//
			////自定义验证器
			////https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go
			//val.RegisterTranslation("is-validuser", trans, func(ut ut.Translator) error {
			//	return ut.Add("is-validuser", "{0} 填写不正确哦", true)
			//}, func(ut ut.Translator, fe validator.FieldError) string {
			//	t, _ := ut.T("is-validuser", fe.Field())
			//	return t
			//})
			break
		}
		c.Set(public.TranslatorKey, trans)
		c.Set(public.ValidatorKey, val)
		c.Next()
	}
}
