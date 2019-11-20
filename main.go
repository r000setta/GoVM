package main

import (
	"GoVM/classfile"
	"GoVM/classpath"
	"GoVM/cmd"
	"fmt"
	"strings"
)

func main() {
	c := cmd.ParseCmd()
	if c.Version {
		fmt.Println("Version 0.0.1")
	} else if c.Help || c.Class == "" {
		cmd.PrintUsage()
	} else {
		startJVM(c)
	}
}

func startJVM(c *cmd.Cmd) {
	cp := classpath.Parse(c.XjreOption, c.CpOption)
	fmt.Printf("classPath:%s class:%s args:%v\n",
		c.CpOption, c.Class, c.Args)
	className := strings.Replace(c.Class, ".", "/", -1)
	cf := loadClass(className, cp)
	//classData,_,err:=cp.ReadClass(className)
	//	//if err != nil {
	//	//	fmt.Printf("Could not load main class %s\n",c.Class)
	//	//	return
	//	//}
	//	//fmt.Printf("class data:%v\n",classData)
	fmt.Println(c.Class)
	printClassInfo(cf)
}

func loadClass(className string, cp *classpath.ClassPath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return cf
}

func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("Version: %v.%v\n", cf.GetMajorVersion())
	fmt.Printf("Constants count: %v\n", len(cf.ConstantPool))
	fmt.Printf("Access flag: 0x%x\n", cf.AccessFlag)
	fmt.Printf("This class: %v\n", cf.ClassName())
	fmt.Printf("Super class: %v\n", cf.SuperClassName())
	fmt.Printf("Interfaces: %v\n", cf.InterfaceNames())
	fmt.Printf("Fields count: %v\n", len(cf.Fields))
	for _, f := range cf.Fields {
		fmt.Printf("  %s\n", f.Name())
	}
	fmt.Printf("Methods count: %v\n", len(cf.Methods))
	for _, m := range cf.Methods {
		fmt.Printf("  %s\n", m.Name())
	}
}
