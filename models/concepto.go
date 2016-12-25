package models


type Concepto struct {
	Id             int           `orm:"column(id);pk"`
	NombreConcepto string        `orm:"column(nombre_concepto)"`
	TipoConcepto   *TipoConcepto `orm:"column(tipo_concepto);rel(fk)"`
}
