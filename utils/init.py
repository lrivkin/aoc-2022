import os
from pathlib import Path

p = Path(__file__).parent.resolve().joinpath("..")
# p.joinpath("..")
for day in range(1, 26):
    p.joinpath(f"day{day}").mkdir(exist_ok=True)
    p.joinpath(f'day{day}/test.txt').touch(exist_ok=True)
    p.joinpath(f'day{day}/input.txt').touch(exist_ok=True)
    p.joinpath(f'day{day}/day{day}.go').touch(exist_ok=True)