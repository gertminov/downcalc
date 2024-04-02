package main

import (
	"fmt"
	"github.com/cqroot/prompt"
	"github.com/cqroot/prompt/input"
	"github.com/urfave/cli/v2"
	"os"
	"path/filepath"
	"strconv"
)

func writeConfig(speed string) {
	home, err := os.UserHomeDir()
	CheckErr(err)

	configDir := filepath.Join(home, ".config", "downcalc")
	err = os.MkdirAll(configDir, 0777)
	CheckErr(err)

	configFile := filepath.Join(configDir, "downloadspeed.txt")
	err = os.WriteFile(configFile, []byte(speed), 0777)
	CheckErr(err)

}

func validateSpeed(speed string) error {
	if _, err := strconv.Atoi(speed); err != nil {
		return fmt.Errorf("Not valid speed")
	} else {
		return nil
	}
}

func askSpeed() int {
	speed, err := prompt.New().Ask("Speed in MBit/s:").Input("100", input.WithValidateFunc(validateSpeed))
	CheckErr(err)
	speedInt, err2 := strconv.Atoi(speed)
	CheckErr(err2)
	return speedInt
}

func Config(cCtx *cli.Context) {
	var speed int
	var err error
	if cCtx.Args().Len() > 0 {
		maybeSpeed := cCtx.Args().Get(0)
		speed, err = strconv.Atoi(maybeSpeed)
		CheckErr(err)
	} else {
		speed = askSpeed()
	}

	writeConfig(strconv.Itoa(speed))
}

func getConfigDir() string {
	homedir, err := os.UserHomeDir()
	CheckErr(err)
	return filepath.Join(homedir, ".config", "downcalc")
}

func GetConfigFile() string {
	return filepath.Join(getConfigDir(), "downloadspeed.txt")
}
