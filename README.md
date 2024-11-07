To start the program, open multiple terminals and run:

- go run Node/Node.go -p <Client Port> -tp <Target Port>
Set <Client Port> to the unique port for each instance.
Set <Target Port> to the <Client Port> of the next terminal instance to link them.

To close the client ring and initiate the process:

In the last terminal, set <Client Port> as the final node's port and <Target Port> as the first node’s <Client Port>.
Add the -s True flag to indicate this is the starting node:
- go run Node/Node.go -p <Client Port> -tp <Target Port> -s True
