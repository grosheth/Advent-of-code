# A = Rock 1
# B = Paper 2 
# C = Scissor 3

# X = Rock
# Y = Paper
# Z = Scissor

# X = LOSS
# Y = DRAW
# Z = WIN

# 3 pts draw
# 6 pts win

# Part 1

with open("input.txt", 'r') as f:
    total = 0
    for i in f:
        if i[2] == "X":
            total += 1
            if i[0] == "C":
                total += 6
            elif i[0] == "A":
                total += 3
        elif i[2] == "Y":
            total += 2
            if i[0] == "A":
                total += 6
            elif i[0] == "B":
                total += 3
        elif i[2] == "Z":
            total += 3
            if i[0] == "B":
                total += 6
            elif i[0] == "C":
                total += 3
    print(f"you have: {total} points")

# Part 2

with open("input.txt", 'r') as f:
    total = 0
    for i in f:
        if i[2] == "X":
            if i[0] == "A": total += 3
            elif i[0] == "B": total += 1
            else: total += 2
        if i[2] == "Y":
            total += 3
            if i[0] == "A": total += 1
            elif i[0] == "B": total += 2
            else: total += 3
        if i[2] == "Z":
            total += 6
            if i[0] == "A": total += 2
            elif i[0] == "B": total += 3
            else: total += 1
    print(f"you have: {total} points")

