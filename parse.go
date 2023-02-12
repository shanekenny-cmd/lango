package main

import (
    "fmt"
)

type Node struct {
    tok Token;
    kids []*Node;
}

func parse_mul() *Node {
    if CurTok.tocat != TokenCats[IDENT] && CurTok.tocat != TokenCats[INTLIT] {
        // this is an error i think
        // we don't have what we expected to
        fmt.Println("Unexpected token", *CurTok)
        Advance()
        return &Node{}
    }

    var term_node Node = Node{*CurTok, make([]*Node, 0)}
    Advance()
    // now we must have a * or / followed by another mul
    if CurTok.tocat == TokenCats[MULOP] || CurTok.tocat == TokenCats[DIVOP] {
        var binexpr_node Node = Node{*CurTok, []*Node{&term_node}}
        Advance()
        var rhs *Node = parse_expr()
        binexpr_node.kids = append(binexpr_node.kids, rhs)
        return &binexpr_node
    }
    return &term_node
}

func parse_expr() *Node {
    if CurTok.tocat == TokenCats[LPAREN] {
        // add support for {} here
        // figure out what their purpose is first
        var tok Token = *CurTok
        var lp_node Node = Node{tok, make([]*Node, 0)}
        Advance()
        var cexpr_node *Node = parse_expr()
        if CurTok.tocat != TokenCats[RPAREN] {
            fmt.Println("Missing RPAREN")
        }
        tok = *CurTok
        var rp_node Node = Node{tok, make([]*Node, 0)}
        var expr_node Node = Node{Token{"PEXP", TokenValue{}}, []*Node{&lp_node, cexpr_node, &rp_node}}
        Advance()
        if CurTok.tocat == TokenCats[ADDOP] || CurTok.tocat == TokenCats[SUBOP] || CurTok.tocat == TokenCats[MULOP] || CurTok.tocat == TokenCats[DIVOP] || CurTok.tocat == TokenCats[EQOP] {
            var binexpr_node Node = Node{*CurTok, []*Node{&expr_node}}
            Advance()
            var rhs *Node = parse_expr()
            binexpr_node.kids = append(binexpr_node.kids, rhs)
            return &binexpr_node
        }
        return &expr_node
    }
    var mul *Node = parse_mul()
    if CurTok.tocat == TokenCats[ADDOP] || CurTok.tocat == TokenCats[SUBOP] {
        var binexpr_node Node = Node{*CurTok, []*Node{mul}}
        Advance()
        var rhs *Node = parse_expr()
        binexpr_node.kids = append(binexpr_node.kids, rhs)
        return &binexpr_node
    } else if CurTok.tocat == TokenCats[ASNOP] {
        if mul.tok.tocat != TokenCats[IDENT] {
            fmt.Println("Cannot assign expressions to eachother")
        }
        var asnexpr_node Node = Node{*CurTok, []*Node{mul}}
        Advance()
        var rhs *Node = parse_expr()
        asnexpr_node.kids = append(asnexpr_node.kids, rhs)
        return &asnexpr_node
    } else if CurTok.tocat == TokenCats[EQOP] {
        var binexpr_node Node = Node{*CurTok, []*Node{mul}}
        Advance()
        var rhs *Node = parse_expr()
        binexpr_node.kids = append(binexpr_node.kids, rhs)
        return &binexpr_node
    }
    return mul
}

func parse_stmt() *Node {
    var lhs *Node = parse_expr()
    if CurTok.tocat == TokenCats[EOF] {
        return lhs
    } else if CurTok.tocat == TokenCats[ASNOP] {
        if lhs.tok.tocat == "PEXP" {
            var val_node *Node = lhs
            for val_node.tok.tocat != TokenCats[IDENT] && len(val_node.kids) > 0 {
                val_node = val_node.kids[1]
            }
            if val_node.tok.tocat != TokenCats[IDENT] {
                fmt.Println("Cannot assign to", val_node.tok.tocat)
                return &Node{}
            }
        } else if lhs.tok.tocat != TokenCats[IDENT] {
            // we have a problem, depending on which type of
            // node we are trying to assign to 
            // this can be expanded into a switch or more else statements
            fmt.Println("Cannot assign to", lhs.tok.tocat)
            return &Node{}
        }
        var asn_node Node = Node{*CurTok, []*Node{lhs}}
        Advance()
        var rhs *Node = parse_expr()
        asn_node.kids = append(asn_node.kids, rhs)
        return &asn_node
    }
    fmt.Println("Cannot parse multiple statements", *CurTok)
    return &Node{}
}

var Ast *Node

func parse() {
    Advance()
    Ast = &Node{Token{}, make([]*Node, 0)}
    var root *Node = parse_stmt()
    Ast.kids = append(Ast.kids, root)
    traverse_ast(Ast)
}

func traverse_ast(ast *Node) {
    for _, kid := range ast.kids {
        traverse_ast(kid)
    }
    fmt.Println(ast)
}
