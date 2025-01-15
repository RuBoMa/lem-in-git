# lem-in

This project is a digital version of an ant farm simulation.

# Objectives
This project is meant to make you code a digital version of an ant farm.

Create a program lem-in that will read from a file (describing the ants and the colony) given in the arguments.

Upon successfully finding the quickest path, lem-in will display the content of the file passed as argument and each move the ants make from room to room.

How does it work?

You make an ant farm with tunnels and rooms.
You place the ants on one side and look at how they find the exit.
You need to find the quickest way to get n ants across a colony (composed of rooms and tunnels).

At the beginning of the game, all the ants are in the room ##start. The goal is to bring them to the room ##end with as few moves as possible.
The shortest path is not necessarily the simplest.
Some colonies will have many rooms and many links, but no path between ##start and ##end.
Some will have rooms that link to themselves, sending your path-search spinning in circles. Some will have too many/too few ants, no ##start or ##end, duplicated rooms, links to unknown rooms, rooms with invalid coordinates and a variety of other invalid or poorly-formatted input. In those cases the program will return an error message ERROR: invalid data format. If you wish, you can elaborate a more specific error message (example: ERROR: invalid data format, invalid number of Ants or ERROR: invalid data format, no start room found).
You must display your results on the standard output in the following format :

number_of_ants
the_rooms
the_links

Lx-y Lz-w Lr-o ...

# Usage:

Run the program with "go run . (filename)"
if you want to see the time program takes to run simply add time in front of go run "time go run . (filename)"

# Visual representations

This image is showing what everything means in the example files.

![alt text](image.png)

# Instructions

**Prerequisites:**

Go installed on your system.

# Authors:

Fanni, Roope & Toft