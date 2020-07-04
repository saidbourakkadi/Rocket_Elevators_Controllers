
//////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////Author : Said Bourakkadi/////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////// Commercial Controller//////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////// 1 battery <- 4 colomns <- 5 cages ////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////
using System;
using System.Collections.Generic;
using System.Threading;
using System.Linq;
using System.Text;

namespace Rockets_Controllers_elevtors
    {
       
    //////////////////////////////////////////////////////////////////////////
    // Buttons : are used to generate pickup Orders from either the 
    //           panel or from a given floor.                                 
    //////////////////////////////////////////////////////////////////////////

    public class CallButton
    {
        public int id;
        public string status;

        public CallButton(int id, string status)
        {
            this.id = id;
            this.status = status;
        }

        public void CallButtonPressed()
        {
            this.status = "Active";// Active, Inactif
        }
    }

    public class FloorButton
    {
        public int id;
        public string status;

        public FloorButton(int id, string status)
        {
            this.id = id;
            this.status = status;
        }
    }

    ///////////////////////////////////////////////////////////////////////////////////////////////
    // Cages  Each cage contain all the necessary methods for                         
    //        Moving up and down  or stopping and                                   
    //        Opening and closing the door                                          
    //        Assure the security (with sensors )                                     
    //         Receive the destination order from battery and manage it.                  
    /////////////////////////////////////////////////////////////////////////////////////////////////

    public class Cage
    {
        public readonly int id;
        public string status;// In-service, Loading, idle.
        public string doors; // Open or Cloded
        public int levelActual = 1;
        public string direction = "up"; //up, down
        public int timer = 0;
        public List<Order> picReq = new List<Order>();
        public List<Order> destReq = new List<Order>();


        public Cage(int id, string status, string doors)
        {
            this.id = id;
            this.status = status;
            this.doors = doors;
        }

        public void CleanUpOrders()
        {
                       for (int y = this.picReq.Count - 1; y >= 0; y--)
            {
                if (this.levelActual == this.picReq[y].pickup)
                {
                    this.picReq[y].status = "Destination";
                }
                if (this.picReq[y].status == "Destination")
                {
                    this.destReq.Add(this.picReq[y]);
                    this.picReq.Remove(this.picReq[y]);
                    Console.WriteLine("Destination now is " + this.destReq[y].destination);
                }
            }
            for (int y = this.destReq.Count - 1; y >= 0; y--)
            {
                if (this.levelActual == this.destReq[y].destination)
                {
                    this.destReq[y].status = "Finished";
                }
                if (this.destReq[y].status == "Finished")
                {
                    this.destReq.Remove(this.destReq[y]);
                }
            }
        }
        // Door Methods //
        public void open_door()
        {
            if (this.status != "In-Service")
            {
                this.doors = "Open";
                Console.WriteLine("Cage " + this.id + " doors are open for 5 seconds");
                this.timer = 5;
                while (this.timer > 0)
                {
                    Console.WriteLine("Closing in " + this.timer + " seconds.");
                    Thread.Sleep(1000);
                    this.timer -= 1;
                   
                }Console.Beep();
                this.close_door();
            }
        }
        public void close_door()
        {
            if (this.timer < 5)
            {
                this.doors = "Closed";
                Console.WriteLine("Cage doors are closed");
                if (this.picReq.Count == 0 && this.destReq.Count == 0)
                {
                    this.status = "Idle";
                }
                else
                {
                    this.status = "Loading";
                }
                 Console.WriteLine("Beep number {0}.", this.timer);
                Console.Beep();
            }
        }
        public void OpenButtonPressed()
        {
            if (this.status != "In-Service")
            {
                this.open_door();
            }
        }
        public void CloseButtonPressed()
        {
            if (this.timer < 5)
            {
                this.close_door();
            }
        }
        

        

        // Movement //
        public void move_down()
        {
            while (this.doors != "Closed")
            {
                this.close_door();
            }
            this.status = "In-Service";
            this.direction = "down";
            Console.WriteLine("Cage " + this.id + "\tDirection is down\tlevelActual " + this.levelActual);
            if (this.levelActual - 1 == 0)
            {
                this.levelActual -= 2;
            }
            else
            {
                this.levelActual -= 1;
            }
            Console.WriteLine("Cage " + this.id + "\t\t\t\tlevelActual now is "  + this.levelActual);
            this.status = "Loading";
        }

        public void move_up()
        {
            while (this.doors != "Closed")
            {
                this.close_door();
            }
            this.status = "In-Service";
            this.direction = "up";
            Console.WriteLine("Cage " + this.id + "\tDirection is up\t\tlevelActual " + this.levelActual);
            if (this.levelActual + 1 == 0)
            {
                this.levelActual += 2;
            }
            else
            {
                this.levelActual += 1;
            }
            Console.WriteLine("Cage " + this.id + "\t\t\t\tlevelActual now is " + this.levelActual);
            this.status = "Loading";
        }
    }
    ///////////////////////////////////////////////////////////////////////////////////////////////
    // Clomns(4):  Each Column object has a list of Cage objects (5)                         
    //                                                 
    //    Receive the destination order from battery and manage it.                    
    /////////////////////////////////////////////////////////////////////////////////////////////////

    public class Column
    {
        public readonly int id;// identifier column
        public readonly string status;
        public readonly List<Cage> cages;//list cages (5)
        public readonly List<int> floors_column; // floors who column deserved

        public Column(int id, List<Cage> cages, List<int> floors_column)
        {
            this.id = id;
            this.status = "Actif";// supposed Actif all times
            this.cages = cages;
            this.floors_column = floors_column;
        }
        
    }

    //////////////////////////////////////////////////////////////////////////
    // Panel : simulates a panel in the reception of the building.    
    //          This panel directs the user to the appropriate column for their 
    //          Ordered floor and sends the appropriate pickup Order.  
    //////////////////////////////////////////////////////////////////////////

    public class Panel
    {
        public readonly List<FloorButton> floorButtons = new List<FloorButton>();

        public Panel()
        {
            for (int x = 0 - Configuration.totalBasements; x < 0; x++)
            {
                floorButtons.Add(new FloorButton(x, "Inactive"));
            }
            for (int x = 1; x <= Configuration.totalFloors; x++)
            {
                floorButtons.Add(new FloorButton(x, "Inactive"));
            }
        }

        // Methods //
        public void OrderElevator(int floorNumber, CageManager cageManager)
        {
            foreach (FloorButton button in floorButtons)
            {
                if (button.id == floorNumber)
                {
                    button.status = "Active";
                }
            }

            Column myColumn = cageManager.takeColumn(1, floorNumber);
            Console.WriteLine("Floor Ordered. \tProceed to column " + myColumn.id);

        }

        // Reports //
        public void GetFloorButtonsStatus()
        {
            for (int x = 0; x < this.floorButtons.Count; x++)
            {
                Console.WriteLine("Floor " + this.floorButtons[x].id + " button is " + this.floorButtons[x].status);
            }
        }
    }


    /////////////////////////////////////////////////////////////////////////////////
    // Floors : The floor object is generated by the Configuration object as a    //
    // list of floors each with a call button equal to the number of             //
    // total floors set by the user.                                            //
    /////////////////////////////////////////////////////////////////////////////

    public class Floor
    {
        public readonly int id;
        public CallButton button;

        public Floor(int id, CallButton button)
        {
            this.id = id;
            this.button = button;
        }
    }


    ////////////////////////////////////////////////////////////////////////
    //                           Order                                    //
    ////////////////////////////////////////////////////////////////////////
    ////////////////////////////////////////////////////////////////////////
    // A Order object is generated each time a Floor or Call button is  //
    // pressed. The Order is queued by the main program before being    //
    // assigned to a cage for treatment.                                  //
    ////////////////////////////////////////////////////////////////////////

    public class Order
    {
        public string status;// pickup
        public string assignment = "Unassigned";//Assigned or Unassigned
        public int pickup; // floor pickup 
        public int destination; // floor destination
        public string direction; //

        public Order(string status, int pickup, int destination, string direction)
        {
            this.status = status;
            this.pickup = pickup;
            this.destination = destination;
            this.direction = direction;
        }
        
    }


    

    ////////////////////////////////////////////////////////////////////////
    //                  MANAGER COLUMNS AND CAGES
    //This object contains all the column and cage objects in the system //
    // Only one CageManager should instantiated and only after Config has //
    // been called during the initial setup.                              //
    ////////////////////////////////////////////////////////////////////////

    public class CageManager
    {
        public List<Column> colList = new List<Column>();

        public CageManager()
        {
            int floorRange;
            int extraFloors;
            if (Configuration.totalBasements > 0)
            {
                if ((Configuration.totalFloors - 1) % (Configuration.totalColumns - 1) != 0)
                {
                    floorRange = ((Configuration.totalFloors - 1) / (Configuration.totalColumns - 1));
                    extraFloors = ((Configuration.totalFloors - 1) % (Configuration.totalColumns - 1));
                }
                else
                {
                    floorRange = ((Configuration.totalFloors - 1) / (Configuration.totalColumns - 1));
                    extraFloors = 0;
                }
            }
            else
            {
                if (Configuration.totalFloors % Configuration.totalColumns != 0)
                {
                    floorRange = Configuration.totalFloors / Configuration.totalColumns;
                    extraFloors = Configuration.totalFloors % Configuration.totalColumns;
                }
                else
                {
                    floorRange = Configuration.totalFloors / Configuration.totalColumns;
                    extraFloors = 0;
                }
            }

            List<Column> colList = this.colList;

            if (Configuration.totalBasements > 0)
            {
                int floorCounter = 2;
                List<int> bColumnFloors = new List<int>();
                for (int x = 0; x < Configuration.totalBasements; x++)
                {
                    if (Configuration.floorList[x].id < 0)
                    {
                        bColumnFloors.Add(Configuration.floorList[x].id);
                    }
                }
                bColumnFloors.Add(1);
                colList.Add(new Column(1, this.GenerateCages(Configuration.cagesPerColumn), bColumnFloors));
                for (int x = 2; x <= Configuration.totalColumns; x++)
                {
                    List<int> floors_column = new List<int>();
                    floors_column.Add(1);
                    if (Configuration.totalColumns - x < extraFloors)
                    {
                        for (int i = 0; i < floorRange + 1; i++)
                        {
                            floors_column.Add(floorCounter);
                            floorCounter++;
                        }
                        colList.Add(new Column(x, this.GenerateCages(Configuration.cagesPerColumn), floors_column));
                    }
                    else
                    {
                        for (int i = 0; i < floorRange; i++)
                        {
                            floors_column.Add(floorCounter);
                            floorCounter++;
                        }
                        colList.Add(new Column(x, this.GenerateCages(Configuration.cagesPerColumn), floors_column));
                    }
                }
            }
            else
            {
                int floorCounter = 2;
                for (int x = 1; x <= Configuration.totalColumns; x++)
                {
                    List<int> floors_column = new List<int>();
                    floors_column.Add(1);
                    for (int i = 0; i < floorRange; i++)
                    {
                        floors_column.Add(floorCounter);
                        floorCounter++;
                    }
                    colList.Add(new Column(x, this.GenerateCages(Configuration.cagesPerColumn), floors_column));
                }
            }
        }

        // Methods //

        // This method loops through all cages in a given column and returns //
        // the cage which can most efficiently fulfill the given Order.    // 
        public Cage takeCage(string direction, int column, int reqFloor)
        {
            Cage currentCage = this.colList[column].cages[0];
            Cage bestCage = this.colList[column].cages[0];
            int x = 0;
            while (x < this.colList[column].cages.Count)
            {
                currentCage = this.colList[column].cages[x];
                if (currentCage.direction == direction && direction == "up" && currentCage.levelActual < reqFloor && (currentCage.status == "In-Service" || currentCage.status == "Loading"))
                {
                    Console.WriteLine("Same direction UP was selected"); 
                    return currentCage; // Returns the cage with the same direction (UP) that has not yet passed the Ordered floor
                }
                else if (currentCage.direction == direction && direction == "down" && currentCage.levelActual > reqFloor && (currentCage.status == "In-Service" || currentCage.status == "Loading"))
                {
                    Console.WriteLine("Same direction DOWN was selected"); 
                    return currentCage; // Returns the cage already going the same direction (DOWN) that has not yet passed the Ordered floor
                }
                else if (currentCage.status == "Idle")
                {
                    bool allCagesAreIdle = true;
                    for (int r = 0; r < this.colList[column].cages.Count; r++)
                    {
                        if (this.colList[column].cages[r].status != "Idle")
                        {
                            allCagesAreIdle = false;
                        }
                    }
                    if (allCagesAreIdle)
                    {
                        for (int i = x + 1; i < this.colList[column].cages.Count; i++)
                        {
                            Cage compareCage = this.colList[column].cages[i];
                            if (compareCage.status == "Idle")
                            {
                                Console.WriteLine("Cage " + bestCage.id + "\tto be compared to " + compareCage.id); 
                                int before = Math.Abs(bestCage.levelActual - reqFloor);
                                int after = Math.Abs(compareCage.levelActual - reqFloor);
                                if (after < before)
                                {
                                    bestCage = compareCage; // Closest idle cage
                                }
                            }
                            Console.WriteLine("Cage " + currentCage.id + " is selected."); 
                        }
                        return bestCage;
                    }
                }
                else
                {
                    for (int i = 0; i < this.colList[column].cages.Count; i++)
                    {
                        if (direction == "up" && this.colList[column].cages[i].destReq.Count < currentCage.destReq.Count)
                        {
                            currentCage = this.colList[column].cages[i];
                        }
                        else if (direction == "down" && this.colList[column].cages[i].picReq.Count < currentCage.picReq.Count)
                        {
                            currentCage = this.colList[column].cages[i];
                        }
                    }
                     
                }
                x++;
            }Console.WriteLine("The most suitable cage is selected");
            return currentCage; 
        }

        // Returns a column where the Ordered floor is served //
        public Column takeColumn(int pickup, int destination)
        {
            bool pickupServed = false;
            bool destServed = false;
            foreach (Column column in this.colList)
            {
                foreach (int id in column.floors_column)
                {
                    if (id == pickup)
                    {
                        pickupServed = true;
                    }
                    if (id == destination)
                    {
                        destServed = true;
                    }
                    if (pickupServed && destServed)
                    {
                        return column;
                    }
                }
            }
            return null;
        }

        // Instantiates cages based off a given number //
        public List<Cage> GenerateCages(int numCages)
        {
            List<Cage> cageList = new List<Cage>();
            for (int x = 1; x <= numCages; x++)
            {
                cageList.Add(new Cage(x, "Idle", "Closed"));
            }
            return cageList;
        }


        // Watch all columns and their cages as well as their current floor and status //
        public void takeCageStatus()
        {
            
            Console.WriteLine($"\n{"| Column ",0} {"| Cage ",0} {"| Status ",0} {"| level Actual ",0} {"| Door status |",0}");
                     
            for (int x = 0; x < this.colList.Count; x++)
            {
                for (int i = 0; i < this.colList[x].cages.Count; i++)
                {
                    Cage currentCage = this.colList[x].cages[i];
                    
                    Console.WriteLine($"\n{this.colList[x].id,5} {currentCage.id,8} {currentCage.status,10} {currentCage.levelActual,10} {currentCage.doors,17}");
                }
            }
        }

        // Returns a string of the floors served by a given column //
        public string Take_Floors_column(Column column)
        {
            string myFloors = string.Join(",", column.floors_column);
            string myString = "Column " + column.id + ": " + myFloors;
            return myString;
        }
    }


    

    ////////////////////////////////////////////////////////////////////////
    // Configuration : This static object generates a configuration
    //               from user input and the corresponding floor list.                         
    ////////////////////////////////////////////////////////////////////////

    public static class Configuration
    {
        public static bool batteryOn;
        public static int totalColumns;
        public static int cagesPerColumn;
        public static int totalFloors;
        public static int totalBasements;

        public static List<Floor> floorList = new List<Floor>();

        
        
        public static int GetIntInput(string prompt, uint minValue)
        {
            Console.WriteLine(prompt);
            int myInput = -1;
            string userInput = Console.ReadLine();
            while (myInput == -1)
            {
                try
                {
                    myInput = Convert.ToInt32(userInput);
                    if (myInput < minValue)
                    {
                        Console.WriteLine("Value cannot be less than " + minValue + ".");
                        myInput = -1;
                        userInput = "";
                    } 
                }
                catch (System.FormatException)
                {
                    if (userInput == "")
                    {
                        Console.WriteLine("Enter a valid number.");
                        userInput = Console.ReadLine();
                    }
                    else
                    {
                        Console.WriteLine(userInput + " is not a valid number.\n Enter a valid number.");
                        userInput = Console.ReadLine();
                    }
                }
            }
            return myInput;
        }

        // To be called once upon startup: Generates a configuration based on user input //
        public static void Config()
        {
            ConsoleKeyInfo letters;
            do
            {
                Console.WriteLine("Activate battery? [y/n]");
                while (Console.KeyAvailable == false)
                {
                    Thread.Sleep(100); // Loop until valid input is entered.
                }

                letters = Console.ReadKey(true);
                if (letters.Key != ConsoleKey.Y && letters.Key != ConsoleKey.N)
                {
                    Console.WriteLine("You pressed the '{0}' key.  make a valid selection.", letters.Key);
                }
                else if (letters.Key == ConsoleKey.N)
                {
                    Console.WriteLine("Startup Aborted!");
                    return;
                }
            } while (letters.Key != ConsoleKey.Y);

            Console.WriteLine("Initializing...");

            // Set total number of columns //
            int totalColumns = GetIntInput("Enter the total number of columns", 1);

            // Set cages per column //
            int cagesPerColumn = GetIntInput("How many cages are installed per column?", 1);

            // Set number of floors //
            int totalFloors = GetIntInput("How many floors (excluding basements) are there in the building?", 2);

            // Set number of basements //
            int totalBasements = GetIntInput("How many basements are there?", 0);

            // Set Configuration Values //
            Configuration.batteryOn = true;
            Configuration.totalColumns = totalColumns;
            Configuration.cagesPerColumn = cagesPerColumn;
            Configuration.totalFloors = totalFloors;
            Configuration.totalBasements = totalBasements;

            // Confirm Setup Conditions //
            Console.WriteLine("\n-------SIMULATION-------");
            Console.WriteLine($"\n{"Categories",-15} {"Value",15}\n");
            Console.WriteLine($"{"Battery",-15} {"On",15}");
            Console.WriteLine($"{"Total Columns",-15} {Configuration.totalColumns,15}");
            Console.WriteLine($"{"Cages Per Column",-15} {Configuration.cagesPerColumn,15}");
            Console.WriteLine($"{"Total Floors",-15} {Configuration.totalFloors,15}");
            Console.WriteLine($"{"Total Basements",-15} {Configuration.totalBasements,15}");
        }

        // To be called after Config: Generates floor objects and adds them to floor list //
        public static void GenerateFloors()
        {
            // Checks if building has basements and adds them to the floor list //
            if (totalBasements > 0)
            {
                for (int x = 0 - totalBasements; x < 0; x++)
                {
                    floorList.Add(new Floor(x, new CallButton(x, "Inactive")));
                }
            }

            // Adds remaining floors to the floor list //
            for (int x = 1; x < 1 + totalFloors; x++)
            {
                floorList.Add(new Floor(x, new CallButton(x, "Inactive")));
            }
        }

        // Reports //
        public static void GetFloorStatus()
        {
            Console.WriteLine("\n-----------------FLOOR STATUS------------------\n");
            for (int x = 0; x < floorList.Count; x++)
            {
                Console.WriteLine(String.Format("{0, -6} {1, 2} {2, -26} {3, -8}", "Floor ", floorList[x].id, ":  Active  //  Call Status: ", floorList[x].button.status));
            }
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
                    Console.WriteLine("Floor button " + floor.button.id + " is active.");
                    if (floor.id > 0)
                    {
                        Order myOrder = new Order("Pickup", floor.button.id, 1, "down");
                        foreach (Order Order in OrderQueue)
                        {
                            if (floor.button.id == Order.pickup && Order.status == "Pickup")
                            {
                                Console.WriteLine("My Order for floor " + floor.button.id + " was not sent.");
                                return;
                            }
                        }
                        OrderQueue.Add(myOrder);
                        Console.WriteLine("My Order for floor " + myOrder.pickup + " was added to the Order list");
                    }
                    else
                    {
                        Order myOrder = new Order("Pickup", floor.id, 1, "up");
                        foreach (Order Order in OrderQueue)
                        {
                            if (floor.id == Order.pickup && Order.status == "Pickup")
                            {
                                Console.WriteLine("My Order for floor " + floor.button.id + " was not sent.");
                                return;
                            }
                        }
                        Console.WriteLine("My Order for floor " + floor.button.id + " was added to the Order list");
                        OrderQueue.Add(myOrder);
                    }
                    floor.button.status = "Inactive";
                    Console.WriteLine("Floor " + floor.button.id + " is " + floor.button.status);
                }
            }

            // Check floor buttons //
            foreach (FloorButton button in myPanel.floorButtons)
            {
                if (button.status == "Active")
                {
                    Console.WriteLine("Panel button " + button.id + " is " + button.status);
                    if (button.id > 0)
                    {
                        Order myOrder = new Order("Pickup", 1, button.id, "up");
                        foreach (Order Order in OrderQueue)
                        {
                            if (myOrder.destination == Order.destination && Order.status == "Pickup")
                            {
                                Console.WriteLine("My Order for floor " + button.id + " was not sent.");
                                return;
                            }
                        }
                        Console.WriteLine("My Order for floor " + button.id + " was added to the Order list");
                        OrderQueue.Add(myOrder);
                    }
                    else
                    {
                        Order myOrder = new Order("Pickup", 1, button.id, "down");
                        foreach (Order Order in OrderQueue)
                        {
                            if (myOrder.destination == Order.destination && Order.status == "Pickup")
                            {
                                Console.WriteLine("My Order for floor " + button.id + " was not sent.");
                                return;
                            }
                        }
                        Console.WriteLine("My Order for floor " + myOrder.pickup + " was added to the Order list");
                        OrderQueue.Add(myOrder);
                    }
                    button.status = "Inactive";
                    Console.WriteLine("Floor " + button.id + " is " + button.status);
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
                    Console.WriteLine("Column " + myColumn.id + " is selected.");
                    Cage myCage = myCageManager.takeCage(Order.direction, myColumn.id - 1, Order.pickup);
                    Order.assignment = "Assigned";
                    myCage.picReq.Add(Order);
                    Console.WriteLine("Cage " + myCage.id + " receive Order for floor " + myCage.picReq[0].pickup);
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
             

            Console.WriteLine("---------Scenario 1------------");
            Console.WriteLine(" Pickup: 7, direction: down");
            Console.WriteLine(" Pickup: 13, direction: up");
            Console.WriteLine(" Pickup: 5, direction: down");
            Console.WriteLine(" Pickup: 2, direction: down");
            Console.WriteLine(" Pickup: 5, direction: down");
            Console.WriteLine(" some one is floor 1 and request destination :17 ");
            Console.WriteLine("---------Scenario 1------------");
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
            Console.WriteLine(" destination: 1, direction: down");
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
            Console.WriteLine("---------Scenario 2------------");
            Console.WriteLine(" Pickup: 21, direction: up");
            Console.WriteLine(" Pickup: 28, direction: up");
            Console.WriteLine(" Pickup: 1, direction: down");
            Console.WriteLine(" Pickup: 24, direction: down");
            Console.WriteLine(" Pickup: 3, direction: down");
            Console.WriteLine(" some one is floor 1 and request destination :33 ");
            Console.WriteLine("---------Scenario 2-----------");
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
            Console.WriteLine("---------Scenario 3------------");
            Console.WriteLine(" Pickup: 1, direction: down");
            Console.WriteLine(" Pickup: 63, direction: up");
            Console.WriteLine(" Pickup: 54, direction: up");
            Console.WriteLine(" Pickup: 50, direction: down");
            Console.WriteLine(" Pickup: 1, direction: up");
            Console.WriteLine(" some one is floor 57 and request destination :1 ");
            Console.WriteLine("---------Scenario 3-----------");
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
            Console.WriteLine("---------Scenario 4------------");
            Console.WriteLine(" Pickup: -5, direction: down");
            Console.WriteLine(" Pickup: 1, direction: up");
            Console.WriteLine(" Pickup: -3, direction: down");
            Console.WriteLine("---------Scenario 4-----------");
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
                Console.WriteLine("Actual configuration? [y/n]");
                
                while (Console.KeyAvailable == false)
                {
                    Thread.Sleep(100); // Loop until valid input is entered.
                }

                letters = Console.ReadKey(true);
                if (letters.Key != ConsoleKey.Y && letters.Key != ConsoleKey.N)
                {
                    Console.WriteLine("You pressed the '{0}' key.  make a valid selection.", letters.Key);
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
                Console.WriteLine("\n=======================");
                Console.WriteLine("\n|     CONFIGURATION    |");
                Console.WriteLine("\n=======================");
                Console.WriteLine($"\n{"Categories",-17} {"Value",15}\n");
                Console.WriteLine($"{"Battery",-17} {"On",15}");
                Console.WriteLine($"{"Total Columns",-17} {Configuration.totalColumns,15}");
                Console.WriteLine($"{"Cages Per Column",-17} {Configuration.cagesPerColumn,15}");
                Console.WriteLine($"{"Total Floors",-17} {Configuration.totalFloors,15}");
                Console.WriteLine($"{"Total Basements",-17} {Configuration.totalBasements,15}");
            }

            // INSTANTIATE FLOORS //
            Configuration.GenerateFloors();

            // INSTANTIATE CAGEMANAGER //
            CageManager myCageManager = new CageManager();

            // INSTANTIATE PANEL //
            Panel myPanel = new Panel();

            while (Configuration.batteryOn)
            {
                int selection = Configuration.GetIntInput("\nSelect a scenario\n([1,2,3,4]\t to EXIT [0])\n", 0);
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
                    Console.WriteLine(selection + " is not a valid selection. Make a valid selection.");
                }
            }
        }
}
}
