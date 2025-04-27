<h1>Deploy main.go with the API</h1>
  First we create application. For example, deploy the John Travolta application with the Rawdat.AI API.
 
 ``curl -ivX POST --data-binary @main.go -H 'API-KEY: '"{MY-API-KEY}"'' "https://rawdat.ai/api/v1/deploy-function?uri=/*&language=go&protocol=https&hostname=third.run.place"``

  - The payload is the source code (@main.go)
  - Set the HTTP API-KEY header to your assigned secret key
  - Query parameter uri is a wildcard /*
  - Query parameter protocol = https
  - Query parameter hostname = juicebox.publicvm.com;

`` Function is succesfully deployed. Entry point is 11.11.XX.XX.`` You can note down your VIP (Virtual IP) for future reference

Connect to the application (https://juicebox.publicvm.com)

