package alert

type Alert struct {
	Text       string       `json:"text"`
	Attachment []Attachment `json:"attachments"`
}

type Attachment struct {
	Color     string  `json:"color"`
	Field     []Field `json:"fields"`
	Footer    string  `json:"footer"`
	TimeStamp int64   `json:"ts"`
	Title     string  `json:"title"`
	Text      string  `json:"text"`
}

type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}
