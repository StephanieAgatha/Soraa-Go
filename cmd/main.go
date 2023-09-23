package main

import (
	"github.com/StephanieAgatha/Soraa-Go/delivery"
	"github.com/StephanieAgatha/Soraa-Go/util/helper"
)

func main() {
	helper.PrintAscii()
	delivery.NewServer().Run()
}
