import requests
from random import choice

term = input("What would you like to search for? ")
url = 'https://icanhazdadjoke.com/search'
res = requests.get(
    url,
    headers = {"Accept": "application/json"},
    params={"term": term}
    ).json()

num_jokes = res["total_jokes"]
results = res["results"]
if num_jokes > 1: 
    print("There are many jokes!")
    print(choice(results)["joke"])
elif num_jokes == 1: 
    print("There is one joke")
    print(results[0]['joke'])
else: 
    print(f"Sorry, couldn't find a joke with your term: {term}")
