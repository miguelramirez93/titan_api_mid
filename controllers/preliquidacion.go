package controllers

import (
	"github.com/astaxie/beego"
	"titan_api_mid/models"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
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
		filtrodatos := "Id:"+strconv.Itoa(v[0].Nomina)
		var temp_nomina []models.Nomina
		if err := getJson("http://"+beego.AppConfig.String("Urlcrud")+":"+beego.AppConfig.String("Portcrud")+"/"+beego.AppConfig.String("Nscrud")+"/nomina?limit=1&query="+filtrodatos, &temp_nomina); err == nil{
			 if(temp_nomina == nil ){
					 c.Data["json"] = "no existe esa nomina..."
					 c.ServeJSON()
				 }else if( temp_nomina[0].Vinculacion.Nombre == "HC"){
						var n *PreliquidacionHcController
						v[0].FechaInicio = time.Now().Local()
						v[0].FechaFin =time.Now().Local()
						resumen := n.Preliquidar(&v[0])
						pr := CargarNovedadesPersona(v[0].PersonasPreLiquidacion[0].IdPersona,&v[0])
						fmt.Println("prueba: ", pr)
						c.Data["json"] = resumen
					  c.ServeJSON()
			}
		}else{
			fmt.Println("error1: ", err)
		}
		}else{
			fmt.Println("error2: ", err)
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

func CargarNovedadesPersona(id_persona int, datos_preliqu *models.DatosPreliquidacion)(reglas string){
	//formato de las fechas para el rango de validez de la(s) novedades
	y1, M1, d1 := datos_preliqu.FechaInicio.Date()
	y2, M2, d2 := datos_preliqu.FechaFin.Date()
	fechadesde := strconv.Itoa(int(y1))+"-"+strconv.Itoa(int(M1))+"-"+strconv.Itoa(int(d1))
	fechahasta := strconv.Itoa(int(y2))+"-"+strconv.Itoa(int(M2))+"-"+strconv.Itoa(int(d2))
	filtrodatos := "Persona:"+strconv.Itoa(id_persona)+",FechaDesde__gt:"+fechadesde+",FechaHasta__lt:"+fechahasta
	//-----------------------------------------------------------------

	//consulta de la(s) novedades que pueda tener la persona para la pre-liquidacion
	var v []models.ConceptoPorPersona
	reglas = "" //inicializacion de la variable donde se inyectaran las novedades como reglas
	if err := getJson("http://"+beego.AppConfig.String("Urlcrud")+":"+beego.AppConfig.String("Portcrud")+"/"+beego.AppConfig.String("Nscrud")+"/concepto_por_persona?limit=0&query="+filtrodatos, &v); err == nil{
		if(v != nil){
			reglas = "econtre"
		}

	}

	//------------------------------------------------------------------------------
	return reglas

}
