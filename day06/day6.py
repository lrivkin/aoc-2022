def find_idx(input: str, num_chars: int):
    seen = {}
    min_idx = 0
    for i in range(len(input)):
        c = input[i]
        if c in seen:
            new_min = seen[c] + 1
            for j in range(min_idx, new_min):
                del seen[input[j]]
            min_idx = new_min

        seen[c] = i
        if i - min_idx > num_chars:
            """drop off the trailing one"""
            del seen[input[min_idx]]
            min_idx += 1

        if len(seen) == num_chars:
            print(f'{num_chars} unique characters \t{i+1} ({"".join(sorted(seen.keys(), key=lambda k: seen[k]))})')
            return

with open("test.txt", "r") as f:
    tests = f.readlines()
    print("Tests")
    for t in tests:
        find_idx(t, 4)
        find_idx(t, 14)

with open("input.txt", "r") as f:
    real_input = f.readlines()
    print(f"\nMy input:")
    find_idx(real_input[0], 4)
    find_idx(real_input[0], 14)
