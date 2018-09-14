#!/bin/bash

set -e

VER=$(git tag | tail -n1)

cd build
list=$(ls)
for d in $list; do
	echo Working on $d
	name=${d}-${VER}
	mv $d $name
	if [ "$d" = "sysmon-client-win32-ia32" ] || [ "$d" = "sysmon-client-win32-x64" ]; then
		out_name=$name.7x
		echo Compressing $d 7z to $out_name
		if [ -f "$out_name" ]; then
			continue
		fi
		7z a $name.7z $name -mx=9
	else
		out_name=$name.tar.xz
		echo Compressing $d to $out_name
		if [ -f "$out_name" ]; then
			continue
		fi
		tar --checkpoint=100 --checkpoint-action=dot --totals -Jcf $name.tar.xz $name
	fi
done

cd ..
echo Moving build to build-$VER
mv build build-$VER

