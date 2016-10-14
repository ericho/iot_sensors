/*
 * Copyright (c) 2016 Intel, Inc. All rights reserved.
 * $COPYRIGHT$
 *
 * Additional copyrights may follow
 *
 * $HEADER$
 */

#include "mysensor.hpp"
#include "grove.hpp"
#include <iostream>
#include <stdexcept>
#include <cstdlib>

upm::GroveTemp* temp;

void mySensor::init()
{
    std::cout << "On init in plugin" << std::endl;
    temp = new upm::GroveTemp(0);
}

void mySensor::sample(dataContainer &dc)
{
    int celsius = temp->value();
    std::cout << "On sample, storing temperature number " << celsius << std::endl;
    // Storing integer data into the dataContainer object
    dc.put("temperature", celsius, "ints");
}

void mySensor::finalize()
{
    std::cout << "On finalize" << std::endl;
    delete temp;
}

export_plugin(mySensor, "MyTempSensor");
