n=int(input())
for i in range(0,n):
          word=input("")
          if len(word)>10:
                    
            print(word[0],len(word)-2,word[-1])
          else:
            print(word)
