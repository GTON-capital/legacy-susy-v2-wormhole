#!/bin/bash

if [ -z $N ]
then
	echo "N is not set"
	exit 1
fi



for ((i=0; i<N; i++))
do
	echo "i is $i"
done
