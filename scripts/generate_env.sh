#!/bin/bash

# generate_password() {
#   openssl rand -base64 12
# }

create_env_file() {
#   PASSWORD=$(generate_password)
  PASSWORD=1234

  cat <<EOF > configs/.env
DB_USERNAME=postgres
DB_PASSWORD=$PASSWORD
DB_HOST=localhost
DB_PORT=5436
DB_NAME=postgres
DB_SSLMODE=disable
EOF

  echo "env file created"
}

create_env_file
