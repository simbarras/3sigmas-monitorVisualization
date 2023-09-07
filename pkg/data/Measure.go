package data

import (
	"fmt"
	"time"
)

type Measure struct {
	Date        time.Time
	Value       float64
	Temperature float64
	Captor      string
	Sensor      string
}

func (m Measure) String() string {
	return m.Date.Format("2006-01-02T15:04:05Z") + " " + m.Captor + " " + m.Sensor + " " + fmt.Sprintf("%f", m.Value) + " " + fmt.Sprintf("%f", m.Temperature)
}
