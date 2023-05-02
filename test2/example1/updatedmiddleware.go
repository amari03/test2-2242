//Filename: main.go

/*
What is middleware? 
When you're building a web application there's probably some shared functionality 
that you want to run for many (or even all) HTTP requests.
organising this shared functionality is to set it up as middleware — self-contained 
code which independently acts on a request before or after your normal application handlers.

Some operations, such as authentication, logging, and cookie validation, are 
often implemented as middleware functions, which act on a request independently 
before or after the regular route handlers.

 In Go a common place to use middleware is between a router
 and application handlers.

 simple demonstration below
*/

package main

import (
    "log"
    "net/http"
)

//write middleware
//The middleware function accepts a handler and returns a handler.
func middlewareA(next http.Handler) http.Handler{
    return  http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
        //this is executed on the way down to the handler
        log.Println("Executing middleware A")
        //above- code to handle the request part
        next.ServeHTTP(w,r) 
        //below- code to handle the response part
        log.Println("Executing middleware A again")//this is executed on the way up to the client
    })
}

func middlewareB(next http.Handler) http.Handler{
    return  http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
        //this is executed on the way down to the handler
        log.Println("Executing middleware B")
        //above- code to handle the request part
        if r.URL.Path == "/mango"{
            return
        }
        next.ServeHTTP(w,r) 
        //below- code to handle the response part
        log.Println("Executing middleware B again")//this is executed on the way up to the client
    })
}

// create a handler
//values string
// key functions
//functin that will deal with the HTTP methods
func ourHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("Executing the hanler...")
    w.Write([]byte("STRAWBERRIES\n"))
    

}
func main() {
    //multiplexer act lik a router
    mux := http.NewServeMux()//multiplexer
    //creating an arbitrarily long handler chain by nesting middleware functions inside each other.
    //we can do this because the function accepts a handler as a parameter and returns a handler.
    mux.Handle("/", middlewareA(middlewareB(http.HandlerFunc(ourHandler))))// "/= endpoint(url)where you wanna go" key "home = value"

    log.Print("starting server on :4000")//value
    err := http.ListenAndServe(":4000", mux)//we create a server
    // the job of the web server listen for request does not process
    // when it gets the request on the right port it will send it to the multiplexer/router and then it will send it to the handler
    log.Fatal(err)
}
