#!/bin/bash


DIST=../../web/dist
RSC=$DIST/static

# resources from web dist
#IMG=$RSC/img
JS=$RSC/js
CSS=$RSC/css
FONTS=$RSC/fonts

clean() {
	echo clean
	rm -rf ./static index.html
}

copy_web_files() {
	echo copy web files
	mkdir static

	mkdir static/css
	cp -rf $CSS/*.css static/css

	mkdir static/js
	cp -rf $JS/*.js static/js
	
	mkdir static/fonts
	cp -rf $FONTS/*.woff2 static/fonts

	cp -rf $DIST/index.html .
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

modify_index() {
	echo modify index
	sed -i "s|/static/|static/|g" index.html
}
#
#add_css_link(){
#	pattern=$1
#	css=$(cd static/css && ls $pattern*.css)
#	if [[ -z $css ]]; then
#		echo add_css_link: no $pattern css file
#		exit 4
#	fi
#	link="<link href=\"static/css/$css\" rel=\"stylesheet\">"
#	sed -i "/<\/title>/ a $link" $index
#}
#modify_index() {
#	index=index.html
#	
#	echo regenerate js link in $index
#	del_js_link
#	add_js_link manifest
#	add_js_link vendor
#	add_js_link app
#
#	echo regenerate css link in $index
#	del_css_link
#	add_css_link app
#}

generate() {
	clean
	copy_web_files
	modify_css
	modify_index
}

#clear_generated_files() {
#	echo clear autogen files
#	rm -rf static
#	rm -rf backup-*
#
#	echo clear links in index.html
#	del_js_link
#	del_css_link
#}

show_help() {
	echo "
Usage: $0 [option]

Options:
  gen:		Generate files taht used to build mobile APP, it requires web static files in $DIST
  clear: 	Clear all files generated automatically
  clean: 	Clear all files generated automatically
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
	clear | clean )
		clean
		exit $?
	;;
	* )
		show_help
		exit $?
esac
