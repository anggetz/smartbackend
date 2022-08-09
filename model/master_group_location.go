package model

type GroupLocation struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	*Base
}

func (g *GroupLocation) TableName() string {
	return "m_group_location"
}
