package main

import(

	"fmt"
	"net/http"
	
)

func firstFunc(w http.ResponseWriter,r *http.Request){
	fmt.Fprint(w,"First message displayed on website")
}

func main(){
	http.HandleFunc("/",firstFunc)
	fmt.Print("Server is running on 3000")
	http.ListenAndServe(":3000",nil)
}	