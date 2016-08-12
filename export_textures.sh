#!/bin/bash

width=50
height=100

#
# It seems like inkscape is a little bit faster than ImageMagick, so we
# try inkscape first and then fall back on ImageMagick
#

if command -v inkscape >/dev/null 2>&1; then
    for i in images/*.svg; do
	    inkscape --export-background-opacity=0 \
		     -w $width -h $height --without-gui \
		     --export-png="images/$(basename "$i" .svg).png" "$i" >/dev/null;
    done
elif command -v convert >/dev/null 2>&1; then
    for i in images/*.svg; do
	    convert -background none "$i" \
		    -geometry "$width"x"$height" \
		    "images/$(basename "$i" .svg).png";
    done
else
	echo "Inkscape or ImageMagick needs to be installed in order to convert SVG files to PNG" 1>&2
	exit 1
fi
