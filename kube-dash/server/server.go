package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	sensor "kube-dash/server/sensor"
	sensorpb "kube-dash/server/sensorpb/protos"

	"google.golang.org/grpc"
)

type server struct {
	Sensor *sensor.Sensor
}

func (s *server) TempSensor(req *sensorpb.SensorRequest, stream sensorpb.Sensor_TempSensorServer) error {
	for {
		time.Sleep(5 * time.Second)

		log.Println("Getting Temp Sensor value.")

		temp := int64(123) //s.Sensor.GetTempSensor()
		log.Println("Sending Temp Sensor value.")
		err := stream.Send(&sensorpb.SensorResponse{Value: temp})

		if err != nil {
			log.Println("Error sending metric message ", err)
		}
	}
	return nil
}

func (s *server) HumiditySensor(req *sensorpb.SensorRequest, stream sensorpb.Sensor_HumiditySensorServer) error {
	for {
		time.Sleep(2 * time.Second)

		log.Println("Getting Humidity Sensor value.")
		humid := int64(456) //s.Sensor.GetHumiditySensor()
		log.Println("Sending Humidity Sensor value.")
		err := stream.Send(&sensorpb.SensorResponse{Value: humid})

		if err != nil {
			log.Println("Error sending metric message ", err)
		}
	}
	return nil
}

var (
	defaultIP   string = ""
	defaultPort string = "8000"
)

func main() {
	ip := defaultIP
	if envIP := os.Getenv("SERVER_ADDRESS"); envIP != "" {
		ip = envIP
	}

	port := defaultPort
	if envPort := os.Getenv("SERVER_PORT"); envPort != "" {
		port = envPort
	}

	sns := sensor.NewSensor()
	sns.StartMonitoring()

	addr := fmt.Sprintf("%s:%s", ip, port)

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Error while listening: %v", err)
	}

	s := grpc.NewServer()
	sensorpb.RegisterSensorServer(s, &server{})

	log.Printf("Starting server on: %s\n", addr)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error while serving : %v", err)
	}
}
