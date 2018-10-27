#!/bin/bash
abigen --sol=execute/execute.sol --pkg=execute --out=execute/execute.go
go build

