# hookah

A simple build automation system

It runs on port 8080. (should become a parameter in the future)

## How it works

Start the buildspace _docker container_ from the `felixfx/buildspace:alpine` image. The buildspace should become a dynamic in the future.

- Run the `pullHandler` script with the `pullParams` as arguments
- Moves into the repo. Then moves into a `buildPath`.
- Run the `hookah` command of the project's **makefile** (make hookah)
- Run the `pushHander` script with the `pushParams` as arguments

All the logs are saved in the build history of each project.

The buildspace's container is then removed.

## Request a build

Make a `HTTP POST` at `/build` with a JSON object in the body. 

The response is the count of currently pending requests (including the one that has just been sent).

Build request's body example:
```
{
  "buildPath": "./pass",
  "projectName": "ProjectA",
  "buildName":  "Some build name",
  "pullHandler": "demo",
  "pullParams": "",
  "pushHandler": "docker",
  "pushParams": "testing/ProjectA:latest"
}
```

You can test the endpoint with script `./test_req.sh "projectA" "pass" "Some build name"`. Asuming Hookah is listening on `127.0.0.1:8080`.

If Hookah receive a build request while busy, the request will simply be queued to be built after the currently running build and all the previously queued requests.

Support for paralle builds could be added eventualy.

## UI

The list of all project is accessible at `/`.

The list of all builds in a projects is accessible at `/project/<projectName>`.

You can click on each builds _"lines"_ to show more details.

No javascript required.

## API

Same as the UI, but you get JSON. Simply set the request header `Accept` to `application/json`.


## FAQ 

Hahaha... Nobody has asked me anything yet. But prevention beats reaction.

### How can I request a build on each Github push?
You need something that convert a Github Webhook request to a _build request_. A script doing _exactly this_ should be available shortly.

### Can I make my builds on somethings else than Alpine Linux?
The buildspace image will become dynamic (_parameterable_) eventually.
I have no plan to support anything that **can't run** inside a Docker container.