package entity

type Task struct {
	Id       int    `json:"id"`
	Time     string `json:"time"`
	Phone    string `json:"phone"`
	DutyName string `json:"duty_name"`
}
