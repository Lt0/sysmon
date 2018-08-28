#!/bin/bash


DIST=../../web/dist
RSC=$DIST/static

# resources from web dist
IMG=$RSC/img
JS=$RSC/js
CSS=$RSC/css

FONTS=$RSC/fonts

# mobile only resources
MANIFEST=manifest.json
UNPACK=unpackage

backup() {
	bak_files="img js css $UNPACK"
	dir=backup-$(date +%y_%m_%d-%H:%M:%S)
	mkdir $dir
	mv $bak_files $dir 2>/dev/null
}

update_web_files() {
	mkdir $UNPACK
	cp -rf $IMG $JS $CSS .
	cp -rf $FONTS $UNPACK
}

modify_css() {
	css_file=$(cd css && ls *.css)
	if [[ -z $css_file ]]; then
		echo modify_css: no css file
		exit 2
	fi

	echo regenerate fonts link in css 
	sed -i "s|/static/fonts|../$UNPACK/fonts|g" css/$css_file
}

add_js_link() {
	pattern=$1

	js=$(cd js && ls $pattern*.js)
	if [[ -z $js ]]; then
		echo modify_index: no $pattern js file
		exit 3
	fi
	link="<script type=\"text/javascript\" src=\"js/$js\"></script>"
	sed -i "/<\/body>/ i $link" $index
}
modify_index() {
	index=index.html
	
	echo regenerate js link in $index
	# remove old js link
	sed -i '/<script type="text\/javascript" src="/d' $index
	# add js link
	add_js_link manifest
	add_js_link vendor
	add_js_link app


	# modify css
	echo regenerate css link in $index
	css_file=$(cd css && ls *.css)
	if [[ -z $css_file ]]; then
		echo modify_index: no css file
		exit 4
	fi

	link="<link href=\"css/$css_file\" rel=\"stylesheet\">"
	sed -i "s|<link href=\"css/.*rel=\"stylesheet\">|$link|" $index
}

generate() {
	backup

	echo get web dist files
	update_web_files
	modify_css
	modify_index
}

clear_generated_files() {
	echo clear autogen files
	rm -rf img js css $UNPACK
	rm -rf backup-*
}

show_help() {
	echo "
Usage: $0 [option]

Options:
  gen:		Generate files taht used to build mobile APP, it requires web static files in $DIST
  clear: 	Clear all files generated automatically
  help:  	Show current msg

Example:
  1. $0 gen
    - generate files to build mobile APP

  2. $0 clear
    - clear files

  3. $0 help
    - show help info
"
}

case $1 in
	gen | generate )
		generate
		exit $?
	;;
	clear )
		echo clear
		clear_generated_files
		exit $?
	;;
	* )
		show_help
		exit $?
esac


if [[ -n $1 ]] && [[ $1 = "clear" ]]; then
	echo clear autogen files
	rm -rf img js css
	rm -rf $UNPACK/fonts
	rm -rf backup-*
	exit 0
fi

