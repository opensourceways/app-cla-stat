package main

import (
	"context"
	"errors"
	"flag"
	"os"
	"time"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"

	"github.com/opensourceways/app-cla-stat/accesstoken/domain"
	commondb "github.com/opensourceways/app-cla-stat/common/infrastructure/mongodb"
	"github.com/opensourceways/app-cla-stat/config"
	"github.com/opensourceways/app-cla-stat/interrupts"
	_ "github.com/opensourceways/app-cla-stat/routers"
	"github.com/opensourceways/app-cla-stat/signing/domain/dp"
	"github.com/opensourceways/app-cla-stat/util"
)

type options struct {
	configFile string
}

func (o *options) Validate() error {
	if o.configFile == "" {
		return errors.New("missing config file")
	}

	return nil
}

func gatherOptions(fs *flag.FlagSet, args ...string) options {
	var o options

	fs.StringVar(
		&o.configFile, "config-file", "", "config file path.",
	)

	if err := fs.Parse(args); err != nil {
		logs.Error(err)
	}

	return o
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	o := gatherOptions(
		flag.NewFlagSet(os.Args[0], flag.ExitOnError),
		os.Args[1:]...,
	)
	if err := o.Validate(); err != nil {
		logs.Error("Invalid options, err:%s", err.Error())

		return
	}

	cfg := loadConfig(o.configFile)
	if cfg == nil {
		return
	}

	startSignSerivce(cfg)
}

func loadConfig(f string) *config.Config {
	cfg, err := config.Load(f)
	err1 := os.Remove(f)

	if err2 := util.MultiErrors(err, err1); err2 != nil {
		logs.Error(err2)

		return nil
	}

	return &cfg
}

func startSignSerivce(cfg *config.Config) {
	dp.Init(&cfg.Domain.DomainPrimitive)
	domain.Init(&cfg.Domain.AccessToken)

	if err := commondb.Init(&cfg.Mongodb.DB); err != nil {
		logs.Error(err)
		return
	}

	defer exitMongoService()

	// must run after init mongodb
	if err := initSigning(cfg); err != nil {
		logs.Error(err)

		return
	}

	run()
}

func exitMongoService() {
	if err := commondb.Close(); err != nil {
		logs.Error(err)
	}
}

func run() {
	defer interrupts.WaitForGracefulShutdown()

	interrupts.OnInterrupt(func() {
		shutdown()
	})

	beego.Run()
}

func shutdown() {
	logs.Info("server shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	if err := beego.BeeApp.Server.Shutdown(ctx); err != nil {
		logs.Error("error to shut down server, err:", err.Error())
	}
	cancel()
}
