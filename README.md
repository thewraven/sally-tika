## Readme


This is a bare-bone API to Apache Zika for Go, a fork of [FunnyMonkey](https://github.com/FunnyMonkey/sally-tika) effort.  

### Setup

Download apache tika's jar file 'tika-app-1.5.jar' http://tika.apache.org/download.html you will also need a working java runtime environment.

Please set up your the location of your java installation and tika.jar ubication on the cmd/main.go file.
Example:

```go
conf := tika.Configuration{
		Java:    "/usr/bin/java",
		TikaApp: "/home/juan/tika.jar",
	}
```