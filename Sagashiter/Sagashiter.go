package Sagashiter

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

type SagashiterStruct struct {
	Content     string
	RegTemplate string
}

type SagashiterInterface interface {
	Tansaku() []string
	IncreaseTime([]string, time.Duration) string
}

func NewAssObj(str string) SagashiterInterface {
	return &SagashiterStruct{Content: str, RegTemplate: "\\d:\\d+:\\d+.\\d+"}
}

func NewSrtObj(str string) SagashiterInterface {
	return &SagashiterStruct{Content: str, RegTemplate: "none"}
}

func (obj *SagashiterStruct) Tansaku() []string {
	r, _ := regexp.Compile(obj.RegTemplate)
	timers := r.FindAllString(obj.Content, -1)
	return timers
}

func (obj *SagashiterStruct) IncreaseTime(timers []string, inc time.Duration) string {
	for _, item := range timers {
		layout := "15:04:05.00"
		newTime, err := time.Parse(layout, item)
		if err != nil {
			fmt.Println(err)
		}
		newTime = newTime.Add(inc)
		timeToReplace := newTime.Format(layout)

		obj.Content = strings.Replace(obj.Content, item, timeToReplace, -1)
	}
	return obj.Content
}
