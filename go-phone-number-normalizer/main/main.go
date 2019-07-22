package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	dat, err := ioutil.ReadFile("/home/glowatsk/go/src/github.com/glowatsk/go-projects/go-phone-number-normalizer/main/phonenumber.txt")
	if err != nil {
		panic(err)
	}
	r, _ := regexp.Compile("[0-9]")
	split := strings.Split(string(dat), "")
	var returnArray []string
	tempString := ""
	for i := 0; i < len(split); i++ {
		match := r.FindString(split[i])
		if len(match) > 0 {
			tempString += match
			if len(tempString) == 10 {
				returnArray = append(returnArray, tempString)
				tempString = ""
			}
		}
	}
	fmt.Println(returnArray)
}
