#!/bin/bash

# echo "Compiling"
# mvn compile

echo "Packaging..."
if mvn package; then
  echo "Running program..."
  echo "------------------------------------------------"
  if [[ $(uname) == "Darwin" ]]; then
    java -jar -XstartOnFirstThread target/mango-0.1.0.jar -Djava.library.path=target/natives-macos
  else
    java -jar target/mango-0.1.0.jar -Djava.library.path=target/natives-linux
  fi
  java -jar target/mango-0.1.0.jar com.micahelias.App
else
  echo "Packaging failed, program will not be run"
fi


