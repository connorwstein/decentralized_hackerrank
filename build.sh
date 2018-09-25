#!/bin/bash
abigen --sol=execute.sol --pkg=main --out=execute.go
go build

