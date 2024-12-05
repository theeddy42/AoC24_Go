def is_safe(levels):
    inc = levels[1] - levels[0]
    for i in range(1, len(levels)):
        cur = levels[i] - levels[i - 1]
        if abs(cur) not in range(1, 4):
            return False
        if inc < 0 and cur > 0 or inc > 0 and cur < 0:
            return False
    return True

safe_levels = list()
with open("input.txt") as file:
    safe1, safe2 = 0, 0
    for line in file:
        levels = [int(level) for level in line.split()]
        if is_safe(levels):
            safe1 += is_safe(levels)
            continue
        for i in range(len(levels)):
            new_levels = levels[:]
            new_levels.pop(i)
            if is_safe(new_levels):
                safe2 += 1
                safe_levels.append(new_levels)
                break

print(safe1)
print(safe1 + safe2)
print(safe_levels)