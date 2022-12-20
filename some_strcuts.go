package main

import (
	"fmt"
	"time"

	boom "github.com/tyjet/buildify/drinks"
	"github.com/tyjet/buildify/food"
)

type NotAStruct string

type AStruct struct {
	F00 string
	B4r int `zero:28`
	B4z fmt.Scanner
}

type BStruct struct {
	Honk  AStruct `json:"h0nk" zero:NewAStructBuilder().Build()`
	Binge food.Bird
	Bride boom.Liquid
}

type CStruct struct {
	BStruct
	time.Ticker
	food.Bird
	boom.Liquid
	Lemons AStruct
	Grapes fmt.Formatter
}
