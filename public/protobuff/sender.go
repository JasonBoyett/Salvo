package protobuff

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/JasonBoyett/salvo/local/util"
	"google.golang.org/protobuf/proto"
)

func FormatMessage(data Message) *Response {
  var response Response
  for _, r := range data.Results {
    response.Result = append(response.Result, ToProtobuff(r))
  }
  response.ErrorCount = int32(data.Fails)

  return &response
}

func SendStdout(w io.Writer, data Message) error {
  message := FormatMessage(data) 
  pbByteArray, err := proto.Marshal(message)
  if err != nil {
    return err
  }
  if _, err = fmt.Fprintf(w, "%s", pbByteArray); err != nil {
    return err
  }
  return nil
}

// SendFile creates a file called salvo.pb in the directory of the file path.
// If the file path points to a file, salvo.pb will be created in the same directory.
// If the file path points to a directory, salvo.pb will be created in that directory.
func SendFile(path string, data Message) error {
  message := FormatMessage(data) 
  pbByteArray , err := proto.Marshal(message)
  if err != nil {
    return err
  }
  
  path, err = util.FindDir(path)
  if err != nil {
    return fmt.Errorf("Error finding base directory")
  }

  file, err := os.Create(path + string(os.PathSeparator) + "salvo.pb")
  if err != nil {
    return err
  }

  if err = ioutil.WriteFile(file.Name(), pbByteArray, 0644); err != nil {
    return err
  }

  return nil
}
