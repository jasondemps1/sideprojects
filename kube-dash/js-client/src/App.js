import React, { useEffect, useState } from 'react';
import './App.css';

import { SensorRequest } from "./sensorpb/sensor_pb"
import { SensorClient } from "./sensorpb/sensor_grpc_web_pb"

var client = new SensorClient('http://localhost:8000')

function App() {
  const [temp, setTemp] = useState(-9999)
  const [humidity, setHumidity] = useState(-9999)

  const getTemp = () => {
    console.log("temp called")

    var sensorRequest = new SensorRequest()
    var stream = client.tempSensor(sensorRequest, {})

    stream.on('data', function(response) {
      setTemp(response.getValue())
    });
  };

  const getHumidity = () => {
    console.log("humidity called")

    var sensorRequest = new SensorRequest()
    var stream = client.humiditySensor(sensorRequest, {})

    stream.on('data', function(response) {
      setHumidity(response.getValue())
    });
  };

  useEffect(()=>{
    getTemp()
  }, []);

  useEffect(()=>{
    getHumidity()
  }, []);

  return (
    <div>
      Temperature : {temp} F
      <br />
      Humidity : {humidity} %
    </div>
  );
}

export default App;
