def isBalanced(s):
    # Write your code here
    if len(s) % 2 != 0:
        return "NO"
    closed = "]})"
    opened = "{(["
    stack = []
    
    if s[0] in closed:
        return "NO"
    m = {"{":"}", "[": "]", "(":")"}
    for i in s:
        print(i)
        if i in opened:
            stack.append(i)
            continue
        if not stack and i in closed:
            return "NO"
        if i == m.get(stack[-1]):
            stack.pop()
        else:
            stack.append(i)
    if len(stack) == 0:
        return "YES"
    return "NO"
