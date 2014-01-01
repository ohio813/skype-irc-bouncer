package main

import "os"
import "log"
import "time"
import "flag"

import gsi "github.com/williballenthin/go-skype-irc"

func main() {
  configFilename := flag.String("config", "gsi.conf", "path to gsi.conf file")
  flag.Parse()

  config, e := gsi.LoadConfig(*configFilename)
  if e != nil {
    log.Fatal(e)
  }

  client, e := gsi.MakeClient(config)
  if e != nil {
    log.Fatal(e)
  }

  username := config.GetIfDefined("Client", "username", "")
  password := config.GetIfDefined("Client", "password", "")
  if e := client.Authenticate(username, password); e != nil {
    log.Fatal(e)
  }

  go func() {
    time.Sleep(2 * time.Second)
    client.TriggerUserTest()
    time.Sleep(2 * time.Second)
    client.DumpUsers(os.Stdout)

    time.Sleep(2 * time.Second)
    client.TriggerGroupTest()
    time.Sleep(2 * time.Second)
    client.DumpGroups(os.Stdout)

    time.Sleep(2 * time.Second)
    client.TriggerChatmessageTest()
    time.Sleep(2 * time.Second)
    client.DumpChatmessages(os.Stdout)

    time.Sleep(2 * time.Second)
    client.TriggerChatTest()
    time.Sleep(2 * time.Second)
    client.DumpChats(os.Stdout)
  }()

  client.Serve()
}
