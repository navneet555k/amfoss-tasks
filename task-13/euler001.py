suma=0
pum=0
lum=0
n=int(input())
for l in range(0,n):
        num=int(input())
        for i in range(1,num):
                if  i%3==0:
                 suma=suma +i
        for j in range (1,num):
                if j%5==0:
                   pum=pum+j
        for k in range(1,num):
                  if k%15==0:
                    lum=lum+k        
        tot=suma  + pum -lum
        print(tot)
