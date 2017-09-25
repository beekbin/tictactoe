#!/bin/bash

make build
ret=$?
if [ $ret -ne 0 ] ; then
    echo "build failed"
    exit 1
fi

./_output/tictac --v 2
