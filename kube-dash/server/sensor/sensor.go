package sensor

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

type SensorType int

const (
	temp SensorType = iota
	humidity
)

func (st SensorType) String() string {
	return [...]string{"temp", "humidity"}[st]
}

// Sensor - Contains a map of sensor data and a Mutex
type Sensor struct {
	Data map[string]int64
	M    sync.RWMutex
}

// NewSensor creates a new sensor object
func NewSensor() *Sensor {
	return &Sensor{
		Data: make(map[string]int64),
	}
}

// SetTempSensor - Sets random temp. sensor value
func (s *Sensor) SetTempSensor() {
	for {
		s.M.Lock()
		s.Data[fmt.Sprintf("%s", SensorType(temp))] = int64(rand.Intn(120))
		s.M.Unlock()

		time.Sleep(5 * time.Second)
	}
}

// SetHumiditySensor - Sets random humidity sensor value
func (s *Sensor) SetHumiditySensor() {
	for {
		s.M.Lock()
		s.Data[fmt.Sprintf("%s", SensorType(humidity))] = int64(rand.Intn(100))
		s.M.Unlock()

		time.Sleep(2 * time.Second)
	}
}

// StartMonitoring - Starts fetching data from fake sensors in Goroutines
func (s *Sensor) StartMonitoring() {
	log.Println("Start monitoring...")

	go s.SetHumiditySensor()
	go s.SetTempSensor()
}

// GetTempSensor - Returns latest temp. sensor data
func (s *Sensor) GetTempSensor() int64 {
	s.M.RLock()
	defer s.M.RUnlock()

	return s.Data[fmt.Sprintf("%s", SensorType(temp))]
}

// GetHumiditySensor - Returns latest temp. sensor data
func (s *Sensor) GetHumiditySensor() int64 {
	s.M.RLock()
	defer s.M.RUnlock()

	return s.Data[fmt.Sprintf("%s", SensorType(humidity))]
}
