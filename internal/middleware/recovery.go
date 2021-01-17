package middleware

//func Recovery() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		defer func() {
//			if err := recover(); err != nil {
//				global.Logger.WithCallersFrames().Errorf(c, "panic recover err: %v", err)
//
//				// TODO Email
//				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
//				c.Abort()
//			}
//		}()
//		c.Next()
//	}
//}
