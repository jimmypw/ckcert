# ckcert

A small utility that checks certificate expiry dates. If there is less than half of the validity period remaining it will alert. 

The exit status
* exitOK      = 0
* exitRENEW   = 1
* exitEXPIRED = 2
* exitERROR   = 3

## Usage

### To check minimum remaining days

`./ckcert -f "path to certificate" -d "minimum days"`

e.g.

`./ckcert -f /tmp/cert.crt -d 100`

This would make sure there are at least 100 days remaining on the certificate

### To check minimum remaining percentage

`./ckcert -f "path to certificate" -p "percentage of validity"`

e.g. 

`./ckcert -f /tmp/cert.crt -p 30`

This will ensure that there is at least 30% of the life of the certificate remaining.