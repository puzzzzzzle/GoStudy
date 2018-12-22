package httpSvr

import "net/http"

func  StartHttpSvr(){
	http.HandleFunc("/",helloWorld)
	http.HandleFunc("/bd",redirectBaidu)
	http.HandleFunc("/q",getQuery)

	http.ListenAndServe(":8000",nil)
}
