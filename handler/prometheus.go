package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/kexirong/repeater/models"
	api "github.com/kexirong/repeater/prometheus_api"
	v1 "github.com/kexirong/repeater/prometheus_api/v1"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/common/model"
	"gorm.io/gorm"
)

func ReadPrometheusDataSource(c echo.Context) error {
	id := c.Param("id")
	if id != "" {
		var pds models.PrometheusDataSource
		if err := models.DB.First(&pds, id).Error; err != nil {
			return err
		}
		return c.JSON(http.StatusOK, pds)
	}

	var pdss []models.PrometheusDataSource
	if err := models.DB.Find(&pdss).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, pdss)
}

func AddPrometheusDataSource(c echo.Context) error {

	var pds models.PrometheusDataSource

	if err := c.Bind(&pds); err != nil {
		return err
	}

	if err := models.DB.Create(&pds).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, echo.Map{"id": pds.ID})

}

func UpdatePrometheusDataSource(c echo.Context) error {
	param := c.Param("id")

	var pds models.PrometheusDataSource
	result := models.DB.First(&pds, param)
	if result.RowsAffected == 0 {
		return result.Error
	}

	err := c.Bind(&pds)
	if err != nil {
		return err
	}

	if err := models.DB.Save(&pds).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, echo.Map{"id": pds.ID})

}

func DeletePrometheusDataSource(c echo.Context) error {
	param := c.Param("id")

	var pds models.PrometheusDataSource
	result := models.DB.First(&pds, param)
	if result.RowsAffected == 0 {
		return result.Error
	}

	if err := models.DB.Unscoped().Delete(&pds).Error; err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func ReadPrometheusLabelSetTarget(c echo.Context) error {
	source_name := c.QueryParam("source")
	var pfds models.PrometheusFakeDataSource
	err := models.DB.First(&pfds, "name = ?", source_name).Error
	if err == nil {
		return c.JSON(http.StatusOK, pfds.LabelSetList)
	} else if err != gorm.ErrRecordNotFound {
		return err
	}
	var pds models.PrometheusDataSource
	if err := models.DB.First(&pds, "name = ?", source_name).Error; err != nil {
		return err
	}
	client, err := api.NewClient(pds.BaseUrl)
	if err != nil {
		return err
	}

	v1api := v1.NewAPI(client)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if pds.Targets == "" {
		pds.Targets = v1.EndpointTargets
	}
	tr, err := v1api.Targets(ctx, pds.Targets)
	if err != nil {
		return err
	}

	var labelSets []model.LabelSet
	for _, at := range tr.Active {
		labelSets = append(labelSets, at.Labels)
	}
	return c.JSON(http.StatusOK, labelSets)
}

func ReadPrometheusAlerts(c echo.Context) error {

	var pdss []models.PrometheusDataSource
	if err := models.DB.Find(&pdss).Error; err != nil {
		return err
	}
	var alerts []v1.Alert
	for _, pds := range pdss {
		client, err := api.NewClient(pds.BaseUrl)
		if err != nil {
			return err
		}

		v1api := v1.NewAPI(client)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if pds.Alerts == "" {
			continue
		}
		ar, err := v1api.Alerts(ctx, pds.Alerts)
		if err != nil {
			return err
		}
		alerts = append(alerts, ar...)
	}

	return c.JSON(http.StatusOK, alerts)
}

func ReadPrometheusRuleSeverity(c echo.Context) error {
	source_name := c.QueryParam("source")

	var pfds models.PrometheusFakeDataSource
	err := models.DB.First(&pfds, "name = ?", source_name).Error
	if err == nil {
		return c.JSON(http.StatusOK, echo.Map{"severity_set": pfds.SeveritySet})
	} else if err != gorm.ErrRecordNotFound {
		return err
	}

	var pds models.PrometheusDataSource
	if err := models.DB.First(&pds, "name = ?", source_name).Error; err != nil {
		return err
	}

	client, err := api.NewClient(pds.BaseUrl)
	if err != nil {
		return err
	}

	v1api := v1.NewAPI(client)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if pds.Rules == "" {
		pds.Rules = v1.EndpointRules
	}
	tr, err := v1api.Rules(ctx, pds.Rules)

	if err != nil {
		return err
	}
	//
	var severitySet = make(map[string]struct{})
	for _, group := range tr.Groups {
		for _, rule := range group.Rules {


			v, ok := rule.(v1.AlertingRule)
			if !ok {
				continue
			}


			if severity, ok := v.Labels["severity"]; ok {
				severitySet[string(severity)] = struct{}{}
			}
		}
	}
	var severityList []string
	for severity := range severitySet {
		severityList = append(severityList, severity)
	}
	return c.JSON(http.StatusOK, echo.Map{"severity_set": severityList})
}

func ReadPrometheusFakeDataSource(c echo.Context) error {
	id := c.Param("id")
	if id != "" {
		var pfds models.PrometheusFakeDataSource
		if err := models.DB.First(&pfds, id).Error; err != nil {
			return err
		}
		return c.JSON(http.StatusOK, pfds)
	}

	var pfdss []models.PrometheusFakeDataSource
	if err := models.DB.Find(&pfdss).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, pfdss)
}

func AddPrometheusFakeDataSource(c echo.Context) error {

	var pfds models.PrometheusFakeDataSource

	if err := c.Bind(&pfds); err != nil {
		return err
	}

	if err := models.DB.Create(&pfds).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, echo.Map{"id": pfds.ID})

}

func UpdatePrometheusFakeDataSource(c echo.Context) error {
	param := c.Param("id")

	var pfds models.PrometheusFakeDataSource
	result := models.DB.First(&pfds, param)
	if result.RowsAffected == 0 {
		return result.Error
	}

	err := c.Bind(&pfds)
	if err != nil {
		return err
	}

	if err := models.DB.Save(&pfds).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, echo.Map{"id": pfds.ID})

}

func DeletePrometheusFakeDataSource(c echo.Context) error {
	param := c.Param("id")

	var pfds models.PrometheusFakeDataSource
	result := models.DB.First(&pfds, param)
	if result.RowsAffected == 0 {
		return result.Error
	}

	if err := models.DB.Delete(&pfds).Error; err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
