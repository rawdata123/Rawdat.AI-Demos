<h1>Deploying the application with the API</h1>
  First create the application. For example, deploy the John Travolta application (John_Travolta.go).
  ``  ``
  
  ``curl -ivX POST --data-binary @main.go -H 'API-KEY: '"{MY-API-KEY}"'' "https://rawdat.ai/api/v1/deploy-function?uri=/*&language=go&protocol=https&hostname=third.run.place"``

  - The payload is the source file of the application we are deploying (John_Travolta.go)
  - Set the HTTP API-KEY header to your assigned secret key for user authentication (Reach out to support@rawdat.ai for an API-Key)
  - Query parameter language is Go since the app is written in go
  - Query parameter uri is a wildcard = /*
  - Query parameter protocol = https
  - Query parameter hostname = juicebox.publicvm.com;

`` Function is succesfully deployed. Entry point is 11.11.XX.XX.`` You can note down your VIP (Virtual IP) for future reference

Connect to the application (https://juicebox.publicvm.com)


<h1>Pushing changes to your created applications</h1>
  To push new changes to your applications in real time, commit your changes and push to the repository.


  ``git add John_Travolta.go ``

  ``git commit -m "New update to John_Travolta app" ``

  ``git push``

  Connect to the application and review the change (https://juicebox.publicvm.com)
