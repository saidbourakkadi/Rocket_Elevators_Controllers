Rocket_Elevators_Controllers
algorithm residential and commercial elevators

Efficiency; is not the most important feature, the most important is to make sure every order is followed, that there is no starvation. 
If someone presses 60 and people keep pressing 1 and 2 it may be efficient to keep going between those floors, but be nice for 60 
to be visited at some point. I think (from personal observation when I was interested in figuring out) that most of them do:

Start going in the direction of the first button pressed, keep track of which direction we are going
When a floor is reached and that button was pressed, stop and open the doors, mark the buttons for this floor as not pressed anymore.
If there are still more floors that we need to visit that are in the same direction, keep going in that direction.  If not and
there are still floors we need to visit, move in that direction.  If not then we're done and will start at 1 when a button is pressed
again. Note that many elevators have buttons "I want to go up" and "I want to go down" next to the doors instead of a single button. 
The algorithm only needs a small change: in 2, if the only button pressed for that floor is one of the buttons next to the door,
only stop and open the doors if we are going in that direction. Possibly keep the button pressed if the doors open because of a button
pressed inside the elevator and it is going in the wrong direction. You never have to figure out an entire path, just in which 
direction to go next.
As long as there’s someone inside or ahead of the elevator who wants to go in the current direction, keep heading in that direction.
Once the elevator has exhausted the requests in its current direction, switch directions if there’s a request in the other direction.
Otherwise, stop and wait for a call.
I didn't have time to test every single scenarioes, but for the moment I focused on the fact that it working and this week-end 
I will performe a better testing.
