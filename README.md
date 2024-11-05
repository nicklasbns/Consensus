
# BSDISYS1KU-En-GO-gruppe
**Run the program**

Open terminal session and cd to project
```
project
│   .gitignore
│   CriticalSection.txt
│   README.md
│   client.go
│   consensus.proto
└───consensus
│   │   consensus.pb.go
│   │   consensus_grpc.pb.go
└───└─────────────────────────
```

Run the clientside part of the program the amount of times that clients are needed
```
go run client.go
```

When all clients have joined and you wish to start the program then
```
go run ping.go
```

| System Requirements |  |
|--|--|
| R1 | x |
| R2 | x |
| R3 | x |

-   R1: Implement a system with a set of peer nodes and a Critical Section that represents a sensitive system operation. Any node may request access to the Critical Section at any time. In this exercise, the Critical Section is emulated, for example, by a print statement or writing to a shared file on the network.
-   R2: (Safety) Only one node may enter the Critical Section at any time.
-   R3: (Liveliness) Every node that requests access to the Critical Section will eventually gain access.
