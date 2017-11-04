package main

import "DbOperation"
import "demo"

func main(){
	demo.BuildDB()
	DbOperation.BuildConnection()
}



