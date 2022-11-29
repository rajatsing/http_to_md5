package main

import (
	"flag"
	"log"

	httpmd "http_to_md5/httptomd"
)

func main() {
	parallel := flag.Int("parallel", 10, "limit the parallel requests")
	flag.Parse()
	urls := flag.Args()
	log.Println("urls are -->	", urls)
	httpmd.Run(*parallel, urls)
}
