# Web

## Why Web?
- Web is used by millions and billions of people
- Web can be used to study, watch videos, or listen to music
- Web is NOT internet

## Internet
- Internet is a communication mechanism of computers
- Web runs on top of Internet

## Web
- Web is a group of information available inside a computer that is connected continuously via internet
- Web is filled with any kinds of information, like text, images, audio, video, etc
- Web runs in an application called Web Server, Web server is an app that is used to store and send informations inside a Web
- Computer -> Via Internet -> Computer(WebServer(Web))
- For example, consume computer connects via internet to YouTube's computer to watch videos

## Web Host
- Web owners usually does not run a Web Server Application in their own computer
- They rent a computer in a data center, owned by a Web Server Provider or Cloud Provider
- Web Server Providers is called Web Host
- For example -> Heroku, Hostinger, Netlify, Vercel

## Domain
- When Web is connected to the internet, they have an address in the form of IP Address
- Because IP Address is hard to memorize, use Domain Name
- For example -> Google.com, YouTube.com, Facebook.com
- Domain Name System -> DNS Server

## Web Browser
- Web Browser is an Application to access Web via the Internet
- We can technically access Web without Web Browser, but the Web Server would send raw machine language
- Web Browser will transform the machine language (HTML, CSS, JS) into something a human can understand

## Client and Server
- Web is an application based on Client and Server
- Client Server is an architecture concept that connects two party, Client System, and Client Server
- They communicate via the internet (if hosted) or the same computer with localhost
- Client -> Asks data (Request) -> Server
- Client <- Returns data (Respond) <- Server
- Benefit
  - Application changes can be done inside the server, without making changes to the client
  - Can be used by many client at the same time, even though there is only 1 server
  - Can be accessed anywhere as long as its connected to the server
- Example:
  - Web is Client and Server
  - Client -> Web Browser (Chrome, Firefox, Opera, Edge, etc)
  - Server -> Web Server (The Code for the Web Program)

## GoLang Web
- Go is popular to create Web, especially Web API (Backend)
- Go also has its own package to create Web, along with a package to implement Unit Testing for Web
- This makes making Web using Go easier, without any additional Library or Framework
- How it works
  - Web Browser perform an HTTP Request to the Web Server
  - GoLang receives the HTTP Request, and executes the request accordingly
  - After Execution is done, GoLang returns the data and renders HTML, CSS, JavaScript as requested
  - GoLang will return the rendered content as an HTTP Response to the Web Browser
  - Web Browser receives the content from the Web Server and renders the content according to the type
  - Web Browser -> Request -> Golang -> Execute Web Server / Render
  - Web Browser <- Render <- Response <- Golang

## Package net/http
- Go has a package called net/http to create a Web
- But if we are building Web application at scale, it is still recommended to use a Framework to make it easier
- Every Web Framework for GoLang is build on top or based on net/http
