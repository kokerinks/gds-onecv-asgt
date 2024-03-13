# gds-onecv-asgt
 
## Link to the hosted API

The server is currently hosted on EC2 Free Tier, under the link [http://13.236.153.51:3000/](http://13.236.153.51:3000/)

E.g to run `GET /api/commonstudents?teacher=<teacher_email>`, use the url:
`http://13.236.153.51:3000/api/commonstudents?teacher=<teacher_email>`

## Instructions for running locally

1. Git clone the repository
2. Start up docker service
3. Rename .env.example to .env, and set a password of your choice
4. On repo directory, run `docker-compose up --build`
5. The server should now be running on `http://localhost:3000`!

To run unit tests:
- Run `go test` on repo directory
  - This should run `unit_test.go`, which contain all unit tests required 

## File Structure

**/controllers:** Handles HTTP Requests and formulate responses.

**/models:** Defines database entities.

**/routes:** Define API routes and link them to controllers.

**/testData:** Contains scripts for testing/seeding data

**/utils:** Provides scripts for utilities/services e.g. database

**Structure was adapted from past experiences working on Express backend applications