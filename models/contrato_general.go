package models

import (
	"time"

)

type ContratoGeneral struct {
	Id                           string              `orm:"pk;column(numero_contrato)"`
	Vigencia                     int                 `orm:"column(vigencia);null"`
	ObjetoContrato               string              `orm:"column(objeto_contrato);null"`
	PlazoEjecucion               int                 `orm:"column(plazo_ejecucion);null"`
	FormaPago                    *Parametros         `orm:"column(forma_pago);rel(fk);null"`
	OrdenadorGasto               *ArgoOrdenadores    `orm:"column(ordenador_gasto);rel(fk);null"`
	ClausulaRegistroPresupuestal bool                `orm:"column(clausula_registro_presupuestal);null"`
	SedeSolicitante              string              `orm:"column(sede_solicitante);null"`
	DependenciaSolicitante       string              `orm:"column(dependencia_solicitante);null"`
	NumeroSolicitudNecesidad     int                 `orm:"column(numero_solicitud_necesidad);null"`
	NumeroCdp                    int                 `orm:"column(numero_cdp);null"`
	Contratista                  *InformacionProveedor  `orm:"column(contratista);rel(fk)"`
	UnidadEjecucion              *Parametros         `orm:"column(unidad_ejecucion);rel(fk);null"`
	ValorContrato                string             `orm:"column(valor_contrato);null"`
	Justificacion                string              `orm:"column(justificacion);null"`
	DescripcionFormaPago         string              `orm:"column(descripcion_forma_pago);null"`
	Condiciones                  string              `orm:"column(condiciones);null"`
	UnidadEjecutora              *UnidadEjecutora    `orm:"column(unidad_ejecutora);rel(fk);null"`
	FechaRegistro                time.Time           `orm:"column(fecha_registro);type(date);null"`
	TipologiaContrato            int                 `orm:"column(tipologia_contrato);null"`
	TipoCompromiso               int                 `orm:"column(tipo_compromiso);null"`
	ModalidadSeleccion           int                 `orm:"column(modalidad_seleccion);null"`
	Procedimiento                int                 `orm:"column(procedimiento);null"`
	RegimenContratacion          int                 `orm:"column(regimen_contratacion);null"`
	TipoGasto                    int                 `orm:"column(tipo_gasto);null"`
	TemaGastoInversion           int                 `orm:"column(tema_gasto_inversion);null"`
	OrigenPresupueso             int                 `orm:"column(origen_presupueso);null"`
	OrigenRecursos               int                 `orm:"column(origen_recursos);null"`
	TipoMoneda                   int                 `orm:"column(tipo_moneda);null"`
	ValorContratoMe              float64             `orm:"column(valor_contrato_me);null"`
	ValorTasaCambio              float64             `orm:"column(valor_tasa_cambio);null"`
	TipoControl                  int                 `orm:"column(tipo_control);null"`
	Observaciones                string              `orm:"column(observaciones);null"`
	Supervisor                   *SupervisorContrato `orm:"column(supervisor);rel(fk);null"`
	ClaseContratista             int                 `orm:"column(clase_contratista);null"`
	Convenio                     string              `orm:"column(convenio);null"`
	NumeroConstancia             int                 `orm:"column(numero_constancia);null"`
	Estado                       bool                `orm:"column(estado);null"`
	ResgistroPresupuestal        int                 `orm:"column(resgistro_presupuestal);null"`
	TipoContrato                 *TipoContrato       `orm:"column(tipo_contrato);rel(fk);null"`
	LugarEjecucion               *LugarEjecucion     `orm:"column(lugar_ejecucion);rel(fk);null"`
	//ActaInicio                   *ActaInicio   `orm:"reverse(one)"`
}
