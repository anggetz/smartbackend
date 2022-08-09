package model

type Blok struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	GroupLocationID int    `json:"id_group_location"`
	GroupLocation   GroupLocation
	*Base
}

func (b *Blok) TableName() string {
	return "m_blok"
}
