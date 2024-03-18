fs = {}; path = ''; total = 0; partial = 0; maximum = 100000; minimum = 0

def create_directory(fs:dict, path:str):
    directory = {path: []}
    fs.update(directory)
    return fs

def create_file(fs:dict, line:list):
    fs[path].append(line[0])
    return fs

def change_directory(path:str):
    if path == '':
        path = path + line[2]
    elif line[2] == "..":
        path = path.split('/'); path.pop(-1); path.pop(-1)
        if path == '':
            path.append('/'); path = ('').join(path)
        else:
            path = ('/').join(path); path = path + '/'
    else:
        path = path + line[2] + '/'
    return path

def join_files():
    for x in fs:
        for y in fs:
            if x != y and x in y:
                for w in fs[y]:
                    fs[x].append(w)
    return fs

# # part 1
# def get_total(total:int, partial:int):
#     for x in fs:
#         for y in fs[x]:
#             partial = partial + int(y)
#         if maximum >= partial:
#             total = total + partial
#         partial = 0
#     return total

# part 2
def get_minimum():
    used = 0
    for y in fs["/"]:
        used = used + int(y)
    free_space = 70000000 - used
    for x in range(30000000):
        if free_space < 30000000:
            free_space += 1
        else:
            minimum = x + 1
            break
    return minimum

def get_total(total:int, partial:int):
    for x in fs:
        for y in fs[x]:
            partial = partial + int(y)
        if partial >= minimum:
            if total == 0:
                total = partial
            elif partial < total:
                total = partial
        partial = 0
    return total

with open("input.txt", "r") as file:
    lines = [line.rstrip('\n') for line in file]
    for line in lines:
        line = line.split(' ')
        if line[1] == "cd" and line[2] == "..":
            path = change_directory(path)
        elif line[1] == "cd":
            path = change_directory(path)
            fs = create_directory(fs, path)
        if line[1] != "ls" and line[0] != "dir" and line[1] != "cd":
            fs = create_file(fs, line)
    fs = join_files()
    minimum = get_minimum()
    total = get_total(total, partial)
    print(total)
