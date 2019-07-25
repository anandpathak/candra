# Candra

<img src="assets/logo.png">

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger)


A CLI tool to simplify the process of ssh into ec2 instances.


### Candra CLI tools will help in running ssh to ec2 servers.


# Installation.

  -  `go get github.com/anandpathak/candra` 
  -  `make build && make install`
  - use the binary generated in build and enjoy!



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
