package main

import (
	"bufio"
	"fmt"
	"os"
	"text/template/parse"

	"github.com/ShunyaNagashige/phone-book/database"
)

func main(){
	
}

func input(){
	scanner:=bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for{
		fmt.Print(">")
		for scanner.Scan(){
			switch(scanner.Text()){
			case "insert":
			}
		}
	}
}

func insert(){
	u:=&database.User{
		Id:
	}
}