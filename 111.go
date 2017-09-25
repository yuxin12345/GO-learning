package main
import (
	"fmt"
	"net/http"
)
type Mymux struct {
}
func (p *Mymux) ServeHTTP(w http.ResponseWriter,r *http.Request){
	if r.URL.Path =="/"{
		say(w,r)
		return
	}
	http.NotFound(w,r)
	return
}
func say(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"nimahai")
}
func main(){
	mux := &Mymux{}
	http.ListenAndServe(":9090",mux)
}