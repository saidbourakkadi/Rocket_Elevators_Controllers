
// variables
// var firstListUp = 0;
// var firstListDown = 0;



//class constructor cages
class Cage {
    constructor(id, column, battery, levelActual, direction) {
        this.battery = battery;
        this.column = column;
        this.id = id;

        this.statusDoor = "closed"; // open or closed
        this.position_door_x = 0; // x = 0 door closed, x=150 door open

        this.timing_door = 3000; // max 3s for open or close
        this.timing_floor = 20000; // max 20 s by floor
        this.listUp = [];        // list order up and atopped step by step
        this.listUpOrder = []; // order up this list
        this.listDown = [];   // list order down and atopped step by step
        this.listDownOrder = []; // order down this list
        this.direction = direction; // up, down, idle
        this.levelActual = levelActual; // level cage floor from 0 to 10
        this.levelCall = [];
        this.levelAsk = [];
        this.sensor_status_cage = "Clear"; // security sensor obstruct door cage 
        this.close_button = "inactif"; // button close door
        this.open_button = "inactif"; // button open door

        console.log("Cage : " + this.id + " levelActual : " + this.levelActual);
    }
    open_door() {

        for (var x = 0; x < 3; x++) {
            this.position_door_x = this.position_door_x + 50;
            this.statusDoor = "open";
            console.log("being closed 3 second : " + this.position_door_x + " status door : " + this.statusDoor);
            this.Timer(1000);
        } this.close_door();



    }

    close_door() {

        for (var x = 0; x < 3; x++) {
            this.position_door_x = this.position_door_x - 50;
            this.statusDoor = "closed";
            console.log("being open 3 second : " + this.position_door_x + " status door : " + this.statusDoor);
            this.Timer(1000);
        }


    }
    Timer(milliseconds) {
        const date = Date.now();
        let currentDate = null;
        do {
            currentDate = Date.now();
        } while (currentDate - date < milliseconds);
    }
    Call_level(call) {
        if (this.levelActual === call) {
            console.log("You are in this floor")
            this.open_door();
        }
        else if (this.levelActual < call) {
            this.listUp.push(call)
            this.order_area_ascendant(this.listUp);
            console.log("we up " + " listUp : " + this.listUp);
            this.move();
        } else {
            this.listDown.push(call)
            this.order_area_descendant(this.listDown);
            console.log("we down" + " listDown : " + this.listDown);
            this.move();
        }
    }
    order_area_ascendant(list) {
        list.sort(function (a, b) { return a - b });
        console.log("after tri listUP: " + this.listUp);
    }

    order_area_descendant(list) {
        list.sort(function (a, b) { return b - a });
        console.log("after tri listDown: " + this.listDown);
    }
    move() {
        var x;
        if (this.direction === "up" && this.listUp != null) {
            if (this.levelActual === this.listUp[0]) {
                this.open_door();

            } else {
                this.moveUp(this.listUp[0]);
            }

        } else if (this.direction === "down" && this.listDown != null) {
            if (this.levelActual === this.listUp[0]) {
                this.open_door();
            } else {
                this.moveDown(this.listDown[0]);
            }
        } else { this.direction = "idle"; }
    } moveUp(e) {
        if (this.listUp[0] != null) {
            this.Timer(3000);
            this.levelActual = this.levelActual + 1;
            console.log("levelActual is: " + this.levelActual);
            if (this.levelActual === e) {
                this.open_door();
                this.listUp.splice(0, 1);
                console.log("splice : " + this.listUp);
                this.move();

            } else {
                this.move();

            }
        } else {
            this.direction = "down";
            this.move();

        } console.log("level actual is :" + this.levelActual);
    }
    moveDown(e) {
        if (this.listDown[0] != null) {
            this.Timer(3000);
            this.levelActual = this.levelActual - 1;
            console.log("levelActual is: " + this.levelActual);
            if (this.levelActual === e) {
                this.open_door();
                this.listDown.splice(0, 1);
                console.log("splice : " + this.listDown);
            }
            else if (this.listDown != null || this.listDown != null) {
                this.move();
            }


        } else if (this.listUp[0] != null) {
            this.direction = "up";
            this.move();

        }
        else {
            this.direction = "idle";
            console.log("level actual is :" + this.levelActual);
        }

    }

    // this compute counter step supposed the elevators 
    //started from floor number 0 until floor number 10
    // on button floor pressed the fonction counterStep for cage1 and cage 2

