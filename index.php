<?php

function isValidBracketSequence($s) {
    $stack = [];
    $bracketMap = [')' => '(', ']' => '[', '}' => '{'];
    $openingBrackets = array_values($bracketMap);

    foreach (str_split($s) as $char) {
        if (in_array($char, $openingBrackets)) {
            array_push($stack, $char);
        } elseif (array_key_exists($char, $bracketMap)) {
            if (empty($stack) || array_pop($stack) !== $bracketMap[$char]) {
                return false;
            }
        }
    }

    return empty($stack);
}

// Test cases
var_dump(isValidBracketSequence("({}[]){}"));         // True
var_dump(isValidBracketSequence("([([([])])])"));     // True
var_dump(isValidBracketSequence("(()]"));              // False
var_dump(isValidBracketSequence(")()("));              // False
var_dump(isValidBracketSequence("{}("));               // False
var_dump(isValidBracketSequence("{(})"));              // False
