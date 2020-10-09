package model

import "time"

type Orders struct {
	Id       uint64 `gorm:"primary_key;column:id" json:"id"`
	Distance int    `json:"distance"`
	Status   string `json:"status"`

	StartLatitude  string `json:"start_latitude"`
	StartLongitude string `json:"start_longitude"`

	EndLatitude  string `json:"end_latitude"`
	EndLongitude string `json:"end_longitude"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *Orders) BeforeSave() (err error) {
	u.UpdatedAt = time.Now()
	return
}
