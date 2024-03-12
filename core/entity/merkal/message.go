package merkal

import (
	"encoding/json"
	"reflect"
	"time"
)

type Message struct {
	Title      string    `json:"title"`
	Text       string    `json:"text"`
	CreateTime time.Time `json:"createTime"`
}

func (this *Message) String() string {
	res, _ := json.Marshal(this)
	return string(res)
}

func (this *Message) Equals(other Message) bool {
	return reflect.DeepEqual(this, other)
}
