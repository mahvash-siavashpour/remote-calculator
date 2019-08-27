package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Operation struct {
	Operand1 string
	Operator string
	Operand2 string
}

func main() {

	url := "http://localhost:8080/"
	//get the operator and the operands from user
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter First Number:")
	num1, _ := reader.ReadString('\n')
	fmt.Println("Enter Operator:")
	op, _ := reader.ReadString('\n')
	fmt.Println("Enter Second Number:")
	num2, _ := reader.ReadString('\n')

	operation := Operation{num1, op, num2}
	js, err := json.Marshal(operation)
	if err != nil {
		fmt.Println("Error in json")
		return
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(js))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("response Body:", string(body))

}
