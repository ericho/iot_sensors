#include "mraa.hpp"
#include <iostream>

#define LM35_RESOLUTION 0.01
#define ADCBITS 1024
#define REFERENCE_VOLTAGE 5

class lm35
{
public:
    lm35(int adcPort);
    ~lm35(){};
    float getTemp(void);
private:
    mraa::Aio *adc;
};
