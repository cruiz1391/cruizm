package cruizm

import (
	"net/http"
	"time"
	"strings"
	"fmt"
)

func init()  {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/static/", assetsHandler)
}

var (
	cacheSince = time.Now().Format(http.TimeFormat)
	cacheUntil = time.Now().AddDate(0, 0, 1).Format(http.TimeFormat)
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.Host, "www") {
		httpHeader := "http"
		if r.TLS != nil {
			httpHeader = "https"
		}
		newUrl := fmt.Sprintf("%s://www.%s%s", httpHeader, r.Host, r.URL.Path)
		if r.URL.RawQuery != "" {
			newUrl += "?" + r.URL.RawQuery
		}
		http.Redirect(w, r, newUrl, http.StatusMovedPermanently)
		return
	}

	w.Header().Set("Cache-Control", "max-age:86400, public")
	w.Header().Set("Last-Modified", cacheSince)
	w.Header().Set("Expires", cacheUntil)

	if r.URL.Path=="/" {
		http.ServeFile(w,r,"index.html")
		return
	} else if r.URL.Path=="/robots.txt" {
		http.ServeFile(w,r,"robots.txt")
		return
	} else if r.URL.Path=="/sitemap.xml" {
		http.ServeFile(w,r,"sitemap.xml")
		return
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Not Found")
}

func assetsHandler(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, r.URL.Path[1:])
}