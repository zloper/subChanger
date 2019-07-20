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

const (
	INC = 10
)

var config struct {
	InputDir string `long:"input"     description:"path to dir with original subs" env:"INPUT_DIR" default:"./"`
	Increase int    `long:"increase"     description:"time to increase subs timing" env:"INCREASE" default:"1"`
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

	fmt.Println("list", inputList)
	for _, f := range inputList {
		fb, _ := ioutil.ReadFile(config.InputDir + "/" + f.Name())
		content := string(fb)

		var obj Sagashiter.SagashiterInterface
		if strings.HasSuffix(f.Name(), ".ass") {
			obj = Sagashiter.NewAssObj(content, f.Name())
		} else if strings.HasSuffix(f.Name(), ".srt") {
			//TODO
			obj = Sagashiter.NewSrtObj(content, f.Name())
		}
		//TODO incTime := config.Increase * time.Second
		incTime := INC * time.Second
		timers := obj.Tansaku()
		res := obj.IncreaseTime(timers, incTime)
		obj.Save()
		//TODO change const on flags
		fmt.Println(res)

	}

}
