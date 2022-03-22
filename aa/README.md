# go-revel-rest
Example of REST API to create and retrieve users with Go and Revel framework

# Install dependencies

This project uses Revel framework. You can install it by:

<code>go get github.com/revel/revel</code>

<code>go get github.com/revel/cmd/revel</code>

# Run

Once Revel framework is installed, you can run the server by:

<code>revel run go-revel-rest</code>

Note that the project must be located under <code>$GOPATH/src/go-revel-rest</code>

# Routes

The API routes are defined in <code>conf/routes</code> file:

<code>GET /users</code> Retrieve all the users

<code>GET /users/{id}</code> Retrieve user with id {id}

<code>PUT /users/</code> Save a new user. <code>id</code> and <code>nickname</code> must be provided as form params.
