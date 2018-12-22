package httpSvr

import (
	"log"
	"net/http"
	"os"
	"strings"
)
var logger = log.New(os.Stdout,"http_",log.Lshortfile|log.Ldate|log.Ltime)
func helloWorld(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path !="/"{
		http.NotFound(w,r)
		return
	}
	w.Write([]byte("hello world!"))

}

func redirectBaidu(w http.ResponseWriter, r *http.Request)  {
	http.Redirect(w,r,"http://www.baidu.com",300)
}

func getQuery(w http.ResponseWriter, r *http.Request)  {
	q := r.URL.Query()
	for k,v:=range q{
		writer:=strings.Builder{}
		writer.WriteString(k)
		writer.WriteString("\t:\t")
		for _,l:=range v{
			writer.WriteString(l)
		}
		writer.Write([]byte("\n"))
		logger.Print(writer.String())
		w.Write([]byte(writer.String()))
	}
}