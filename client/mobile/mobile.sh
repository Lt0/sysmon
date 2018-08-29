#!/bin/bash


DIST=../../web/dist
RSC=$DIST/static

# resources from web dist
#IMG=$RSC/img
#JS=$RSC/js
#CSS=$RSC/css
#FONTS=$RSC/fonts

# mobile only resources
MANIFEST=manifest.json

backup() {
#	bak_files="img js css fonts"
	bak_files="static"
	dir=backup-$(date +%y_%m_%d-%H:%M:%S)
	mkdir $dir
	mv $bak_files $dir 2>/dev/null
}

update_web_files() {
	cp -rf $RSC .
	rm -rf static/PWA
}

modify_css() {
	css_file=$(cd static/css && ls *.css)
	if [[ -z $css_file ]]; then
		echo modify_css: no css file
		exit 2
	fi

	echo regenerate fonts link in css 
	sed -i "s|/static/fonts|../fonts|g" static/css/$css_file
}

del_js_link() {
	index=index.html
	sed -i '/<script type="text\/javascript" src="/d' $index
}
del_css_link() {
	index=index.html
	sed -i '/<link href="static\/css/d' $index
}
add_js_link() {
	pattern=$1

	js=$(cd static/js && ls $pattern*.js)
	if [[ -z $js ]]; then
		echo add_js_link: no $pattern js file
		exit 3
	fi
	link="<script type=\"text/javascript\" src=\"static/js/$js\"></script>"
	sed -i "/<\/body>/ i $link" $index
}
add_css_link(){
	pattern=$1
	css=$(cd static/css && ls $pattern*.css)
	if [[ -z $css ]]; then
		echo add_css_link: no $pattern css file
		exit 4
	fi
	link="<link href=\"static/css/$css\" rel=\"stylesheet\">"
	sed -i "/<\/title>/ a $link" $index
}
modify_index() {
	index=index.html
	
	echo regenerate js link in $index
	del_js_link
	add_js_link manifest
	add_js_link vendor
	add_js_link app

	echo regenerate css link in $index
	del_css_link
	add_css_link app
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
	rm -rf static
	rm -rf backup-*

	echo clear links in index.html
	del_js_link
	del_css_link
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
