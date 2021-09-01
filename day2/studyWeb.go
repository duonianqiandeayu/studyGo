package main

import "net/http"


func main(){
	http.HandleFunc("/",someFunc)
	// http.ListenAndServe(":5000",mMux)
	print("started!")
	
}

func someFunc(w http.ResponseWriter, req *http.Request){
	w.Write([]byte("hello go web app"))
}

func aboutFunc(){

}