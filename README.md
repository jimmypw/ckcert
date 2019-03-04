# ckcert

A small utility that checks certificate expiry dates.

The exit status
* OK      = 0
* RENEW   = 1
* ERROR   = 3

## installation

`go install github.com/jimmypw/ckcert`

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