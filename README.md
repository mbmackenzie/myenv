# myenv

Configure and manage environment variables and files that are commonly shared across projects.

### Install

TODO

### Usage

#### Configure variables

```
myenv create var DB_USERNAME mbm2228 
myenv create var APIKEY abcdefg1234567

myenv create ref SOME_SCRIPT ~/files/somescript.sh
myenv create ref CLIENT_DRIVER ~/files/CLIENT.jar
```

#### In a project

```
myenv add var DB_USERNAME 
myenv add var APIKEY --rename APP_APIKEY

myenv add ref SOME_SCRIPT
myenv add ref CLIENT_DRIVER --copy-over
```

This creates a file (or updates it) `.env` that looks like:

```
DB_USERNAME=mbm2228
APP_APIKEY=abcdefg1234567

SOME_SCRIPT=~/files/somescript.sh
CLIENT_DRIVER=CLIENT.jar
```

#### Create common configs

```
myenv config create DS_PROJECT
myenv config add DS_PROJECT DB_USERNAME
myenv config add DS_PROJECT APIKEY

myenv config add DS_PROJECT
```

