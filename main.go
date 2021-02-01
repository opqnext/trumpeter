package main

import "trumpeter/cmd"

//var host string
//var port int
//
//func init() {
//	flag.Usage = func() {
//		fmt.Printf("Usage of %s:\n", os.Args[0])
//		fmt.Printf("Young and talented trumpeter🎺 \n")
//		flag.PrintDefaults()
//	}
//
//	flag.StringVar(&host, "h", "localhost", "specify host to use. defaults to localhost.")
//	flag.IntVar(&port, "p", 6379, "specify port to use. defaults to 6379.")
//}

func main() {

	cmd.Execute()

	//flag.Parse()
	//
	//app := httpd.New()
	//app.Get("/", indexHandlerFunc)
	//app.Run()
	//
	//fmt.Printf("host = %s \n", host)
	//fmt.Printf("port = %d \n", port)
	//db.InitRedis(host, port)
	//db.Pop()

	//web.Run()
}
