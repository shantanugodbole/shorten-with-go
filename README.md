# URL Shortener Microservice made with Go!

The main objective of this project was to mainly learn Go by implementing and understading why it is such a widely used language at the backend for large and complex systems.

## Project Overview

The project albeit small has 3 main parts to it:

1. Store:
    This directory makes use of Redis - **RE**mote **D**ictionary **S**erver. Although Redis is a fairly well-established SaaS, this is my first time implementing it. The Redis instance is hosted in the Redis cloud and connected to this service via the redis-go library. 
2. Shortener:
    The actual underlying mechanism of shortening the URL is based on cryptography taught in undergrad CS classes and utilizes SHA256 inbuilt library present in Go. To fix the length of the shortened URL, we only slice out the first 8 characters. We also hash it using a unique User ID which is set manually for now. This ensures that no users can create a shortened link which maps to a different URL
3. Handler:
    The handler serves as the main part of the service which makes use of the store and shortener code to save and create shortened URLs


Additonally, testing scripts for both store and shortener module were written to make sure that the code performs as expected. Leveraged testing library to carry out this task.

