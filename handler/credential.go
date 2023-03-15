package handler

import (
	"github.com/kexirong/repeater/dingtalk"
	"github.com/kexirong/repeater/lark"
	"net/http"

	"github.com/kexirong/repeater/email"
	"github.com/kexirong/repeater/models"
	"github.com/kexirong/repeater/wecom"
	"github.com/labstack/echo/v4"
)

func ReadStmpCredentials(c echo.Context) error {

	var scs []models.StmpCredential
	if err := models.DB.Find(&scs).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, scs)
}

func AddStmpCredential(c echo.Context) error {

	var sc models.StmpCredential
	if err := c.Bind(&sc); err != nil {
		return err
	}

	if sc.Name == "" {
		sc.Name = "default"
	}

	if sc.Name != "default" {
		var count int64
		models.DB.Model(&models.StmpCredential{}).Where("name = ?", "default").Count(&count)
		if count == 0 {
			return echo.NewHTTPError(http.StatusBadRequest, "please create a default entry first")
		}
	}
	if err := models.DB.Create(&sc).Error; err != nil {
		return err
	}
	email.StmpCredentialRegist(sc)
	return c.JSON(http.StatusCreated, echo.Map{"id": sc.ID})

}

func DeleteStmpCredential(c echo.Context) error {
	param := c.Param("id")

	var sc models.StmpCredential
	result := models.DB.First(&sc, param)
	if result.RowsAffected == 0 {
		return result.Error
	}

	if err := models.DB.Unscoped().Delete(&sc).Error; err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)

}
// 企业微信操作接口
func ReadWecomCredentials(c echo.Context) error {

	var wcs []models.WecomCredential
	if err := models.DB.Find(&wcs).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, wcs)

}

func AddWecomCredential(c echo.Context) error {
	var wc models.WecomCredential
	if err := c.Bind(&wc); err != nil {
		return err
	}

	if wc.Name == "" {
		wc.Name = "default"
	}

	if wc.Name != "default" {
		var count int64
		models.DB.Model(&models.WecomCredential{}).Where("name = ?", "default").Count(&count)
		//if count == 0 {
		//	return echo.NewHTTPError(http.StatusBadRequest, "please create a default entry first")
		//}
	}

	if err := models.DB.Create(&wc).Error; err != nil {
		return err
	}
	wecom.WecomCredentialRegist(wc)
	return c.JSON(http.StatusCreated, echo.Map{"id": wc.ID})

}

func DeleteWecomCredential(c echo.Context) error {
	param := c.Param("id")

	var wc models.WecomCredential
	result := models.DB.First(&wc, param)
	if result.RowsAffected == 0 {
		return result.Error
	}

	if err := models.DB.Unscoped().Delete(&wc).Error; err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)

}
// 钉钉应用操作接口

func ReadDingCredentials(c echo.Context) error {

	var dcs []models.DingCredential
	if err := models.DB.Find(&dcs).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, dcs)

}

func AddDingCredential(c echo.Context) error {
	var dc models.DingCredential
	if err := c.Bind(&dc); err != nil {
		return err
	}

	if dc.Name == "" {
		dc.Name = "default"
	}

	//if dc.Name != "default" {
	//	var count int64
	//	models.DB.Model(&models.DingCredential{}).Where("name = ?", "default").Count(&count)
	//	if count == 0 {
	//		return echo.NewHTTPError(http.StatusBadRequest, "please create a default entry first")
	//	}
	//}

	if err := models.DB.Create(&dc).Error; err != nil {
		return err
	}
	//wecom.WecomCredentialRegist(dc)
	dingtalk.WecomCredentialRegist(dc)
	return c.JSON(http.StatusCreated, echo.Map{"id": dc.ID})

}


func DeleteDingCredential(c echo.Context) error {
	param := c.Param("id")

	var dc models.DingCredential
	result := models.DB.First(&dc, param)
	if result.RowsAffected == 0 {
		return result.Error
	}

	if err := models.DB.Unscoped().Delete(&dc).Error; err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)

}

// 飞书接口操作


func ReadLarkCredentials(c echo.Context) error {
	var lcs []models.LarkCredential
	if err := models.DB.Find(&lcs).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, lcs)

}

func AddLarkCredential(c echo.Context) error {
	var lc models.LarkCredential
	if err := c.Bind(&lc); err != nil {
		return err
	}

	if lc.Name == "" {
		lc.Name = "default"
	}

	//if dc.Name != "default" {
	//	var count int64
	//	models.DB.Model(&models.DingCredential{}).Where("name = ?", "default").Count(&count)
	//	if count == 0 {
	//		return echo.NewHTTPError(http.StatusBadRequest, "please create a default entry first")
	//	}
	//}

	if err := models.DB.Create(&lc).Error; err != nil {
		return err
	}
	//wecom.WecomCredentialRegist(dc)
	lark.LarkCredentialRegist(lc)
	return c.JSON(http.StatusCreated, echo.Map{"id": lc.ID})

}

func DeleteLarkCredential(c echo.Context) error {
	param := c.Param("id")

	var lc models.LarkCredential
	result := models.DB.First(&lc, param)
	if result.RowsAffected == 0 {
		return result.Error
	}

	if err := models.DB.Unscoped().Delete(&lc).Error; err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)

}
