package main

import (
    "fmt"
    "gopkg.in/maciekmm/messenger-platform-go-sdk.v4"
    "net/http"
    "log"
    "flag"
)

var mess *messenger.Messenger
/* {
    VerifyToken: "thoughtbuzztoken",
    AppSecret: "",
    AccessToken: "",
}*/

func main() {
    verifyToken := flag.String("verifyToken", "", "")
    appSecret := flag.String("appSecret", "", "")
    accessToken := flag.String("accessToken", "", "")
    port := flag.String("port", "5646", "")
    flag.Parse()
    mess = &messenger.Messenger{
        VerifyToken:  *verifyToken,
        AppSecret: *appSecret,
        AccessToken: *accessToken,
    }
    mess.MessageReceived = MessageReceived
    http.HandleFunc("/webhook/automobo", mess.Handler)
    fmt.Printf("started on port:%s",*port)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", *port), nil))
}

func MessageReceived(event messenger.Event, opts messenger.MessageOpts, msg messenger.ReceivedMessage) {
    fmt.Println("SenderID: ",opts.Sender.ID)
    fmt.Println("RecipientID: ",opts.Recipient.ID)
    fmt.Println("Text: ",msg.Text)
    profile, err := mess.GetProfile(opts.Sender.ID)
    if err != nil {
        fmt.Println(err)
        return
    }
    resp, err := mess.SendSimpleMessage(opts.Sender.ID, fmt.Sprintf("Hello, %s %s, %s", profile.FirstName, profile.LastName, msg.Text))
    if err != nil {
        fmt.Println(err)
    }
    fmt.Printf("%+v", resp)
}

