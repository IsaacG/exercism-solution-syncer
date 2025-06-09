def is_paired(input_string):
  stack = []
  r = {')': '(', ']': '[', '}': '{'}
  for c in input_string:
    if c in '[{(':
      stack.append(c)
    if c in ']})':
      if len(stack) == 0:
        return False
      if stack[-1] == r[c]:
        stack.pop()
      else:
        return False
  return len(stack) == 0


# vim:ts=2:sw=2:expandtab
