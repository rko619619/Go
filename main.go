package main

import (
	"os"
	"github.com/sirupsen/logrus"
	"github.com/kelseyhightower/envconfig"
)
	

type Celsius float32
type Fahrenheit float32
type EnvConfig struct {   
	LogFile string `envconfig:"LOG_FILE"`
}

func main()  {
	var eConf EnvConfig
	envconfig.Process("", &eConf)
	eConf.LogFile = "file.log" 
	file, _ := os.OpenFile(eConf.LogFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)

	var log = logrus.New() 
	log.Out = file        

	c := Celsius(32)
	logrus.Info("Температура по фаренгейту - ", toFahrenheit(c))
	
}

func toFahrenheit(t Celsius) Fahrenheit {
	var temp Fahrenheit

	temp = Fahrenheit((t * 9 / 5) + 32)

	return temp
}