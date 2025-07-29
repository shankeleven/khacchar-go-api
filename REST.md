REST = “Use HTTP like it was meant to be used.”
rest principles:
Resources identified by URIs
Each thing you want to interact with (a user, post, comment) has a URL.

Statelessness
Every request contains all the info needed to process it.
The server does not store client state between requests.

Client-server separation
The client (frontend or app) and server are separate and communicate over HTTP.

Uniform interface
Always use the same operations (GET, POST, etc.) in consistent ways.

Representations
You interact with a representation of the resource (like JSON, XML).
The server sends you the state of the resource, not internal logic.

Cacheable
Responses should include info about whether they can be cached.





no matter how you store the data , rest would allow the client to ask the data in his preferred format ,
your db could store data in tables and rows and deliver the content to client as json
hence the word representational state transfer
hence http is very well suited to implement rest
but REST is just a specification , you could use it over http , tcp or any other protocol , just a guideline (say)
it just happens to gel well with http(most commonly used) , hence they became synonymous
http verbs (get, put, post, delete) made it buttery smooth
delete /users/1 } for a type user , we'd like to delete resource 1.
also we could multiplex it on same resource(get put post delete /users/1) , this makes it way more intuitive
gels well with http because:
    http clients like curl , postman , requests etc
    big ecosystem already in place

