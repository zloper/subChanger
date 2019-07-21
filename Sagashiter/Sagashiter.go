package Sagashiter

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

type SagashiterStruct struct {
	Content     string
	RegTemplate string
	Name        string
	IsAss       bool
}

type SagashiterInterface interface {
	Tansaku() []string
	IncreaseTime([]string, time.Duration) string
	Save()
}

func NewAssObj(content, name string) SagashiterInterface {
	return &SagashiterStruct{Content: content, Name: name, IsAss: true, RegTemplate: `(\d:\d+:\d+)\.\d+`}
}

func NewSrtObj(str, name string) SagashiterInterface {
	return &SagashiterStruct{Content: str, Name: name, RegTemplate: "\\d+:\\d+:\\d+", IsAss: false}
}

func (obj *SagashiterStruct) Tansaku() []string {
	r := regexp.MustCompile(obj.RegTemplate)
	timers := r.FindAllString(obj.Content, -1)
	return timers
}

func (obj *SagashiterStruct) IncreaseTime(timers []string, inc time.Duration) string {
	for _, item := range timers {
		layout := "15:04:05"
		newTime, err := time.Parse(layout, item)
		if err != nil {
			fmt.Println(err)
		}
		newTime = newTime.Add(inc)
		timeToReplace := newTime.Format(layout)

		if obj.IsAss {
			timeToReplace = timeToReplace[1:]
			item = strings.Split(item, ".")[0]
		}

		obj.Content = strings.Replace(obj.Content, item, timeToReplace, -1)
	}
	return obj.Content
}

func (obj *SagashiterStruct) Save() {
	err := os.MkdirAll("./output", os.ModePerm)
	if err != nil {
		panic(err)
	}

	f, err := os.Create("./output/" + obj.Name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString(obj.Content)
	if err != nil {
		panic(err)
	}
}
