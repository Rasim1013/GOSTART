package types

type User struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Msisdn string `json:"msisdn"`
}

type UserDetail struct {
	User
	StatusID  int    `json:"statusId"`
	Trpl_id   int    `json:"trpl_id"`
	Trpl_name string `json:"trpl_name"`
	Birthday  string `json:"birthday"`
}
