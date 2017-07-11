#!coding=utf-8

'''

def decorate_A(function):
    def wrap_function(*args, **kwargs):
        kwargs['str'] = 'hello'
        return function(*args, **kwargs)
    return wrap_function



@decorate_A
def print_message_A(*args, **kwargs):
    print(kwargs['str'])


print_message_A()


'''


def decorate_B(function):
    def wrap_function(*args, **kwargs):
        str = 'hello'
        return function(str, *args, **kwargs)
    return wrap_function


@decorate_B
def print_message_B(str, *argsg, **kwargs):
    print(str)

print_message_B()



