package go_format_indo

import (
	"github.com/paimanbandi/rupiah"
	"reflect"
	"strconv"
	"strings"
)

func FormatFloat64Struct(value interface{}, types interface{}, mode string, len int) interface{} {
	values := reflect.ValueOf(value).Elem() //untuk tipe data dari field struct nya, kayak string dari field nama
	typess := reflect.ValueOf(types).Type() //untuk nama dari field struct nya, kayak field nama
	for i := 0; i < len; i++ {
		for j := 0; j < values.NumField(); j++ {
			f := values.Field(j)
			if f.Kind().String() == "float64" {
				values.FieldByName(strings.TrimSuffix(typess.Field(j).Name, "Float")).SetString(rupiah.FormatFloat64ToRp(f.Float()))
			}
		}
	}

	return values
}
func FormatTanggal(tanggal string) string {
	res := strings.SplitAfter(tanggal, "-")
	tanggal = res[2] + " " + MonthIndo(strings.Trim(res[1], "-")) + " " + strings.Trim(res[0], "-")
	tanggal = strings.Trim(tanggal, "-")
	return tanggal
}

func MonthIndo(monthstring string) string {
	var bulan string
	month, _ := strconv.Atoi(monthstring)
	switch month {
	case 1:
		bulan = "Januari"
	case 2:
		bulan = "Februari"
	case 3:
		bulan = "Maret"
	case 4:
		bulan = "April"
	case 5:
		bulan = "Mei"
	case 6:
		bulan = "Juni"
	case 7:
		bulan = "Juli"
	case 8:
		bulan = "Agustus"
	case 9:
		bulan = "September"
	case 10:
		bulan = "Oktober"
	case 11:
		bulan = "November"
	case 12:
		bulan = "Desember"
	}
	return bulan
}
