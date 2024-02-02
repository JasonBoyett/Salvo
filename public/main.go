package main

import (
  "flag"
  "os"
)

func main() {
  args := os.Args[1:]
  basePath :=os.Args[0]

  sendType := flag.String(
    "send", 
    "protobuf", 
    "Indicates weather salvo will send and receive json or a protocol buffer",
    )
  out := flag.String(
    "out", 
    "std", 
    "Indicates weather the output will be sent via standard output or be written to a file",
    )
}

