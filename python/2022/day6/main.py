with open("input.txt", "r") as file:
    signal = file.readline()

def get_marker(signal, marker_length):
    read, count = signal[:marker_length], marker_length
    for i in range(marker_length, len(signal)):
        if len(set(read[-marker_length:])) == marker_length:
            break

        read += signal[i]
        count += 1

    return count


# count = get_marker(signal, 4)

count = get_marker(signal, 14)
print(count)