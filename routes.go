package cruizm

import (
	"net/http"
)

func init()  {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/static/", assetsHandler)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path=="/" {
		http.ServeFile(w,r,"index.html")
		return
	}
	http.Redirect(w,r,"/", http.StatusTemporaryRedirect)
}

func assetsHandler(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, r.URL.Path[1:])
}