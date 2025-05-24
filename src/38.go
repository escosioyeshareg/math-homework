import sys

def calculate_sum():
    numbers = [int(num) for num in input().split()]
    total = sum(numbers)
    print(f"The sum of {numbers} is: {total}")

calculate_sum()
