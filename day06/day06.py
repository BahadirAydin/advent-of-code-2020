f = open("input.txt","r")
lines = f.readlines()

curr = set()
sum = 0
flag = True
for i in range(len(lines)):
    line = lines[i]
    print(curr,sum)
    if line == '\n':
        sum += len(curr)-1
        curr = set()
        flag = True
    else:
        if flag:
            for c in line:
                curr.add(c)
            flag = False
        else:
            for char in curr.copy():
                if char not in line:
                    curr.remove(char)

print(sum)
