package models

import "encoding/json"

// KafkaRecord define kafka record
type KafkaRecord struct {
	Key   string
	Value KafkaValue
}

// KafkaValue define the interface of kafka value
type KafkaValue interface {
	Value() []byte
}

//xxxKafkaMsg kafka msg
type xxxKafkaMsg struct {
	xxxResultID int64
}

func (r xxxKafkaMsg) Value() []byte {
	marshal, _ := json.Marshal(r)
	return marshal
}
