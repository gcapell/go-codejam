#!/usr/bin/env python

def f(d,b):
	print 'f({d}, {b})'.format(**locals())
	if b == 1:
		return d
	if d == 1:
		return 1
	return 1 + f(d-1, b) + f(d-1, b-1)

print f(5,3)
print
print f(3,3)
print f(14,3)
