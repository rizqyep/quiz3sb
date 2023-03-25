package entities

type SegitigaQuery struct {
	Alas   float64 `form:"alas"`
	Tinggi float64 `form:"tinggi"`
	Hitung string  `form:"hitung"`
}
