package classpath

import (
	"os"
	"path/filepath"
)

type ClassPath struct {
	BootClassPath Entry
	ExtClassPath Entry
	UserClassPath Entry
}

func Parse(jreOption, cpOption string) *ClassPath {
	cp:=&ClassPath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

func (self *ClassPath) parseUserClasspath(cpOption string) {
	if cpOption=="" {
		cpOption="."
	}
	self.UserClassPath=newEntry(cpOption)
}

func (self *ClassPath) ReadClass(className string) ([]byte,Entry,error) {
	className=className+".class"
	if data, entry, err := self.BootClassPath.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := self.ExtClassPath.readClass(className);err==nil {
		return data, entry, err
	}
	return self.UserClassPath.readClass(className)
}

func (self *ClassPath) String() string {
	return self.UserClassPath.String()
}

func (self *ClassPath) parseBootAndExtClasspath(jreOption string) {
	jreDir:=getJreDir(jreOption)
	jreLibPath:=filepath.Join(jreDir,"lib","*")
	self.BootClassPath=newWildcardEntry(jreLibPath)

	jreExtPath:=filepath.Join(jreDir,"lib","ext","*")
	self.ExtClassPath=newWildcardEntry(jreExtPath)
}

func getJreDir(jreOption string) string {
	if jreOption!="" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh,"jre")
	}
	panic("can not find jre folder")
}

func exists(path string) bool {
	if _,err:=os.Stat(path);err!=nil{
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}