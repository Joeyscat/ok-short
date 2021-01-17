package middleware

//func Translations() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		uni := ut.New(en.New(), zh.New(), zh_Hant_TW.New())
//		locale := c.GetHeader("locale")
//		trans, _ := uni.GetTranslator(locale)
//		v, ok := binding.Validator.Engine().(*validator.Validate)
//		if ok {
//			switch locale {
//			case "zh":
//				_ = zh_translations.RegisterDefaultTranslations(v, trans)
//				break
//			case "en":
//				_ = en_translations.RegisterDefaultTranslations(v, trans)
//				break
//			default:
//				_ = zh_translations.RegisterDefaultTranslations(v, trans)
//				break
//			}
//			c.Set("trans", trans)
//		}
//		c.Next()
//	}
//}
