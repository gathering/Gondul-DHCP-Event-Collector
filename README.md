# Gondul DHCP Event Collector

dhcpd runs 'on commit' for every lease being handed out to clients  
this will push the lease info to the gondul api

# Examples
See the [examples](examples) folder for the config needed to support this.

# Installing
Download latest release to /usr/local/sbin/

# Sample data
Below is a sample of the data generated that is sendt to the Gondul collector api

```
{
  "src": "dhcp",
  "metadata": {
    "server": "dhcp.gathering.org"
  },
  "data": {
    "clientip": "2001:db8:db8:e11:0:0:0:1337",
    "clientmac": "06:c1:a8:ab:e7:47",
    "clientname": "WIN-589NNPG7EGD",
    "time": "2019-02-13T22:08:34.781014993+01:00"
  }
}
```

# Known bugs
Juniper has a bug in its dhcpv6 option 18, so we dont get circuit id for v6.  
This is reported to juniper (link) and we are waiting for a fix.  
Until this is fixed we exclude sending circuit id to the gondul api if we dont receive any dhcpv6 option18 info.  

The setting for this is in the edge-switch.conf, but disabled.
```
inactive: dhcpv6-option18 {
    use-option-82;
}
```

