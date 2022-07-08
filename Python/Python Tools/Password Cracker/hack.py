from itertools import product
import time
import random

def printf(format, *args):
    print(format % args)

def main():
    first_time = time.time()
    choices = 'abcdefghijklmnopqrstuvwxyz0123456789!@#$&^*ABCDEFGHIJKLMNOPQRSTUVWXYZ'
    random_password = ''.join(random.choice(choices) for _ in range(9))
    print('The random password is',random_password)
    for passcode in product(choices, repeat=9):
        if ''.join(passcode) == random_password:
            printf('Found password: %s', ''.join(passcode))
            break
    last_time = time.time()
    printf('Time taken: %.2f seconds', last_time - first_time)

main()
