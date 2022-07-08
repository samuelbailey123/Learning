# import statements
from time import sleep
from functions.killer import kill
from selenium import webdriver
from selenium.webdriver.chrome.service import Service
from webdriver_manager.chrome import ChromeDriverManager

# main
if __name__ == '__main__':
       driver = webdriver.Chrome(service=Service(ChromeDriverManager().install()))
       url = '''https://i.prefinery.com/projects/wynknsfn/users/new?display=inline&version=2&_ga=2.144163317.1546534149.1656341654-1539424012.1656341654&creation_location=https%3A%2F%2Fwww.wonsulting.com%2Fresumai%23join-waitlist&creation_location_title=ResumAI%20by%20Wonsulting&referrer=https%3A%2F%2Fwww.google.com%2F&referral_token=bailey407n'''
       for _ in range(10):
              driver.get(url)
              sleep(2)
              print('Website Loaded')
              kill(driver)
              driver.back()
              sleep(2)
              