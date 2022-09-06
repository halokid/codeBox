digits = input("input digits: ")
num = input("output num: ")

# todo: why here initial 2 seconds for `msec`? because the scence is
# todo: below, such as input num '210`:
# todo: 1. use one finger press key `2`, because use one finger, this takes 1 second
# todo: 2. move finger to key `1`
# todo: 3. move finger to key `0`
# todo: 4. then press key 'enter` to finish the operation, takes 1 second
# todo: so msec initial `2` seconds
total_sec = 2
double = False
tmp_index = 0
for char in num:
    char_index = digits.index(char)
    print("char_index -->>>", char_index)
    if double:
        tmp_sec = abs(int(char_index) - int(tmp_index))
        print("tmp_sec -->>>", tmp_sec)
        total_sec += tmp_sec

    tmp_index = char_index
    print("tmp_index -->>>", tmp_index)
    double = True
    print("==========================")
    
print(total_sec)