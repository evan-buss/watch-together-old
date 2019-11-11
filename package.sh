#!/bin/bash

# Build complete application into a binary file.

cd web || exit
npm run build
cd ..
packr2
go build
packr2 clean