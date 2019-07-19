package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"subChanger/Sagashiter"
	"time"
)

const (
	INPUT = "./input"
	INC   = 10 * time.Second
)

func main() {
	if _, err := os.Stat(INPUT); os.IsNotExist(err) {
		os.Mkdir(INPUT, os.ModePerm)
	}

	inputList, err := ioutil.ReadDir(INPUT)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("list", inputList)
	for _, f := range inputList {
		fb, _ := ioutil.ReadFile(INPUT + "/" + f.Name())
		content := string(fb)

		var obj Sagashiter.SagashiterInterface
		if strings.HasSuffix(f.Name(), ".ass") {
			obj = Sagashiter.NewAssObj(content, f.Name())
		} else if strings.HasSuffix(f.Name(), ".srt") {
			//TODO
			obj = Sagashiter.NewSrtObj(content, f.Name())
		}

		timers := obj.Tansaku()
		res := obj.IncreaseTime(timers, INC)
		obj.Save()
		//TODO change const on flags
		fmt.Println(res)

	}

}
