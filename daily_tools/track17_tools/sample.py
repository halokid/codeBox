#!/usr/bin/env python
#!coding=utf-8


import tracker17
import json


t = tracker17.Tracker()
r = t.track("JJD0002258743111220")
print r


print "\n..........................\n"

print json.dumps(r)
