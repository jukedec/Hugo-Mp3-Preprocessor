#!/bin/bash

# Parameter 1 is the source folder that mp3s are stored in, Parameter 2 is the temporry folder where site is initially set up. Paramter 3 is the output directory of finished site.


# /Users/frigginglorious/code/Hugo-Mp3-Pre/bash.sh /Users/frigginglorious/project/testMusic/Jeff\ Rosenstock\ -\ 2020\ DUMP /Users/frigginglorious/project/sites/temp/brickHouse /Users/frigginglorious/project/sites/brickHouse
# send hugo new site param to be the output for the 

# cd "$(dirname "$0")"

try() { "$@" || die "cannot $*"; }

runDir="$(dirname "$0")"
cd runDir

siteDir=$2
# siteDir=${siteDir:3}
echo "I'm starting in folder:" 
pwd
echo "new site in $siteDir"

# scriptDir="$(dirname "$0")"
# cd $scriptDir
echo "Script location is $0 in folder:"
pwd

hugo new site $siteDir

OUTPUT=$($runDir/main "$1" "$2" "$4")
# OUTPUT=$($0/..//main "$1" "$2" "$4")

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

LASTLINE=${OUTPUT##*$'\n'}

echo $LASTLINE

mv -v public $LASTLINE
mv -v $LASTLINE $3

echo 'SITE IS AT?:'
echo $3

# hugo server -D --verbose
echo "${OUTPUT}"
