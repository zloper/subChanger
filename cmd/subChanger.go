package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"io/ioutil"
	"os"
	"strings"
	"subChanger/Sagashiter"
	"time"
)

var config struct {
	InputDir string `long:"input"     description:"path to dir with original subs" env:"INPUT_DIR" default:"./input"`
	Increase int64  `long:"increase"     description:"time to increase subs timing" env:"INCREASE" default:"1"`
}

func main() {
	_, err := flags.Parse(&config)
	if err != nil {
		fmt.Println(err)
	}
	if _, err := os.Stat(config.InputDir); os.IsNotExist(err) {
		os.Mkdir(config.InputDir, os.ModePerm)
	}

	inputList, err := ioutil.ReadDir(config.InputDir)
	if err != nil {
		fmt.Println(err)
	}

	for _, f := range inputList {
		fb, _ := ioutil.ReadFile(config.InputDir + "/" + f.Name())
		content := string(fb)

		var obj Sagashiter.SagashiterInterface
		if strings.HasSuffix(f.Name(), ".ass") {
			obj = Sagashiter.NewAssObj(content, f.Name())
		} else if strings.HasSuffix(f.Name(), ".srt") {
			obj = Sagashiter.NewSrtObj(content, f.Name())
		}

		if obj != nil {
			incTime := time.Duration(config.Increase) * time.Second
			timers := obj.Tansaku()
			res := obj.IncreaseTime(timers, incTime)
			obj.Save()
			fmt.Println(res)
		}

	}

}
