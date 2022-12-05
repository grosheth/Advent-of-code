import string

total = 0
alphabet = list(string.ascii_letters)
common = ""

# Part 1

# with open("input.txt", 'r') as f:
#     for i in f:
#         first_half,second_half = [], []
#         half = (len(i) -1) / 2; counter = len(i) -1; count = 0
#         while count != half:
#             first_half.append(i[count]);count += 1
#         while count != counter:
#             second_half.append(i[count]);count += 1
#         for x in first_half:
#             for y in second_half:
#                 if x == y:
#                     common = x
#         temp = alphabet.index(common)
#         total += temp +1
#         print(total,common)

# Part 2
line_count = 0; block1,block2,block3 = [],[],[]
with open("input.txt", 'r') as f:
    for i in f:
        counter = len(i) -1; count = 0;
        if line_count == 0:
            while count != counter:
                block1.append(i[count]);count += 1
                if count == counter: line_count +=1
        if line_count == 1:
            while count != counter:
                block2.append(i[count]);count += 1
                if count == counter: line_count +=1
        if line_count == 2:
            while count != counter: 
                block3.append(i[count]);count += 1
                if count == counter: line_count +=1
        if line_count == 3:
            for x in block1:
                for y in block2:
                    if x == y:
                        for w in block3:
                            if x == w:
                                common = x ;line_count = 0 ; temp = alphabet.index(common); total += temp +1; block1,block2,block3 = [],[],[]
                                print(total,common)
                                break

