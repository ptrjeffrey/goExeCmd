package main

import (
	"flag"
	"fmt"
	"os/exec"
)

func main() {
	// 定义一个名为 "file" 的命令行参数，类型为字符串，默认值为 ""，并附加了说明信息
	filePtr := flag.String("file", "", "fileName")
	// 解析命令行参数
	flag.Parse()
	//
	fmt.Println("使用说明: cp2server.exe --file=<要替换的文件名>")
	if *filePtr == "" {
		fmt.Println("没有文件")
		return
	}
	// 输出解析后的参数值
	fmt.Println("文件名称:", *filePtr)

	dist_rlmt := "k3d-k3s-v-local-reg-server-0:/mnt/wsl/twinverse-edge-pv/rlmt-agent-sys-1g"

	// str_dt := "k3d-k3s-volumn-local-reg-1-server-0:/data/twinverse-edge-pv/rlmt-mock-data-1g/config" //
	cmd := exec.Command("docker", "cp", *filePtr, dist_rlmt)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(stdout))
	fmt.Println("finished")
}
