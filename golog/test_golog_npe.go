package golog

import (
	"fmt"
	"strconv"
	models "titan_api_mid/models"

	. "github.com/mndrix/golog"
)

func CargarReglasPE(reglas string,pensionado models.InformacionPensionado) (rest []models.Respuesta) {
	 fmt.Println("pensionado")
	 	var resultado []models.Respuesta
	 temp := models.Respuesta{}
	 var lista_descuentos []models.ConceptosResumen
	 var cedulaProveedor = strconv.Itoa(pensionado.InformacionProveedor)
	 var valorpension= strconv.Itoa(pensionado.ValorPensionAsignada)
	 reglas = reglas + "valor_mesada(" + cedulaProveedor + "," + valorpension + ",1).";
	 fmt.Println(reglas)
	m := NewMachine().Consult(reglas)

	valor := m.ProveAll("aporte_fondoSoli(" + cedulaProveedor +",W).")

	for _, solution := range valor {
		Valor, _ := strconv.ParseFloat(fmt.Sprintf("%s", solution.ByName_("W")), 64)
		temp_conceptos := models.ConceptosResumen{Nombre: "pagoBruto",
			Valor: fmt.Sprintf("%.0f", Valor),
		}
		lista_descuentos = append(lista_descuentos, temp_conceptos)
		temp.Conceptos = &lista_descuentos
		resultado = append(resultado, temp)
}

	return resultado
}
