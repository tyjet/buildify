// Code generated by "buildify some_strcuts.go CStruct"; DO NOT EDIT.

package main

import (
	fmt "fmt"
	boom "github.com/tyjet/buildify/drinks"
	food "github.com/tyjet/buildify/food"
	time "time"
)

type CStructBuilder struct {
	grapes  fmt.Formatter
	bStruct BStruct
	ticker  time.Ticker
	bird    food.Bird
	liquid  boom.Liquid
	lemons  AStruct
}

func (b *CStructBuilder) BStruct(bStruct BStruct) *CStructBuilder {
	b.bStruct = bStruct
	return b
}

func (b *CStructBuilder) Ticker(ticker time.Ticker) *CStructBuilder {
	b.ticker = ticker
	return b
}

func (b *CStructBuilder) Bird(bird food.Bird) *CStructBuilder {
	b.bird = bird
	return b
}

func (b *CStructBuilder) Liquid(liquid boom.Liquid) *CStructBuilder {
	b.liquid = liquid
	return b
}

func (b *CStructBuilder) Lemons(lemons AStruct) *CStructBuilder {
	b.lemons = lemons
	return b
}

func (b *CStructBuilder) Grapes(grapes fmt.Formatter) *CStructBuilder {
	b.grapes = grapes
	return b
}

func NewCStructBuilder() *CStructBuilder {
	return &CStructBuilder{}
}

func (b *CStructBuilder) Build() CStruct {
	return CStruct{
		BStruct: b.bStruct,
		Ticker:  b.ticker,
		Bird:    b.bird,
		Liquid:  b.liquid,
		Lemons:  b.lemons,
		Grapes:  b.grapes,
	}
}
