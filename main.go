package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
)

// Configuration is the set of parameters required
// to run the tika standalone jar
type Configuration struct {
	TikaApp string
	Java    string
}

const tikaMeta = "--metadata"
const tikaHTML = "--html"
const tikaText = "--text"
const tikaExtract = "--extract"

var expCleaner = regexp.MustCompile(`\n([^\s])`)

func (conf Configuration) processFile(file *os.File) (string, error) {
	defer file.Close()
	cmdText := exec.Command(conf.Java, "-jar", conf.TikaApp, tikaText, file.Name())
	text, err := cmdText.Output()
	if err != nil {
		return "", err
	}
	result := expCleaner.ReplaceAllString(string(text), " $1")
	return result, nil
}

func main() {
	conf := defaultConfig()
	f, err := os.Open("data.pdf")
	if err != nil {
		fmt.Println("error reading file ", err)
	}
	data, err := conf.processFile(f)
	if err != nil {
		fmt.Println("error running tika ", err)
	}
	fmt.Println(data)
}

func defaultConfig() Configuration {
	conf := Configuration{
		Java:    "/usr/bin/java",
		TikaApp: "/home/juan/tika.jar",
	}
	return conf
}
