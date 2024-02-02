package json

import (
  "encoding/json"
  "fmt"
  "os"
  "io"

  "github.com/JasonBoyett/salvo/local/util"
)

func send(data Message) ([]byte, error) {
  json, err := json.Marshal(data)
  if err != nil {
    return []byte{}, err
  }
  return json, nil
}

func SendStdout(w io.Writer, data Message) error {
  json, err := send(data)
  if err != nil {
    return err
  }
  fmt.Fprint(w, string(json))
  return nil
}

// SendFile creates a file called salvo.json in the directory of the file path.
// If the file path points to a file, salvo.json will be created in the same directory.
// If the file path points to a directory, salvo.json will be created in that directory.
func SendFile(data Message, path string) error {
  base, err := os.Open(path) 
  if err != nil {
    return err
  }
  baseInfo, err := base.Stat()
  if err != nil {
    return err
  }
  if !baseInfo.IsDir() {
    path, err = util.FindDir(path)
    if err != nil {
      return fmt.Errorf("Error finding base directory")
    }
  }
  fmt.Printf("path before panic: %s\n", path)
  if path[len(path) - 1] != os.PathSeparator {
    path = path + string(os.PathSeparator)
  }

  path = path + "salvo.json"
  json, err := send(data)
  if err != nil {
    return err
  }

  file, err := os.Create(path)
  if err != nil {
    return err
  }

  defer file.Close()
  _, err = file.Write(json)
  if err != nil {
    return err
  }

  return nil
}

