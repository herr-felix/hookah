build() {
  PROJECT_DIR=$( /bin/sh /opt/handlers/pull/$PULL_HANDLER.sh $PULL_PARAMS | tail -1 )
  echo "MOVING to $PROJECT_DIR"
  cd $PROJECT_DIR
  if [ "$?" -ne "0" ]; then
    echo "CAN'T MOVE TO PROJECT '$PROJECT_DIR'"
    return 1
  fi

  echo "MOVING TO $BUILD_PATH"
  cd $BUILD_PATH
  if [ "$?" -ne "0" ]; then
    echo "CAN'T MOVE TO BUILD PATH '$BUILD_PATH' INSIDE PROJECT '$PROJECT_DIR'"
    return 1
  fi

  make hookah
  if [ "$?" -ne "0" ]; then
    echo "BUILD FAILED"
    return 1
  fi

  # If there is a push handler
  if [ $PUSH_HANDLER != "" ]; then
    /bin/sh "/opt/handlers/push/$PUSH_HANDLER.sh" $PUSH_PARAMS
    if [ "$?" -ne "0" ]; then
      echo "PUSH FAILED"
      return 1
    fi
  fi
}

build 2>&1 # TODO putting 'ts' here without affecting the status code
if [ "$?" -ne "0" ]; then
  exit 1
fi
