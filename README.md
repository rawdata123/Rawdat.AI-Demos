<h1>Try deploying main.go with the API</h1>
  curl -ivX POST --data-binary @main.go -H 'API-KEY: '"{MY-API-KEY}"'' "https://rawdat.ai/api/v1/deploy-function?uri=/*&language=go&protocol=https&hostname=third.run.place" &

  - The payload is the source code (@main.go)
  - Set the API-KEY header to your assigned secret key
  - Query parameter uri = /*
  - Query parameter protocol = https
  - Query parameter hostname = juicebox.publicvm.com

Connect to the application (https://third.run.place)

