## Task

Write a function that checks whether a string is a correct sequence of brackets.
3 types of brackets are supported - [], {}, and ()
Input parameter - a string of any length.
Result - a boolean.
True if each opening bracket has a corresponding closing bracket of the same type and brackets are placed in the right order.
False if not.

({}[]){} - true
([([([])])]) - true
(()] - false
)()( - false
{}( - false
{(}) - false

## Resolved in both Python & Go (Golang)
