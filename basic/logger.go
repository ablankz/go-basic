package basic

import (
	"io"
	"log"
	"os"
)

func Logger() {
	file, err := os.Create("log.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	flags := log.Lshortfile | log.LstdFlags
	warnLogger := log.New(io.MultiWriter(file, os.Stderr), "WARN: ", flags)
	errorLogger := log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", flags)

	warnLogger.Println("warning A")

	// ログを書き込んだあと、強制終了
	errorLogger.Fatalln("critical error")
}
