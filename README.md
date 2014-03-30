george
======

# IN DEVELOPMENT

George is kinda like an old school switch board operator. He keeps an eye on where things are from the point of view of a node and routes requests to available services.

A node sits in a different environment to any form of service discovery and is in the best position to make decisions about which of the services it knows about it can connect to. 

#### For example..
A webserver sits on a network happily speaking to a redis server, possibly on a different cloud, happily making queries (not writing, that's a whole different problem).  It also knows about another redis server, a slave of the primary one. The network drops between the clouds - OH NO! George steps in, having been talking to the primary server forever and tells the web app that the slave is now the one that it wants to talk to. simples.

This is a separate solution to the redis server being generally unavailable - that's what something like Serf is for - a centrally held config sent to nodes.

The main point is that between multiple networks, and even nodes, you can never guarantee the environment is the same as a central controller. A service one node is having trouble with, might be generally available to others - simply switching to alternate endpoints might be the simplest way of solving a temporary network glitch.


## Example
```go
package main

import (
	"github.com/matzhouse/george/george"
)

func main() {

	george.Start()

}
```
