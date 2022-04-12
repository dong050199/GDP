package model

type Scada struct {
	Command  string `json:"command,omitempty"`
	Order_ID int    `json:"order_id,omitempty"`
}
