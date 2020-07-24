package model

import framework "github.com/aripstheswike/swikefw"

func GetAllProvinsi() ([]map[string]interface{}, error) {
	db := framework.Database{}
	defer db.Close()

	db.Select("id_prov id, nama label").From("provinsi")
	return db.Result()
}

func GetKota(idprov string) ([]map[string]interface{}, error) {
	db := framework.Database{}
	defer db.Close()

	db.Select("id_kab id, nama label").From("kabupaten").Where("id_prov", idprov)
	return db.Result()
}

func GetKecamatan(idkab string) ([]map[string]interface{}, error) {
	db := framework.Database{}
	defer db.Close()

	db.Select("id_kec id, nama label").From("kecamatan").Where("id_kab", idkab)
	return db.Result()
}

func GetKelurahan(idkec string) ([]map[string]interface{}, error) {
	db := framework.Database{}
	defer db.Close()

	db.Select("id_kel id, nama label").From("kelurahan").Where("id_kec", idkec)
	return db.Result()
}
