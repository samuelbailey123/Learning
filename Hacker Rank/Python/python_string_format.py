# Complete the print_formatted function in the editor below.
# print_formatted has the following parameters:
# int number: the maximum value to print
def print_formatted(number):
    # your code goes here
    w = len(bin(number)[2:])
    for i in range(1, number+1):
        print(str(i).rjust(w, ' '), str(oct(i)[2:]).rjust(w, ' '), str(hex(i)[2:].upper()).rjust(w, ' '), str(bin(i)[2:]).rjust(w, ' '))