package quiz

type Question struct {
	Question    string   `json:"question"`
	Code        string   `json:"code,omitepmty"`
	Options     []string `json:"options"`
	Answer      []string `json:"answer"`
	Explanation string   `json:"explanation"`
}
