package database

// A Attachment represents a key-value object used for storing ODRecord.
type Messages map[string] []interface{}
// {"text": "Welcome to the Chatfuel Rockets!"}
type Text map[string]string

type AttachmentMedia struct {
    Attachment`json:"attachment,omitempty"`
  }

type Attachment struct{
  Type string `json:"type,omitempty"`
  Payload map[string]interface{} `json:"payload,omitempty"`
}
