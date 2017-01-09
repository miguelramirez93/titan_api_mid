package controllers

import (
	"github.com/astaxie/beego"
	"titan_api_mid/models"
//	"strconv"
	"titan_api_mid/golog"
//	"fmt"

)

// PreliquidacionHcController operations for PreliquidacionHc
type PreliquidacionFpController struct {
	beego.Controller
}

// Post ...
// @Title Create
// @Description create PreliquidacionHc
// @Param	body		body 	models.PreliquidacionHc	true		"body for PreliquidacionHc content"
// @Success 201 {object} models.PreliquidacionHc
// @Failure 403 body is empty
// @router / [post]
func (c *PreliquidacionFpController) Preliquidar(datos *models.DatosPreliquidacion , reglasbase string) (res []models.Respuesta) {
	//declaracion de variables

	var resumen_preliqu []models.Respuesta
	for i := 0; i < len(datos.PersonasPreLiquidacion); i++ {
				 //consulta que envie ID de proveedor en datos y retorne el salario, para que sea enviado a CargarReglas
			 	temp := golog.CargarReglasFP(reglasbase,datos.Preliquidacion.Nomina.Periodo)
				resultado := temp[len(temp)-1]
				resultado.NumDocumento  =float64(datos.PersonasPreLiquidacion[i].IdPersona)
				resumen_preliqu = append(resumen_preliqu, resultado)
			}

			return resumen_preliqu

}
