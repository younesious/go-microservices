package main

import (
	"os"
	"testing"
)

var testApp Config

func TestMain(m *testing.M) {

	os.Exit(m.Run())
}
