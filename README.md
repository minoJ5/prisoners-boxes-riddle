<h1>Simulating the prisoners boxes riddle</h1>
This is an attempt at a smulating the prisoners boxes riddle by leveraging Go's concurrency feature. 
<h2>Short Description of the Riddle</h2>
The idea is that the (100) prisoners go into a room with (100) boxes numbered from 1 to 100 but the contents of the boxes have been shuffeled, i. e, box number 5 may contain the number 7 and so on. 
The prosoners can only open half of the boxes. 
If all the prisoners find their numbers in the shuffeled boxes, they are going to be set free. 
Succeeding in this with random opening of the boxes has a probability of (1/2)^100. 
The optimal way to do this, which will result in a probability of approx. 0.307 is for every prisoner to open the box with their number on it and if they don't find their number in the box they go to the box with the index corresponding to the number ib the opened box. 
<h2>Simulation</h2>
The way this go program works is: For every experiment, create a shuffled array, open the boxes with the method described above and then check if the expirement succeeded. 
<h2>Result</h2>
Running 10,000 experiments resulted in emperical probability P of approx. P = 0.305 
