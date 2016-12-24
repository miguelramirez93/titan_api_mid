package controllers

import (
	"github.com/astaxie/beego"
	"titan_api_mid/models"
)

// PreliquidacionHcController operations for PreliquidacionHc
type PreliquidacionHcController struct {
	beego.Controller
}

// Post ...
// @Title Create
// @Description create PreliquidacionHc
// @Param	body		body 	models.PreliquidacionHc	true		"body for PreliquidacionHc content"
// @Success 201 {object} models.PreliquidacionHc
// @Failure 403 body is empty
// @router / [post]
func (c *PreliquidacionHcController) Preliquidar(*models.DatosPreliquidacion) {

}
