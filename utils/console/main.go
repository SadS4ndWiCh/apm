package console

import (
  "os"
  "fmt"
  "bufio"
  "regexp"
)

func Ask(message string) (bool, error) {
  fmt.Print(message)

  reader := bufio.NewReader(os.Stdin)
  input, err := reader.ReadString('\n')

  if err != nil { return false, err }

  validation := regexp.MustCompile(`^[Yy]$`)
  ok := validation.MatchString(string([]byte(input)[0]))

  return ok, nil
}
