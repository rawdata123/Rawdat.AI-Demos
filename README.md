<h1>Composing applications in the sky with Rawdat.AI</h1>
  <p>First create the application. You can chose any source file from the repositiory. For example, deploy the John Travolta application (John_Travolta.go).<br></p>  
  

  ``curl -ivX POST --data-binary @John_Travolta.go -H 'API-KEY: '"{MY-API-KEY}"'' "https://rawdat.ai/api/v1/deploy-function?uri=/*&language=go&protocol=https&hostname=juicebox.publicvm.com"``

   
  ``Function is succesfully deployed. Entry point is 11.11.XX.XX.``

  - The payload is the source file (John_Travolta.go)
  - Set the API-KEY header to your assigned secret key for authentication (Reach out to support@rawdat.ai for an API-Key)
  - Query parameter <strong>language</strong> is go since we are deploying a go application
  - Query parameter <strong>uri</strong> is a wildcard = /*
  - Query parameter <strong>protocol</strong> = https
  - Query parameter <strong>hostname</strong> = juicebox.publicvm.com.  

Connect to the application (https://juicebox.publicvm.com)

<strong>juicebox.publicvm.com is just an example. You should use your own hostname when composing the app.</strong> 

<h1>Pushing changes to your created applications</h1>
  To push new changes to your applications in real time, commit your changes and push to the repository.

  ```
  git add John_Travolta.go 
  git commit -m "New update to John_Travolta app"
  git push
  ```

Connect to the application and review the change (https://juicebox.publicvm.com)
