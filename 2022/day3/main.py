import string

total = 0
alphabet = list(string.ascii_letters)

with open("input.txt", 'r') as f:
    for i in f:
        half = len(i) / 2
        for x in i:
            for y in i:
                if y == x:
                    common = x
                    break
        temp = alphabet.index[common]
        total += temp