package models

type Nomina struct {
	Id          int              `orm:"column(id);pk"`
	Vinculacion *TipoVinculacion `orm:"column(vinculacion);rel(fk)"`
	Nombre      string           `orm:"column(nombre)"`
	Descripcion string           `orm:"column(descripcion)"`
	TipoNomina  string           `orm:"column(tipo_nomina)"`
	Estado      string           `orm:"column(estado)"`
	Periodo     string           `orm:"column(periodo);null"`
}
