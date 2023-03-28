package main

import (
	"fmt"
	"time"
)


func main(){
	fmt.Println("THIS IS MY FIRST PROJECT AND IT IS ABOUT PRACTICING WHAT I HAVE LEARNT USING FOR,SWTCH ,AND THE LIKES")


	name_of_footballers := []string{
		"Ronaldo",
		 "Messi",
		  "Neymar",
		   "Mbappe",
		    "Saka",
			 "Jesus", 
			 "Benzema",
			  "AY BUMBUM",
			}
	
	i := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(i)

	s := []struct {
		i int
		s string
	}{
		{1, "Ronaldo"},
		{2, "Messi"},
		{3, "Neymar"},
		{4, "Mbappe"},

	}
	fmt.Println(s)

	fmt.Println("Names of footballers from the best to the worst:  \n", name_of_footballers)
	fmt.Println("Selecting the name of the best two players")
	var f []string = name_of_footballers[:2]
	fmt.Println(f)

	fmt.Println("When Will Blacklist Season 10 episode 5 be released?")
	this_Month := time.Now().Month()
	switch time.April {
	case this_Month + 0:
		fmt.Println("This Month.")
	case this_Month + 1:
		fmt.Println("Next Month/April.")
	case this_Month + 2:
		fmt.Println("May")
	default:
		fmt.Println("Not Anytime soon!! Sorry!!!")
	}


}