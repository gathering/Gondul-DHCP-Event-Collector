# Gondul DHCP Event Collector

dhcpd runs 'on commit' for every lease beeing handed out to clients  
this will push the lease info to the gondul api

# Examples
See the [examples](examples) folder for the config needed to support this


# Known bugs
Juniper has a bug in its dhcpv6 option 18, so we dont get circuit id for v6.  
This is reported to juniper (link) and we are waiting for a fix.  
Until this is fixed we exclude sending circuit id to the gondul api if we dont receive any dhcpv6 option18 info.  
