package main

import (
    "fmt"
)

const (
    NUM = iota
    CON
    STR
    ERR
)

type res_cat string
var result_cats = [...]res_cat{"NUM", "CON", "STR", "ERR"}
type result struct {
    num int;
    con bool;
    res string;
    typo res_cat;
}

var context map[string]result

func eval_node(Ast *Node) result {
    var res result = result{}
    switch Ast.tok.tocat {
        case "": 
            // this is empty and we return the result of it's child
            if len(Ast.kids) == 0 {
                res.typo = result_cats[ERR]
                res.res = "too few arguments for " + Ast.tok.value.lexeme
                return res
            }
            return eval_node(Ast.kids[0])
        case "PEXP":
            if len(Ast.kids) == 0 {
                res.typo = result_cats[ERR]
                res.res = "too few arguments for " + Ast.tok.value.lexeme
                return res
            }
            return eval_node(Ast.kids[1])
        case TokenCats[IDENT]:
            pot, ok := context[Ast.tok.value.lexeme]
            if !ok {
                res.typo = result_cats[ERR]
                res.res = Ast.tok.value.lexeme + " has not been assigned yet"
                return res
            }
            res = pot
            res.typo = result_cats[NUM]
        case TokenCats[INTLIT]:
            res.num = Ast.tok.value.numval
            res.typo = result_cats[NUM]
        case TokenCats[ADDOP]:
            if len(Ast.kids) == 0 {
                res.typo = result_cats[ERR]
                res.res = "too few arguments for " + Ast.tok.value.lexeme
                return res
            }
            res.num = eval_node(Ast.kids[0]).num + eval_node(Ast.kids[1]).num
            res.typo = result_cats[NUM]
        case TokenCats[SUBOP]:
            if len(Ast.kids) == 0 {
                res.typo = result_cats[ERR]
                res.res = "too few arguments for " + Ast.tok.value.lexeme
                return res
            }
            res.num = eval_node(Ast.kids[0]).num - eval_node(Ast.kids[1]).num
            res.typo = result_cats[NUM]
        case TokenCats[MULOP]:
            if len(Ast.kids) == 0 {
                res.typo = result_cats[ERR]
                res.res = "too few arguments for " + Ast.tok.value.lexeme
                return res
            }
            res.num = eval_node(Ast.kids[0]).num * eval_node(Ast.kids[1]).num
            res.typo = result_cats[NUM]
        case TokenCats[DIVOP]:
            if len(Ast.kids) == 0 {
                res.typo = result_cats[ERR]
                res.res = "too few arguments for " + Ast.tok.value.lexeme
                return res
            }
            var lhs result = eval_node(Ast.kids[0])
            var rhs result = eval_node(Ast.kids[1])
            if rhs.num == 0 {
                res.typo = result_cats[ERR]
                res.res = "division by zero"
                return res
            }
            res.num = lhs.num / rhs.num
            res.typo = result_cats[NUM]
        case TokenCats[ASNOP]:
            if len(Ast.kids) == 0 {
                res.typo = result_cats[ERR]
                res.res = "too few arguments for " + Ast.tok.value.lexeme
                return res
            }
            var lhs *Node = Ast.kids[0]
            var rhs result = eval_node(Ast.kids[1])
            context[lhs.tok.value.lexeme] = rhs
            res.num = 0
            res.res = "SUCCESS"
            res.typo = result_cats[STR]
        case TokenCats[EQOP]:
            res.typo = result_cats[CON]
            var lhs result = eval_node(Ast.kids[0])
            var rhs result = eval_node(Ast.kids[1])
            if lhs.typo == rhs.typo {
                switch lhs.typo {
                    case result_cats[NUM]:
                        res.con = lhs.num == rhs.num
                    case result_cats[CON]:
                        res.con = lhs.con == rhs.con
                    default:
                        res.typo = result_cats[ERR]
                        res.res = "cannot compare lhs & rhs, different types"
                }
            } else {
                res.typo = result_cats[ERR]
                res.res = "cannot compare lhs & rhs, different types"
            }
        default:
            return result{}
    }
    return res
}

func eval() {
    fmt.Println("Evaluating...")
    var res result = eval_node(Ast)
    switch res.typo {
        case result_cats[STR]:
            fmt.Println(res.res)
        case result_cats[NUM]:
            fmt.Println(res.num)
        case result_cats[CON]:
            fmt.Println(res.con)
        case result_cats[ERR]:
            fmt.Println("ERROR:", res.res)
        default:
            fmt.Println("unreachable: result has invalid type:", res.typo)
    }
}
