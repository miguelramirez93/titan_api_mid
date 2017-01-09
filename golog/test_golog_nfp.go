package golog

import (
  "fmt"
  "strconv"
  "titan_api_mid/models"
  . "github.com/mndrix/golog"
)

func CargarReglasFP(reglas string, periodo string)  (rest []models.Respuesta){

      var resultado []models.Respuesta
      temp := models.Respuesta{}
      var lista_descuentos []models.ConceptosResumen

      m := NewMachine().Consult(reglas)
      valor_salario := m.ProveAll("sb(1000000,15,V).")
      for _, solution := range valor_salario {
        Valor,_ := strconv.ParseFloat(fmt.Sprintf("%s", solution.ByName_("V")), 64)

        temp_conceptos := models.ConceptosResumen {Nombre : "pagoBruto" ,
                                                   Valor : fmt.Sprintf("%.0f", Valor),
                                                 }
        lista_descuentos = append(lista_descuentos,temp_conceptos)
        temp.Conceptos = &lista_descuentos
        resultado = append(resultado,temp)

      }
      return resultado
}
