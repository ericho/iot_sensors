#include "lm35Sensor.hpp"


lm35::lm35(int adcPort)
{
    adc = new mraa::Aio(adcPort);
    if (!adc) {
	exit(1);
    }
}

float lm35::getTemp()
{
    return (float)adc->read() * REFERENCE_VOLTAGE / ADCBITS / LM35_RESOLUTION;
}

int main(int argc, char *argv[])
{
    lm35 *sensor;
    sensor = new lm35(0);
    std::cout << sensor->getTemp() << "\n";
}
