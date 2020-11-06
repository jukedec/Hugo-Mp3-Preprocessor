#!/bin/bash

# Parameter 1 is the source folder that mp3s are stored in, Parameter 2 is the folder you want the site created in
# Example:

# ./bash.sh ~/Music/myBand ~/project/hugoStuff/sites/myBandSite
# ./bash.sh ~/Music/Setlers\ -\ Katana\ EP ~/project/hugoStuff/sites/setlers123
# send hugo new site param to be the output for the 

cd "$(dirname "$0")"



siteDir=$2
# siteDir=${siteDir:3}
echo "I'm starting in folder:" 
pwd
echo "new site in $siteDir"

# scriptDir="$(dirname "$0")"
# cd $scriptDir
echo "im in $scriptDir"
echo "or in $0"
pwd

hugo new site $siteDir

OUTPUT=$(./main "$1" "$2" | tail -1)

# siteDir = $1
# siteDir = ${siteDir:3}


mkdir $3
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
hugo -D --noTimes #create full site, then serve test

mv -v public $OUTPUT
mv -v $OUTPUT $3

echo 'SITE IS AT?:'
echo $3

# hugo server -D --verbose
echo "${OUTPUT}"
