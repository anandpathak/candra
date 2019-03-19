# Candra

<img src="assets/logo.png">

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger)


A cli tool to simplify the process of ssh into ec2 instances.

    When DevOps don't know that reboot exist in aws and always stop and starts server.
    When you have too many server to work with specially manually configuring each.
    When you get frustated with always changing server.

### Candra cli tools will help in running ssh to ec2 servers.


# Installation.

  -  `go get github.com/anandpathak/candra` 
  -  `make build && make install`
  - use the binary generated in build and enjoy!

This text you see here is *actually* written in Markdown! To get a feel for Markdown's syntax, type some text into the left window and watch the results in the right.

## Comamands
 - get the list of available command
```
candra --help
```
 - configure cli 
```
candra config add
```
 - list configuration 
```
candra config list
```
 - search for ec2 server 
```
candra search
    --flags 
        -t  aws describe instance filter tag name
        -v aws describe instance filter tag value
```

### Tech

 - this is build using golang and cobra and viper framework


License
----

MIT


**Free Software, Hell Yeah!**
