pa = 0 #initial points of alice
pb = 0#intial points of bob
atriplet= input().split() #taking 3 inputs in one line
btriplet = input().split()

for i in range(len(atriplet )):
          
    if int(atriplet[i])>int(btriplet [i]):#comparing values place by place
        pa+= 1
        
    if int(atriplet[i])<int(btriplet[i]):
        pb += 1
print(pa,pb)
