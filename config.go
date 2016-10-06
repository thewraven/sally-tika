package tika

import (
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

//ProcessFile gets the text plain content of a file
func (conf Configuration) ProcessFile(file *os.File) (string, error) {
	defer file.Close()
	cmdText := exec.Command(conf.Java, "-jar", conf.TikaApp, tikaText, file.Name())
	text, err := cmdText.Output()
	if err != nil {
		return "", err
	}
	trimmedText := strings.Replace(string(text), "\n", "", -1)
	return trimmedText, nil
}
