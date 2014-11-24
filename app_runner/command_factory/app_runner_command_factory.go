package command_factory

import (
	"fmt"
	"io"
	"strings"

	"github.com/codegangsta/cli"
)

type appRunner interface {
	StartDockerApp(name, startCommand, dockerImagePath string) error
	ScaleDockerApp(name string, instances int) error
	StopDockerApp(name string) error
}

type AppRunnerCommandFactory struct {
	appRunner appRunner
	output    io.Writer
}

func NewAppRunnerCommandFactory(appRunner appRunner, output io.Writer) *AppRunnerCommandFactory {
	return &AppRunnerCommandFactory{appRunner, output}
}

func (commandFactory *AppRunnerCommandFactory) MakeStartDiegoAppCommand() cli.Command {

	var startFlags = []cli.Flag{
		cli.StringFlag{
			Name:  "docker-image, i",
			Usage: "the docker image to run",
		},
		cli.StringFlag{
			Name:  "start-command, c",
			Usage: "the command to run in the context of the docker image (ie the start command for the app)",
		},
	}

	var startCommand = cli.Command{
		Name:        "start",
		ShortName:   "s",
		Description: "Start a docker app on diego",
		Usage:       "diego-edge-cli start APP_NAME -i DOCKER_IMAGE -c START_COMMAND",
		Action:      commandFactory.startDiegoApp,
		Flags:       startFlags,
	}

	return startCommand
}

func (commandFactory *AppRunnerCommandFactory) MakeScaleDiegoAppCommand() cli.Command {

	var scaleFlags = []cli.Flag{
		cli.IntFlag{
			Name:  "instances, i",
			Usage: "the number of instances to scale to",
		},
	}

	var scaleCommand = cli.Command{
		Name:        "scale",
		Description: "Scale a docker app on diego",
		Usage:       "diego-edge-cli scale APP_NAME --instances NUM_INSTANCES ",
		Action:      commandFactory.scaleDiegoApp,
		Flags:       scaleFlags,
	}

	return scaleCommand
}

func (commandFactory *AppRunnerCommandFactory) MakeStopDiegoAppCommand() cli.Command {

	var scaleCommand = cli.Command{
		Name:        "stop",
		Description: "Stop a docker app on diego",
		Usage:       "diego-edge-cli stop APP_NAME",
		Action:      commandFactory.stopDiegoApp,
	}

	return scaleCommand
}

func (commandFactory *AppRunnerCommandFactory) startDiegoApp(c *cli.Context) {
	startCommand := c.String("start-command")
	dockerImage := c.String("docker-image")
	name := c.Args().First()

	if name == "" || dockerImage == "" || startCommand == "" {
		commandFactory.incorrectUsage()
		return
	} else if !strings.HasPrefix(dockerImage, "docker:///") {
		commandFactory.incorrectUsage()
		commandFactory.say("Docker Image should begin with: docker:///")
		return
	}

	err := commandFactory.appRunner.StartDockerApp(name, startCommand, dockerImage)

	if err != nil {
		commandFactory.say(fmt.Sprintf("Error Starting App: %s", err))
		return
	}

	commandFactory.say("App Staged Successfully")
}

func (commandFactory *AppRunnerCommandFactory) scaleDiegoApp(c *cli.Context) {
	instances := c.Int("instances")
	appName := c.Args().First()
	if appName == "" {
		commandFactory.incorrectUsage()
		return
	} else if instances == 0 {
		commandFactory.say(fmt.Sprintf("Error Scaling to 0 instances - Please stop with: diego-edge-cli stop cool-web-app"))
		return
	}

	err := commandFactory.appRunner.ScaleDockerApp(appName, instances)

	if err != nil {
		commandFactory.say(fmt.Sprintf("Error Scaling App: %s", err))
		return
	}

	commandFactory.say("App Scaled Successfully")
}

func (commandFactory *AppRunnerCommandFactory) stopDiegoApp(c *cli.Context) {
	appName := c.Args().First()
	if appName == "" {
		commandFactory.incorrectUsage()
		return
	}

	err := commandFactory.appRunner.StopDockerApp(appName)

	if err != nil {
		commandFactory.say(fmt.Sprintf("Error Stopping App: %s", err))
		return
	}

	commandFactory.say("App Stopped Successfully")
}

func (commandFactory *AppRunnerCommandFactory) incorrectUsage() {
	commandFactory.say("Incorrect Usage\n")
}

func (commandFactory *AppRunnerCommandFactory) say(output string) {
	commandFactory.output.Write([]byte(output))
}