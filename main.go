package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	session, err := discordgo.New()
	if err != nil {
		fmt.Println("Error in create session")
		panic(err)
	}

	discordToken := loadToken()
	if discordToken == nil {
		panic("no discord token exists.")
	}
	session.Token = discordToken

	session.AddHandler(onMessageCreate)

	err = session.Open()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	fmt.Println("booted!!!")

	<-sc
	return
}

func onMessageCreate(session *discordgo.Session, event *discordgo.MessageCreate) {
	if event.Author.Bot {
		return
	}
	/* メッセージを受け取った際の処理 */
}

func loadToken() (token string) {
	token = os.Getenv("DISCORD_TOKEN")
	return
}
