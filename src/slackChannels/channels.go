package main

import(
  "fmt"
  "github.com/nlopes/slack"
)

type ChannelCreatedEvent struct {
    Type           string             `json:"type"`
    Channel        ChannelCreatedInfo `json:"channel"`
    EventTimestamp string             `json:"event_ts"`
}

type ChannelCreatedInfo struct {
    ID        string `json:"id"`
    IsChannel bool   `json:"is_channel"`
    Name      string `json:"name"`
    Created   int    `json:"created"`
    Creator   string `json:"creator"`
}

func main(){
  api := slack.New("xoxp-254723404592-256259256790-270170005381-f12427611965933f636e0066af4acb8f")
  channels, err := api.GetChannels(false)
  if err != nil {
    fmt.Printf("%s\n", err)
    return
  }

  for _, channel := range channels {
		fmt.Println(channel.Name)
		// channel is of type conversation & groupConversation
		// see all available methods in `conversation.go`
	}


}
 //create slack channel
  // func (api) CreateChannel(channel string) (*Channel, error){
  //   // fmt.Println("Creating a slack channel")
  // }