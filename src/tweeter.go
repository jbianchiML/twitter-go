package main

import (
	"fmt"

	"github.com/abiosoft/ishell"
	"github.com/jbianchiML/twitter-go/src/domain"
	"github.com/jbianchiML/twitter-go/src/service"
)

func main() {

	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Type your username: ")

			user := c.ReadLine()

			c.Print("Type your tweet: ")

			text := c.ReadLine()

			tweet := domain.NewTweet(user, text)

			_, err := service.PublishTweet(tweet)

			if err == nil {
				c.Print("Tweet sent\n")
			} else {
				c.Print("Error publishing tweet:", err)
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweet",
		Help: "Shows a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweet := service.GetTweets()

			//c.Println(tweet)

			fmt.Printf("%v\n", tweet)
			return
		},
	})

	shell.Run()

}
