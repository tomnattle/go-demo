import time
h = int(time.strftime("%M", time.localtime()))
if(h == 0):
    print(55)
else:
    if(h%5 == 0):
        h -= 1
    print(int(h) - int(h) % 5)
# 1