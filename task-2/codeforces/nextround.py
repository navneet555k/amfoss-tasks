l=input().split()
for i in range(len(l[0])):
    points=input().split()
count=0
for i in range(int(l[0])):
    if int(points[i])>=int(points[int(l[1])-1]) :
        count=count+1    
print (count)
