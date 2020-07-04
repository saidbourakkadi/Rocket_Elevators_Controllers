

import time
print("Welcome, my name is Elevator Sam, I m a pleasure to give you a good services")


class TimerError(Exception):
    """A custom exception used to report errors in use of Timer class"""


class Timer:
    def __init__(self):
        self._start_time = None

    def start(self):
        """Start a new timer"""
        if self._start_time is not None:
            raise TimerError(f"Timer is running. Use .stop() to stop it")

        self._start_time = time.perf_counter()

    def stop(self):
        """Stop the timer, and report the elapsed time"""
        if self._start_time is None:
            raise TimerError(f"Timer is not running. Use .start() to start it")

        elapsed_time = time.perf_counter() - self._start_time
        self._start_time = None
        print(f"Elapsed time: {elapsed_time:0.4f} seconds")


# Class Elevators constructor
class Cage:
    def __init__(self, id, levelActual, direction):
        self.id = id
        self.battery = 1  # sorry not used no time
        self.column = 1  # sorry not used no time
        self.statusDoor = "closed"  # sorry not used no time
        self.position_door_x = 0  # close: 0 / open : 150
        self._start_time = time.perf_counter()
        self.timer = 0
        self.timing_door = 3000
        self.timing_floor = 20000
        self.destination = []
        self.listUp = []
        self.listDown = []
        self.listPoubelle = []
        self.counterStepCage = 0
        self.direction = direction
        self.levelActual = levelActual
        self.sensor_status_cage = "Clear"  # sorry not used no time
        self.close_button = "inactif"  # sorry not used no time
        self.open_button = "inactif"  # sorry not used no time

    def display_status_cage(self):
        print('status_cage: ' + str(self.id) + ' , level Actual: ' +
              str(self.levelActual) + ' , direction: ' + str(self.direction))

    def call_level(self, levelCall):

        print('level call: '+str(levelCall))
        if levelCall == -1 or levelCall == 0:
            self.listPoubelle.append(levelCall)

        elif self.levelActual > levelCall and (levelCall != 0):
            if levelCall in self.listDown:
                print('this number is already Actif')
            else:
                self.listDown.append(levelCall)
                self.listDown.sort(reverse=True)

            #print("level actual is biggest then the level call so we down on "+str(step)+' steps')
        elif self.levelActual < levelCall and (levelCall != 0):
            if levelCall in self.listUp:
                print('this number is already Actif')
            else:
                self.listUp.append(levelCall)
                self.listUp.sort(reverse=False)
            #print("level actual is lower then the level call so we up on "+str(step)+' steps')
        else:
            print(
                "level actual is the same then the call so we open the door for you you arrived")
            open_close_door(self)
        print('listUp: ' + str(self.listUp))
        print('listDown: ' + str(self.listDown))

    def move_up(self):
        self.direction = 'up'
        for list in self.listUp:
            print("level actual is: " + str(self.levelActual))
            print("the direction in this cage is : " + self.direction)
            print("we go in the floor: " + str(list))
            self.timer = 0
            while self.levelActual < list:
                #print("For each floor it took 5 seconds  to go up ")
                while self.timer < 5:
                    #print("For each floor it took " + str(self.timer) + " seconds  to go up ")
                    self.timer += 1
                    time.sleep(1)
                self.levelActual += 1
                print("Display actualy: " + str(self.levelActual) + " th floor")
            self.open_close_door()
        del self.listUp[:]
        print("the listUp is empty now :" + str(self.listUp))
        #print("the direction in this cage is : " + self.direction)

    def move_down(self):
        self.direction = 'down'
        for list in self.listDown:
            print("level actual is: " + str(self.levelActual))
            print("the direction in this cage is : " + self.direction)
            print("we go in the floor: " + str(list))
            self.timer = 0
            while self.levelActual > list:
                #print("For each floor it took 5 seconds  to go up ")
                while self.timer < 5:
                    #print("For each floor it took " + str(self.timer) + " seconds  to go up ")
                    self.timer += 1
                    time.sleep(1)
                self.levelActual -= 1
                print("Display actualy: " + str(self.levelActual) + " th floor")
            self.open_close_door()
        del self.listDown[:]
        self.direction = "idle"
        print("the listDown is empty now :" + str(self.listDown))
        print("the direction in this cage is : " + self.direction)

    def open_close_door(self):
        print("we are arrived")
        print("opening in 3 seconds")
        self.position_door_x = 150
        self.timer = 0

        while self.timer < 3:
            #print("opening in 3" + str(self.timer) + " seconds")
            self.position_door_x -= 50
            print("position door: "+str(self.position_door_x))
            self.timer += 1
            time.sleep(1)

        print("The door is opened / position_door_x is at:" +
              str(self.position_door_x))

        self.position_door_x = 0
        self.other_order_after_opening()

        while self.timer > 0:
            self.position_door_x += 50
            print("position door: "+str(self.position_door_x))
            self.timer -= 1
            time.sleep(1)
        print("The door is closed / position_door_x is at:" +
              str(self.position_door_x))

    def other_order_after_opening(self):
        dest = 1
        while dest != 0 and dest > 0 and dest <= 10:
            destination = int(
                input('Wait for the next order [1 to 10], if not press 0 for continue: '))
            self.call_level(destination)
            dest = destination

    def counter_step(self, levelAsk, directionAsk):
        print("the compute start for the cage: "+str(self.id) +
              " levelAsk is : "+str(levelAsk)+" direction ask is: "+str(directionAsk))
        diff = self.levelActual - levelAsk
        sum = self.levelActual + levelAsk
        self.counterStepCage = 0
        if self.direction == "idle":
            self.counterStepCage = abs(self.levelActual - levelAsk)
            #print("levelAsk : " + str(levelAsk) + " direction : " + str(directionAsk) + " counterstep : " + str(counterStepCage))
            return self.counterStepCage
        elif directionAsk == "up" and self.direction == "up":
            if levelAsk >= self.levelActual:
                self.counterStepCage = diff * (-1)
                #print("levelAsk : " + levelAsk + " direction : " + directionAsk + " counterstep : " + counterStepCage)
                print("counterstep de la cage:"+str(self.id) +
                      " est " + str(self.counterStepCage))
                return self.counterStepCage
            else:
                self.counterStepCage = (20 - diff)
                #print("levelAsk : " + levelAsk + " direction : " + directionAsk + " counterstep : " + counterStepCage)
                print("counterstep de la cage:"+str(self.id) +
                      " est " + str(self.counterStepCage))
                return self.counterStepCage
        elif directionAsk == "up" and self.direction == "down":
            self.counterStepCage = sum
            #print("levelAsk : " + levelAsk + " direction : " + directionAsk + " counterstep : " + counterStepCage)
            print("counterstep de la cage:"+str(self.id) +
                  " est " + str(self.counterStepCage))
            return self.counterStepCage
        elif directionAsk == "down" and self.direction == "up":
            self.counterStepCage = (20 - sum)
            #print("levelAsk : " + levelAsk + " direction : " + directionAsk + " counterstep : " + counterStepCage)
            print("counterstep de la cage:"+str(self.id) +
                  " est " + str(self.counterStepCage))
            return self.counterStepCage
        elif directionAsk == "down" and self.direction == "down":
            if levelAsk <= self.levelActual:
                self.counterStepCage = diff
                #print("levelAsk : " + levelAsk + " direction : " + directionAsk + " counterstep : " + counterStepCage)
                print("counterstep de la cage:"+str(self.id) +
                      " est " + str(self.counterStepCage))
                return self.counterStepCage
            else:
                self.counterStepCage = sum
                #print("levelAsk : " + levelAsk + " direction : " + directionAsk + " counterstep : " + counterStepCage)
                print("counterstep de la cage:"+str(self.id) +
                      " est " + str(self.counterStepCage))
                return self.counterStepCage


