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
    fmt.Println(err)
  }

	session.Token = loadToken()

  session.AddHandler(onMessageCreate)

  err = session.Open()
  defer session.Close()
  if err != nil {
    fmt.Println(err)
  }

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
