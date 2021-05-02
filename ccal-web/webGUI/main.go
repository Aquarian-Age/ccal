package main

import (
	"fmt"
	"github.com/webview/webview"
	"log"
	"os/exec"
	"runtime"
)

func main() {
	os := runtime.GOOS
	switch os {
	case "linux":
		//启动服务
		cmd := exec.Command("/usr/bin/sh", "/home/xuan/ccal/web/listen.sh")
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}

		debug := true
		w := webview.New(debug)
		defer w.Destroy()
		w.SetTitle("农历择日")
		w.SetSize(1366, 800, webview.HintNone)
		w.Navigate("http://127.0.0.1:9090")
		w.Run()

	case "windows":
		fmt.Println("不支持的系统")
		//windows()
	}
}
/*
func windows() {
	cmd := exec.Command("cmd")
	in := bytes.NewBuffer(nil)
	cmd.Stdin = in //绑定输入
	var out bytes.Buffer
	cmd.Stdout = &out //绑定输出
	go func() {
		// start stop restart
		in.WriteString("nssm restart uplusSVCWB\n") //写入你的命令，可以有多行，"\n"表示回车
	}()
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(cmd.Args)
	err = cmd.Wait()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
	}
	rt := out.String()
	fmt.Println(rt)

	if strings.ContainsAny(rt, "成功") {
		fmt.Fprintf(w, "{\"msg\":\"%s\"}", "操作成功")
	} else {
		fmt.Fprintf(w, "{\"msg\":\"%s\"}", rt)
	}
}*/

/*func filePath() {
	folderPath, err := osext.ExecutableFolder()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(folderPath)
}*/
