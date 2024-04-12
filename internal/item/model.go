package item

type Item struct {
	ID        string `json:"id" bson:"_id, omitempty"`
	Number    string `json:"number" bson:"number"`
	Mqtt      string `json:"mqtt" bson:"mqtt"`
	Invid     string `json:"invid" bson:"invid"`
	UnitGuid  string `json:"unitguid" bson:"unitguid"`
	MsgId     string `json:"msgid" bson:"masgid"`
	Text      string `json:"text" bson:"text"`
	Context   string `json:"ctx" bson:"ctx"`
	Class     string `json:"class" bson:"class"`
	Level     string `json:"level" bson:"level"`
	Area      string `json:"area" bson:"area"`
	Addr      string `json:"addr" bson:"addr"`
	Block     string `json:"block" bson:"block"`
	Type      string `json:"type" bson:"type"`
	Bit       string `json:"bit" bson:"bit"`
	InvertBit string `json:"invbit" bson:"invbit"`
}

type ItemDTO struct {
	Number    string `json:"number"`
	Mqtt      string `json:"mqtt"`
	Invid     string `json:"invid"`
	UnitGuid  string `json:"unitguid"`
	MsgId     string `json:"msgid"`
	Text      string `json:"text"`
	Context   string `json:"ctx"`
	Class     string `json:"class"`
	Level     string `json:"level"`
	Area      string `json:"area"`
	Addr      string `json:"addr"`
	Block     string `json:"block"`
	Type      string `json:"type"`
	Bit       string `json:"bit"`
	InvertBit string `json:"invbit"`
}
