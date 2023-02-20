#!/bin/bash

# echo "Compiling"
# mvn compile

echo "Packaging..."
if mvn package; then
  echo "Running program..."
  echo "------------------------------------------------"
  java -jar target/mango-0.1.0.jar com.micahelias.App
else
  echo "Packaging failed, program will not be run"
fi

