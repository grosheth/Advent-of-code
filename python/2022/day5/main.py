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
        with open("input2.txt", "r") as file:
            for line in file:
                line = line.strip("\n")
                if line != "":
                    col_data.append(line[l])
        stacks[x] = col_data; l += 4; col_data = []
    print(stacks)
    return stacks


def cleanup_stacks(stacks):
    new_stacks = []
    for x in stacks:
        for y in stacks[x]:
            if y in string.ascii_uppercase:
                new_stacks.append(y)
            stacks[x] = new_stacks
        new_stacks = []
    return stacks

# Part 1
# def move_cargo(stacks):
#     with open("input.txt", "r") as file:
#         for line in file:
#             if "move" in line:
#                 for x in ["\n", "from", "to", "move"]:
#                     line = line.replace(x, "")
#                 line = line.split(" ")
#                 amount = line[1]; cargo_from = line[3]; cargo_to = line[5]
#                 for x in range(int(amount)):
#                     cargo = stacks[int(cargo_from)][0]; stacks[int(cargo_from)].pop(0); stacks[int(cargo_to)].insert(x,cargo)
#     return stacks

# part 2
def move_cargo(stacks):
    with open("input.txt", "r") as file:
        for line in file:
            if "move" in line:
                for x in ["\n", "from", "to", "move"]:
                    line = line.replace(x, "")
                line = line.split(" ")
                amount = line[1]; cargo_from = line[3]; cargo_to = line[5]
                for x in range(int(amount)):
                    cargo = stacks[int(cargo_from)][0]; stacks[int(cargo_from)].pop(0); stacks[int(cargo_to)].insert(x,cargo)
    return stacks

def get_top_crate():
    letter_list = []
    for x in range(1,len(stacks)+1):
        letter_list.append(stacks[x][0])
    answer = ''.join(str(x) for x in letter_list)
    print(answer)



stacks, num_stacks = create_stacks()
stacks = add_to_stacks(stacks)
stacks = cleanup_stacks(stacks)
stacks = move_cargo(stacks)
get_top_crate()