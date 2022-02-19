package model

type Provinsi struct {
	ID       uint       `gorm:"primarykey"`
	Name     string     `json:"name"`
	KabKotas []KabKotas `gorm:"foreignKey:ProvId"`
	Spasial  Spasial    `gorm:"foreignKey:IdDaerah"`
}

type KabKotas struct {
	ID         uint        `gorm:"primarykey"`
	ProvId     uint        `json:"prov_id"`
	Name       string      `json:"name"`
	Kecamatans []Kecamatan `gorm:"foreignKey:IdKab"`
	Spasial    Spasial     `gorm:"foreignKey:IdDaerah"`
}

type Kecamatan struct {
	ID         uint        `gorm:"primarykey"`
	IdKab      uint        `json:"id_kab"`
	Name       string      `json:"name"`
	Kelurahans []Kelurahan `gorm:"foreignKey:IdKec"`
	Spasial    Spasial     `gorm:"foreignKey:IdDaerah"`
}

type Kelurahan struct {
	ID      uint    `gorm:"primarykey"`
	IdKec   uint    `json:"id_kec"`
	Name    string  `json:"name"`
	Spasial Spasial `gorm:"foreignKey:IdDaerah"`
}

type Spasial struct {
	ID       uint   `gorm:"primarykey"`
	IdDaerah uint   `json:"id_daerah"`
	Center   string `json:"center"`
}

type User struct {
	ID       uint   `gorm:"primarykey"`
	Username string `json:"username"`
	Password string `json:"password"`
}
