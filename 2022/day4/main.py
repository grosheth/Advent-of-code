
total = 0; total_partial= 0
# Part 1
with open("input.txt", "r") as f:
    for i in f:
        no_newline = i.replace("\n", ""); splitted = no_newline.split(","); first = splitted[0].split("-"); second = splitted[1].split("-")

        first_cover = [num for num in range(int(first[0]), int(first[1]) + 1)]
        second_cover = [num for num in range(int(second[0]), int(second[1]) + 1)]

        if all(x in first_cover for x in second_cover) or all(x in second_cover for x in first_cover):
            total += 1

print(total)


# Part 2
with open("input.txt", "r") as f:
    for i in f:
        no_newline = i.replace("\n", ""); splitted = no_newline.split(","); first = splitted[0].split("-"); second = splitted[1].split("-")

        first_cover = [num for num in range(int(first[0]), int(first[1]) + 1)]
        second_cover = [num for num in range(int(second[0]), int(second[1]) + 1)]

        if any(x in first_cover for x in second_cover) or any(x in second_cover for x in first_cover):
            total_partial += 1

print(total)