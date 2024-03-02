package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/JasonBoyett/salvo/public/json"
	"github.com/JasonBoyett/salvo/public/protobuff"
	"google.golang.org/protobuf/proto"
)

func main() {
	Main()
}

func Main() {
	isFile := flag.Bool("file", false, "is file")
	isJson := flag.Bool("json", false, "is json")
	optsString := flag.String("opts", "", "options")
	filePath := flag.String("path", "", "file path")
	flag.Parse()
	if *filePath == "" && *isFile {
		fmt.Println("no file path provided")
		os.Exit(1)
	}
	if *optsString == "" && !*isFile {
		fmt.Println("no options provided")
		fmt.Println("options: ", *optsString)
		os.Exit(1)
	}

	if *isFile {
		handleFile(*filePath, *optsString, *isJson)
	} else {
		handleStdOut(*filePath, *optsString, *isJson)
	}
}

func handleFile(filePath string, optsString string, isJson bool) {
	if isJson {
		message, err := json.Receive([]byte(optsString))
		if err != nil {
			fmt.Println("error from receiver: ", err)
		}
		err = json.SendFile(message, filePath)
		if err != nil {
			fmt.Println("error error from file write: ", err)
		}
	} else {
		pbuff := &protobuff.Options{}
		err := proto.Unmarshal([]byte(optsString), pbuff)
		if err != nil {
			fmt.Println("error unmarshaling protobuff: ", err)
		}
		message := protobuff.Receive(pbuff)
		err = protobuff.SendFile(filePath, message)
		if err != nil {
			fmt.Println("error error from file write: ", err)
		}
	}
}

func handleStdOut(filePath string, optsString string, isJson bool) {
	if isJson {
		message, err := json.Receive([]byte(optsString))
		if err != nil {
			fmt.Println("error from receiver: ", err)
		}
		err = json.SendStdout(os.Stdout, message)
		if err != nil {
			fmt.Println("error error from printer: ", err)
		}
	} else {
		pbuff := &protobuff.Options{}
		err := proto.Unmarshal([]byte(optsString), pbuff)
		if err != nil {
			fmt.Println("error unmarshaling protobuff: ", err)
		}
		message := protobuff.Receive(pbuff)
		err = protobuff.SendStdout(os.Stdout, message)
		if err != nil {
			fmt.Println("error error from printer: ", err)
		}
	}
}
