# Dining Philosophers Golang


### features
* it's not  inspired from any reference. just a simple solution and it works based on tests.
* used random time to show philosophers stop eating on variable time.
* 2 channels(eat, think) communicate with each other to run the simulation to infinite
* there is only a time limit cap to make sure the simulation finish. 
* at the start loop all philosophers and find open chops and let eat those with open chops
* when any philosopher stops eating chooses a neighbor with less eat count.
* the sim doesn't need any exclusive locking as it's handed by logic. 
* locks are bad based on waste of cpu cycles
* results showed my no lock and no philosopher starvation. 

### warnings

* concentrated on solving the problem
* based on simplicity not much care on architecture and principles