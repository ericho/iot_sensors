#include "mysensor.hpp"
#include "grove.hpp"
#include "hcsr04.hpp"
#include "jhd1313m1.hpp"
#include <iostream>
#include <stdexcept>
#include <cstdlib>
#include <sstream>

#define TRIGGER_PIN 		5
#define ECHO_PIN 		6
#define TEMP_ANALOG_PIN 	0
#define LIGHT_ANALOG_PIN	1


upm::GroveTemp* temp;
upm::HCSR04* dist;
upm::GroveLight* light;
upm::Jhd1313m1* lcd;

void mySensor::init()
{
    std::cout << "On init in plugin" << std::endl;
    temp = new upm::GroveTemp(TEMP_ANALOG_PIN);
    dist = new upm::HCSR04(TRIGGER_PIN, ECHO_PIN);
    light = new upm::GroveLight(LIGHT_ANALOG_PIN);
    lcd = new upm::Jhd1313m1(0, 0x3e, 0x62);
    lcd->clear();
    lcd->write("Initialized");
}

void mySensor::sample(dataContainer &dc)
{
    lcd->clear();
    lcd->setColor(99, 33, 99);
    lcd->write("Sampling");
    int celsius = temp->value();
    int lux = light->value();
    double distance = dist->getDistance(CM);

    lcd->clear();
    std::stringstream m;
    m << celsius << " C " << "L=" << lux << " " << distance;
    lcd->write(m.str().c_str());
    // Storing integer data into the dataContainer object
    dc.put("temperature", celsius, "ints");
    dc.put("light", lux, "ints");
    dc.put("distance", distance, "doubles");
    lcd->setColor(32, 99, 32);
}

void mySensor::finalize()
{
    std::cout << "On finalize" << std::endl;
    delete temp;
    delete dist;
    delete light;
}

export_plugin(mySensor, "MyTempSensor");
