fs = {}; path = ['']
# fs = {"/": {"a": {}}}
def create_directory(fs:dict, line:list, path:str):
    directory = {line[2]: {}}
    if path[0] != '':
        # print(path)
        for y in fs:
            print(y, fs)
            fs[y].update(directory)
    else:
        fs.update(directory); path.pop(0)
    path.append(line[2])
    # print(path , fs)
    return fs, path

# def create_file(fs:dict, line:list):
#     directory = {line[2]: {}}
#     fs.update(directory)
#     print(fs)

def change_directory(path:str):
    path.append(line[2])
    return path

with open("input.txt", "r") as file:
    lines = [line.rstrip('\n') for line in file]
    for line in lines:
        line = line.split(' ')
        if line[1] == "cd" and line[1] == "..":
            path = change_directory(path)
        elif line[1] == "cd":
            fs, path = create_directory(fs, line, path)
        # if line[1] == "ls":
        #     fs = create_file(fs, line)
        