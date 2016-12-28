package models

import (
	"time"
)

type DatosPreliquidacion struct {
	IdPreliquidacion 			int
	Nomina           			int
	Periodo								float64
	FechaInicio      			time.Time
	FechaFin         			time.Time
	PersonasPreLiquidacion		[]PersonasPreliquidacion
}
