left_list = []
right_list = []

with open("input.txt", "r") as file: 
    for line in file:
        left, right = map(int, line.strip().split())
        left_list.append(left)
        right_list.append(right)

cnt = {}

for num in right_list:
    if num in cnt:
        cnt[num] += 1
    else:
        cnt[num] = 1

ans = 0
for num in left_list:
    if num in cnt:
        ans += num * cnt[num]

print(ans)