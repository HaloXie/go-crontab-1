package main

import "runtime"

func initEnv()  {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
func main()  {
	//初始化线程
	initEnv()

	//启动Api Http服务

}