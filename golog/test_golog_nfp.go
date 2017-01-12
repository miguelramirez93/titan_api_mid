package golog

import (
  "fmt"
  "strconv"
  "titan_api_mid/models"
  . "github.com/mndrix/golog"
)

func CargarReglasFP(reglas string, informacion_cargo []models.FuncionarioCargo, dias_laborados float64,periodo string)  (rest []models.Respuesta){

      var resultado []models.Respuesta
      temp := models.Respuesta{}
      var lista_descuentos []models.ConceptosResumen
      asignacion_basica_string := strconv.Itoa(informacion_cargo[0].Asignacion_basica)
      id_cargo_string := strconv.Itoa(informacion_cargo[0].Id)
      dias_laborados_string := strconv.Itoa(int(dias_laborados))
      var total_devengado float64;

      m := NewMachine().Consult(reglas)
      valor_salario := m.ProveAll("sb("+asignacion_basica_string+",15,V).")
      for _, solution := range valor_salario {
        Valor,_ := strconv.ParseFloat(fmt.Sprintf("%s", solution.ByName_("V")), 64)
        total_devengado  = total_devengado + Valor;
        temp_conceptos := models.ConceptosResumen {Nombre : "salarioBase" ,
                                                   Valor : fmt.Sprintf("%.0f", Valor),
                                             }

        codigo := m.ProveAll("codigo_concepto("+temp_conceptos.Nombre+",C).")

        for _, cod := range codigo{
            temp_conceptos.Id , _ = strconv.Atoi(fmt.Sprintf("%s", cod.ByName_("C")))

        }

        lista_descuentos = append(lista_descuentos,temp_conceptos)
        temp.Conceptos = &lista_descuentos
        resultado = append(resultado,temp)

      }



      valor_gastos_representacion := m.ProveAll("gr("+asignacion_basica_string+",15,2016,"+id_cargo_string+",V).")
      for _, solution := range   valor_gastos_representacion {
        Valor,_ := strconv.ParseFloat(fmt.Sprintf("%s", solution.ByName_("V")), 64)
        total_devengado  = total_devengado + Valor;
        temp_conceptos := models.ConceptosResumen {Nombre : "gastosRep" ,
                                                   Valor : fmt.Sprintf("%.0f", Valor),
                                             }

       codigo := m.ProveAll("codigo_concepto("+temp_conceptos.Nombre+",C).")

        for _, cod := range codigo{
              temp_conceptos.Id , _ = strconv.Atoi(fmt.Sprintf("%s", cod.ByName_("C")))

        }
        lista_descuentos = append(lista_descuentos,temp_conceptos)
        temp.Conceptos = &lista_descuentos
        resultado = append(resultado,temp)

      }


      valor_prima_antiguedad := m.ProveAll("prima_ant("+asignacion_basica_string+",15,2016,"+dias_laborados_string+",V).")
      for _, solution := range     valor_prima_antiguedad {
        Valor,_ := strconv.ParseFloat(fmt.Sprintf("%s", solution.ByName_("V")), 64)
        total_devengado  = total_devengado + Valor;
        temp_conceptos := models.ConceptosResumen {Nombre : "primaAnt" ,
                                                   Valor : fmt.Sprintf("%.0f", Valor),
                                             }
       codigo := m.ProveAll("codigo_concepto("+temp_conceptos.Nombre+",C).")

        for _, cod := range codigo{
          temp_conceptos.Id , _ = strconv.Atoi(fmt.Sprintf("%s", cod.ByName_("C")))

        }

        lista_descuentos = append(lista_descuentos,temp_conceptos)
        temp.Conceptos = &lista_descuentos
        resultado = append(resultado,temp)

      }

      valor_bonificacion_servicios := m.ProveAll("bon_ser("+asignacion_basica_string+",15,2016,"+dias_laborados_string+","+id_cargo_string+",V).")
      for _, solution := range     valor_bonificacion_servicios {
        Valor,_ := strconv.ParseFloat(fmt.Sprintf("%s", solution.ByName_("V")), 64)
        total_devengado  = total_devengado + Valor;
        temp_conceptos := models.ConceptosResumen {Nombre : "bonServ" ,
                                                   Valor : fmt.Sprintf("%.0f", Valor),
                                             }

        codigo := m.ProveAll("codigo_concepto("+temp_conceptos.Nombre+",C).")

        for _, cod := range codigo{
              temp_conceptos.Id , _ = strconv.Atoi(fmt.Sprintf("%s", cod.ByName_("C")))
        }

        lista_descuentos = append(lista_descuentos,temp_conceptos)
        temp.Conceptos = &lista_descuentos
        resultado = append(resultado,temp)

      }

      valor_prima_secretarial := m.ProveAll("prima_secretarial("+asignacion_basica_string+",2016,"+dias_laborados_string+",V).")
      for _, solution := range     valor_prima_secretarial {
        Valor,_ := strconv.ParseFloat(fmt.Sprintf("%s", solution.ByName_("V")), 64)
        total_devengado  = total_devengado + Valor;
        temp_conceptos := models.ConceptosResumen {Nombre : "primaSecr" ,
                                                   Valor : fmt.Sprintf("%.0f", Valor),
                                                  }
      codigo := m.ProveAll("codigo_concepto("+temp_conceptos.Nombre+",C).")

       for _, cod := range codigo{
          temp_conceptos.Id , _ = strconv.Atoi(fmt.Sprintf("%s", cod.ByName_("C")))

        }
        lista_descuentos = append(lista_descuentos,temp_conceptos)
        temp.Conceptos = &lista_descuentos
        resultado = append(resultado,temp)

      }

      total_devengado_string := strconv.Itoa(int(total_devengado))
      fmt.Println(total_devengado_string)
      valor_salud := m.ProveAll("salud_fun("+total_devengado_string+",2016,V).")
      for _, solution := range     valor_salud {
        Valor,_ := strconv.ParseFloat(fmt.Sprintf("%s", solution.ByName_("V")), 64)
        temp_conceptos := models.ConceptosResumen {Nombre : "salud" ,
                                                   Valor : fmt.Sprintf("%.0f", Valor),
                                             }
        codigo := m.ProveAll("codigo_concepto("+temp_conceptos.Nombre+",C).")

        for _, cod := range codigo{
          temp_conceptos.Id , _ = strconv.Atoi(fmt.Sprintf("%s", cod.ByName_("C")))
          }
        lista_descuentos = append(lista_descuentos,temp_conceptos)
        temp.Conceptos = &lista_descuentos
        resultado = append(resultado,temp)

      }


      return resultado
}
