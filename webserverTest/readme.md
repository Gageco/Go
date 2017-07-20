# webserverTest

To run you need to run 

```go run *.go```

This will then start a webserver at http://localhost:8080. At this point you need to use a curl or Rest service to interface with the API fully

## /timers

**Function:** creates a timer for an amount of time

**Post:** 

```
{"length": n}	   //n is timer length in seconds
```

**Get:**

```
{"id": 1, "length": 10, "finished": true}
```

An ID is generated for each Post Request sent and then you can see if the timer is finished or not

## /randomText

**Function:** Returns a random string from an array

**Get:** Returns random string 

## /stateChange

**Function:** Change the state of a specific id with a post request

**Post:**

```
{"state": false, "description": "your description"}
```

adds a new item with your given state and description, will return an id

```
{"id": 1, "state": true}
```

by using the id you can update the state of the id

**Get:** returns all items with id, state, and description

## /getID

**Function:** Allows you go get a seed which can be used in /useID. Sorta an authorization test

**Post:**

```
{"user": "gage", "pass": "coprivnicar"}
```

returns an integer which is used in /useID

## /useID

**Function:** Use the ID retrieved from /getID to succesfully access this

**Post:** 

```
{"usercode": yourInt}
```

returns success if it was the correct int or fail if it was not

