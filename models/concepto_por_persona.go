package models

import (

	"time"

)

type ConceptoPorPersona struct {
	ValorNovedad  int64     `orm:"column(valor_novedad)"`
	EstadoNovedad int64     `orm:"column(estado_novedad)"`
	FechaDesde    time.Time `orm:"column(fecha_desde);type(date)"`
	FechaHasta    time.Time `orm:"column(fecha_hasta);type(date)"`
	NumCuotas     int64     `orm:"column(num_cuotas)"`
	Persona       int       `orm:"column(persona)"`
	Concepto      *Concepto `orm:"column(concepto);rel(fk)"`
	Nomina        int       `orm:"column(nomina)"`
	Id            int       `orm:"column(id);pk"`
}
