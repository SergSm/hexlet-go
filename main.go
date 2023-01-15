package main

import (
	"github.com/fatih/color"
	"github.com/sirupsen/logrus" //path to a package
)

func main() {
	logrus.Println("hello1")
	color.Cyan("this is Cyan color")
	color.Magenta("this is Magenta color")
	color.Magenta("haha ")
}
