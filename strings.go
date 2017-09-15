package main
import "fmt"
import "strings"
func main() {
	a := "aabbCCDD"
	b := "a"
	fmt.Printf("%d\n", strings.Count(a, b))
	c :=strings.ToLower(a)
	d :=strings.ToUpper(a)
	fmt.Printf("%s,%s\n",c,d)
	e := strings.Repeat("a",2)
	fmt.Printf("%s\n",e)
	f := strings.Fields(a)
	fmt.Printf("%s\n",f)
	g:="I want play computer"
	h :=strings.Split(g,"a")
	fmt.Printf("%s\n",h)
	i := strings.Join(h,"|")
	fmt.Printf("%s\n",i)
}