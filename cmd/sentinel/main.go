package main

import (
	"flag"
	"fmt"
	"github.com/jsfaint/sentinel"
	"github.com/jsfaint/sentinel/config"
	log "gopkg.in/clog.v1"
	"path/filepath"
)

const (
	name        = "sentinel"
	description = "A sentinel for WanKeYun"
	version     = "v0.1"
	mail        = "jsfaint@gmail.com"
)

type options struct {
	summary *bool
	check   *bool
	draw    *bool
	log     *string
}

var (
	opt options
	cfg config.Config
)

func init() {
	fmt.Println(name, version, mail)
	fmt.Println(description)

	opt.summary = flag.Bool("summary", false, "Show summary")
	opt.check = flag.Bool("check", false, "Check device status")
	opt.draw = flag.Bool("draw", false, "Draw the incoming coins")
	opt.log = flag.String("log", "", "Logging output with different log level(trace, info, warn, error, fatal)")

	flag.Parse()
}

func newLog(level string) (err error) {
	var l log.LEVEL

	switch level {
	case "trace":
		l = log.TRACE
	case "info":
		l = log.INFO
	case "warn":
		l = log.WARN
	case "error":
		l = log.ERROR
	case "fatal":
		l = log.FATAL
	default:
		l = log.WARN
	}

	name := filepath.Join(".", name+".log")

	err = log.New(log.CONSOLE, log.ConsoleConfig{
		Level:      log.TRACE,
		BufferSize: 0,
	})

	//If the log levl is nil, we don't log it into file
	if level == "" {
		return
	}

	err = log.New(log.FILE, log.FileConfig{
		Level:      l,
		BufferSize: 0,
		Filename:   name,
		FileRotationConfig: log.FileRotationConfig{
			Rotate:  true,
			Daily:   true,
			MaxSize: 10240000,
			MaxDays: 7,
		},
	})

	return err
}

func main() {
	if err := newLog(*opt.log); err != nil {
		fmt.Println(err)
	}
	defer log.Shutdown()

	//Get config from config.json
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(0, "%v", err)
	}

	if err := cfg.Check(); err != nil {
		log.Fatal(0, "%v", err)
	}

	sentinel.SetToken(cfg.Token)

	//Walk through the configs, support multiple account
	var users []*sentinel.UserReq
	for _, u := range cfg.Accounts {
		//skip if phone or pwd is null
		if u.Phone == "" || u.Pass == "" {
			continue
		}

		user := sentinel.NewUser(u.Phone, u.Pass)

		if err := user.Login(); err != nil {
			log.Error(0, "Login fail %v", err)
			continue
		}

		users = append(users, user)
	}

	//If some parameter were given, run the specfied task else run cron job
	if *opt.summary || *opt.check || *opt.draw {
		//Refresh all data
		refresh(users)

		if *opt.summary {
			summary(users)
		}

		if *opt.check {
			checkStatus(users)
		}

		if *opt.draw {
			withDraw(users)
		}

		return
	}

	if err := startCrontable(users); err != nil {
		return
	}
	defer stopCrontable()

	select {}
}
