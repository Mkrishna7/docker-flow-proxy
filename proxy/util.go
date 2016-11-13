package proxy

import (
	"io/ioutil"
	"log"
	"os/exec"
)

var cmdRunHa = func(cmd *exec.Cmd) error {
	return cmd.Run()
}
var readConfigsFile = ioutil.ReadFile
var writeFile = ioutil.WriteFile
var ReadFile = ioutil.ReadFile
var logPrintf = log.Printf
var readPidFile = ioutil.ReadFile
var readConfigsDir = ioutil.ReadDir