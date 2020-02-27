/*
   @Time : 2020/2/22 20:50
   @Author : wangbo
   @File : auth
*/
package handle

import (
	"SkyDrive/utils"
	"fmt"
	"net/http"
)

func HTTPInterceptor(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		username := r.Form.Get("username")
		token := r.Form.Get("token")
		fmt.Println(username, token)
		if len(username) < 3 || !utils.IsTokenVaild(token) {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("拒绝访问"))
			return
		}
		h.ServeHTTP(w, r)
	}
}
