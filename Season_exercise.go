package main
import "fmt"
func season(SeasonInt int)(SeasonEnglish string) {
	switch{
	case SeasonInt==1:
		return "January"

	case SeasonInt==2:
		return "February"
		
	case SeasonInt==3:
		return " March"
		
	case SeasonInt==4:
		return " April"
		
	case SeasonInt==5:
		return "May"
		
	case SeasonInt==6:
		return " June"
		
	case SeasonInt==7:
		return " July"
		
	case SeasonInt==8:
		return " August"
		
	case SeasonInt==9:
		return " September"
		
	case SeasonInt==10:
		return " October"
		
	case SeasonInt==11:
		return "November"
		
	case SeasonInt==12:
		return " December"
	}
	return
}
func main(){
	a := season(7)
   fmt.Println(a)
}