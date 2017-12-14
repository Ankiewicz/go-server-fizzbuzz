package main

import (
	"net/http"
	"io"
	"strconv"
	"strings"
	)

func fizzBuzz(w http.ResponseWriter, r *http.Request) {
  x := 100

  for i := 0; i <= x; i++ {
     if i % 3 == 0 && i % 5 == 0{
        fizzbuzz := strings.Join([]string{"FIZZBUZZ -- Number: ", strconv.Itoa(i),  " is a mulitple of 3 & 5\n"}, "")
        io.WriteString(w, fizzbuzz)
    } else if i % 5 == 0{
      	buzz :=  strings.Join([]string{"BUZZ -- Number: ", strconv.Itoa(i), " is a mulitple of 5\n"}, "")
	io.WriteString(w, buzz)
    } else if i % 3 == 0{
     	fizz := strings.Join([]string{"FIZZ -- Number: ",  strconv.Itoa(i),  " is a mulitple of 3\n"}, "")
 	io.WriteString(w, fizz)
    }
  }
}

func main() {
	http.HandleFunc("/", fizzBuzz)
	http.ListenAndServe(":80", nil)
}


