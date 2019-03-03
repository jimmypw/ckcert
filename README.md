# ckcert

A small utility that checks certificate expiry dates. If there is less than half of the validity period remaining it will alert. 

The exit status
* exitOK      = 0
* exitRENEW   = 1
* exitEXPIRED = 2
* exitERROR   = 3
