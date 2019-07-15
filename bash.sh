#!/bin/bash
# Param 1 is folder you want site created in, param 2 is source mp3 folder.
# send hugo new site param to be the output for the 
hugo new site $1
./main $2
cd $1 
cd themes
git clone https://github.com/htr3n/hyde-hyde
cd ../
hugo server -D --verbose --renderToDisk