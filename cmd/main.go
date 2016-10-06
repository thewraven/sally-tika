package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/thewraven/sally-tika"
)

var (
	filename = flag.String("name", "", "The name of the file to be processed")
)

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
	data, err := conf.ProcessFile(f)
	if err != nil {
		fmt.Println("error running tika ", err)
		return
	}
	fmt.Println(data)
}

func defaultConfig() tika.Configuration {
	conf := tika.Configuration{
		Java:    "/usr/bin/java",
		TikaApp: "/home/juan/tika.jar",
	}
	return conf
}
