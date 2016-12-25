package models



type Descuentos struct {
	Id                    int                    `pk;orm:"column(id)"`
	Nombre                string                 `orm:"column(nombre);null"`
	Base                  string               `orm:"column(base);null"`
	Valor                 string                `orm:"column(valor);null"`
	DetallePreliquidacion *DetallePreliquidacion `orm:"column(detalle_preliquidacion);rel(fk)"`
}
