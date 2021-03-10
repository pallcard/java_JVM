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
	cpOption         string   //用户类路径
	XjreOption       string   // 指定jre目录
	class            string   //类
	args             []string //参数
}

func parseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = printUsag //parse解析失败时调用
	// 绑定一个flag到变量上
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.BoolVar(&cmd.verboseClassFlag, "verbose", false, "enable verbose output")
	flag.BoolVar(&cmd.verboseClassFlag, "verbose:class", false, "enable verbose output")
	flag.BoolVar(&cmd.verboseInstFlag, "verbose:inst", false, "enable verbose output")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")
	flag.Parse()        // 把用户传递的命令行参数解析为对应变量的值
	args := flag.Args() //函数返回没有被解析的命令行参数
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}

func printUsag() {
	fmt.Printf("Usage: %s [-option] class [args ...]\n", os.Args[0])
}