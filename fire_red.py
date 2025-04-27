def application(env, start_response):
    start_response('200 OK', [('Content-Type', 'text/html')])
    
    html_content = """
    <!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Fire Red Application</title>
        <style>
            body {
                background-color: #ff0000;
                color: white;
                font-family: Arial, sans-serif;
                text-align: center;
                padding: 50px;
            }
            .container {
                background: rgba(255, 255, 255, 0.1);
                padding: 20px;
                border-radius: 10px;
                display: inline-block;
            }
            h1 {
                font-size: 2.5em;
                text-transform: uppercase;
            }
            .btn {
                background: white;
               color: #ff0000;
                padding: 10px 20px;
                font-size: 1.2em;
                border: none;
                border-radius: 5px;
                cursor: pointer;
                transition: 0.3s;
            }
            .btn:hover {
                background: black;
                color: white;
            }
        </style>
    </head>
    <body>
        <div class="container">
            <h1>Welcome to Fire Red App</h1>
            <p>Your ultimate red-hot experience starts here.</p>
            <button class="btn">Get Started</button>
        </div>
    </body>
    </html>
    """    
    return [html_content.encode('utf-8')]

