package khachhar

import (
	"log/slog"
	"net/http"
)

// this file would ofcourse return handlerfunction

func New() http.HandlerFunc{

			slog.Info("Adding a khachhar")
	  return func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("Pranaam!")) }
}


/*


http is stateles, which would mean a request is complete in itelf
not expecting the server to know anything ,hence it is very easy to use in distributed systems.
this is a stateless request handler, which means it does not maintain any state between requests
stateful request handler would mean it maintains some state between requests, usually using a database or maintaning it on client side

there are different http headers but the one thing that amazed me was how extensible they make http
http is a protocol, which means it is a set of rules that define how data is transmitted over the internet, in application layer

a request looks like :
GET / HTTP/1.1
headers......
blank line to indicate end of headers
{ json body if any, usually in POST or PUT requests }

 a response looks like :
HTTP/1.1 200 OK
headers......
blank line to indicate end of headers
 { json body if any, usually in GET or POST requests }


methods(get,post,put(complete replacement),patch(partial replacement or append action),delete) are used to define the intent of the interaction
idempotency is a key concept in http, which means that a request can be made multiple times without changing the result(Get , Put, Delete are idempotent, Post is not idempotent)
post is not idempotent because it creates a new resource every time it is called, while put replaces the resource if it exists or creates it if it does not exist

PUT /users/123
Body: {
    "name": "Alice",
    "email": "alice@example.com"
} now whether you call it once or multiple times, the user with id 123 will always have the same name and email(same state basically)
POST /users
Body: {
    "name": "Alice",
    "email": "alice@example.com"
}Every time you call this:
It may create a new user (like /users/124, /users/125, etc.)
The server might generate a new ID, timestamp, log entry, or trigger downstream effects.

then there is OPTIONS method which is used to describe the communication options for the target resource, it is used to check what methods are allowed on a resource, it is not used to get or set any data
usually used in CORS (Cross-Origin Resource Sharing) to check if the server allows requests from a different origin, it is a preflight request that checks if the actual request is safe to send
cors allows the server to specify who can access its resources, what methods are allowed, and what headers can be used in the request
CORS has primarily two types of request flows: simple requests and preflight requests

in simple requests, the browser sends the request directly to the server, while in preflight requests, the browser sends an OPTIONS request to the server to check if the actual request is safe to send
if the response includes the appropriate cors headers i.e. Access-Control-Allow-Origin, Access-Control-Allow-Methods, Access-Control-Allow-Headers, etc., the browser will send the actual request
these are GET POST HEAD requests with no custom headers or authorization headers
and with content type of `application/x-www-form-urlencoded`, `multipart/form-data`, or `text/plain`
simple means simple nothing that requires authentication or anything that requires the server to check if the request is safe to send

and in a pre-flight request flow : read about this in detail
1. The browser sends an OPTIONS request to the server with the `Origin` header to check if the server allows requests from that origin.
the method is not GET, POST, or HEAD, and the request includes custom headers or uses methods like PUT or DELETE.
or 2. the request includes some non-simple headers (like `authorization`, `x-custom-header`, etc.)
3. the request has content type other than `application/x-www-form-urlencoded`, `multipart/form-data`, or `text/plain`.
// most of our requests involve json data , so most of them are preflight requests
these include Custom headers like Authorization, X-Custom-Token A Content-Type like application/json

the response then is a 204 No Content or 200 OK with the appropriate CORS headers, indicating that the server allows requests from that origin and the methods and headers that are allowed.
no  the access-control-max-age header is used to specify how long the results of a preflight request can be cached by the browser, so that it does not have to send a preflight request for every request
this saves time and resources for these extra requests with OPTIONS method
write a mock and study using burp suite

so basically the request asks for some headers and some methods and the server responds with the allowed methods and headers, if the request is allowed
access-control-request-method/header and access-control-allow-methods/headers are the headers used in preflight requests and responses respectively

If you're using curl, Postman, or a server-side language (like Go, Python), you can make cross-origin requests freely — no CORS errors!
But in the browser, CORS restrictions apply strictly.so it is the browser that enforces CORS, not the server
A webpage can only access resources (data, cookies, DOM, etc.) from the same origin (protocol + domain + port).
The request still goes to the server, even if CORS fails.But the browser hides the response from your frontend code.


NOW we go to next part , which is status codes/response codes
these just indicate the status of the request, whether it was successful or not, and if not, what went wrong without needing us to sneak in the body of the response
for example 401 is unauthorized access , which would mean we ask the user to login again,
or 403 is forbidden access, which would mean the user does not have permission to access the resource,
or 404 is not found, which would mean the resource does not exist,
or 500 is internal server error, which would mean something went wrong on the server side
also this standard is followed so we can have streamlined interactions without depending on which stack , the server is built upon
 reponse codes are 3 digit numbers, the first digit indicates the class of the response, and the last two digits indicate the specific response code within that class
1xx - Informational: The request was received, continuing process. // commonly used with large file uploads or upgrading some protocols like websocket
2xx - Success: The request was successfully received, understood, and accepted.
	200 OK: The request has succeeded. something successfully retrieved or processed.
	201 Created: The request has been fulfilled and resulted in a new resource being created.
	204 No Content: The server successfully processed the request, but there is no content to return.
3xx - Redirection: Further action needs to be taken in order to complete the request.
	301 Moved Permanently: The requested resource has been assigned a new permanent URI. // suppose you had a route /users and you moved it to /people, then you would use this code
	302 Found: The requested resource resides temporarily under a different URI. // same but this is temporary, so later on, the client should use the original URI for future requests
	304 Not Modified: The resource has not been modified s ince the last request, so the client can use its cached version.
4xx - Client Error: The request contains bad syntax or cannot be fulfilled.' // we mostly deal with these errors
	400 Bad Request: The server cannot process the request due to client error (e.g., malformed request syntax).
 	401 Unauthorized: The request requires user authentication. // the user is not logged in or the token is invalid
	403 Forbidden: The server understood the request, but refuses to authorize it. // the user is logged in but does not have permission to access the resource
	404 Not Found: The server has not found anything matching the Request-URI. // the resource does not exist
	409 Conflict: The request could not be completed due to a conflict with the current state of the resource. // for example, trying to create a user with an email that already exists
5xx - Server Error: The server failed to fulfill an apparently valid request.
// but we do not show these errors to the user for obvious security reasons, we just log them and show a generic error message
	500 Internal Server Error: The server encountered an unexpected condition that prevented it from fulfilling the request. // something went wrong on the server side
	502 Bad Gateway: The server, while acting as a gateway or proxy, received an invalid response from the upstream server.
	503 Service Unavailable: The server is currently unable to handle the request due to temporary overloading or maintenance of the server. // the server is down or overloaded
	504 Gateway Timeout: The server, while acting as a gateway or proxy, did not receive a timely response from the upstream server.


HTTP caching is a mechanism that allows the browser to store copies of resources (like HTML pages, images, scripts, etc.)
so that it can quickly retrieve them without having to make a new request to the server.(for static data obviously, else we'd have to implement cache invalidation strategies)
cache-control is a header that is used to specify how the resource should be cached by the browser and other intermediaries (like proxies, CDNs, etc.)
it contains fields like max-age=10seconds for non-static data , eliminating the need for those cache invalidation strategies
E-tag (a hash) is a unique identifier for a specific version of a resource, it is used to check if the resource has changed since the last request
similar to what dropbox does, when uploading chunks.
now for the ensuring that the data we cache is still in sync with the server, we use the If-None-Match header in the request , that includes the E-tag of the cached resource
the server then checks if the E-tag matches the current version of the resource, if it does, it returns a 304 Not Modified response , else normal 200 OK with the new resource
similarly there is a tag called if-modified-since, which is used to check if the resource has been modified since the last request


Content-Negotiation is a mechanism that allows the client and server to agree on the format of the response, based on the capabilities of the client and server
for example, the client can specify the Accept header to indicate the format it prefers (like JSON, XML, HTML, etc.)
the server would try to respond with the format that the client prefers, if it supports it, otherwise it would respond with the default format (fallback format)
there are other important headers in negotiation like encoding headers for example gzip, deflate, br, etc. which are used to compress the response body to reduce the size of the response
Accept-language header is used to specify the preferred language of the response, so that the server can respond with the appropriate language
this is how we can have a single endpoint that can respond with different formats and languages, based on the capabilities of the client and server

http based compression also lies under this umbrella, which is used to reduce the size of the response body, so that it can be transmitted faster over the network
and the client can decompress it to get the original response body
this is done using the Content-Encoding header, which specifies the encoding used to compress the response body
it goes from 24mb to 4mb, which is a significant reduction in size // try doing it once


persistent connections are used to keep the connection open between the client and server, so that multiple requests can be sent over the same connection
this reduces the overhead of establishing a new connection for each request, which can be expensive in terms of time and resources
keep-alive is a header that is used to specify that the connection should be kept open for a certain period of time, so that multiple requests can be sent over the same connection
this is done using the Connection header, which specifies the connection options for the request
we don't have to do anything special to implement this, it is done by default in most web servers and clients, and we don't work with it much as default values usually work fine

Handling large requests and responses is another important aspect of HTTP, especially when dealing with file uploads or downloads
1 . Multipart requests are used to send large files in chunks, so that the server can process them without running out of memory
    Chunked transfer encoding is used to send large responses in chunks, so that the client can start processing the response before the entire response is received
there is a delimiter/boundary like the hash tamrakar added , signifying the end of each chunk, so that the client can know when to stop reading the response
2 .Streaming responses are used to send large responses in a continuous stream, so that the client can start processing the response as soon as it recives the first chunk
  content-type is text/event-stream for this, which is used to send server-sent events (SSE) to the client and not simple text/plain
 connection : keep-alive is used to keep the connection open for a certain period of time, so that multiple chunks can be sent over the same connection

reads about ssl(outdatesd) , tls(newer and more secure) {these are what make http https}

*/
