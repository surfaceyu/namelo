package main

import (
	"log"
	"net/http"
	"os"

	"github.com/surfaceyu/namelo"
	"github.com/surfaceyu/namelo/component"
	"github.com/surfaceyu/namelo/examples/demo/tadpole/logic"
	"github.com/surfaceyu/namelo/serialize/json"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "tadpole"
	app.Author = "nano authors"
	app.Version = "0.0.1"
	app.Copyright = "nano authors reserved"
	app.Usage = "tadpole"

	// flags
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "addr",
			Value: ":23456",
			Usage: "game server address",
		},
	}

	app.Action = serve

	app.Run(os.Args)
}

func serve(ctx *cli.Context) error {
	components := &component.Components{}
	components.Register(logic.NewManager())
	components.Register(logic.NewWorld())

	// register all service
	options := []nano.Option{
		nano.WithComponents(components),
		nano.WithSerializer(json.NewSerializer()),
		nano.WithCheckOriginFunc(func(_ *http.Request) bool { return true }),
	}

	//nano.EnableDebug()
	log.SetFlags(log.LstdFlags | log.Llongfile)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	addr := ctx.String("addr")
	nano.ListenWS(addr, options...)
	return nil
}
