data = []
with open('data') as f:
    for row in f:
        data.append([int(x) for x in row.split()])

chksum = sum(max(line) - min(line) for line in data)
print(chksum)