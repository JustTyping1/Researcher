import webbrowser
from subprocess import call

call(["./scraper"])

file = open("data.txt", "r")

lines = file.readlines()

print(lines[0])
print(lines[1])

webbrowser.open(lines[len(lines)-1])

