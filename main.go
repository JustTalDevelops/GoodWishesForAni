
package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/corpix/uarand"
	"math/rand"
	"time"
)

// ALLAH GANG FOREVER
func main() {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	lines, err := ReadLines("sentences.txt")
	if err != nil {
		fmt.Println("sentences.txt does not exist!")
		return
	}

	token, err := ReadLine("token.txt")
	if err != nil {
		fmt.Println("token.txt does not exist!")
		return
	}

	dg, err := discordgo.New(token)
	dg.UserAgent = uarand.GetRandom()
	if err != nil {
		fmt.Println("error creating Discord session: ", err)
		return
	}

	fmt.Println("time to spam ani")

	channel, err := dg.UserChannelCreate("272412535410393098")
	if err != nil {
		panic(err)
	}

	for {
		dg.ChannelMessageSend(channel.ID, lines[seededRand.Intn(len(lines))])
		time.Sleep(1 * time.Hour)
	}
}