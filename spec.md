; an expression is an addition
E   -> IDN = ADD
E   -> ADD == ADD
E   -> ADD          
E   -> LP E RP
LP  -> (
RP  -> )

; an addition is a multiplication that is optionally followed
; by +- and another addition
ADD -> MUL T  
T   -> PM ADD
T   -> ε
PM  -> +
PM  -> -

; a multiplication is an integer that is optionally followed
; by */ and another multiplication
MUL -> IDN G
MUL -> INT G
G   -> MD MUL
G   -> ε
MD  -> *
MD  -> /

; an integer is a digit that is optionally followed by an integer
INT -> DIGIT J
J   -> INT
J   -> ε

; digits
DIGIT -> 0
DIGIT -> 1
DIGIT -> 2
DIGIT -> 3
DIGIT -> 4
DIGIT -> 5
DIGIT -> 6
DIGIT -> 7
DIGIT -> 8
DIGIT -> 9
