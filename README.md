# Tesla Automated Sentry

This tool auto-enables Sentry-mode on your Tesla when it detects your Tesla is connected to a charger.

## How does it work
It polls the Tesla API to check if your Tesla is awake or sleeping, if sleeping it doesn't wake the car. If it's online and charging (or done charging) it will verify if Sentry mode is enabled, if not it will turn it on.

## Setup
Running this tool is super easy. Just make sure you set 2 environment variables containing your Tesla Credentials : 

1. TESLA_SENTRY_EMAIL
2. TESLA_SENTRY_PASSWORD

```bash
For Linux : 
export TESLA_SENTRY_EMAIL="elon@tesla.com" && export TESLA_SENTRY_PASSWORD="roadster"
```

Run the tool by excecuting the .exe (Windows) or Binary for Linux every 5 minutes.

```bash
For Linux Cronjob : 
*/5 * * * * /home/<username>/teslaautomatedsentry
```

If you want to use the daemon mode of the tool, which runs the Tesla Automated Sentry part every 2 minutes, pass the --daemon true flag : 

```bash
For Linux : 
./teslaautomatedsentry --daemon true
For Windows:
teslaautomatedsentry.exe --daemon true
```

## Docker
You can also run this tool in a simple Docker container.
Just run this command : 
```
docker run -ti -d tesla-automated-sentry --env TESLA_SENTRY_EMAIL=<emailaddress> --env TESLA_SENTRY_PASSWORD=<password>
```
This will start the container in daemon mode!

## Copyright & License

Copyright (c) 2020 Dennis Brouwer. Released under the terms of the MIT license. See LICENSE for details.