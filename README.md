# Backend Exercise

**Please limit your time working on this exercise to four hours.**

Please create a Golang, Java, or Node.js application implementing a CRUD web service for `posts`. 

- The API should support `CREATE`, `READ`, `UPDATE`, `DELETE` of the `posts` resource with validation.
- Support fetching multiple posts, a single post, and filtering posts by user.
- Use an in-memory database or native data structures of your choice for persistence. 
- Cover all the important behavior of the code with automated tests. Feel free to choose between integration tests and unit tests.

#### Additional topics

If you have extra time, we'd love to see what you create around any of the topics below:

- Authentication (Basic, OAuth2, JWT)
- Use a persistent DB (postgres, mongo)
- Deploy the application (Heroku, GCP, AWS)

### Posts

``` javascript
[
  {
    "userId": 1,
    "id": 1,
    "title": "Node is awesome",
    "body": "Node.js is a JavaScript runtime built on Chrome's V8 JavaScript engine."
  },
  {
    "userId": 1,
    "id": 2,
    "title": "Spring Boot is cooler",
    "body": "Spring Boot makes it easy to create stand-alone, production-grade Spring based Applications that you can "just run"."
  },
  {
    "userId": 2,
    "id": 3,
    "title": "Go is faster",
    "body": "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software."
  },
  {
    "userId": 3,
    "id": 4,
    "title": "'What about me?' -Rails",
    "body": "Ruby on Rails makes it much easier and more fun. It includes everything you need to build fantastic applications, and you can learn it with the support of our large, friendly community."
  }
]
```

### Users

```javascript
[
  {
    "id": 1,
    "name": "Ryan Dahl",
    "email": "node4lyfe@example.com",
    "expertise": "Node"
  },
  {
    "id": 2,
    "name": "Rob Pike",
    "email": "gofarther@example.com",
    "expertise": "Go"
  },
  {
    "id": 3,
    "name": "DHH",
    "email": "magic@example.com",
    "expertise": "Rails"
  }
]
```

#TODO:
This is just a list of things that still need to be implemented for this project which I intended to complete but ran out of time.
- Unit Testing to achieve at least 75% coverage
- Implementation of the dynamo.go file which should be the data access layer to the DynamoDB storage
- Structured logging for easier debugging
- Buffered channel implementation for multiple entry to avoid overloading dependent resources
- Dockerfile for composing the image and container to hold this service
- CloudFormation to spin up AWS resources: Dynamo, EKS assets, etc.
- creation of more error primitives
- console flags/environment variables for configuring dependencies and just general setup for the service
- GoReport Card
- Deployment onto AWS
- Docker Compose for local testing of Dynamo and also dependency management
- Swagger API Documentation
- Go Modules dependencies handling(forgot to go mod init before shipping tsk tsk tsk)
- Error handling
- Correlation/request ID with context
- Smoke, Integration, Load, and Performance Tests
- General documentation
- Go Linter
- Analysis of package structure layout