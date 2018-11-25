#!/bin/bash

release="2018g"
file="timezones.geojson.zip"

wget "https://github.com/evansiroky/timezone-boundary-builder/releases/download/${release}/${file}"
unzip "${file}"

mkdir "reduced"
mapshaper -i dist/combined.json -simplify visvalingam 20% -o reduced/reduced.json

go build -o builder tzshapefilegen/main.go
./builder -file reduced/reduced.json

rm "${file}"
rm "dist/combined.json" "reduced/reduced.json" "builder"
rmdir "dist" "reduced"