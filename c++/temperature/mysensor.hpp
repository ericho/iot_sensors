#include "orcm/common/udsensors.h"

class mySensor : public UDSensor
{
public:
    mySensor(){};
    virtual ~mySensor(){};
    void init();
    void sample(dataContainer &dc);
    void finalize();
};
