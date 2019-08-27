package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Operation struct {
	Operand1 string
	Operator string
	Operand2 string
}

func calculate(operand1, operator, operand2 string) string {

	n1, _ := strconv.Atoi(strings.Split(operand1, "\r\n")[0])
	n2, _ := strconv.Atoi(strings.Split(operand2, "\r\n")[0])

	fmt.Println(n1 + n2)
	var res int
	if operator == "+\r\n" {
		res = n1 + n2
	} else if operator == "-\r\n" {
		res = n1 - n2
	} else if operator == "*\r\n" {
		res = n1 * n2
	} else if operator == "/\r\n" {
		res = n1 / n2
	} else {
		fmt.Println("Error in input")
		return ""
	}

	return strconv.Itoa(res)

}

func prepareRespond(w http.ResponseWriter, r *http.Request) {

	//request, _ := ioutil.ReadAll(r.Body)
	////request := r.Body
	//fmt.Println(string(request))

	decoder := json.NewDecoder(r.Body)
	var t Operation
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	log.Println(t.Operand1 + t.Operator + t.Operand2)
	_, _ = w.Write([]byte(calculate(t.Operand1, t.Operator, t.Operand2)))

}
func main() {
	http.HandleFunc("/", prepareRespond)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error Occurred!")
	}
}
