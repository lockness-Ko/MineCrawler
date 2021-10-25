import os
from random import randint as ri
import sys
import numpy as np

##############################################
# HELPER STUFF!!

HEADER = '\033[95m'
OKBLUE = '\033[94m'
OKCYAN = '\033[96m'
OKGREEN = '\033[92m'
WARNING = '\033[93m'
FAIL = '\033[91m'
ENDC = '\033[0m'
BOLD = '\033[1m'
UNDERLINE = '\033[4m'
INFO = ENDC+OKBLUE+'[-] '
LOSE = ENDC+FAIL+BOLD+'[!] '
WIN = ENDC+OKGREEN+'[*] '

def cls():
    	print("\033[H\033[2J")

def pos(x, y):
    	os.system("echo -n '\033["+str(int(y/2))+";"+str(int(x))+"H'")
     
wall_vert = "║"
wall_horiz = "═"
wall_tl = "╔"
wall_tr = "╗"
wall_bl = "╚"
wall_br = "╝"
exit = "╬"
char_player = "옷"
char_robot = "🤖"

##############################################

def makeRoom(x, y, width, height):
    height = int(height/2)
    for i in range(height):
        pos(x, y+i*2)
        if i == 0:
            print(wall_tl+wall_horiz*width+wall_tr)
        elif i == height - 1:
            print(wall_bl+wall_horiz*width+wall_br)
        else:
            print(wall_vert+" "*width+wall_vert)

def printPlayer(x,y):
    pos(x,y)
    print(char_player)

def playerOutsideBox(px,py,bx,by,bw,bh):
    if (px <= bx or px >= bx+bw) or (py <= by or py >= by+bh):
        return True
    return False

objects = []

class Object():
    def __init__(self, x,y,name,char):
        self.x = x
        self.y = y
        self.name = name
        self.char = char
    def draw(self):
        if self.name == "bad":
            self.x += ri(-1,1)*2
            self.y += ri(-1,1)*2
        pos(self.x, self.y)
        print(self.char)

def main():
    windows = []

    player_x = 10
    player_y = 10
    player_health = 100
    player_kills = 0
    player_gold = 0

    terminal_size = os.get_terminal_size()
    
    for i in range(5):
        objects.append(Object(ri(2,48),ri(2,48),"gold","🪙"))
    
    for i in range(5):
        objects.append(Object(ri(2,48),ri(2,48),"bad",char_robot))
    
    while True:
        cls()
        makeRoom(0,0,50,50)
        
        for ind, i in enumerate(np.array(objects)):
            if abs(player_y-i.y)<=2 and abs(player_x-i.x)<=2:
                if i.name == "gold":
                    player_gold += 1
                    objects.pop(ind)
                if i.name == "bad":
                    player_health -= 10
            i.draw()
        
        printPlayer(player_x, player_y)
        if playerOutsideBox(player_x, player_y, 0, 0, 50, 50):
            player_health -= 50
        if player_health <= 0:
            print("You Died!")
            a = input("Would you like to revive for 50 gold?\n>").upper()
            if a == "Y" or a == "YES":
                if player_gold >= 50:
                    player_gold -= 50
                    player_health = 100
                else:
                    print("Not enough gold!")
                    print("Bye!")
                    print()
                    print("Health: "+str(player_health))
                    print("Kills: "+str(player_kills))
                    print("Gold: "+str(player_gold))
                    sys.exit(0)
            else:
                print("Bye!")
                print()
                print("Health: "+str(player_health))
                print("Kills: "+str(player_kills))
                print("Gold: "+str(player_gold))
                sys.exit(0)
                    


        ##############################################
        # STATS
        pos(terminal_size[0] - 10, 0)
        print("Stats: ")
        pos(terminal_size[0] - 10, 4)
        print("Health: "+str(player_health))
        pos(terminal_size[0] - 10, 6)
        print("Kills: "+str(player_kills))
        pos(terminal_size[0] - 10, 8)
        print("Gold: "+str(player_gold))
        #
        ##############################################
        
        ##############################################
        # MOVEMENT
        pos(0, terminal_size[1]*2)
        a = input(">")
        # input()
        if a=="w":
            player_y-=2
        elif a=="s":
            player_y+=2
        elif a=="a":
            player_x-=2
        elif a=="d":
            player_x+=2
        elif a=="q":
            cls()
            print("Bye!")
            print()
            print("Health: "+str(player_health))
            print("Kills: "+str(player_kills))
            print("Gold: "+str(player_gold))
            sys.exit(0)
        #
        ##############################################

if __name__ == '__main__':
    main()