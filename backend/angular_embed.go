//go:build !debug
// +build !debug

package main

import (
	"net/http"
)

var AngularHandler = http.FileServer(http.Dir("../frontend/dist"))
