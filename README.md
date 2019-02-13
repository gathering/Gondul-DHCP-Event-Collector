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
    "server": "dhcp.tg.lol"
  },
  "data": [
    {
      "clientip": "127.27.36.162",
      "clientmac": "0e:74:14:1f:ce:e2",
      "clientname": "WIN-8KE6TTQA821",
      "leasetime": 120,
      "circuitid": "ge-0/0/1.0:1011",
      "time": "2019-02-13T22:21:27.195685448+01:00"
    }
  ]
}
{
  "src": "dhcp",
  "metadata": {
    "server": "dhcp.tg.lol"
  },
  "data": [
    {
      "clientip": "2001:db8:db8:e11:0:0:0:1994",
      "clientmac": "1a:61:8b:60:be:e6",
      "clientname": "WIN-5BPB8IE9S48",
      "time": "2019-02-13T22:21:27.521113806+01:00"
    }
  ]
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