# function choice the cage with the shortest way
def Choice_cage(levelAsk, directionAsk):
    counterCage1 = cage1.counter_step(levelAsk, directionAsk)
    counterCage2 = cage2.counter_step(levelAsk, directionAsk)
    if counterCage1 <= counterCage2:
        print("the Controller choice : cage1")
        if directionAsk == "up" and levelAsk <= 10 and levelAsk >= 0:
            cage1.listUp.append(levelAsk)
        elif directionAsk == "dwon" and levelAsk <= 10 and levelAsk >= 0:
            cage1.listDown.append(levelAsk)
    else:
        print("the Controller choice : cage2")
        if directionAsk == "up" and levelAsk <= 10 and levelAsk >= 0:
            cage2.listUp.append(levelAsk)
        elif directionAsk == "dwon" and levelAsk <= 10 and levelAsk >= 0:
            cage2.listDown.append(levelAsk)


# batterie
# Please give the instance to my object cage and cage2
# with this://def __init__(self,id, levelActual, direction):
# After that follow the instructions bellow
cage1 = Cage(1, 5, "up")
cage2 = Cage(2, 1, "up")

# display  status cage:id
cage1.display_status_cage()
cage2.display_status_cage()
# Manage the callAsk buttons

ask = 1
while ask != 0 and ask >= 0 and ask <= 10:
    ask = int(
        input('Enter level Ask number [0,10] ? if you finished enter 0 to execute: '))
    destination = input(
        'Enter your destination if you finished enter 11 to execute: ')
    Choice_cage(ask, destination)

    print("cage1 listUp: "+str(cage1.listUp))
    print("cage1 lisDown: "+str(cage1.listDown))
    print("cage2 listUp: "+str(cage2.listUp))
    print("cage2 listUp: "+str(cage2.listDown))
# variables
id = 1

# choice your cage
while id != 1 and id != 2:
    id = int(input('Enter cage number [1,2] ? : '))
    print("you are in the cage number: "+str(id))


# function call inside
# the first call is decided for the first direction cage
dest = 1
if id == 1:
    while dest > 0 and dest <= 10:
        dest = int(
            input('Enter your destination [1 to 10] ? if you finished enter 0: '))
        cage1.call_level(dest)
    cage1.move_up()
    cage1.move_down()


if id == 2:
    while dest != 0 and dest > 0 and dest <= 10:
        dest = int(
            input('Enter your destination [1 to 10] ? if you finished enter 0: '))
        cage2.call_level(dest)
    cage2.move_up()
    cage2.move_down()
