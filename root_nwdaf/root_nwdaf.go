package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/free5gc/version"
	"nwdaf.com/logger"
	"nwdaf.com/service"
)

var NWDAF = &service.NWDAF{}
var appLog *logrus.Entry

func init() {
	appLog = logger.AppLog
}

func main() {
	fmt.Println("====== [Root-NWDAF] Start Root-NWDAF ======")
	app := cli.NewApp()
	app.Name = "root-nwdaf"
	appLog.Infoln(app.Name)
	appLog.Infoln("Root NWDAF version: ", version.GetVersion())
	app.Usage = "-free5gccfg common configuration file -nwdafcfg nwdaf configuration file"
	app.Action = action
	app.Flags = NWDAF.GetCliCmd()
	if err := app.Run(os.Args); err != nil {
		appLog.Errorf("Root NWDAF error: %v", err)
		return
	}
	fmt.Println("====== [Root-NWDAF] End Root-NWDAF ======")
}

func action(c *cli.Context) error {
	if err := NWDAF.Initialize(c); err != nil {
		logger.CfgLog.Errorf("%+v", err)
		return fmt.Errorf("Failed to initialize !!")
	}
	NWDAF.Start()
	return nil
}
