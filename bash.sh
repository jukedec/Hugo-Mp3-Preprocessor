#!/bin/bash
# Param 1 is folder you want site created in, param 2 is source mp3 folder.
# send hugo new site param to be the output for the 

siteDir=$1
# siteDir=${siteDir:3}
echo "I'm starting in folder:" 
pwd
echo "new site in $siteDir"
hugo new site $siteDir
./main "$1" "$2" 
# siteDir = $1
# siteDir = ${siteDir:3}

cd $siteDir
echo "I'm in folder:" 
pwd
# cd themes
# rm -rf hyde-hyde
# git clone https://github.com/htr3n/hyde-hyde themes/import
# git clone https://github.com/fncnt/vncnt-hugo themes/import
# git clone https://github.com/MarcusVirg/forty themes/import
git clone https://github.com/frigginglorious/hyde-hyde themes/import
# cd ../
hugo server -D --verbose --renderToDisk