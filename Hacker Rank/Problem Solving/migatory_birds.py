# Complete the 'migratoryBirds' function below.
# The function is expected to return an INTEGER.
# The function accepts INTEGER_ARRAY arr as parameter.
def migratoryBirds(arr):
    # Write your code here
    bird_count = [0] * 5
    for bird in arr:
        bird_count[bird - 1] += 1
    return bird_count.index(max(bird_count)) + 1