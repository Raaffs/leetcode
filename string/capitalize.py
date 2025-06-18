def capitalize(sentence: str):
    formatted = ""
    i = 0
    while i < len(sentence):
        if i == 0:
            formatted += sentence[i].upper()
            i += 1
            continue
        if sentence[i] in [".", "?"] and i + 2 < len(sentence):
            formatted += sentence[i]
            formatted += sentence[i + 1]  # usually a space
            formatted += chr(ord(sentence[i+2])-32)
            i += 3
        else:
            formatted += sentence[i]
            i += 1
    print(formatted)

sentence = "hello, how are you? if you're ok replys yes. otherwise reply no."
capitalize(sentence)
