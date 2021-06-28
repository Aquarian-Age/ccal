package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strings"
)

func checkSystemName()  {
	cmd:=exec.Command("lsb_release","-si")
	out,_:=cmd.Output()
	fmt.Printf("Your System is: %s",string(out))
	osname:= strings.ToLower(string(out))
	osname= strings.Trim(osname,"\n")
	if !strings.EqualFold(osname,"opensuse"){
		fmt.Println("Your system is not authorized! bGlhbmd6aUB5YW5kZXguY29tCg==")
		os.Exit(1)
	}
}
//检查编码
func checkInfo(md5s, machNumber string) {

	machN, md5 := readInfo("auth.info")
	if strings.EqualFold(md5s, md5) && strings.EqualFold(machN, machNumber) {
		fmt.Println("OK")
	} else if !strings.EqualFold(md5s, md5) || !strings.EqualFold(machN, machNumber) {
		fmt.Println("Your system is not authorized! bGlhbmd6aUB5YW5kZXguY29tCg==")
		os.Exit(1)
	}
}

//go build -o authTool-linux -ldflags="-s -w" -tags timetzdata .
//生成硬件信息的编码
func auth() {
	mac, err := getMACAddress()
	if err != nil {
		fmt.Println(err)
	}
	s, _ := getMacnineID()
	n := getNumCPU()
	sn := fmt.Sprintf("%d", n)
	s = s + mac + sn
	md5s := GetMd5String(s, true, true)
	writeToFile(md5s)
}
func readInfo(file string) (string, string) {
	f := "/tmp/"+file
	b, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Println(err)
	}
	machineOSNumber := string(b[:1]) //机器编码 0win 1linux 2osx
	md5s := string(b[22:])
	md5s = strings.Trim(md5s, "\n") //去掉末尾换行
	//fmt.Printf("readmachine:%s %d\n", machineOSNumber, len(machineOSNumber))
	//fmt.Printf("readmd5s: %s %d\n", md5s, len(md5s))
	return machineOSNumber, md5s
}

//输出log
func writeToFile(info string) {
	fileName := "/tmp/auth.info"
	logFile, err := os.Create(fileName)
	defer logFile.Close()
	if err != nil {
		log.Fatalln("create file err!")
	}
	//创建日志
	debugLog := log.New(logFile, "1 ", log.LstdFlags)
	debugLog.Println(info)
}

//机器id 这个系统重装之后会改变 后面数字是编码 0win 1linux 2osx
func getMacnineID() (string, string) {
	cmd := exec.Command("cat", "/etc/machine-id")
	b, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	machiNumber := "1"
	//fmt.Printf("machine:%s %d\n", machiNumber, len(machiNumber))
	return (string(b)), machiNumber
}

//用户信息
func getGuid() (*user.User, error) {
	u, err := user.Current()
	return u, err
}

//cpu最大数量
func getNumCPU() int {
	return runtime.NumCPU()
}

//获取mac
func getMACAddress() (string, error) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		panic(err.Error())
	}
	mac, macerr := "", errors.New("Can not get the MAC Addr")
	for i := 0; i < len(netInterfaces); i++ {
		//fmt.Println(netInterfaces[i])
		if (netInterfaces[i].Flags&net.FlagUp) != 0 && (netInterfaces[i].Flags&net.FlagLoopback) == 0 {
			addrs, _ := netInterfaces[i].Addrs()
			for _, address := range addrs {
				ipnet, ok := address.(*net.IPNet)
				//fmt.Println(ipnet.IP)
				if ok && ipnet.IP.IsGlobalUnicast() {
					// 如果IP是全局单拨地址，则返回MAC地址
					mac = netInterfaces[i].HardwareAddr.String()
					return mac, nil
				}
			}
		}
	}
	return mac, macerr
}

//生成32位md5字串
func GetMd5String(s string, upper bool, half bool) string {
	h := md5.New()
	h.Write([]byte(s))
	result := hex.EncodeToString(h.Sum(nil))
	if upper == true {
		result = strings.ToUpper(result)
	}
	if half == true {
		result = result[8:24]
	}
	//fmt.Printf("result:%s %d\n", result, len(result))
	return result
}
