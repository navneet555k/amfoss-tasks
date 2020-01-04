n,m,a =map(int,input().split())
if   n*m%a**2==0:
     p=n*m//a**2
     print (p)
elif n*m<a*a:
       p=1
       print(p)
else:
          p=(n*ma**2)+2
          print(p)
