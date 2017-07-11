#coding=utf-8

def fuck(fn):
    print "fuck %s!" % fn.__name__[::-1].upper()
    print fn.__name__

@fuck
def wfg():
    pass



