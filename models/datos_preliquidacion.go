package models

import (
	"time"
)

type DatosPreliquidacion struct {
	IdPreliquidacion 			int
	Nomina           			string
	FechaInicio      			time.Time
	FechaFin         			time.Time
	PersonasPreLiquidacion		*[]PersonasPreliquidacion
}
