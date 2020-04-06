Hi!

Welcome to the VDart Digital's
Golang coding challenge
Author: W. Allan Edwards

Overview
Client has a very demanding environment that requires strong technology delivery
personnel. This is a realistic coding exercise designed to validate a candidate’s
potential for our development needs.

Primary Goal
Deliver a working client and server in Go that leverages GRPC.

Extra Points Goal
Implement a docker solution that runs the entire application within a single shell
script. Encapsulate the entire solution in docker compose and be able to run the
solution in containers.

Deliverable Requirements

    ● 1 Golang workspace, standard Go environment setup
    ○ Workspace named demo-workspace (root directory)
    ○ Use standard src, pkg, bin directories, set gopath to root
    ● 2 executables in src workspace
    ○ 1 go console application that calls the grpc server with a client
    ○ 1 go console application that implements a grpc server with 1 hello
    world endpoint

    ○ On the server side implement 1 grpc service that implements 1
    endpoint, the endpoint should simply return an object that contains 1
    string, hell world

    ● Deliver 1 unit test with the server side solution
    ● Documentation
    ○ 1 architecture document of the solution, use visio or lucid chart to
    output a png file you package in a zip
    ○ 1 document describing how to setup and make your solution run

Expected Implementation Time

There is no time limit for this exercise. For the minimal solution we expect 2 to 3
hours. For the full on solution and creative add ons, we expect 3+ hours.
Success Criteria

    ● Can unzip the deliverable file on mac, windows, or linux and run the solution
    in Goland, Visual Studio Code, or command line. The expectation is a
    standard Go environment setup.
    ● If you went for extra points, I should be able to open the zip and run the
    shell script and run docker-compose up to execute the client calling the go
    lang server.
    ● Your documentation is well written and makes good sense

Be Creative if you Want
This document contains a baseline solution to implement. If you want to be
creative and add additional capabilities and disciplines into your solution, by all
means do it and document what they are on the deliverable. The detail we left out
on this document was intentional to see what all people might creatively produce.
    
    Cool ideas win extra points!
    
    ● Examples of creative ideas
    ○ Implement a mock and more complex transmission data structure
    ○ Implement a kubernetes compatible script that I can run against
    Kubernetes for this solution
    ○ Implement something more complex on data transmission such as
    binay blogs
    ○ Use lucid chart to develop more extensive documentation

Instructions:
Please send it as a Zip File or Share the Github link
Please complete the challenge within 24 hours timeframe.
The whole exercise is the primary scenario based.