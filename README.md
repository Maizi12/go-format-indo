# go-format-indo

Bismillahirrahmanirrahim

Memformat float64 di struct menjadi string dan memasukkan string itu ke field string di struct tersebut
contoh:
type Value struct {
	HasilJual                   string  `json:"hasil_jual_string"`
	UangPinjaman                string  `json:"uang_pinjaman_string"`
	JasaSimpan                  string  ` json:"jasa_simpan_string"`
	BiayaPemeliharaan           string  `json:"biaya_pemeliharaan_string"`
	BiayaPenjualan              string  `json:"biaya_penjualan_string"`
	TotalKewajiban              string  `json:"total_kewajiban_string"`
	UangKelebihanPenjualan      string  `json:"uang_kelebihan_penjualan_string"`
	HasilJualFloat              float64 `json:"hasil_jual"`
	UangPinjamanFloat           float64 `json:"uang_pinjaman"`
	JasaSimpanFloat             float64 ` json:"jasa_simpan"`
	BiayaPemeliharaanFloat      float64 `json:"biaya_pemeliharaan"`
	BiayaPenjualanFloat         float64 `json:"biaya_penjualan"`
	TotalKewajibanFloat         float64 `json:"total_kewajiban"`
	UangKelebihanPenjualanFloat float64 `json:"uang_kelebihan_penjualan"`
}

FormatFloat64Struct(&Value,value,"",n)

n=panjang dari slice nya struct,kalau bukan slice masukkan 1.

Function ini akan looping ke dalam struct nya, mengambil dan mengubah value dari field float ke string dan memasukkan nya ke field string, misalnya dari HasilJualFloat kemudian diubah dan dimasukkan ke HasilJual

FormatTanggal(tanggal string) string
2023-01-13 jadi 13 Januari 2023 yyyy-mm-dd(mengikuti format tgl mysql)
