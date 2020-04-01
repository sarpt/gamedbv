package serv

import "net/http"

const corsOptionsMethod string = "OPTIONS"

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Access-Control-Allow-Origin", "*")
		resp.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		resp.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Method")

		if req.Method == corsOptionsMethod {
			return
		}

		next.ServeHTTP(resp, req)
	})
}
