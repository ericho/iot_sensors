# Setup instructions

The versions of mraa and upm should be upgraded. For that. run the following in the edison board. 

```
opkg remove upm-dev
opkg remove upm
opkg remove libmraa0
echo "src mraa-upm http://iotdk.intel.com/repos/3.5/intelgalactic/opkg/i586" > /etc/opkg/mraa-upm.conf
opkg update
opkg install mraa
opkg install upm
```

Compile a simple example.

```
#include <unistd.h>
#include <iostream>
#include <iomanip>
#include "grove.hpp"

int
main(int argc, char **argv)
{
//! [Interesting]

    // Create the temperature sensor object using AIO pin 0
    upm::GroveTemp* temp = new upm::GroveTemp(0);
    std::cout << temp->name() << std::endl;

    // Read the temperature ten times, printing both the Celsius and
    // equivalent Fahrenheit temperature, waiting one second between readings
    for (int i=0; i < 10; i++) {
        int celsius = temp->value();
        int fahrenheit = (int) (celsius * 9.0/5.0 + 32.0);
        printf("%d degrees Celsius, or %d degrees Fahrenheit\n",
                celsius, fahrenheit);
        sleep(1);
    }

    // Delete the temperature sensor object
    delete temp;
//! [Interesting]

    return 0;
}
```

And build with : 

```
g++ -lupm-grove -I /usr/include/upm/ temp.cpp -o temp
```
