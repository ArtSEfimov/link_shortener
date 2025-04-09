package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

func Chain(middlewares ...Middleware) Middleware {
	middleware := func(next http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}
		return next
	}
	return middleware
}

// Python analogue for a better understanding

//	def chain(*handlers):
//		def handler(next_handler, *args):
//			for func in reversed(handlers):
//				next_handler = func(next_handler, *args)
//
//			return next_handler
//
//		return handler
//
