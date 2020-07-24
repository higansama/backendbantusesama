package model

import framework "github.com/aripstheswike/swikefw"

func GetAllCampaign(filter map[string]interface{}) ([]map[string]interface{}, error) {
	db := framework.Database{}
	defer db.Close()

	db.Select("*").From("campaign")

	r, e := db.Result()
	return r, e
}

func SortAllCampaign(sort string) ([]map[string]interface{}, error) {
	db := framework.Database{}
	defer db.Close()

	db.Select("c.nama_campaign, c.id, date_created tanggal, kab.nama kota, kec.nama kecamatan, kel.nama kelurahan, c.foto").From("campaign c")
	db.Join("kecamatan kec", "c.kecamatan = kec.id_kec", "")
	db.Join("kelurahan kel", "c.kelurahan = kel.id_kel", "")
	db.Join("kabupaten kab", "c.kota = kab.id_kab", "")
	if sort == "terbaru" {
		db.OrderBy("date_created", "DESC")
	}

	r, e := db.Result()
	return r, e
}

func GetActionCampaignByIDAndValueAction(filter map[string]interface{}) ([]map[string]interface{}, error) {
	db := framework.Database{}
	defer db.Close()

	db.Select("c.nama_campaign, c.latar_cerita, c.kota, c.kecamatan, c.kelurahan, a.date_created tanggal_mulai ").From("user_action_campaign a")
	db.Join("campaign c", "c.id = a.id_campaign", "")
	db.Where("id_user", filter["id_user"])
	db.Where("value_action", filter["action_value"])
	r, e := db.Result()
	return r, e
}

func DetailCampaign(idcampaign string) ([]map[string]interface{}, error) {
	db := framework.Database{}
	defer db.Close()

	db.Select("c.nama_campaign, c.latar_cerita, c.kota, c.kecamatan, c.kelurahan").From("campaign c")
	db.Where("c.id", idcampaign)
	r, e := db.Result()
	return r, e
}

func GetKomentarCampaign(idcampaign string) ([]map[string]interface{}, error) {
	db := framework.Database{}
	defer db.Close()

	db.Select("*").From("komentar k")
	db.Join("campaign c", "c.id = k.id_campaign", "")
	db.Join("user u", "u.id = k.id_user", "")
	db.Where("k.id_campaign", idcampaign)
	r, e := db.Result()
	return r, e
}
