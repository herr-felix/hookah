git version

build() {
  PROJECT_DIR=$( /bin/sh /opt/handlers/pull/$PULL_HANDLER.sh | tail -1 )
  echo "Moving to $PROJECT_DIR"
  cd $PROJECT_DIR

  echo "Moving to $BUILD_PATH"
  cd $BUILD_PATH
  if [ "$?" -ne "0" ]; then
    echo "Moving to '$BUILD_PATH' failed"
    return 1
  fi

  make all
  if [ "$?" -ne "0" ]; then
    echo "Build fail"
    return 1
  fi

	docker build . -t demo_go # REPLACE BY THE PUSH HANDLER
  if [ "$?" -ne "0" ]; then
    echo "Docker build failed"
    return 1
  fi
}

build demo 2>&1 # TODO putting 'ts' here without affecting the status code
if [ "$?" -ne "0" ]; then
  exit 1
fi
