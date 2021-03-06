package setup_cli

import (
	"os"
	"os/signal"

	"github.com/codegangsta/cli"
	"github.com/cloudfoundry-incubator/lattice/cli/config"
	"github.com/cloudfoundry-incubator/lattice/cli/config/config_helpers"
	"github.com/cloudfoundry-incubator/lattice/cli/config/persister"
	"github.com/cloudfoundry-incubator/lattice/cli/exit_handler"
	"github.com/pivotal-golang/lager"

	"github.com/cloudfoundry-incubator/lattice/cli/cli_app_factory"
	"github.com/cloudfoundry-incubator/lattice/cli/config/target_verifier"
	"github.com/cloudfoundry-incubator/lattice/cli/config/target_verifier/receptor_client_factory"
	"github.com/cloudfoundry-incubator/lattice/cli/output"
)

const (
	latticeCliHomeVar = "LATTICE_CLI_HOME"
	timeoutVar        = "LATTICE_CLI_TIMEOUT"
)

func NewCliApp() *cli.App {
	config := config.New(persister.NewFilePersister(config_helpers.ConfigFileLocation(ltcConfigRoot())))

	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt)
	exitHandler := exit_handler.New(signalChan, os.Exit)
	go exitHandler.Run()

	targetVerifier := target_verifier.New(receptor_client_factory.MakeReceptorClient)
	app := cli_app_factory.MakeCliApp(os.Getenv(timeoutVar), ltcConfigRoot(), exitHandler, config, logger(), targetVerifier, output.New(os.Stdout))
	return app
}

func logger() lager.Logger {
	logger := lager.NewLogger("ltc")
	var logLevel lager.LogLevel

	if os.Getenv("LTC_LOG_LEVEL") == "DEBUG" {
		logLevel = lager.DEBUG
	} else {
		logLevel = lager.INFO
	}

	logger.RegisterSink(lager.NewWriterSink(os.Stderr, logLevel))
	return logger
}

func ltcConfigRoot() string {
	if os.Getenv(latticeCliHomeVar) != "" {
		return os.Getenv(latticeCliHomeVar)
	}

	return os.Getenv("HOME")
}
