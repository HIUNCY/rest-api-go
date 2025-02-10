package model

type Transaction struct {
	IDTabungan int    `gorm:"column:id_tabungan;primaryKey"`
	NIK        string `gorm:"column:nik;size:25;not null" json:"nik"`
	Setor      int    `gorm:"column:setor;not null" json:"setor"`
	Tarik      int    `gorm:"column:tarik;not null" json:"tarik"`
	Jenis      string `gorm:"column:jenis;type:enum('setor','tarik');not null" json:"jenis"`
}
