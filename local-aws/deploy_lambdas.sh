# #!/bin/bash

./build_lambdas.sh

terraform apply --var-file="terraform.tfvars" -auto-approve