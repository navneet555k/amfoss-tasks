
n=int(input())
arr=[]
prime=[]
for i in range(n):
    temp=[]
    a=int(input())
    arr.append(a)
    for k in range(1,a):
        if(a%k==0):
            temp.append(k)
    prime.append(temp)
for i in range(len(prime)):
    if(len(prime[i])==1):
        print(arr[i])
    else:
        print(max(prime[i]))
