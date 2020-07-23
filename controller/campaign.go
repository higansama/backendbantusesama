package controller

import (
	"backend_skripsi/lib"
	"backend_skripsi/model"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateCampaign(c *gin.Context) {
	user_pembuat := c.PostForm("id_user")
	nama_campaign := c.PostForm("nama_campaign")
	latar_cerita := c.PostForm("latar_cerita")
	kota := c.PostForm("kota")
	kecamatan := c.PostForm("kecamatan")
	kelurahan := c.PostForm("kelurahan")
	date_expired := c.PostForm("date_expired")
	nama_jalan := c.PostForm("nama_jalan")
	foto := c.PostForm("foto")

	campaignInsert := map[string]interface{}{
		"user_pembuat":  user_pembuat,
		"kelurahan":     kelurahan,
		"latar_cerita":  latar_cerita,
		"kota":          kota,
		"kecamatan":     kecamatan,
		"nama_campaign": nama_campaign,
		"date_expired":  date_expired,
		"nama_jalan":    nama_jalan,
		"foto":          foto,
	}
	err := model.InsertDataToTable("campaign", campaignInsert)
	if err != nil {
		fmt.Println(err.Error())
		lib.ReturnError(c, err)
		return
	}
	lib.Json(c, 200, "success", gin.H{})
}

func GetAllCampaign(c *gin.Context) {
	user_pembuat := c.Query("id_user")
	nama_campaign := c.Query("nama_campaign")
	latar_cerita := c.Query("latar_cerita")
	kota := c.Query("kota")
	kecamatan := c.Query("kecamatan")
	kelurahan := c.Query("kelurahan")
	date_expired := c.Query("date_expired")

	sort := c.Query("sort")
	data := []map[string]interface{}{}

	if sort != "" {
		data, err := model.SortAllCampaign(sort)
		if err != nil {
			lib.ReturnError(c, err)
			return
		}

		lib.Json(c, 200, "success", data)
	}

	filter := map[string]interface{}{
		"user_pembuat":  user_pembuat,
		"kelurahan":     kelurahan,
		"latar_cerita":  latar_cerita,
		"kota":          kota,
		"kecamatan":     kecamatan,
		"nama_campaign": nama_campaign,
		"date_expired":  date_expired,
	}

	data, err := model.GetAllCampaign(filter)
	if err != nil {
		lib.ReturnError(c, err)
		return
	}

	lib.Json(c, 200, "success", data)
}

func FollowCampaign(c *gin.Context) {
	id_user := c.PostForm("id_user")
	id_campaign := c.PostForm("id_campaign")
	action := "0"

	actionData := map[string]interface{}{
		"id_user":      id_user,
		"id_campaign":  id_campaign,
		"value_action": action,
	}
	err := model.InsertDataToTable("user_action_campaign", actionData)
	if err != nil {
		lib.ReturnError(c, err)
		return
	}

	lib.Json(c, 200, "success", gin.H{})
}
func LikeCampaign(c *gin.Context) {
	id_user := c.PostForm("id_user")
	id_campaign := c.PostForm("id_campaign")
	action := "0"

	actionData := map[string]interface{}{
		"id_user":      id_user,
		"id_campaign":  id_campaign,
		"value_action": action,
	}
	err := model.InsertDataToTable("user_action_campaign", actionData)
	if err != nil {
		lib.ReturnError(c, err)
		return
	}

	lib.Json(c, 200, "success", gin.H{})
}
func RememberCampaign(c *gin.Context) {
	id_user := c.PostForm("id_user")
	id_campaign := c.PostForm("id_campaign")
	action := "1"

	actionData := map[string]interface{}{
		"id_user":      id_user,
		"id_campaign":  id_campaign,
		"value_action": action,
	}
	err := model.InsertDataToTable("user_action_campaign", actionData)
	if err != nil {
		lib.ReturnError(c, err)
		return
	}

	lib.Json(c, 200, "success", gin.H{})
}
func ActionCampaign(c *gin.Context) {
	id_user := c.PostForm("id_user")
	id_campaign := c.PostForm("id_campaign")
	action := "2"

	actionData := map[string]interface{}{
		"id_user":      id_user,
		"id_campaign":  id_campaign,
		"value_action": action,
	}
	err := model.InsertDataToTable("user_action_campaign", actionData)
	if err != nil {
		lib.ReturnError(c, err)
		return
	}

	lib.Json(c, 200, "success", gin.H{})
}
func ShareCampaign(c *gin.Context) {
	id_user := c.PostForm("id_user")
	id_campaign := c.PostForm("id_campaign")
	action := "3"

	actionData := map[string]interface{}{
		"id_user":      id_user,
		"id_campaign":  id_campaign,
		"value_action": action,
	}
	err := model.InsertDataToTable("user_action_campaign", actionData)
	if err != nil {
		lib.ReturnError(c, err)
		return
	}

	lib.Json(c, 200, "success", gin.H{})
}

func GetFollowCampaign(c *gin.Context) {
	id_user := c.Query("userid")
	filter := map[string]interface{}{
		"id_user":      id_user,
		"action_value": "0",
	}
	data, err := model.GetActionCampaignByIDAndValueAction(filter)
	if err != nil {
		lib.ReturnError(c, err)
		return
	}
	lib.Json(c, 200, "success", data)
}
func GetLikeCampaign(c *gin.Context) {
	id_user := c.Query("userid")
	filter := map[string]interface{}{
		"id_user":      id_user,
		"action_value": "1",
	}
	data, err := model.GetActionCampaignByIDAndValueAction(filter)
	if err != nil {
		lib.ReturnError(c, err)
		return
	}
	lib.Json(c, 200, "success", data)
}
func GetRememberCampaign(c *gin.Context) {
	id_user := c.Query("userid")
	filter := map[string]interface{}{
		"id_user":      id_user,
		"action_value": "2",
	}
	data, err := model.GetActionCampaignByIDAndValueAction(filter)
	if err != nil {
		lib.ReturnError(c, err)
		return
	}
	lib.Json(c, 200, "success", data)
}
func GetActionCampaign(c *gin.Context) {
	id_user := c.Query("userid")
	filter := map[string]interface{}{
		"id_user":      id_user,
		"action_value": "3",
	}
	data, err := model.GetActionCampaignByIDAndValueAction(filter)
	if err != nil {
		lib.ReturnError(c, err)
		return
	}
	lib.Json(c, 200, "success", data)
}
func GetShareCampaign(c *gin.Context) {
	id_user := c.Query("userid")
	filter := map[string]interface{}{
		"id_user":      id_user,
		"action_value": "4",
	}
	data, err := model.GetActionCampaignByIDAndValueAction(filter)
	if err != nil {
		lib.ReturnError(c, err)
		return
	}
	lib.Json(c, 200, "success", data)
}

func TambahKomentar(c *gin.Context) {
	id_user := c.PostForm("id_user")
	id_campaign := c.PostForm("id_campaign")
	isi_komentar := c.PostForm("isi_komentar")

	dataKomentar := map[string]interface{}{
		"id_user":      id_user,
		"id_campaign":  id_campaign,
		"isi_komentar": isi_komentar,
	}

	err := model.InsertDataToTable("komentar", dataKomentar)
	if err != nil {
		lib.ReturnError(c, err)
		return
	}
	lib.Json(c, 200, "success", gin.H{})
}

func DetailCampaign(c *gin.Context) {
	id_campaign := c.Query("idcampaign")
	datacampaign, err := model.DetailCampaign(id_campaign)
	if err != nil {
		lib.ReturnError(c, err)
		return
	}

	komentar, err := model.GetKomentarCampaign(id_campaign)
	if err != nil {
		lib.ReturnError(c, err)
		return
	}

	dataReturn := map[string]interface{}{
		"campaign": datacampaign,
		"komentar": komentar,
	}
	lib.Json(c, 200, "succes", dataReturn)
}
