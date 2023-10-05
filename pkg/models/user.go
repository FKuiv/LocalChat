package models

type User struct {
	ID       string `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Local_Ip string `json:"local_ip"`
}
