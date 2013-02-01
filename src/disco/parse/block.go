// Copyright 2013 Mark Stahl. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can found in the BSD-LICENSE file.

package parse

import (
	"disco/ast"
	"disco/scan"
)

// block := 
//	LBRACE [statements] RBRACE
//
func (p *Parser) parseBlock() (b *ast.Block) {
	p.expect(scan.LBRACE)

	b = &ast.Block{Exprs: p.parseStatements()}

	p.expect(scan.RBRACE)

	return
}

// statements :=
//	[expression [PERIOD expression]*]
//
func (p *Parser) parseStatements() []*ast.Expr {
	var stmts []*ast.Expr

	return stmts
}
