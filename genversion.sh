#!/bin/sh
. ./VERSION
echo "package edsm" > version.go
echo "" >> version.go
echo "const (" >> version.go
echo "	Vmajor   uint16 = "$major >> version.go
echo "	Vminor   uint16 = "$minor >> version.go
echo "	Vbugfix  uint16 = "$bugfix >> version.go
echo "	Vdate    string = \""`date "+%F %T %Z"`"\"" >> version.go
echo "	Vquality string = \""$quality"\"" >> version.go
echo ")" >> version.go
