package main

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	helpFlag         bool // 帮助信息
	versionFlag      bool //版本信息
	verboseClassFlag bool
	verboseInstFlag  bool
	cpOption         string //用户类路径
	XjreOption       string // 指定jre目录
	class            string
	args             []string
}

func parseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = printUsag
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}

func printUsag() {
	fmt.Printf("Usage: %s [-option] class [args ...]\n", os.Args[0])
}
