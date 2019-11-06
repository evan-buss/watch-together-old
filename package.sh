#!/bin/bash

cd web
npm run build
cd ..
packr2
go build
packr2 clean