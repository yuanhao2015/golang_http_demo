package utils

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func WriteLog(r *http.Request, t time.Time, match string, pattern string) {

	if Conf.ReadStr("website", "loglevel") != "prod" {

		d := time.Now().Sub(t)

		l := fmt.Sprintf("[ACCESS] | % -10s | % -40s | % -16s | % -10s | % -40s |", r.Method, r.URL.Path, d.String(), match, pattern)

		log.Println(l)
	}
}
