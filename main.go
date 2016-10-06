package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
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

func (conf Configuration) processFile(file *os.File) (string, error) {
	defer file.Close()
	cmdText := exec.Command(conf.Java, "-jar", conf.TikaApp, tikaText, file.Name())
	text, err := cmdText.Output()
	if err != nil {
		return "", err
	}
	trimmedText := strings.Replace(string(text), "\n", "", -1)
	return trimmedText, nil
}

var filename = flag.String("name", "", "The name of the file to be processed")

func main() {
	flag.Parse()
	if *filename == "" {
		fmt.Println("-name flag required")
	}
	conf := defaultConfig()
	f, err := os.Open(*filename)
	if err != nil {
		fmt.Println("error reading file ", err)
		return
	}
	data, err := conf.processFile(f)
	if err != nil {
		fmt.Println("error running tika ", err)
		return
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
