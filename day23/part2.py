# Copy pasted from stackoverflow
def is_prime(n):
    if n == 2 or n == 3: return True
    if n < 2 or n%2 == 0: return False
    if n < 9: return True
    if n%3 == 0: return False
    r = int(n**0.5)
    f = 5
    while f <= r:
        if n%f == 0: return False
        if n%(f+2) == 0: return False
        f +=6
    return True

def run():
    a = 1
    b = c = d = e = f = g = h = 0
    b = 81
    c = b
    if a != 0:
        b = 100 * b
        b = b + 100000
        c = b
        c = c + 17000
    while 1:
        f = 1
        d = 2
        # while 1:
        #     e = 2
        #     while 1:
        #         pass
        #         g = d
        #         g = g * e
        #         g = g - b
        #         if g != 0:
        #             f = 0
        #         e = e + 1
        #         g = e
        #         g = g - b
        #         if g == 0:
        #             break
        #     d = d + 1
        #     g = d
        #     g = g - b
        #     if g == 0:
        #         break
        # This 2 loops are actually just a Prime Check! I honsetly figured it out myself!
        f = int(is_prime(b))
        if f == 0:
            h = h + 1
        g = b
        g = g - c
        print(locals())
        if g != 0:
            b = b + 17
        else:
            return h

if __name__ == '__main__':
    print("H={}".format(run()))


