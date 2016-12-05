package main

import (
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "GoTweet"
	app.Usage = "glide sample project: simple tweet command"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "ck",
			Usage: "Consumer Key",
		},
		cli.StringFlag{
			Name:  "cs",
			Usage: "Consumer Secret",
		},
		cli.StringFlag{
			Name:  "at",
			Usage: "Access Token",
		},
		cli.StringFlag{
			Name:  "as",
			Usage: "Access Token Secret",
		},
		cli.StringFlag{
			Name:  "m",
			Usage: "teet message",
		},
	}

	app.Action = func(c *cli.Context) error {
		ck := c.String("ck")
		cs := c.String("cs")
		at := c.String("at")
		as := c.String("as")
		m := c.String("m")
		if len(ck) == 0 || len(cs) == 0 || len(at) == 0 || len(m) == 0 {
			return cli.NewExitError("argument error: -h, show help", -1)
		}

		anaconda.SetConsumerKey(ck)
		anaconda.SetConsumerSecret(cs)
		api := anaconda.NewTwitterApi(at, as)

		api.PostTweet(m, nil)
		return nil
	}

	app.Run(os.Args)
}
