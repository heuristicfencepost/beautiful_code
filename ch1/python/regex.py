def match(regex,text):
    if regex[0] == '^':
        return matchhere(regex[1:],text)
    for key in range(len(text)):
        if matchhere(regex,text[key:]) == 1:
            return 1
    return 0

def matchhere(regex,text):
    if len(regex) == 0:
        return 1
    if len(regex) > 1 and regex[1] == '*':
        return matchstar(regex[0],regex[2:],text)
    if regex[0] == '$' and len(regex) == 1:
        if (len(text)) == 0:
            return 1
        else:
            return 0
    if len(text) > 0 and (regex[0] == '.' or regex[0] == text[0]):
        return matchhere(regex[1:],text[1:])
    return 0

def matchstar(c,regex,text):
    for key in range(len(text)):
        if matchhere(regex,text[key:]) == 1:
            return 1
        if text[key] != c and text[key] != '.':
            break
    return 0

