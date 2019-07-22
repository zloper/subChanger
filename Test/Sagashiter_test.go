package Test

import (
	"io/ioutil"
	"os"
	"strings"
	"subChanger/Sagashiter"
	"testing"
	"time"
)

const assText = `Text
Dialogue: 10,0:00:08.77,0:00:13.87,Default,,0,0,0,,Well well? \NAnd you call this test?`

const srtText = `1
00:01:34,719 --> 00:01:37,847
Why read file? Why not string?
00:01:39,345 --> 00:01:42,211
Because I can!`

func TestShort(t *testing.T) {
	obj := Sagashiter.NewAssObj(assText, "someTestFile.ass")
	timers := obj.Tansaku()
	if !("0:00:08.77" == timers[0]) {
		t.Error("wrong time get")
	}
	afterChange := obj.IncreaseTime(timers, time.Duration(55)*time.Second)
	if !strings.Contains(afterChange, "01:03.77") {
		t.Error("Wrong time in final result")
	}
}

func TestAss(t *testing.T) {
	res, err := Reader("fl.ass", assText)
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(res, ":18.77,0:00:23.87,") {
		t.Error("Wrong result, time not add")
	}

}

func TestSrt(t *testing.T) {
	res, err := Reader("fl.srt", srtText)
	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(res, "01:49,345 --> 00:01:52,21") {
		t.Error("Wrong result, time not add", res)
	}

}

func Reader(flName, flContent string) (string, error) {
	err := ioutil.WriteFile(flName, []byte(flContent), os.ModePerm)
	if err != nil {
		return "", err
	}

	content, err := ioutil.ReadFile(flName)
	if err != nil {
		return "", err
	}

	strContent := string(content)

	var obj Sagashiter.SagashiterInterface
	if strings.HasSuffix(flName, ".ass") {
		obj = Sagashiter.NewAssObj(strContent, flName)
	} else {
		obj = Sagashiter.NewSrtObj(strContent, flName)
	}

	timers := obj.Tansaku()
	incTime := time.Duration(10) * time.Second
	res := obj.IncreaseTime(timers, incTime)
	return res, nil
}
