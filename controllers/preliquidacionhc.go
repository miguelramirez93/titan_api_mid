package controllers

import (
	"github.com/astaxie/beego"
	"titan_api_mid/models"
	"strconv"
	"titan_api_mid/golog"
	"fmt"
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
func (c *PreliquidacionHcController) Preliquidar(datos *models.DatosPreliquidacion) (res []models.Respuesta) {
	//declaracion de variables
	dominio := "1"
	var v []models.Predicado //carga reglas del ruler
	var predicados []models.Predicado //variable para inyectar reglas
	var datos_contrato []models.ActaInicio
	//var datos_novedades []models.ConceptoPorPersona
	var resumen_preliqu []models.Respuesta
	var meses_contrato float64
	//var periodo_liquidacion float64
	var reglasbase string
	var reglasinyectadas string
	var reglas string
	var filtrodatos string
	var idDetaPre interface{}
	//-----------------------
	//carga de reglas desde el ruler
	reglasbase = CargarReglasBase(v , dominio)//funcion general para dar formato a reglas cargadas desde el ruler
	//-----------------------------
	//carga de informacion de los empleados a partir del id de persona Natural (en este momento id proveedor)

	for i := 0; i < len(datos.PersonasPreLiquidacion); i++ {
		filtrodatos = "NumeroContrato.Contratista.Id:"+strconv.Itoa(datos.PersonasPreLiquidacion[i].IdPersona)
		if err := getJson("http://"+beego.AppConfig.String("Urlcrud")+":"+beego.AppConfig.String("Portcrud")+"/"+beego.AppConfig.String("Nscrud")+"/acta_inicio?limit=1&query="+filtrodatos, &datos_contrato); err == nil{

			a,m,d := diff(datos_contrato[0].FechaInicio,datos_contrato[0].FechaFin)
			//al,ml,dl := diff(datos.FechaInicio,datos.FechaFin)
			meses_contrato = (float64(a*12))+float64(m)+(float64(d)/30)
			//periodo_liquidacion = (float64(al*12))+float64(ml)+(float64(dl)/30)
			predicados = append(predicados,models.Predicado{Nombre:"valor_contrato("+strconv.Itoa(datos_contrato[0].NumeroContrato.Contratista.Id)+","+strconv.FormatFloat(datos_contrato[0].NumeroContrato.ValorContrato, 'f', -1, 64)+"). "} )
			predicados = append(predicados,models.Predicado{Nombre:"duracion_contrato("+strconv.Itoa(datos_contrato[0].NumeroContrato.Contratista.Id)+","+strconv.FormatFloat(meses_contrato, 'f', -1, 64)+","+datos.Preliquidacion.Nomina.Periodo+"). "} )
			reglasinyectadas = FormatoReglas(predicados)
			reglasinyectadas = reglasinyectadas + CargarNovedadesPersona(datos_contrato[0].NumeroContrato.Contratista.Id, datos)
			reglas =  reglasinyectadas + reglasbase
			fmt.Println("num: ",reglas)
			temp := golog.CargarReglas(reglas,datos.Preliquidacion.Nomina.Periodo)

			resultado := temp[len(temp)-1]
			resultado.NumDocumento = datos_contrato[0].NumeroContrato.Contratista.NumDocumento
			//se guardan los conceptos calculados en la nomina
			for _, descuentos := range *resultado.Conceptos{
				valor,_ := strconv.ParseInt(descuentos.Valor,10,64)
				detallepreliqu := models.DetallePreliquidacion{Concepto: &models.Concepto{Id : descuentos.Id} , Persona: datos_contrato[0].NumeroContrato.Contratista.Id,Preliquidacion : datos.Preliquidacion.Id, ValorCalculado:valor  }
				if err := sendJson("http://"+beego.AppConfig.String("Urlcrud")+":"+beego.AppConfig.String("Portcrud")+"/"+beego.AppConfig.String("Nscrud")+"/detalle_preliquidacion","POST",&idDetaPre ,&detallepreliqu); err == nil {

				}else{
					beego.Debug("error: ", err)
				}
			}
			//------------------------------------------------
			resumen_preliqu = append(resumen_preliqu, resultado)
				fmt.Println("append: ", resumen_preliqu[0].Nombre_Cont)
			predicados = nil;
			reglas = ""
			reglasinyectadas = ""
		}else{
			fmt.Println("error: ", err)
		}

	}
		//-----------------------------
		return resumen_preliqu
}
