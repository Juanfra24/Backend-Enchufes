# GoTest

Para ejecutar:


Install the Heroku CLI

Download and install the Heroku CLI.

If you haven't already, log in to your Heroku account and follow the prompts to create a new SSH public key.

$ heroku login

Clone the repository

Use Git to clone test-enchufe's source code to your local machine.

$ heroku git:clone -a test-enchufe$ cd test-enchufe


Para ejecutar cambios:

Deploy your changes

Make some changes to the code you just cloned and deploy them to Heroku using Git.

$ git add .
$ git commit -am "make it better"
$ git push heroku master


