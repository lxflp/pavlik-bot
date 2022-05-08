package main

import (
	"fmt"
	volume "github.com/itchyny/volume-go"
	tele "gopkg.in/telebot.v3"
	"log"
	"strconv"
	"time"
)

func main() {
	pref := tele.Settings{
		Token:  "",
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello!")
	})
	// Command: /start <PAYLOAD>
	b.Handle("/volume", func(c tele.Context) error {
		fmt.Println(c.Message().Payload) // <PAYLOAD>
		volume1, err := strconv.Atoi(c.Message().Payload)
		if err != nil {
			return err
		}
		err = volume.SetVolume(volume1)
		c.Send(fmt.Sprintf("Громкость изменена на %v", volume1))

		return err
	})
	fmt.Println("bot is run")
	b.Start()
}
