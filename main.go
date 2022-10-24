package main

import (
	"fmt"

	"github.com/ameyalambat128/discord-pingbot/bot"
	"github.com/ameyalambat128/discord-pingbot/config"
)

func main(){
	err := config.ReadConfig()

	if err != nil{
		fmt.Println(err.Error())
		return
	} 
	
	bot.Start()
    <-make(chan struct{})
	return 
}
