# hookah

A simple build automation pipeline

## How it works

- Receive a repo Git URL to clone
- Run in a docker container
- > Clone the `repo`
- > Execute all the *steps* (setup, test, build) in the `makefile`
- > Build the `dockerfile` and tag the image as `<Repo's name>:<commit hash>` and `<Repo's name>:latest`
- > Push to `registry`
- Remove the docker container
- Profit...
