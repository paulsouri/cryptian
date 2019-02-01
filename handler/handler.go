package handler

import (
	_ "encoding/json"
	_ "errors"
	"net/http"
  . "github.com/paulsouri/cryptian/database"
)


func NewsHandler(w http.ResponseWriter, r *http.Request) {
  attachment:= Attachment{
    Type:"audio",
    Payload:map[string]interface{}{ "url": "https://rockets.chatfuel.com/assets/welcome.png"},
  }
  attachment_media:= AttachmentMedia{attachment}
	Json(w, http.StatusOK, Messages{"messages":[]interface{}{attachment_media}})
}


func MeetupHandler(w http.ResponseWriter, r *http.Request) {
  Json(w, http.StatusOK, map[string]string{"message": "User created:"})
}
