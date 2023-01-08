package main 

import (
	"github.com/sirupsen/logrus"  //path to a package 
 	"github.com/fatih/color"
)	

func main(){
	logrus.Println("hello1")
	color.Cyan("this is Cyan color")
	color.Magenta("this is Magenta color")
}
