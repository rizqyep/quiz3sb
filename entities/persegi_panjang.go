package entities

type PersegiPanjangQuery struct {
	Panjang int    `form:"panjang"`
	Lebar   int    `form:"lebar"`
	Hitung  string `form:"hitung"`
}
