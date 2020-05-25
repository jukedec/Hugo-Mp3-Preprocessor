#!/bin/bash

# Parameter 1 is the source folder that mp3s are stored in, Parameter 2 is the folder you want the site created in
# Example:
# ./bash.sh ~/project/hugoStuff/sites/setlers ~/Music/Setlers\ -\ Katana\ EP
# ./bash.sh ~/Music/myBand ~/project/hugoStuff/sites/myBandSite
# send hugo new site param to be the output for the 

siteDir=$2
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
hugo -D #create full site, then serve test
hugo server -D --verbose
