import gevent.monkey
gevent.monkey.patch_socket()

import gevent
import urllib.request
import simplejson as json

def fetch(pid):
    response = urllib.request.urlopen('http://quan.suning.com/getSysTime.do')
    result = response.read()
    json_result = json.loads(result)
    datetime = json_result['sysTime2']

    print('Process %s: %s' % (pid, datetime))
    return json_result['sysTime2']

def synchronous():
    for i in range(1,10):
        fetch(i)

def asynchronous():
    threads = []
    for i in range(1,10):
        threads.append(gevent.spawn(fetch, i))
    gevent.joinall(threads)

print('Synchronous:')
synchronous()

print('Asynchronous:')
asynchronous()

