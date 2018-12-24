package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var reg = regexp.MustCompile(`\st=(\d+)`)

const interval = time.Second * 2

func main() {
	path := getFilePath()
	go runCheckLoop(path)
}

func runCheckLoop(path string) {
	for true {

		file, e := ioutil.ReadFile(path)
		err(e)
		match := reg.FindSubmatch(file)
		if len(match) > 1 {
			tempString := (string(match[1][:2]) + "." + string(match[1][2:]))
			temp, e := strconv.ParseFloat(tempString, 32)
			err(e)
			temp = (temp * 1.8) + 32.0
			temp = math.Floor(temp*10) / 10
			fmt.Println(temp)
		} else {
			err(errors.New("hello"))
		}
		sleepInterval()
	}
}

func getFilePath() string {
	files, e := ioutil.ReadDir("/sys/bus/w1/devices/")
	err(e)
	fileName := ""
	for _, f := range files {
		if strings.HasPrefix(f.Name(), "28-") {
			fileName = f.Name()
			break
		}
	}
	return "/sys/bus/w1/devices/" + fileName + "/w1_slave"
}

func sleepInterval() {
	time.Sleep(interval)
}

func err(e error) {
	if e != nil {
		panic("Ohh my goodness!")
	}
}
