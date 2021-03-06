package main

import (
	"flag"
	"log"
	"os"

	"github.com/lucas-clemente/quic-go"
)

var (
	addr    = flag.String("addr", "127.0.0.1:4443", "addr")
	harFile = flag.String("har", "", "har")
)

func initiate() error {
	cl, err := dial(*addr)
	if err != nil {
		return err
	}
	defer cl.close()
	cl.initiate("")
	cl.wg.Wait()
	return nil
}

func initDB() error {
	fin, err := os.Open(*harFile)
	if err != nil {
		return err
	}
	defer fin.Close()
	return db.load(fin)
}

func main() {
	flag.Parse()
	quic.SetLogLevel("info")

	err := initDB()
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(initiate())
}
