package main

import (
	"GoVM/classpath"
	"GoVM/cmd"
	"fmt"
	"strings"
)

func main() {
	c:=cmd.ParseCmd()
	if c.Version {
		fmt.Println("Version 0.0.1")
	}else if c.Help || c.Class == "" {
		cmd.PrintUsage()
	}else {
		startJVM(c)
	}
}

func startJVM(c *cmd.Cmd)  {
	cp:=classpath.Parse(c.XjreOption,c.CpOption)
	fmt.Printf("classPath:%s class:%s args:%v\n",
		c.CpOption,c.Class,c.Args)
	className:=strings.Replace(c.Class,".","/",-1)
	classData,_,err:=cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not load main class %s\n",c.Class)
		return
	}
	fmt.Printf("class data:%v\n",classData)
}