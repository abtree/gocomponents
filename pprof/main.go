package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpu", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	env := os.Environ()
	procAttr := &os.ProcAttr{
		Env: env,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}
	pid, err := os.StartProcess("F:/Git/usr/bin/ls.exe", []string{"ls"}, procAttr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("The process id is %v", pid)
}
