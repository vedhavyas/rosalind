from __future__ import division

def fact(n, c):
    if n<= 0:
        return c

    return fact(n-1, c*n)

def comb(n, r):
    return fact(n, 1)/(fact(r, 1)*fact(n-r, 1))

def get_probability(k, m, n):
    tp = comb(k+m+n, 2)
    cc = comb(k, 2) + comb(k, 1)*comb(m, 1) + comb(k, 1)*comb(n, 1) + 0.5*comb(m, 1)*comb(n, 1) + 0.75*comb(m, 2)
    return cc/tp

print get_probability(15, 30, 17)