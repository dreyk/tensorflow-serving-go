package main

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"os"
	"fmt"
	"bufio"
	"flag"
	"os/exec"
)

var baseDir string
var basePackage string
var workdir string
func main() {
	flag.StringVar(&baseDir,"i","pkg","Import dir")
	flag.StringVar(&basePackage,"p","","Package dir")
	flag.Parse()
	workdir,_ = os.Getwd()
	scan_and_gen(baseDir)
}

func scan_and_gen(dir string){
	files,err := ioutil.ReadDir(dir)
	if err!=nil{
		panic(err)
	}
	doGen := false
	for _,f := range files{
		if f.IsDir(){
			scan_and_gen(filepath.Join(dir,f.Name()))
		}
		if strings.HasSuffix(f.Name(),".proto"){
			doGen = true
			fixProtoFile(dir,f.Name())
		}
		if strings.HasSuffix(f.Name(),".import") || strings.HasSuffix(f.Name(),".go"){
			if f.Name()!="gen.go"{
				os.Remove(filepath.Join(dir,f.Name()))
			}
		}
	}
	if doGen{
		//filesDir := strings.TrimPrefix(strings.TrimPrefix(dir,baseDir),"/")
		fmt.Printf("protoc --proto_path=%s %s/*.proto --go_out=plugins=grpc:.\n",baseDir,dir)
		args := []string{fmt.Sprintf("--proto_path=%s", baseDir)}
		for _,f := range files {
			if strings.HasSuffix(f.Name(),".proto") {
				args = append(args,filepath.Join(dir,f.Name()))
			}
		}
		args = append(args,"--go_out=plugins=grpc:.")
		cmd := exec.Command("protoc",
			args...)
		cmd.Dir = workdir
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		err = cmd.Run()
		if err != nil {
			//fmt.Println(string((out)))
			panic(err)
		}
	}
}

func fixProtoFile(dir,file string){
	absFile := filepath.Join(dir,file)
	protoFile,err := os.Open(absFile)
	if err!=nil{
		panic(err)
	}
	scanner := bufio.NewScanner(protoFile)
	lines := []string{}
	for scanner.Scan(){
		line := scanner.Text()
		if !strings.HasPrefix(line,"option go_package"){
			lines = append(lines,line)
		}
	}
	protoFile.Close()
	protoFile,err = os.Create(absFile)
	if err!=nil{
		panic(err)
	}
	defer protoFile.Close()
	for _,l := range lines{
		protoFile.WriteString(l)
		protoFile.WriteString("\n")
	}
	p := strings.TrimPrefix(strings.TrimPrefix(dir,baseDir),"/")
	fmt.Fprintf(protoFile,`option go_package = "%s/%s";`,basePackage,p)
}