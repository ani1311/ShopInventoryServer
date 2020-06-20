import random

file = open("testItems.csv","w")

for i in range(20):
    name= "item" + str(i)
    bc = str(int((random.random()*10000) % 10000))
    file.write('"' + name + '",' + bc + ',' + str(random.random()*1000) + '\n')

file.close()

