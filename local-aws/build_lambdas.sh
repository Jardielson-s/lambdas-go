# #!/bin/bash

while IFS= read -r lambda
do
  echo "lambda: $lambda"

  mkdir -p "$lambda"
  cd "$lambda"

  if [ ! -f "../../src/infra/functions/${lambda}/${lambda}.go" ]; then
    echo "Error: Go source file for ${lambda} not found!"
    exit 1
  fi

  GOOS=linux CGO_ENABLED=0 go build -ldflags '-s -w' -o bootstrap "../../src/infra/functions/${lambda}/${lambda}.go"

  if [ ! -f bootstrap ]; then
    echo "Build failed for ${lambda}. Exiting..."
    exit 1
  fi

  echo "Successfully built ${lambda} binary!"
  cd ..
done < "lambdas.txt"
echo "All Lambda binaries have been successfully built."
