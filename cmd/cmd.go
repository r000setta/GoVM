package cmd

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	XjreOption string
	Help     bool
	Version  bool
	CpOption string
	Class    string
	Args     []string
}

func ParseCmd() *Cmd {
	cmd:=&Cmd{}
	flag.Usage=PrintUsage
	flag.BoolVar(&cmd.Help,"help",false,"print help")
	flag.BoolVar(&cmd.Help,"?",false,"print help")
	flag.BoolVar(&cmd.Version,"version",false,"print version")
	flag.StringVar(&cmd.CpOption,"classpath","","classpath")
	flag.StringVar(&cmd.CpOption,"cp","","classpath")
	flag.StringVar(&cmd.XjreOption,"Xjre","","path to jre")
	flag.Parse()
	args:=flag.Args()
	if len(args)>0{
		cmd.Class =args[0]
		cmd.Args =args[1:]
	}
	return cmd
}

func PrintUsage() {
	fmt.Printf("Usage:%s [-option] class [args...]\n",os.Args[0])
}