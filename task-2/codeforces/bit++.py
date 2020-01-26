x = 0
for i in range(int(input())):
    char = input()
    if char[0] == '+' or char[1] == '+':
        x+=1
    else:
        x-=1
print(x)
