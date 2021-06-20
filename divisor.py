import time

x = int(input())

start = time.time()

answer_list = []

for i in range(1, x+1):
  if x % i == 0:
    answer_list.append(i)

print("answer: ", answer_list)

print(f"elapsed_time:{round((time.time() - start), 1)}[sec]")