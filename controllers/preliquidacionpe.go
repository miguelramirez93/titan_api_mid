package controllers

import (
	"fmt"
	"titan_api_mid/golog"
	"titan_api_mid/models"

	"github.com/astaxie/beego"
)

// PreliquidacionpeController operations for Preliquidacionpe
type PreliquidacionpeController struct {
	beego.Controller
}

func (c *PreliquidacionpeController) Preliquidar(datos *models.DatosPreliquidacion, reglasbase string) (res []models.Respuesta) {
	//	var predicados []models.Predicado //variable para inyectar reglas
	var resumen_preliqu []models.Respuesta
	var pensionados []models.InformacionPensionado // arreglo de informacion_pensionado

	//	var reglasinyectadas string
	//var reglas string
	fmt.Println("afdsgfd")
	for i := 0; i < len(datos.PersonasPreLiquidacion); i++ {
		fmt.Println("afdsgfasfdghdsfsgfhsdd")
		filtrodatos := models.InformacionPensionado{Id: datos.PersonasPreLiquidacion[i].IdPersona}
		if err := sendJson("http://"+beego.AppConfig.String("Urlcrud")+":"+beego.AppConfig.String("Portcrud")+"/"+beego.AppConfig.String("Nscrud")+"/informacion_pensionado", "POST", &pensionados, &filtrodatos); err == nil {
			temp := golog.CargarReglasPE(reglasbase, pensionados[i])
			fmt.Println(temp)
		}
	}
	return resumen_preliqu
}
