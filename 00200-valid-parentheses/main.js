function isValid(s) {
  const stack = [];
  // console.log("isValid", s);
  for (let i = 0; i < s.length; i++) {
    // console.log("Char at", i, s[i]);
    if (s[i] == "(" || s[i] == "[" || s[i] == "{") {
      stack.push(s[i]);
    } else {
      if (stack.length == 0) {
        return false;
      }
      const top = stack.pop();
      if (top == "[" && s[i] != "]") {
        return false;
      }
      if (top == "(" && s[i] != ")") {
        return false;
      }
      if (top == "{" && s[i] != "}") {
        return false;
      }
    }
  }

  return stack.length == 0;
}

s = "()";
console.log("Example 1", isValid(s));

s = "()[]{}";
console.log("Example 2", isValid(s));

s = "(]";
console.log("Exampel 3", isValid(s));

s = "[";
console.log("Exampel 4", isValid(s));
