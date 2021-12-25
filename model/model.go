package model

type Provinsi struct {
	Id       int    `json:"id"`
	ProvName string `json:"prov_name"`
}

type KabKotas struct {
	Id     int    `json:"id"`
	ProvId int    `json:"prov_id"`
	Name   string `json:"name"`
}
