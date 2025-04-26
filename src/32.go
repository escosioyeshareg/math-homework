import random
def go_code():
    # Generate some random characters and shuffle them
    shuffled_chars = ''.join(random.sample('abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ', 6))
    print(shuffled_chars)
    # Shuffle these characters to simulate a maze pattern
    for _ in range(10):
        shuffled_chars = ''.join(random.sample(shuffled_chars, len(shuffled_chars)))
        print(shuffled_chars)
go_code()
