package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func help() {
	fmt.Println("使用说明: dtTool.exe  <cmd> --file=<要替换的文件名>  --target=<目标位置>")
	fmt.Println("支持的指令: cp  <cmd> --file=<要替换的文件名>")
}

func cp(cmStr string, sourcePtr *string, targetPtr *string) ([]byte, error) {
	if *sourcePtr == "" {
		fmt.Println("没有文件")
		return nil, nil
	}
	// 输出解析后的参数值
	fmt.Println("文件名称:", *sourcePtr)

	dist_rlmt := "k3d-k3s-v-local-reg-server-0:/mnt/twinverse-edge-pv/rlmt-edge-api-1g"

	// str_dt := "k3d-k3s-volumn-local-reg-1-server-0:/data/twinverse-edge-pv/rlmt-mock-data-1g/config" //
	cmd := exec.Command("docker", cmStr, *sourcePtr, dist_rlmt)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return stdout, nil
}

func main() {
	// 定义一个名为 "file" 的命令行参数，类型为字符串，默认值为 ""，并附加了说明信息
	filePtr := flag.String("file", "", "fileName")
	cmdPtr := flag.String("cmd", "", "cmdName")
	targetPtr := flag.String("target", "", "targetName")
	l := len(os.Args)
	if l == 1 {
		fmt.Println("没有指令，需要先输入指令")
		help()
		return
	}

	// 解析命令行参数
	flag.Parse()

	cmStr := *cmdPtr
	fmt.Printf("cmd = ", cmStr, " filePtr = ", *filePtr, "targetPtr = ", *targetPtr)
	//
	// if *filePtr == "" {
	// 	fmt.Println("没有文件")
	// 	return
	// }
	// // 输出解析后的参数值
	// fmt.Println("文件名称:", *filePtr)

	// dist_rlmt := "k3d-k3s-v-local-reg-server-0:/mnt/twinverse-edge-pv/rlmt-edge-api-1g"

	// // str_dt := "k3d-k3s-volumn-local-reg-1-server-0:/data/twinverse-edge-pv/rlmt-mock-data-1g/config" //
	// cmd := exec.Command("docker", cmStr, *filePtr, dist_rlmt)
	// stdout, err := cmd.Output()

	if "cp" == cmStr {
		stdout, err := cp(cmStr, filePtr, targetPtr)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println(string(stdout))
		fmt.Println("finished", cmStr)
	}
}
