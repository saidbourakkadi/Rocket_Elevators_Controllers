
//////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////Author : Said Bourakkadi/////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////// Commercial Controller//////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////// 1 battery <- 4 colomns <- 5 cages ////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)


   
       
    //////////////////////////////////////////////////////////////////////////
    // Buttons : are used to generate pickup Orders from either the 
    //           panel or from a given floor.                                 
    //////////////////////////////////////////////////////////////////////////
	// A CallButton is generated per Floor object
type CallButton struct {
	id     int
	status string
}

// NewCallButton is a CallButton factory function (constructor)
func NewCallButton(id int, status string) CallButton {
	b := CallButton{}
	b.id = id
	b.status = status
	return b
}
    type CallButton struct {
    
        id int;
        status string;
	}
    func CallButton(id int id, status string) CallButton{
			b := CallButton{}
            this.id = id;
			this.status = status;
			return b
        }

    func (b *CallButton)CallButtonPressed(){
        
        this.status = "Active";// Active, Inactif
    }
    

    type FloorButton struct {
         id int;
         status string;
	}
    func FloorButton(int id, string status) FloorButton  {
			b := FloorButton{}
            b.id = id;
			b.status = status;
			return b
    }
    

    ///////////////////////////////////////////////////////////////////////////////////////////////
    // Cages  Each cage contain all the necessary methods for                         
    //        Moving up and down  or stopping and                                   
    //        Opening and closing the door                                          
    //        Assure the security (with sensors )                                     
    //         Receive the destination order from battery and manage it.                  
    /////////////////////////////////////////////////////////////////////////////////////////////////

    type Cage struct {
        id int;
        status string;// In-service, Loading, idle.
        doors string; // Open or Cloded
        levelActual int;
        direction string;//up, down
        timer int;
        picReq []Order;
		destReq []Order;
		
		
	}

    func Cage(id int, status string, doors string) Cage{
		c := Cage{}
        c.id = id;
        c.status = status;
		c.doors = doors;
		c.curFloor = 1
		c.direction = "Up"
		c.timer = 0
		return c
	}
	
    func (c *Cage) CleanUpOrders() {
        for y := (len(c.picReq) - 1) ; y >= 0; y-- {
            if c.levelActual == c.picReq[y].pickup {
                    c.picReq[y].status = "Destination";
            }
            if c.picReq[y].status == "Destination"{
				c.destReq = append(c.destReq, c.picReq[y])//c.destReq?
                c.picReq = append(c.picReq[:y], c.picReq[y+1:]...)
                fmt.Println("Destination now is " + strconv.Itoa(c.destReq[y].destination));
            }
            }
        for y := len(c.destReq - 1); y >= 0; y--{
            if (c.levelActual == c.destReq[y].destination){
                    c.destReq[y].status = "Finished";
            }
            if (c.destReq[y].status == "Finished"){
					c.destReq = append(c.destReq[:y], c.destReq[i+1:]...)
            }
        }
	}
	
        //  Methods Door//
        func (c *Cage)open_door() {
            if (c.status != "In-Service"){
                c.doors = "Open";
                fmt.Println("Cage " + strconv.Itoa(c.id) + " doors are open for 5 seconds")
                c.timer = 5;
                for c.timer > 0 {
                    fmt.Println("Closing in " + strconv.Itoa(c.timer) + " seconds.")
                    Thread.Sleep(1* tie.Second);
                    c.timer--
                   
                }
                c.close_door();
            }
		}
		
        func (c *Cage)close_door(){
            if c.timer < 5{
                c.doors = "Closed"
                fmt.Println("Cage doors are closed")
                if len(c.picReq) == 0 && len(c.destReq) == 0 {
                    c.status = "Idle";
                }else{
                    c.status = "Loading";
                }
                 
			}    
            
        }
        func (c *Cage)OpenButtonPressed() {
            if c.status != "In-Service" {
                c.open_door();
            }
        }
        func (c *Cage)CloseButtonPressed() {
            if (c.timer < 5){
                c.close_door();
            }
        }
        

        

        // Movement //
        func (c *Cage) move_down(){
            for c.doors != "Closed"{
                c.close_door();
            }
            c.status = "In-Service"
            c.direction = "down"
            fmt.Println("Cage " + strconv.Itoa(close.id)  + "\tDirection is down\tlevelActual " + strconv.Itoa(this.levelActual))
            if c.levelActual - 1 == 0 {
                c.levelActual -= 2;
            }else{
                c.levelActual -= 1;
            }
            fmt.Println("Cage " + strconv.Itoa(c.id) + "\t\t\t\tlevelActual now is "  + strconv.Itoa(c.levelActual))
            c.status = "Loading"
        }

        func (c *Cage) move_up(){
            for c.doors != "Closed"{
                c.close_door()
            }
            c.status = "In-Service"
            c.direction = "up"
            fmt.Println("Cage " + strconv.Itoa(this.id) + "\tDirection is up\t\tlevelActual " + strconv.Itoa(this.levelActual))
            if c.levelActual + 1 == 0{
                c.levelActual += 2
            } else{
                this.levelActual += 1
            }
            fmt.Println("Cage " + strconv.Itoa(this.id) + "\t\t\t\tlevelActual now is " + strconv.Itoa(this.levelActual))
            c.status = "Loading";
        }
    
    ///////////////////////////////////////////////////////////////////////////////////////////////
    // Clomns(4):  Each Column object has a list of Cage objects (5)                         
    //                                                 
    //    Receive the destination order from battery and manage it.                    
    /////////////////////////////////////////////////////////////////////////////////////////////////

    type Column struct
    {

        id int // identifier column
        status string
        cages []Cage //list cages (5)
        floors_column []int // floors who column deserved
	}
    func Column(id int, cages []Cage, floors_column []int){
		c := Column{}
        c.id = id
        c.status = "Actif"// supposed Actif all times
        c.cages = cages
		c.floors_column = floors_column
		return c
    }
        
    

    //////////////////////////////////////////////////////////////////////////
    // Panel : simulates a panel in the reception of the building.    
    //          This panel directs the user to the appropriate column for their 
    //          Ordered floor and sends the appropriate pickup Order.  
    //////////////////////////////////////////////////////////////////////////

    type Panel struct
    {
        floorButtons []FloorButton
	}
    func Panel() Panel{
		p :=Panel{}
        for x := 0 - Configuration.totalBasements; x < 0; x++{
				
			p.floorButtons = Append(p.floorButtons,FloorButton(x,"Inactive"))
        }
        for x := 1; x <= Configuration.totalFloors; x++{
				
			p.floorButtons = Append(p.floorButtons,FloorButton(x,"Inactive"))
        }
	}
	
	

        // Methods //
    func (p * Panel) OrderElevator(floorNumber int, cageManager CageManager ) {
        for _,button := range p.floorButtons{ 
            if button.id == floorNumber{
                button.status = "Active";
            }
        }

        myColumn := cageManager.takeColumn(1, floorNumber); //var mycolumn ?
        fmt.Println("Floor Ordered. \tProceed to column " + strconvert.Itoa(myColumn.id))

    }

        // Reports //
        func (p Panel)GetFloorButtonsStatus(){
            for _, x := 0; x < (ln(p.floorButtons) - 1); x++ {
                fmt.Println("Floor " + strconv.Itoa(c.floorButtons[x].id) + " button is " + c.floorButtons[x].status)
            }
        }
    }


    /////////////////////////////////////////////////////////////////////////////////
    // Floors : The floor object is generated by the Configuration object as a    //
    // list of floors each with a call button equal to the number of             //
    // total floors set by the user.                                            //
    /////////////////////////////////////////////////////////////////////////////

    type Floor struct
    {
        id int
        button CallButton
	}
    func  Floor(id int, button CallButton) Floor {
		f := Floor{}
        f.id = id
		f.button = button
		return f
    }
    


    ////////////////////////////////////////////////////////////////////////
    //                           Order                                    //
    ////////////////////////////////////////////////////////////////////////
    ////////////////////////////////////////////////////////////////////////
    // A Order object is generated each time a Floor or Call button is  //
    // pressed. The Order is queued by the main program before being    //
    // assigned to a cage for treatment.                                  //
    ////////////////////////////////////////////////////////////////////////

    type Order struct
    {
        status string// pickup
        assignment string //Assigned or Unassigned
        pickup int // floor pickup 
        destination int // floor destination
        direction string //
	}
    func Order( status string,  pickup int,  destination int,  direction string) Order {
		o := Order{}
		o.status = status
		o.assignment = "Unassigned" 
        o.pickup = pickup
        o.destination = destination
		o.direction = direction
		return o
    }
        
       

    ////////////////////////////////////////////////////////////////////////
    //                  MANAGER COLUMNS AND CAGES
    //This object contains all the column and cage objects in the system //
    // Only one CageManager should instantiated and only after Config has //
    // been called during the initial setup.                              //
    ////////////////////////////////////////////////////////////////////////

    type  CageManager struct
    {
        colList []Column
	}
    func CageManager()CageManager{
		c := CageManager{}
        var floorRange int
		var extraFloors int
		var floorCounter int
        if myConfiguration.totalBasements > 0 {
			if (myConfiguration.totalFloors-1)%(myConfiguration.totalColumns-1) != 0 {
				floorRange = (myConfiguration.totalFloors - 1) / (myConfiguration.totalColumns - 1)
				extraFloors = (myConfiguration.totalFloors - 1) % (myConfiguration.totalColumns - 1)
			} else {
				floorRange = (myConfiguration.totalFloors - 1) / (myConfiguration.totalColumns - 1)
				extraFloors = 0
			}
			floorCounter = 2
			var bColumnFloors []int
			for i := 0; i < myConfiguration.totalBasements; i++ {
				if myConfiguration.floorList[i].id < 0 {
					bColumnFloors = append(bColumnFloors, myConfiguration.floorList[i].id)
				}
			}
			bColumnFloors = append(bColumnFloors, 1)
			c.colList = append(c.colList, NewColumn(1, "Active", c.GenerateCages(myConfiguration.cagesPerColumn), bColumnFloors))
			for i := 2; i <= myConfiguration.totalColumns; i++ {
				var floorsServed []int
				floorsServed = append(floorsServed, 1)
				if myConfiguration.totalColumns-i < extraFloors {
					for j := 0; j < floorRange+1; j++ {
						floorsServed = append(floorsServed, floorCounter)
						floorCounter++
					}
					c.colList = append(c.colList, NewColumn(i, "Active", c.GenerateCages(myConfiguration.cagesPerColumn), floorsServed))
				} else {
					for j := 0; j < floorRange; j++ {
						floorsServed = append(floorsServed, floorCounter)
						floorCounter++
					}
					c.colList = append(c.colList, NewColumn(i, "Active", c.GenerateCages(myConfiguration.cagesPerColumn), floorsServed))
				}
			}
		} else {
			if myConfiguration.totalFloors%myConfiguration.totalColumns != 0 {
				floorRange = myConfiguration.totalFloors / myConfiguration.totalColumns
				extraFloors = myConfiguration.totalFloors % myConfiguration.totalColumns
			} else {
				floorRange = myConfiguration.totalFloors / myConfiguration.totalColumns
				extraFloors = 0
			}
			floorCounter = 2
			for i := 1; i <= myConfiguration.totalColumns; i++ {
				var floorsServed []int
				floorsServed = append(floorsServed, 1)
				for j := 0; j < floorRange; j++ {
					floorsServed = append(floorsServed, floorCounter)
					floorCounter++
				}
				c.colList = append(c.colList, NewColumn(i, "Active", c.GenerateCages(myConfiguration.cagesPerColumn), floorsServed))
			}
		}
		return c
	}

        // Methods //

        // This method loops through all cages in a given column and returns //
        // the cage which can most efficiently fulfill the given Order.    // 
        func (c *CageManager) takeCage(direction string,  column int,  reqFloor int)* Cage{
            var currentCage = c.colList[column].cages[0]
            var bestCage = c.colList[column].cages[0]
             x := 0
            for x < c.colList[column].len(cages.Count){
                currentCage = c.colList[column].cages[x]
                if currentCage.direction == direction && direction == "up" && currentCage.levelActual < reqFloor && (currentCage.status == "In-Service" || currentCage.status == "Loading"){
                    fmt.Println("Same direction UP was selected")
                    return currentCage // Returns the cage with the same direction (UP) that has not yet passed the Ordered floor
                }else if currentCage.direction == direction && direction == "down" && currentCage.levelActual > reqFloor && (currentCage.status == "In-Service" || currentCage.status == "Loading"){
                    fmt.Println("Same direction DOWN was selected")
                    return currentCage // Returns the cage already going the same direction (DOWN) that has not yet passed the Ordered floor
                }else if currentCage.status == "Idle"{
                     allCagesAreIdle := true
                    for r := 0; r < len(c.colList[column].cages); r++{
                        if c.colList[column].cages[r].status != "Idle"{
                            allCagesAreIdle = false;
                        }
                    }
                    if allCagesAreIdle{
                        for i := x + 1; i < len(c.colList[column].cages); i++{
                            var compareCage = c.colList[column].cages[i]
                            if compareCage.status == "Idle"{
                                fmt.Println("Cage " + convert.Itoa(bestCage.id) + "\tto be compared to " + convert.Itoa(compareCage.id)); 
                                before := Abs(bestCage.levelActual - reqFloor);
                                after := Abs(compareCage.levelActual - reqFloor);
                                if (after < before){
                                    bestCage = compareCage; // Closest idle cage
                                }
                            }
                            fmt.Println("Cage " + convert.Itoa(currentCage.id) + " is selected."); 
                        }
                        return bestCage;
                    }
                }else{
                    for i := 0; i < len(c.colList[column].cages); i++{
                        if direction == "up" && len(c.colList[column].cages[i].destReq) < len(currentCage.destReq){
                            currentCage = c.colList[column].cages[i]
                        }else if direction == "down" && len(c.colList[column].cages[i].picReq) < len(currentCage.picReq){
                            currentCage = c.colList[column].cages[i]
                        }
                    }
                     
                }
                x++;
			}
			fmt.Println("The most suitable cage is selected")
            return currentCage; 
        }

        // Returns a column where the Ordered floor is served //
        func (c *CageManager)takeColumn(pickup int,  destination int){
            pickupServed := false
            destServed := false
            for _, column := range c.colList {
				for _, id := range column.floors_column {
                
                    if id == pickup{
                    
                        pickupServed = true
                    }
                    if id == destination{
                        destServed = true
                    }
                    if pickupServed && destServed{
                    
                        return &column
                    }
                }
            }
            return nil
        }

        // Instantiates cages based off a given number //
        func (c CageManager) GenerateCages(numCages int){
            var cageList []Cage
            for x := 1; x <= numCages; x++{
				cageList = append(cageList, Cage(i, "Idle", "Closed"))
                
            }
            return cageList;
        }


        // Watch all columns and their cages as well as their current floor and status //
        func (c CageManager)takeCageStatus(){
            
            fmt.Println("\n| Column |\t Cage |\t Status |\t level Actual |\t Door status |")
                     
            for x := 0; x < len(c.colList); x++{
                for i := 0; i < len(c.colList[x].cages); i++{
                    var currentCage Cage= c.colList[x].cages[i]
                    
                    fmt.Println(strconv.Itoa(c.colList[x].id)+"\t"+ strconv.Itoa(currentCage.id)+"\t"+currentCage.status+currentCage.levelActual+"\t"+strconv.Itoa(currentCage.doors)
                }
            }
		}
		// Returns a string of the floors served by a given column //
		func (c CageManager) GetFloorsServed(column Column) string {
			var myFloors []string
			for _, floor := range column.floorsServed {
				myFloors = append(myFloors, strconv.Itoa(floor))
			}
			floorString := strings.Join(myFloors, ",")
			colString := strconv.Itoa(column.id)
			myString := "Column " + colString + ": " + floorString
			return myString
		}
        
        
    


    

    ////////////////////////////////////////////////////////////////////////
    // Configuration : This static object generates a configuration
    //               from user input and the corresponding floor list.                         
    ////////////////////////////////////////////////////////////////////////

    type Configuration struct {
		batteryOn      bool
		totalColumns   int
		cagesPerColumn int
		totalFloors    int
		totalBasements int
		floorList      []Floor
	}

    var myConfiguration = Configuration{}

// GenerateFloors is to be called after Config: Generates Floor structs and adds them to the floorList
func (c *Configuration) GenerateFloors() {
	// Checks if building has basements to add to the floor list
	if c.totalBasements > 0 {
		for i := 0 - c.totalBasements; i < 0; i++ {
			c.floorList = append(c.floorList, Floor(i, CallButton(i, "Inactive")))
		}
	}
	// Adds remaining floors
	for i := 1; i < 1+c.totalFloors; i++ {
		c.floorList = append(c.floorList, Floor(i, CallButton(i, "Inactive")))
	}
}    

        
// a revoir
    //    func takeIntInput( prompt string,  minValue uint){
    //         fmt.Println(prompt);
    //         myInput := -1
    //         userInput := Console.ReadLine();
    //         for myInput == -1{
    //             try{
    //                 myInput = Convert.ToInt32(userInput);
    //                 if myInput < minValue
    //                 {
    //                     fmt.Println("Value cannot be less than " + strconv.Itoa(minValue) + ".")
    //                     myInput = -1
    //                     userInput = ""
    //                 } 
    //             }
    //             catch System.FormatException
    //             {
    //                 if userInput == ""
    //                 {
    //                     fmt.Println("Enter a valid number.");
    //                     userInput = Console.ReadLine();
    //                 }
    //                 else
    //                 {
    //                     fmt.Println(userInput + " is not a valid number.\n Enter a valid number.");
    //                     userInput = Console.ReadLine();
    //                 }
    //             }
    //         }
    //         return myInput;
    //     }
	func takeIntInput(s string) int {
		reader := bufio.NewReader(os.Stdin)
	
		for {
			fmt.Printf("%s: ", s)
	
			input, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			cleanedInput := strings.Replace(input, "\r\n", "", -1)
			myInt, err1 := strconv.Atoi(cleanedInput)
			if err1 != nil {
				fmt.Printf(cleanedInput + " is not a valid number. Please enter a valid number.\n")
			} else if myInt < 0 {
				fmt.Printf("Value cannot be less than zero. Please enter a valid number.\n")
			} else {
				return myInt
			}
		}
	}
// Gets a y or n response from the user
func askForConfirmation(s string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s [y/n]: ", s)

		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		response = strings.ToLower(strings.TrimSpace(response))

		if response == "y" || response == "yes" {
			return true
		} else if response == "n" || response == "no" {
			return false
		} else {
			fmt.Printf(response + "  make a valid selection\n")
		}
	}
}
		// To be called once upon startup: Generates a configuration based on user input //
		
        // func Config(){
        //     ConsoleKeyInfo, letters string
        //     do
        //     {
        //         fmt.Println("Activate battery? [y/n]");
        //         while (Console.KeyAvailable == false)
        //         {
        //             Thread.Sleep(100); // Loop until valid input is entered.
        //         }

        //         letters = Console.ReadKey(true)
        //         if letters.Key != ConsoleKey.Y && letters.Key != ConsoleKey.N
        //         {
        //             fmt.Println("You pressed the '{0}' key.  make a valid selection.", letters.Key)
        //         }
        //         else if letters.Key == ConsoleKey.N{
        //             fmt.Println("Startup Aborted!")
        //             return
        //         }
        //     } for letters.Key != ConsoleKey.Y{
		// 		fmt.Println("Initializing...")
		// 	}
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
			
			
            // Set total number of columns //
             totalColumns int= takeIntInput("Enter the total number of columns", 1)

            // Set cages per column //
             cagesPerColumn int= takeIntInput("How many cages are installed per column?", 1)

            // Set number of floors //
             totalFloors int= takeIntInput("How many floors (excluding basements) are there in the building?", 2)

            // Set number of basements //
             totalBasements int= takeIntInput("How many basements are there?", 0)

            // Set Configuration Values //
            Configuration.batteryOn = true
            Configuration.totalColumns = totalColumns
            Configuration.cagesPerColumn = cagesPerColumn
            Configuration.totalFloors = totalFloors
            Configuration.totalBasements = totalBasements

            // Confirm Setup Conditions //
            
			fmt.Println("\n-------SIMULATION-------")
            fmt.Println("\nCategories\t Value")
            fmt.Println("Battery\t On")
            fmt.Println("Total Columns\t"+ strconv.Itoa(Configuration.totalColumns))
            fmt.Println("Cages Per Column\t"+strconv.Itoa(Configuration.cagesPerColumn))
            fmt.Println("Total Floors\t"+strconv.Itoa(Configuration.totalFloors))
            fmt.Println("Total Basements\t"+strconv.Itoa(Configuration.totalBasements))
        }

        // To be called after Config: Generates floor objects and adds them to floor list //
       func  (c *Configuration)GenerateFloors(){
            // Checks if building has basements and adds them to the floor list //
            if totalBasements > 0 {
                for x := 0 - totalBasements; x < 0; x++{
					
					c.floorList = append(c.floorList, Floor(i, CallButton(x, "Inactive")))
                }
            }

            // Adds remaining floors to the floor list //
            for x := 1; x < 1 + totalFloors; x++{
                c.floorList= append(c.floorList, Floor(x, CallButton(x, "Inactive")))
            }
        }
		

        // Reports //
        func (c *Configuration) GetFloorStatus(){
            fmt.Println("\n-----------------FLOOR STATUS------------------\n");
            for x := 0; x < len(floorList); x++{
				ids := strconv.Itoa(c.floorList[i].id)

                fmt.Println("Floor " + ids + ":  Active  /\t  Call Status: " + c.floorList[i].button.status)
		   
			}
        }
    


    //////////////////////////////////////////////////////////////////////////////////
    ///           Principal function :     Main                                   ///
    ////////////////////////////////////////////////////////////////////////////////

    class Program
    {
        
        // Check all buttons and add Orders to the queue //
        static List<Order> OrderQueue = new List<Order>();
        static void OrderGenerator(Panel myPanel)
        {
            // Checks call buttons //
            foreach (Floor floor in Configuration.floorList)
            {
                if (floor.button.status == "Active")
                {
                    fmt.Println("Floor button " + floor.button.id + " is active.");
                    if (floor.id > 0)
                    {
                        Order myOrder = new Order("Pickup", floor.button.id, 1, "down");
                        foreach (Order Order in OrderQueue)
                        {
                            if (floor.button.id == Order.pickup && Order.status == "Pickup")
                            {
                                fmt.Println("My Order for floor " + floor.button.id + " was not sent.");
                                return;
                            }
                        }
                        OrderQueue.Add(myOrder);
                        fmt.Println("My Order for floor " + myOrder.pickup + " was added to the Order list");
                    }
                    else
                    {
                        Order myOrder = new Order("Pickup", floor.id, 1, "up");
                        foreach (Order Order in OrderQueue)
                        {
                            if (floor.id == Order.pickup && Order.status == "Pickup")
                            {
                                fmt.Println("My Order for floor " + floor.button.id + " was not sent.");
                                return;
                            }
                        }
                        fmt.Println("My Order for floor " + floor.button.id + " was added to the Order list");
                        OrderQueue.Add(myOrder);
                    }
                    floor.button.status = "Inactive";
                    fmt.Println("Floor " + floor.button.id + " is " + floor.button.status);
                }
            }

            // Check floor buttons //
            foreach (FloorButton button in myPanel.floorButtons)
            {
                if (button.status == "Active")
                {
                    fmt.Println("Panel button " + button.id + " is " + button.status);
                    if (button.id > 0)
                    {
                        Order myOrder = new Order("Pickup", 1, button.id, "up");
                        foreach (Order Order in OrderQueue)
                        {
                            if (myOrder.destination == Order.destination && Order.status == "Pickup")
                            {
                                fmt.Println("My Order for floor " + button.id + " was not sent.");
                                return;
                            }
                        }
                        fmt.Println("My Order for floor " + button.id + " was added to the Order list");
                        OrderQueue.Add(myOrder);
                    }
                    else
                    {
                        Order myOrder = new Order("Pickup", 1, button.id, "down");
                        foreach (Order Order in OrderQueue)
                        {
                            if (myOrder.destination == Order.destination && Order.status == "Pickup")
                            {
                                fmt.Println("My Order for floor " + button.id + " was not sent.");
                                return;
                            }
                        }
                        fmt.Println("My Order for floor " + myOrder.pickup + " was added to the Order list");
                        OrderQueue.Add(myOrder);
                    }
                    button.status = "Inactive";
                    fmt.Println("Floor " + button.id + " is " + button.status);
                }
            }
        }

        // Assign each Order to any elevator for they move to destination//
        static void AssignElevator(CageManager myCageManager)
        {
            foreach (Order Order in OrderQueue)
            {
                if (Order.assignment == "Unassigned")
                {
                    Column myColumn = myCageManager.takeColumn(Order.pickup, Order.destination);
                    fmt.Println("Column " + myColumn.id + " is selected.");
                    Cage myCage = myCageManager.takeCage(Order.direction, myColumn.id - 1, Order.pickup);
                    Order.assignment = "Assigned";
                    myCage.picReq.Add(Order);
                    fmt.Println("Cage " + myCage.id + " receive Order for floor " + myCage.picReq[0].pickup);
                    myCage.picReq.OrderBy(o => o.pickup);
                }
            }
        }

        // Move all elevators towards next destination or pickup
        static void move_elevators(CageManager myCageManager)
        {
            if (Configuration.totalBasements > 0)
            {
                foreach (Cage cage in myCageManager.colList[0].cages)
                {
                    if (cage.picReq.Count != 0)
                    {
                        if (cage.levelActual != cage.picReq[0].pickup && cage.levelActual > cage.picReq[0].pickup)
                        {
                            cage.move_down();
                        }
                        else if (cage.levelActual != cage.picReq[0].pickup && cage.levelActual < cage.picReq[0].pickup)
                        {
                            cage.move_up();
                        }
                        else if (cage.levelActual == cage.picReq[0].pickup)
                        {
                            cage.open_door();
                            cage.picReq[0].status = "Destination";
                            cage.CleanUpOrders();
                        }
                    }
                    if (cage.picReq.Count == 0 && cage.destReq.Count != 0)
                    {
                        if (cage.levelActual != cage.destReq[0].destination && cage.levelActual > cage.destReq[0].destination)
                        {
                            cage.move_down();
                        }
                        if (cage.levelActual != cage.destReq[0].destination && cage.levelActual < cage.destReq[0].destination)
                        {
                            cage.move_up();
                        }
                        else if (cage.levelActual == cage.destReq[0].destination)
                        {
                            cage.open_door();
                            cage.destReq[0].status = "Finished";
                            cage.CleanUpOrders();
                        }
                    }
                }
                for (int x = 1; x < myCageManager.colList.Count; x++)
                {
                    foreach (Cage cage in myCageManager.colList[x].cages)
                    {
                        if (cage.picReq.Count != 0)
                        {
                            if (cage.levelActual != cage.picReq[0].pickup && cage.levelActual > cage.picReq[0].pickup)
                            {
                                cage.move_down();
                            }
                            else if (cage.levelActual != cage.picReq[0].pickup && cage.levelActual < cage.picReq[0].pickup)
                            {
                                cage.move_up();
                            }
                            else if (cage.levelActual == cage.picReq[0].pickup)
                            {
                                cage.open_door();
                                cage.picReq[0].status = "Destination";
                                cage.CleanUpOrders();
                            }
                        }
                        if (cage.picReq.Count == 0 && cage.destReq.Count != 0)
                        {
                            if (cage.levelActual != cage.destReq[0].destination && cage.levelActual > cage.destReq[0].destination)
                            {
                                cage.move_down();
                            }
                            if (cage.levelActual != cage.destReq[0].destination && cage.levelActual < cage.destReq[0].destination)
                            {
                                cage.move_up();
                            }
                            else if (cage.levelActual == cage.destReq[0].destination)
                            {
                                cage.open_door();
                                cage.destReq[0].status = "Finished";
                                cage.CleanUpOrders();
                            }
                        }
                    }
                }
            }
            else
            {
                foreach (Column column in myCageManager.colList)
                {
                    foreach (Cage cage in column.cages)
                    {
                        if (cage.picReq.Count != 0)
                        {
                            if (cage.levelActual != cage.picReq[0].pickup && cage.levelActual > cage.picReq[0].pickup)
                            {
                                cage.move_down();
                            }
                            else if (cage.levelActual != cage.picReq[0].pickup && cage.levelActual < cage.picReq[0].pickup)
                            {
                                cage.move_up();
                            }
                            else if (cage.levelActual == cage.picReq[0].pickup)
                            {
                                cage.open_door();
                                cage.picReq[0].status = "Destination";
                                cage.CleanUpOrders();
                            }
                        }
                        if (cage.picReq.Count == 0 && cage.destReq.Count != 0)
                        {
                            if (cage.levelActual != cage.destReq[0].destination && cage.levelActual > cage.destReq[0].destination)
                            {
                                cage.move_down();
                            }
                            if (cage.levelActual != cage.destReq[0].destination && cage.levelActual < cage.destReq[0].destination)
                            {
                                cage.move_up();
                            }
                            else if (cage.levelActual == cage.destReq[0].destination)
                            {
                                cage.open_door();
                                cage.destReq[0].status = "Finished";
                                cage.CleanUpOrders();
                            }
                        }
                    }
                }
            }
        }

        // Checks the OrderQueue for Finished Orders that need removed
        static void CleanUpQueue()
        {
            for (int x = OrderQueue.Count - 1; x >= 0; x--)
            {
                if (OrderQueue[x].status == "Finished")
                {
                    OrderQueue.Remove(OrderQueue[x]);
                }
            }
        }

        static void LoopTest(Panel testPanel, CageManager testManager)
        {
            OrderGenerator(testPanel);
            AssignElevator(testManager);
            move_elevators(testManager);
            CleanUpQueue();
        }

        static void Scenario1(Panel myPanel, CageManager myCageManager)
        {
            
            
            myCageManager.colList[1].cages[0].levelActual = 10;
            myCageManager.colList[1].cages[1].levelActual = 5;
            myCageManager.colList[1].cages[2].levelActual = 19;
            myCageManager.colList[1].cages[3].levelActual = 15;
            myCageManager.colList[1].cages[4].levelActual = 8;
             

            fmt.Println("---------Scenario 1------------");
            fmt.Println(" Pickup: 7, direction: down");
            fmt.Println(" Pickup: 13, direction: up");
            fmt.Println(" Pickup: 5, direction: down");
            fmt.Println(" Pickup: 2, direction: down");
            fmt.Println(" Pickup: 5, direction: down");
            fmt.Println(" some one is floor 1 and request destination :17 ");
            fmt.Println("---------Scenario 1------------");
            //Order(string status, int pickup, int destination, string direction)
            OrderQueue.Add(new Order("Destination", 0, 7, "down"));
            OrderQueue[0].assignment = "Assigned";
            myCageManager.colList[1].cages[0].destReq.Add(OrderQueue[0]);
            OrderQueue.Add(new Order("Destination", 0, 13, "up"));
            OrderQueue[1].assignment = "Assigned";
            myCageManager.colList[1].cages[1].destReq.Add(OrderQueue[1]);
            OrderQueue.Add(new Order("Destination", 0, 5, "down"));
            OrderQueue[2].assignment = "Assigned";
            myCageManager.colList[1].cages[2].destReq.Add(OrderQueue[2]);
            OrderQueue.Add(new Order("Destination", 0, 2, "down"));
            OrderQueue[3].assignment = "Assigned";
            myCageManager.colList[1].cages[3].destReq.Add(OrderQueue[3]);
            fmt.Println(" destination: 1, direction: down");
            OrderQueue.Add(new Order("Destination", 0, 5, "down"));
            OrderQueue[4].assignment = "Assigned";
            myCageManager.colList[1].cages[4].destReq.Add(OrderQueue[4]);
            LoopTest(myPanel, myCageManager);
            OrderQueue.Add(new Order("Pickup", 1, 17, "up"));
            while (OrderQueue.Count > 0)
            {
                LoopTest(myPanel, myCageManager);
            }
            myCageManager.takeCageStatus();
        }

        static void Scenario2(Panel myPanel, CageManager myCageManager)
        {
            myCageManager.colList[2].cages[0].levelActual =5;
            myCageManager.colList[2].cages[1].levelActual = 1;
            myCageManager.colList[2].cages[2].levelActual = 17;
            myCageManager.colList[2].cages[3].levelActual = 30;
            myCageManager.colList[2].cages[4].levelActual = 40;
            fmt.Println("---------Scenario 2------------");
            fmt.Println(" Pickup: 21, direction: up");
            fmt.Println(" Pickup: 28, direction: up");
            fmt.Println(" Pickup: 1, direction: down");
            fmt.Println(" Pickup: 24, direction: down");
            fmt.Println(" Pickup: 3, direction: down");
            fmt.Println(" some one is floor 1 and request destination :33 ");
            fmt.Println("---------Scenario 2-----------");
            //public Order(string status, int pickup, int destination, string direction)
            OrderQueue.Add(new Order("Destination", 0, 21, "up"));
            OrderQueue[0].assignment = "Assigned";
            myCageManager.colList[2].cages[0].destReq.Add(OrderQueue[0]);
            OrderQueue.Add(new Order("Destination", 0, 28, "up"));
            OrderQueue[1].assignment = "Assigned";
            myCageManager.colList[2].cages[1].destReq.Add(OrderQueue[1]);
            OrderQueue.Add(new Order("Destination", 0, 1, "down"));
            OrderQueue[2].assignment = "Assigned";
            myCageManager.colList[2].cages[2].destReq.Add(OrderQueue[2]);
            OrderQueue.Add(new Order("Destination", 0, 24, "down"));
            OrderQueue[3].assignment = "Assigned";
            myCageManager.colList[2].cages[3].destReq.Add(OrderQueue[3]);
            OrderQueue.Add(new Order("Destination", 0, 3, "down"));
            OrderQueue[4].assignment = "Assigned";
            myCageManager.colList[2].cages[4].destReq.Add(OrderQueue[4]);
            OrderQueue.Add(new Order("Pickup", 1, 33, "up"));
            while (OrderQueue.Count > 0)
            {
                LoopTest(myPanel, myCageManager);
            }
            myCageManager.takeCageStatus();
        }

        static void Scenario3(Panel myPanel, CageManager myCageManager)
        {
            myCageManager.colList[3].cages[0].levelActual = 58;
            myCageManager.colList[3].cages[1].levelActual = 50;
            myCageManager.colList[3].cages[2].levelActual = 46;
            myCageManager.colList[3].cages[3].levelActual = 1;
            myCageManager.colList[3].cages[4].levelActual = 60;
            fmt.Println("---------Scenario 3------------");
            fmt.Println(" Pickup: 1, direction: down");
            fmt.Println(" Pickup: 63, direction: up");
            fmt.Println(" Pickup: 54, direction: up");
            fmt.Println(" Pickup: 50, direction: down");
            fmt.Println(" Pickup: 1, direction: up");
            fmt.Println(" some one is floor 57 and request destination :1 ");
            fmt.Println("---------Scenario 3-----------");
            //public Order(string status, int pickup, int destination, string direction)
            OrderQueue.Add(new Order("Destination", 0, 1, "down"));
            OrderQueue[0].assignment = "Assigned";
            myCageManager.colList[3].cages[0].destReq.Add(OrderQueue[0]);
            OrderQueue.Add(new Order("Destination", 0, 63, "up"));
            OrderQueue[1].assignment = "Assigned";
            myCageManager.colList[3].cages[1].destReq.Add(OrderQueue[1]);
            OrderQueue.Add(new Order("Destination", 0, 54, "up"));
            OrderQueue[2].assignment = "Assigned";
            myCageManager.colList[3].cages[2].destReq.Add(OrderQueue[2]);
            OrderQueue.Add(new Order("Destination", 0, 50, "up"));
            OrderQueue[3].assignment = "Assigned";
            myCageManager.colList[3].cages[3].destReq.Add(OrderQueue[3]);
            OrderQueue.Add(new Order("Destination", 0, 1, "up"));
            OrderQueue[4].assignment = "Assigned";
            myCageManager.colList[3].cages[4].destReq.Add(OrderQueue[4]);
            LoopTest(myPanel, myCageManager);
            OrderQueue.Add(new Order("Pickup", 57, 1, "down"));
            while (OrderQueue.Count > 0)
            {
                LoopTest(myPanel, myCageManager);
            }
            myCageManager.takeCageStatus();
        }
       
        static void Scenario4(Panel myPanel, CageManager myCageManager)
        {
            myCageManager.colList[0].cages[0].levelActual = -4;
            myCageManager.colList[0].cages[1].levelActual = 1;
            myCageManager.colList[0].cages[2].levelActual = -3;
            myCageManager.colList[0].cages[3].levelActual = -6;
            myCageManager.colList[0].cages[4].levelActual = -1;
            fmt.Println("---------Scenario 4------------");
            fmt.Println(" Pickup: -5, direction: down");
            fmt.Println(" Pickup: 1, direction: up");
            fmt.Println(" Pickup: -3, direction: down");
            fmt.Println("---------Scenario 4-----------");
            myCageManager.colList[0].cages[2].status = "Loading";
            myCageManager.colList[0].cages[3].status = "Loading";
            myCageManager.colList[0].cages[4].status = "Loading";
            myCageManager.colList[0].cages[2].direction = "down";
            myCageManager.colList[0].cages[3].direction = "up";
            myCageManager.colList[0].cages[4].direction = "down";
            //public Order(string status, int pickup, int destination, string direction)
            OrderQueue.Add(new Order("Destination", 0, -5, "down"));
            OrderQueue[0].assignment = "Assigned";
            myCageManager.colList[0].cages[2].destReq.Add(OrderQueue[0]);
            OrderQueue.Add(new Order("Destination", 0, 1, "up"));
            OrderQueue[1].assignment = "Assigned";
            myCageManager.colList[0].cages[3].destReq.Add(OrderQueue[1]);
            OrderQueue.Add(new Order("Destination", 0, -3, "down"));
            OrderQueue[2].assignment = "Assigned";
            myCageManager.colList[0].cages[4].destReq.Add(OrderQueue[2]);
            LoopTest(myPanel, myCageManager);

            while (OrderQueue.Count > 0)
            {
                LoopTest(myPanel, myCageManager);
            }
            myCageManager.takeCageStatus();
        }
        
        static void Main(string[] args)
        {
            Console.Title= "Rockets_Controllers_elevtors  Author : Said Bourakkadi";
            
            bool useDemoConfig = true;
            ConsoleKeyInfo letters;
            do
            {
                fmt.Println("Actual configuration? [y/n]");
                
                while (Console.KeyAvailable == false)
                {
                    Thread.Sleep(100); // Loop until valid input is entered.
                }

                letters = Console.ReadKey(true);
                if (letters.Key != ConsoleKey.Y && letters.Key != ConsoleKey.N)
                {
                    fmt.Println("You pressed the '{0}' key.  make a valid selection.", letters.Key);
                }
                else if (letters.Key == ConsoleKey.N)
                {
                    continue;
                    
                }
            } while (letters.Key != ConsoleKey.Y);

            // ================================CONFIGURATION====================== //
            if (useDemoConfig)
            {
                Configuration.batteryOn = true;
                Configuration.totalColumns = 4;
                Configuration.cagesPerColumn = 5;
                Configuration.totalFloors = 60;
                Configuration.totalBasements = 6;

                // CONFIRM SETUP //
                fmt.Println("\n=======================");
                fmt.Println("\n|     CONFIGURATION    |");
                fmt.Println("\n=======================");
                fmt.Println($"\n{"Categories",-17} {"Value",15}\n");
                fmt.Println($"{"Battery",-17} {"On",15}");
                fmt.Println($"{"Total Columns",-17} {Configuration.totalColumns,15}");
                fmt.Println($"{"Cages Per Column",-17} {Configuration.cagesPerColumn,15}");
                fmt.Println($"{"Total Floors",-17} {Configuration.totalFloors,15}");
                fmt.Println($"{"Total Basements",-17} {Configuration.totalBasements,15}");
            }

            // INSTANTIATE FLOORS //
            Configuration.GenerateFloors();

            // INSTANTIATE CAGEMANAGER //
            CageManager myCageManager = new CageManager();

            // INSTANTIATE PANEL //
            Panel myPanel = new Panel();

            while (Configuration.batteryOn)
            {
                int selection = Configuration.takeIntInput("\nSelect a scenario\n([1,2,3,4]\t to EXIT [0])\n", 0);
                if (selection == 1)
                {
                    Scenario1(myPanel, myCageManager);
                }
                else if (selection == 2)
                {
                    Scenario2(myPanel, myCageManager);
                }
                else if (selection == 3)
                {
                    Scenario3(myPanel, myCageManager);
                }
                else if (selection == 4)
                {
                    Scenario4(myPanel, myCageManager);
                }
                else if (selection == 0)
                {
                    Configuration.batteryOn = false;
                }
                else
                {
                    fmt.Println(selection + " is not a valid selection. Make a valid selection.");
                }
            }
        }
}
}
