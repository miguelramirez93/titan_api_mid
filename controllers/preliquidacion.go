package controllers

import (
	"github.com/astaxie/beego"
	"titan_api_mid/models"
	"encoding/json"
	"fmt"
)

// PreliquidacionController operations for Preliquidacion
type PreliquidacionController struct {
	beego.Controller
}

// URLMapping ...
func (c *PreliquidacionController) URLMapping() {
	c.Mapping("Preliquidar", c.Preliquidar)
}

// Post ...
// @Title Create
// @Description create Preliquidacion
// @Param	body		body 	models.Preliquidacion	true		"body for Preliquidacion content"
// @Success 201 {object} models.Preliquidacion
// @Failure 403 body is empty
// @router / [post]
func (c *PreliquidacionController) Preliquidar() {
	var v []models.DatosPreliquidacion
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if(v[0].Nomina == "HC"){
			var n *PreliquidacionHcController
			resumen := n.Preliquidar(&v[0])
			c.Data["json"] = resumen
		  c.ServeJSON()
		}
	}else{
		fmt.Println("error: ", err)
	}
}

func FormatoReglas(v []models.Predicado)(reglas string){
	var arregloReglas = make([]string, len(v))
	reglas = ""
	//var respuesta []models.FormatoPreliqu
	for i := 0; i < len(v); i++ {
		arregloReglas[i] = v[i].Nombre
	}

	for i := 0; i < len(arregloReglas); i++ {
		reglas = reglas + arregloReglas[i] + " "
	}
	return
}
