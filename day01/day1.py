with open('day1-input.txt') as f:
    file = f.read()

file_split = file.split('\n\n')
elves = [[int(x) for x in e.split('\n')] for e in file_split]
cals = sorted([sum(x) for x in elves], reverse=True)
print(sum(cals[0:3]))