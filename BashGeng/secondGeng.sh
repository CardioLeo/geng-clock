# Author: William Z Chadwick
# Date Created: 03/17/2023
# Date Modified: 03/14/2023
# Description: A minimal test, take 2, of my geng clock application
# Usage: geng-clock

#!/bin/bash

# x=1; while [ $x -le 5 ]; do echo "Welcome $x times"; x=$(( $x + 1 )); done

# geng=8640;

geng=5;
i=0;

echo "geng-clock started";

while [ $i -le $geng ];
do
        echo $i
        sleep 1
        i=$(( $i + 1 ))
	# add case statements here for when i = different divisions
	# of geng numbers, such as 1/2 geng, 1/4 geng, 1 full geng,
	# or multiples of geng 
done
echo "geng-clock process over"
