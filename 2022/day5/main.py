import string

def create_stacks():
    with open("input.txt", "r") as file:
        lines = [line.rstrip('\n') for line in file]
        for line in lines:
            line = line.strip(" ")
            if line != "" and line[0].isdigit() :
                num_stacks = int(line[-1])
        stacks = {i+1: [] for i in range(num_stacks)}
    return stacks, num_stacks


def add_to_stacks(stacks):
    col_data = []; l = 1
    for x in stacks:
        with open("input.txt", "r") as file:
            for line in file:
                line = line.strip("\n")
                if line != "":
                    col_data.append(line[l])
        stacks[x] = col_data; l += 4; col_data = []
    return stacks


def cleanup_stacks(stacks):
    new_stacks = []
    for x in stacks:
        for y in stacks[x]:
            if y in string.ascii_uppercase or y == ' ':
                new_stacks.append(y)
            stacks[x] = new_stacks
        new_stacks = []


def move_cargo(stacks):
    with open("input.txt", "r") as file:
        for line in file:
            if "move" in line:
                for x in ["\n", "from", "to", "move"]:
                    line = line.replace(x, "")
                line = line.split(" ")
                amount = line[1]; cargo_from = line[3]; cargo_to = line[5]


stacks, num_stacks = create_stacks()
stacks = add_to_stacks(stacks)
stacks = cleanup_stacks(stacks)
move_cargo(stacks)