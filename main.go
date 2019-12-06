package main

import (
	"bufio"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
)

// our main function
func main() {
	stringtoken := tokenRead()

	d, err := discordgo.New("Bot " + stringtoken)
	if err != nil {
		fmt.Println("failed to create discord session", err)
	}

	bot, err := d.User("@me")
	if err != nil {
		fmt.Println("failed to access account", err)
	}
	botID := bot.ID
	fmt.Println(botID)

	d.AddHandler(handleCmd)
	err = d.Open()
	if err != nil {
		fmt.Println("unable to establish connection", err)
	}

	defer d.Close()

	<-make(chan struct{})
}

//handle command
func handleCmd(d *discordgo.Session, msg *discordgo.MessageCreate) {
	user := msg.Author
	if user.Bot {
		return
	}

	content := msg.Content

	if content == "!test" {
		d.ChannelMessageSend(msg.ChannelID, "Testing..")
	}

	fmt.Printf("Message: %+v\n", msg.Message)
}

func tokenRead() string {
	//this fn requires user to have token stored in a token file
	file, err := os.Open("token")

	if err != nil {
		fmt.Println(err)
		return "error"
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	read := scanner.Scan()

	if read {
		fmt.Println("read byte array: ", scanner.Bytes())
		fmt.Println("read string: ", scanner.Text())
	}
	stringtoken := scanner.Text()
	return stringtoken
}

//	d, err := discordgo.New("Bot <your token here>")

//}

// our command handler function
//func handleCmd() {
//}
//}
