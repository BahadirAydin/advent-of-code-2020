f = open("input.txt","r")
lines = f.readlines()

curr = set()
sum = 0
for line in lines:
    if line == '\n':
        sum += len(curr)-1
        curr = set()
    else:
        for char in line:
            curr.add(char)

print(sum)
