# Architectural Decisions

## Netbox 

### Reasoning

In order to save time implementing software for managing IP addresses ([IPAM](https://en.wikipedia.org/wiki/IP_address_management)) and organizing our datacenter infrastructure ([DCIM](https://en.wikipedia.org/wiki/Data_center_management#Data_Center_Infrastructure_Management)), we decided to take advantage of third-party software solutions.

We found that DCIM and IPAM were often combined into the same piece of software. We looked at the following offers on the market:

- [netbox](https://github.com/digitalocean/netbox)
- [phpipam](https://phpipam.net/)
- [racktables](https://www.racktables.org/)
- several smaller and older projects

After looking at the solutions, it became pretty clear that only Netbox would satisfy the expectations we had. Main advantages over the other products were:

- Versatile data model that can actually map our Metal Cloud datacenter (especially VRFs and Blade Chassis, i.e. Blade Servers can be children of a chassis and are located in device bays)
- Complete REST API with Swagger + authentication
- Vibrant community, many Github stars
- The UI in comparison to other products was pleasing to the eye

The main issues we had on our journey with Netbox turned out to be:

- The generated Swagger Go client does not work at all (issues with type conversions, which cannot be resolved easily)
- Data model basically provides us with everything we need, but some stuff is called differently or is not there
- Performance (a lot of client calls are required in order to achieve what we want)

In order to tackle the main issue regarding the Swagger Client, we decided to introduce another microservice. The [netbox-api-proxy](https://git.f-i-ts.de/cloud-native/metal/netbox-api-proxy).

## On How Netbox is Being Utilized in Metal

The Netbox is a shadow inventory of our datacenter infrastructure and manages the IP addresses that are assigned to metal devices. We made the following decisions of what Netbox actually is being used for:

- IPAM -> Managing IP addresses (that will be assigned to the devices)
- Receive information from the `metal-api` on devices and persist them (including serial, site and rack location, network interfaces, owners...)
- Browsing the inventory with the Netbox web UI to get an overview over the datacenter

If the Netbox is not present, device registration, allocation and release does not function.

It is not used for:

- Querying things (like requesting all devices to find out who the owners are), the `metal-api` does not rely on data stored in the Netbox for queries
- Persisting the entire data relevant to the clients of the `metal-api` (e.g. SSH keys, OOB management credentials, ... are not stored in the Netbox). Some data is not really considered in the Netbox data model and we neither want to adopt their model nor do we want the Netbox to adopt our model.