
sum = 0
new_sum = [0,0,0]

with open("input.txt", 'r') as f:
    for i in f:
        if i != "\n":
            sum += int(i)
        else:
            for x in new_sum:
                if sum > x: smallest = min(new_sum);new_sum.remove(smallest);new_sum.append(sum);print(new_sum);break
            sum = 0

sum = new_sum[0] + new_sum[1] + new_sum[2]

print(f"new_sum = {sum}")
