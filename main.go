package main

import(
	"fmt"
	"log"
	"net/http"
)



func main(){

	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		log.Println("recieved request")
		fmt.Fprint(w, "Hello from Docker!!")
	})


	log.Println("start server")
	server := &http.Server{Addr: ":8080"}
	if err := server.ListenAndServe(); err != nil{
		log.Println(err)
	}


}



