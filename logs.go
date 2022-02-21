package main

import (
	"fmt"

	"github.com/fatih/color"
)

type Logger struct {

}

func (self Logger) error (print string, err ...error){
	msg := color.New(color.FgWhite, color.BgRed).SprintFunc()
	if err != nil{
		fmt.Println(msg(" ERROR "), print, err)
	} else {
		fmt.Println(msg(" ERROR "), print)
	}
}

func (self Logger) warning (print string, err ...error){
	msg := color.New(color.FgBlack, color.BgYellow).SprintFunc()
	if err != nil{
		fmt.Println(msg(" WARN  "), print, err)
	} else {
		fmt.Println(msg(" WARN  "), print)
	}
}

func (self Logger) info (print string){
	msg := color.New(color.FgBlack, color.BgBlue).SprintFunc()
	fmt.Println(msg(" INFO  "), print)
}

func (self Logger) hashrate (print string){
	msg := color.New(color.FgBlack, color.BgGreen).SprintFunc()
	fmt.Println(msg(" HASH  "), print)
}

func (self Logger) guessed (print string){
	msg := color.New(color.FgBlack, color.BgCyan).SprintFunc()
	fmt.Println(msg(" GUESS "), print)
}

func (self Logger) job (print []string){
	msg := color.New(color.FgBlack, color.BgMagenta).SprintFunc()
	if(len(print) > 1){
		fmt.Println(msg(" JOB   "), "pref:", print[0], "target:", print[1], "diff:", print[2])
	} else {
		fmt.Println(msg(" JOB   "), print[0])
	}
}

func (self Logger) server (print string){
	msg := color.New(color.FgBlack, color.BgHiBlue).SprintFunc()
	fmt.Println(msg(" SERV  "), print)
}
