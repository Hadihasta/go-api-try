package main

import "log"
func main(){
	err := godotenv.Load(".env")
	if err != nil{
		log.Fatal.(err)
	}
}