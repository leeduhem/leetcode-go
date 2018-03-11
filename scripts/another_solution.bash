#!/usr/bin/env bash

set -x
set -e

if [ $# -ne 1 ]; then
	echo "Usage: $0 id"
	exit 1
fi

id=$(printf "%.3d" $1)

topdir=$(dirname $(dirname $(readlink -e $0)))

cd "${topdir}"/src

if [ ! -d p$id ]; then
	echo "[ERROR] Source code directory does not exist: p$id"
	exit 1
fi

lastdir=$(ls -dr p${id}* | head -1)
suffix=1
if echo $lastdir | grep -q '_'; then
	suffix=$(expr $(echo $lastdir | grep -oP '(?<=_)\d+') + 1)
fi

newdir=p${id}_${suffix}
cp -r p${id} ${newdir}

cd ${newdir}
rename "s/p${id}/p${id}_${suffix}/" *

git add .
sed -i "s/p${id}/p${id}_${suffix}/" *
