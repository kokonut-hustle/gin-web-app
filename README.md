## What is Gin?

Gin is a high-performance micro-framework that can be used to build web applications and microservices. It makes it simple to build a requires handling pipeline from modular, reusable pieces. It does this by allowing you to write middleware that can be plugged into one or more request handlers or groups of request handlers.

## Why Gin?

One of the best features of Go is that it's built-in net/http library that allows you to create an HTTP server effortlessly. However, it is also less flexible and requires some boilerplate code to implement.

There is no built-in support in Go to handle routes based on a regular expression or a "pattern". You can write code to add this functionality. However, as the number of your applications grows, it is quite likely that you will either repeat such code everywhere or create a library for reuse.

This is the crux of what Gin offers. It contains a set of commonly used functionalities, e.g. routing, middleware support, rendering, that reduce boilerplate code and make writing web applications simpler.

## Design application functionality

The application we'll build is a simle article manager. It will be able to show articles in HTML, JSON and XML as needed. This will allow us to illustrate how Gin can be used to design traditional web applications, API servers, and microservices.

To achieve this, we will make use of the following functionalities offered by Gin:
- Routing - to handle various URLs,
- Custom rendering - to handle the response format, and
- Middleware.

### Routing

Routing is one of the core features that all modern frameworks provide. Any web page or an API endpoint is accessed by a URL. Frameworks use routes to handle requests to these URLs. If a URL is http://www.example.com/some/random/route, the route will be /some/random/route.

Gin offers a fast router that’s easy to configure and use. Apart from handling specified URLs, Gin routers can also handle patterns and grouped URLs.

In our application, we will:

Serve the index page at route / (HTTP GET request),
Group article related routes under the /article route. Serve the article page at /article/view/:article_id (HTTP GET request). Take note of the :article_id part in this route. The : at the beginning indicates that this is a dynamic route. This means that :article_id can contain any value and Gin will make this value available in the route handler.

### Rendering

A web application can render a response in various formats like HTML, text, JSON, XML or other formats. API endpoints and microservices typically respond with data, commonly in JSON format but also in any other desired format.

In the next section, we’ll see how we can render different types of responses without duplicating any functionality. We will primarily respond to a request with an HTML template. However, we will also define two endpoints that can respond with JSON or XML data.

### Middleware

In the context of a Go web application, middleware is a piece of code that can be executed at any stage while handling an HTTP request. It is typically used to encapsulate common functionality that you want to apply to multiple routes. We can use middleware before and/or after an HTTP request is handled. Some common uses of middleware include authorization, validation, etc.

If middleware is used before a request is handled, any changes it makes to the request will be available in the main route handler. This is handy if we want to implement some validations on certain requests. On the other hand, if the middleware is used after the route handler, it will have a response from the route handler. This can be used to modify the response from the route handler.

Gin allows us to write middleware that implements some common functionality that needs to be shared while handling multiple routes. This keeps the codebase small, separates concerns and improves code maintainability.

We want to ensure that some pages and actions, eg. creating an article, logging out, are available only to users who are logged in. We also want to ensure that some pages and actions, eg. registering, logging in, are available only to users who aren’t logged in.

If we were to put this logic in every route, it would be quite tedious, repetitive and error-prone. Luckily, we can create middleware for each of these tasks and reuse them in specific routes.

We will also create middleware that will be applied to all routes. This middleware (setUserStatus) will check whether a request is from an authenticated user or not. It will then set a flag that can be used in templates to modify the visibility of some of the menu links based on this flag.
