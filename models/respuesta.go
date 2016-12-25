package models

type Respuesta struct {
	Id	           int
	Nombre_Cont         string
	NumDocumento         int64
	Valor_bruto string
	Valor_neto  string
	Descuentos *[]Descuentos
	Novedades *[]NovedadAplicada
}
type FormatoPreliqu struct {
	//Contrato   *ContratoGeneral
	Respuesta *Respuesta
}
