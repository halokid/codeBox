#coding=utf-8
num = [2, -5, 9, 7, -2,  5, 3, 1, 0, -3, 8]

positive_num_cnt = 0
positive_num_sum = 0

for i in range(len(num)):
    if num[i] > 0:
        positive_num_cnt += 1
        positive_num_sum += num[i]


if positive_num_cnt > 0:
    average = positive_num_sum / positive_num_cnt




print average
