package iparser

import parser "github.com/tencent/gengine/internal/iantlr/alr"

type GengineParserVisitor struct {
	parser.BasegengineVisitor
}

func NewGengineParserVisitor() *GengineParserVisitor {
	return &GengineParserVisitor{}
}
