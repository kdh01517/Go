package main
import "fmt"
func main(){
	var s string
	
    for true {
        check := 1
        fmt.Scan(&s)
        if s == "0" {
            break
        }
        if len(s) == 1 {
        	fmt.Println("yes")
            continue
        }
	    for i := 0;  i <= len(s)/2; i++ {
	    	if s[i] == s[len(s)-1-i] {
	    		continue
		    } else {
		    	check = 0
		    	break
	    	}
	    }
	    if check == 1 {
	    	fmt.Println("yes")
	    } else {
	    	fmt.Println("no")
	    }
    }
}