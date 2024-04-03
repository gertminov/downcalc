package main

import (
	"errors"
	"fmt"

	"os"
	"strconv"
	"time"

	"github.com/cqroot/prompt"
	"github.com/docker/go-units"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

type byteSize int64

func CheckErr(err error) {
	if err != nil {
		if errors.Is(err, prompt.ErrUserQuit) {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		} else {
			panic(err)
		}
	}
}

func saveSpeed() bool {
	save, err := prompt.New().Ask("Choose:").
		Choose([]string{"Yes", "No"})

	CheckErr(err)
	return save == "Yes"
}

func askSize() string {
	sizeStr, err := prompt.New().Ask("File Size:").Input("100gb")
	CheckErr(err)
	return sizeStr
}

func parseFileSize(humanSize string) byteSize {
	size, err := units.FromHumanSize(humanSize)
	CheckErr(err)
	size = int64(float64(size) / 0.931323)
	return byteSize(size)
}

func calculateTime(speedMbits int, sizeBytes byteSize) int64 {
	inputBits := int64(sizeBytes * 8)
	speedBits := int64(speedMbits * 1000000)
	return inputBits / speedBits
}

func main() {
	app := &cli.App{
		Name:  "downcalc",
		Usage: "calculate download time of file size",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Cofigure downcalc",
			},
			&cli.IntFlag{
				Name:     "speed",
				Aliases:  []string{"s"},
				Usage:    "internet speed in MBit/s",
				FilePath: GetConfigFile(),
			},
		},
		Action: func(cCtx *cli.Context) error {
			if cCtx.Bool("config") {
				Config(cCtx)
			}

			speed := cCtx.Int("speed")
			if speed == 0 {
				fmt.Println("No internet speed provided, please input your internet speed")
				speed = askSpeed()
				if saveSpeed() {
					writeConfig(strconv.Itoa(speed))
				}
			}
			fileSize := cCtx.Args().Get(0)
			waitForKeyPress := false
			if fileSize == "" {
				fmt.Println("No file size given to calculate. Please enter file size like so: 300mb, 45gb")
				fileSize = askSize()
				if fileSize != "" {
					waitForKeyPress = true
				}
			}
			inputByteSize := parseFileSize(fileSize)
			durationSeconds := calculateTime(speed, inputByteSize)

			duration := time.Duration(durationSeconds) * time.Second
			hours := int(duration.Hours())
			minutes := int(duration.Minutes()) % 60
			seconds := int(duration.Seconds()) % 60

			fmt.Printf("Download time with %d MBit/s for %v:  ", speed, fileSize)
			color.Green("%dh:%dm:%ds\n", hours, minutes, seconds)

			if waitForKeyPress {
				yellow := color.New(color.FgYellow).SprintFunc()
				fmt.Printf("Press %s to exit", yellow("ENTER"))
				fmt.Scanln()
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
