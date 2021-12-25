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

type Kecamatan struct {
	Id    int    `json:"id"`
	IdKab int    `json:"id_kab"`
	Name  string `json:"name"`
}

type Kelurahan struct {
	Id    int    `json:"id"`
	IdKec int    `json:"id_kec"`
	Name  string `json:"name"`
}
