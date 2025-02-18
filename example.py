def is_valid_bracket_sequence(s: str) -> bool:
    stack = []
    bracket_map = {')': '(', ']': '[', '}': '{'}
    opening_brackets = set(bracket_map.values())
    
    for char in s:
        if char in opening_brackets:
            stack.append(char)
        elif char in bracket_map:
            if not stack or stack.pop() != bracket_map[char]:
                return False
    
    return len(stack) == 0

# Test cases
print(is_valid_bracket_sequence("({}[]){}"))         # True
print(is_valid_bracket_sequence("([([([])])])"))     # True
print(is_valid_bracket_sequence("(()]"))              # False
print(is_valid_bracket_sequence(")()("))              # False
print(is_valid_bracket_sequence("{}("))               # False
print(is_valid_bracket_sequence("{(})"))              # False
