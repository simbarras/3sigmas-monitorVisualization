package data

import (
	"fmt"
	"time"
)

type Measure interface {
	String() string
	Tags() map[string]string
	Fields() map[string]interface{}
	Measurement() string
	Date() time.Time
}

// Measure for senseive data -------------------------------------------------------------------------------------------
type SenseiveMeasure struct {
	DateTime    time.Time
	Value       float64
	Temperature float64
	Captor      string
	Sensor      string
}

func (m SenseiveMeasure) String() string {
	return m.DateTime.Format("2006-01-02T15:04:05Z") + " " + m.Captor + " " + m.Sensor + " " + fmt.Sprintf("%f", m.Value) + " " + fmt.Sprintf("%f", m.Temperature)
}

func (m SenseiveMeasure) Tags() map[string]string {
	return map[string]string{
		"type": m.Sensor,
	}
}

func (m SenseiveMeasure) Fields() map[string]interface{} {
	return map[string]interface{}{
		"value":       m.Value,
		"temperature": m.Temperature,
	}
}

func (m SenseiveMeasure) Measurement() string {
	return m.Captor
}

func (m SenseiveMeasure) Date() time.Time {
	return m.DateTime
}

// Measure for trimble data -------------------------------------------------------------------------------------------
type TrimbleMeasure struct {
	DateTime           time.Time
	Captor             string
	Northing           float64
	Easting            float64
	Elevation          float64
	ReferenceNorthing  float64
	ReferenceEasting   float64
	ReferenceElevation float64
}

func (m TrimbleMeasure) String() string {
	return m.DateTime.Format("2006-01-02T15:04:05Z") + " " + fmt.Sprintf("%f", m.Northing) + " " + fmt.Sprintf("%f", m.Easting) + " " + fmt.Sprintf("%f", m.Elevation)
}

func (m TrimbleMeasure) Tags() map[string]string {
	return map[string]string{}
}

func (m TrimbleMeasure) Fields() map[string]interface{} {
	return map[string]interface{}{
		"northing":           m.Northing,
		"easting":            m.Easting,
		"elevation":          m.Elevation,
		"referenceNorthing":  m.ReferenceNorthing,
		"referenceEasting":   m.ReferenceEasting,
		"referenceElevation": m.ReferenceElevation,
	}
}

func (m TrimbleMeasure) Measurement() string {
	return m.Captor
}

func (m TrimbleMeasure) Date() time.Time {
	return m.DateTime
}
