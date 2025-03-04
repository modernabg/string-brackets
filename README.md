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

## Resolved in Python, Go (Golang) and PHP

### Python explenation

The function uses a stack to keep track of opening brackets.

It iterates through each character in the string:
- If it's an opening bracket, it gets pushed onto the stack.
- If it's a closing bracket, the function checks if the stack is empty or if the top of the stack matches the corresponding opening bracket.

Finally, it returns `True` if the stack is empty (all brackets matched), otherwise `False`.

### Go (Golang) explanation

The function uses a slice as a stack to keep track of opening brackets.

It iterates through each character in the string:
- If it's an opening bracket, it gets added to the stack.
- If it's a closing bracket, it checks if the stack is empty or if the top of the stack matches the corresponding opening bracket.

Finally, it returns `true` if the stack is empty (all brackets matched), otherwise `false`.

### PHP explanation

The function uses an array as a stack to track opening brackets.

It iterates through each character in the string:
- If it’s an opening bracket, it pushes it onto the stack.
- If it’s a closing bracket, it checks if the stack is empty or if the top of the stack matches the corresponding opening bracket.

Finally, it returns `true` if the stack is empty (indicating all brackets matched), otherwise `false`.
