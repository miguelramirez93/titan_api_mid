package controllers

import (
	"github.com/astaxie/beego"
	"titan_api_mid/models"
	"encoding/json"
	"fmt"
	"strconv"

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
	var v models.DatosPreliquidacion
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
			  if( v.Preliquidacion.Nomina.TipoNomina.Nombre == "HC"){
						var n *PreliquidacionHcController
						resumen := n.Preliquidar(&v)
						//pr := CargarNovedadesPersona(v[0].PersonasPreLiquidacion[0].IdPersona,&v[0])
						//fmt.Println("prueba: ", pr)
						c.Data["json"] = resumen
					  c.ServeJSON()
			}

		}else{
			fmt.Println("error2: ", err)
		}

}
func CargarReglasBase (v []models.Predicado , dominio string)(reglas string){
	//carga de reglas desde el ruler
	var reglasbase string = ``
	var datos_conceptos []models.Concepto
	if err := getJson("http://"+beego.AppConfig.String("Urlcrud")+":"+beego.AppConfig.String("Portcrud")+"/"+beego.AppConfig.String("Nscrud")+"/concepto?limit=0", &datos_conceptos); err == nil{
		for _ , datos := range datos_conceptos {
				reglasbase = reglasbase + `codigo_concepto(`+datos.NombreConcepto+`,`+strconv.Itoa(datos.Id)+`).` + "\n"
		}
	}else{

	}

	if err := getJson("http://"+beego.AppConfig.String("Urlruler")+":"+beego.AppConfig.String("Portruler")+"/"+beego.AppConfig.String("Nsruler")+"/predicado?limit=0&query=Dominio.Id:"+dominio, &v); err == nil {
		reglasbase = reglasbase + FormatoReglas(v)//funcion general para dar formato a reglas cargadas desde el ruler
	}else{

	}

	//-----------------------------
	return reglasbase
}

func FormatoReglas(v []models.Predicado)(reglas string){
	var arregloReglas = make([]string, len(v))
	reglas = ""
	//var respuesta []models.FormatoPreliqu
	for i := 0; i < len(v); i++ {
		arregloReglas[i] = v[i].Nombre
	}

	for i := 0; i < len(arregloReglas); i++ {
		reglas = reglas + arregloReglas[i] + "\n"
	}
	return
}

func CargarNovedadesPersona(id_persona int, datos_preliqu *models.DatosPreliquidacion)(reglas string){
	//formato de las fechas para el rango de validez de la(s) novedades
	y1, M1, d1 := datos_preliqu.Preliquidacion.FechaInicio.Date()
	y2, M2, d2 := datos_preliqu.Preliquidacion.FechaFin.Date()


	fechadesde := strconv.Itoa(int(y1))+"-"+strconv.Itoa(int(M1))+"-"+strconv.Itoa(int(d1))
	fechahasta := strconv.Itoa(int(y2))+"-"+strconv.Itoa(int(M2))+"-"+strconv.Itoa(int(d2))
	
	filtrodatos := "Persona:"+strconv.Itoa(id_persona)+",FechaDesde__gte:"+fechadesde+",FechaHasta__lte:"+fechahasta+",EstadoNovedad:1,Nomina:"+strconv.Itoa(datos_preliqu.Preliquidacion.Nomina.Id)
	//-----------------------------------------------------------------

	//consulta de la(s) novedades que pueda tener la persona para la pre-liquidacion
	var v []models.ConceptoPorPersona
	reglas = "" //inicializacion de la variable donde se inyectaran las novedades como reglas
	if err := getJson("http://"+beego.AppConfig.String("Urlcrud")+":"+beego.AppConfig.String("Portcrud")+"/"+beego.AppConfig.String("Nscrud")+"/concepto_por_persona?limit=0&query="+filtrodatos, &v); err == nil{
		if(v != nil){
			for i := 0; i < len(v); i++ {
				reglas = reglas + "concepto("+strconv.Itoa(id_persona)+","+v[i].Concepto.Naturaleza+", "+v[i].Tipo+", "+v[i].Concepto.NombreConcepto+", "+strconv.FormatFloat(v[i].ValorNovedad, 'f', -1, 64)+", "+datos_preliqu.Preliquidacion.Nomina.Periodo+"). "+"\n"
			}

		}

	}

	//------------------------------------------------------------------------------
	return reglas

}
