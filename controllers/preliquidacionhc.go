package controllers

import (
	"github.com/astaxie/beego"
	"titan_api_mid/models"
	"strconv"
	"titan_api_mid/golog"
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
func (c *PreliquidacionHcController) Preliquidar(datos *models.DatosPreliquidacion) {
	//declaracion de variables
	dominio := "1"
	var v []models.Predicado //carga reglas del ruler
	var predicados []models.Predicado //variable para inyectar reglas
	var datos_contrato []models.ActaInicio
	var meses_contrato float64
	//var periodo_liquidacion float64
	var reglasbase string
	var reglasinyectadas string
	var reglas string
	var filtrodatos string
	//-----------------------
	//carga de reglas desde el ruler
	if err := getJson("http://"+beego.AppConfig.String("Urlruler")+":"+beego.AppConfig.String("Portruler")+"/"+beego.AppConfig.String("Nsruler")+"/predicado?limit=0&"+dominio, &v); err == nil {
		reglasbase = FormatoReglas(v)//funcion general para dar formato a reglas cargadas desde el ruler
	}else{
		c.Data["json"] =  "no se pudo generar la preliquidacion motor de reglas no encontrado"
		c.ServeJSON()
	}
	//-----------------------------
	//carga de informacion de los empleados a partir del id de persona Natural (en este momento id proveedor)
	for i := 0; i < len(datos.PersonasPreLiquidacion); i++ {
		filtrodatos = "NumeroContrato.Contratista.Id:"+strconv.Itoa(datos.PersonasPreLiquidacion[i].IdPersona)
		if err := getJson("http://"+beego.AppConfig.String("Urlcrud")+":"+beego.AppConfig.String("Portcrud")+"/"+beego.AppConfig.String("Nscrud")+"/acta_inicio?limit=1&query="+filtrodatos, &datos_contrato); err == nil{
			a,m,d := diff(datos_contrato[0].FechaInicio,datos_contrato[0].FechaFin)
			//al,ml,dl := diff(datos.FechaInicio,datos.FechaFin)
			meses_contrato = (float64(a*12))+float64(m)+(float64(d)/30)
			//periodo_liquidacion = (float64(al*12))+float64(ml)+(float64(dl)/30)
			predicados = append(predicados,models.Predicado{Nombre:"valor_contrato('"+datos_contrato[0].NumeroContrato.Contratista.NomProveedor+"',"+datos_contrato[0].NumeroContrato.ValorContrato+"). "} )
			predicados = append(predicados,models.Predicado{Nombre:"duracion_contrato('"+datos_contrato[0].NumeroContrato.Contratista.NomProveedor+"',"+strconv.FormatFloat(meses_contrato, 'f', -1, 64)+","+strconv.FormatFloat(datos.Periodo, 'f', -1, 64)+"). "} )
			reglasinyectadas = FormatoReglas(predicados)
			reglas = reglasbase + reglasinyectadas
			temp := golog.CargarReglas(reglas,strconv.FormatFloat(datos.Periodo, 'f', -1, 64))
			predicados = nil;
			reglasinyectadas = ""
		}else{

		}
	}
		//-----------------------------
}
