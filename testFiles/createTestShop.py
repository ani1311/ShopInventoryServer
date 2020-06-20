import random

file = open("testShops.csv","w")

for i in range(20):
    shopNO = "shop" + str(i)
    passNO = "pass" + str(i)
    lat = str(21.139293 + random.random()/100)
    lon = str(79.0992126 + random.random()/100)
    file.write('"' + shopNO + '",' + lat + ',' + lon + ',"' + passNO + '"\n')

file.close()
