#!/bin/bash

T=$(date)
echo "$T">> /home/linaro/date-service.txt 2>&1 &
