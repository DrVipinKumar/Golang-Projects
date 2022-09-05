import time
mylist=[]
start=time.time()
for i in range(1000):
    mylist.append(i)
    print(str(mylist[i]))

end = time.time()
print(f"{(end-start)*1000} ms")