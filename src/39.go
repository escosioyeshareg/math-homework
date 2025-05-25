def is_prime(n):
    if n <= 1:
        return False
    for i in range(2, int(n**0.5) + 1):
        if n % i == 0:
            return False
    return True

def gcd(a, b):
    while a != 0:
        temp = a
        a = b % a
        b = temp
    return b

def lcm(a, b):
    return abs((a*b) // gcd(a, b))

# Example usage
num1 = 15
num2 = 30
result = is_prime(num1)
print(f"{num1} and {num2} are {result}.")
