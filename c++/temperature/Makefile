SENSORPLUGIN=temp
PLUGINPREFIX=libudplugin_
CXX=g++
SENSYSPATH=/opt/orcm
PLUGININSTALLPATH=$(SENSYSPATH)/lib/openmpi
CXXFLAGS=-O3 -fPIC -I$(SENSYSPATH)/include/openmpi -I/usr/include/upm
LDFLAGS=-L$(SENSYSPATH)/lib -shared -rdynamic -lsensysplugins -lupm-grove -lupm-hcsr04 -lupm-i2clcd

all:
	$(CXX) $(CXXFLAGS) $(LDFLAGS) mysensor.cpp -o $(PLUGINPREFIX)$(SENSORPLUGIN).so
install: all
	install $(PLUGINPREFIX)$(SENSORPLUGIN).so $(PLUGININSTALLPATH)

clean:
	-rm *.so
