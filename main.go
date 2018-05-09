package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

// build number set at compile-time
var build = "0"

// Version set at compile-time
var Version string

func main() {

	// Load env-file if it exists first
	// if filename, found := os.LookupEnv("PLUGIN_ENV_FILE"); found {
	// 	_ = godotenv.Load(filename)
	// }

	app := cli.NewApp()
	app.Name = "Drone JSON2File"
	app.Usage = "Convert a json string to specify type file"
	app.Copyright = "Copyright (c) 2018 Yu-Shun Wu"
	app.Authors = []cli.Author{
		{
			Name:  "Yu-Shun Wu",
			Email: "wys1203@gmail.com",
		},
	}
	app.Action = run
	app.Version = Version

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "input",
			Usage:  "A base64 string of input",
			EnvVar: "PLUGIN_INPUT,INPUT",
		},
		cli.StringFlag{
			Name:   "output.path",
			Usage:  "Specify the output path",
			EnvVar: "PLUGIN_PATH,PATH",
			Value:  ".",
		},
		cli.BoolFlag{
			Name:   "debug",
			Usage:  "debug mode",
			EnvVar: "PLUGIN_DEBUG",
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println("json2file error: ", err)
		os.Exit(1)
	}
}

func run(c *cli.Context) error {
	plugin := Plugin{
		Config: Config{
			Input: c.String("input"),
			Path:  c.String("output.path"),
			Debug: c.Bool("debug"),
		},
		Writer: os.Stdout,
	}
	return plugin.Exec()
}
