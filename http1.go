package main
import(
	"fmt"
	"net/http"
	"strings"
	"log"
)
func say(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path",r.URL.Path)
	fmt.Println("scheme",r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k,v:=range r.Form{
		fmt.Println("key:",k)
		fmt.Println("value:",strings.Join(v,""))
	}
	fmt.Fprintln(w,"aksdfahdhfighowieang")
	fmt.Fprintln(w,"dfafdwfeaf")
}
func main(){
	http.HandleFunc("/",say)
	err := http.ListenAndServe(":9999",nil)
	if err != nil{
		log.Fatal("dfla",err)
	}
}