    counterStep(levelAsk, directionAsk) {
        console.log("the compute start");
        var diff = this.levelActual - levelAsk;
        var sum = this.levelActual + levelAsk;

        var counterStepCage = 0;
        if (this.direction === "idle") {
            counterStepCage = Math.abs(this.levelActual - levelAsk);
            console.log("levelAsk : " + levelAsk + " direction : " + directionAsk + " counterstep : " + counterStepCage);


            return counterStepCage, levelAsk, directionAsk;
        }

        else if (directionAsk === "up" && this.direction === "up") {

            if (levelAsk >= this.levelActual) {
                counterStepCage = diff * (-1);
                console.log("levelAsk : " + levelAsk + " direction : " + directionAsk + " counterstep : " + counterStepCage);

                return counterStepCage, levelAsk, directionAsk;

            } else {
                counterStepCage = (20 - diff);
                console.log("levelAsk : " + levelAsk + " direction : " + directionAsk + " counterstep : " + counterStepCage);


                return counterStepCage, levelAsk, directionAsk;
            }


        } else if (directionAsk == "up" && this.direction == "down") {
            counterStepCage = sum;
            console.log("levelAsk : " + levelAsk + " direction : " + directionAsk + " counterstep : " + counterStepCage);

            return counterStepCage, levelAsk, directionAsk;
        } else if (directionAsk == "down" && this.direction == "up") {
            counterStepCage = (20 - sum);

            console.log("levelAsk : " + levelAsk + " direction : " + directionAsk + " counterstep : " + counterStepCage);

            return counterStepCage, levelAsk, directionAsk;
        } else if (directionAsk == "down" && this.direction == "down") {
            if (levelAsk <= this.levelActual) {
                counterStepCage = diff;
                console.log("levelAsk : " + levelAsk + " direction : " + directionAsk + " counterstep : " + counterStepCage);

                return counterStepCage, levelAsk, directionAsk;

            } else {
                counterStepCage = sum;
                console.log("levelAsk : " + levelAsk + " direction : " + directionAsk + " counterstep : " + counterStepCage);

                return counterStepCage, levelAsk, directionAsk;
            }
        }


    }

}
// choice the bet cage at asking

battery_compute_BestTimingCage = function BestTimingCage(cage_1, cage_2) {
    if (cage_1 <= cage_2) {
        console.log("the Controller choice : cage1");
        return "cage1";

    } else {
        console.log("the Controller choice : cage2");
        return "cage2";
    }
};




//////////////////////////////////////////////////////////////////
// compute function calculeted shorter steps for each elevator, //
//and send order(levelAsk, direction) to the selected elevator //
////////////////////////////////////////////////////////////////

function compute(levelAsk, directionAsk) {
    console.log("Status Cage1 :");
    var cage_1 = cage1.counterStep(levelAsk, directionAsk);
    console.log("Status Cage2 :");
    var cage_2 = cage2.counterStep(levelAsk, directionAsk);
    battery_compute_BestTimingCage(cage_1, cage_2)
    if (cage1 && directionAsk === "up") {
        cage1.listUp.push(levelAsk);

        console.log("listUp cage1 : " + cage1.listUp);
        this.cage1.move();
    } else if (cage1 && directionAsk === "down") {
        cage1.listDown.push(levelAsk);

        console.log("listDown cage1 : " + cage1.listDown);
        this.cage1.move();
    } else if (cage2 && directionAsk === "up") {
        cage2.listUp.push(levelAsk);

        console.log("listUp cage2 : " + cage2.listUp);
        this.cage2.move();
    } else if (cage2 && directionAsk === "down") {
        cage2.listDown.push(levelAsk);

        console.log("listDown cage2 : " + cage2.listDown);
        this.cage2.move();
    }


};
/////////////////////////////
/////some Scenario //////////
////////////////////////////////////////////////////////
// constructor(id,column,battery,levelActual,direction) 
// initialised cages //
/////////////////////////////////////////    
cage1 = new Cage("cage1", 1, 1, 5, "up");
cage2 = new Cage("cage2", 1, 1, 2, "up");
///////////////////////////////////////////

//call inside the elevator
//1er scenario : 
// cage1.Call_level(5);

// cage1.Call_level(3);
// cage1.Call_level(7);
// cage1.Call_level(4);
// compute(3,'up'); 

//call outside the elevator
//compute(levelAsk, direction)
compute(7, "up");




















