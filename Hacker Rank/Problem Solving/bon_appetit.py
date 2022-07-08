# Complete the 'bonAppetit' function below.
#
# The function accepts following parameters:
#  1. INTEGER_ARRAY bill
#  2. INTEGER k
#  3. INTEGER b
#

def bonAppetit(bill, k, b):
    # Write your code here
    bill.pop(k)
    if sum(bill) / 2 == b:
        print("Bon Appetit")
    else:
        print(int(b - sum(bill) / 2))