
[![CircleCI](https://circleci.com/gh/felts94/QueueUsingStacks.svg?style=svg)](https://circleci.com/gh/felts94/QueueUsingStacks)

# QueueUsingStacks
This is an implementation of a "generic" queue using only stacks written in golang. Technically go does not allow generic types so I used \*[]interface{} which is flexible but heavy overhead on runtime

This implementation is used to drive the backend of the queue-app: [live](http://kfelter.com:8080/) | [code](https://github.com/felts94/queue-app)
