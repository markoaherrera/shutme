package main

import (
  "fmt"
  // "syscall"
  "os"
  "os/exec"
  "time"
  "strconv"
)

func main() {

  args := os.Args[1:]

  sleepTime := args[0]

  minutes, convErr := strconv.Atoi(sleepTime)
  if convErr != nil { panic(convErr) }

  minute := 1
  for ; minute <= minutes; {
    fmt.Printf("\rMinute %d", minute)
    time.Sleep(1 * time.Minute)
    minute += 1
  }
  fmt.Printf("\nPowering off, bye!\n")

  cmd := exec.Command("poweroff")
  cmdErr := cmd.Run()
  if cmdErr != nil {
    panic(cmdErr)
  }

}
