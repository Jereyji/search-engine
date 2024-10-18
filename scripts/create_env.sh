  cat <<EOF > deployments/.env
DB_USERNAME=postgres
DB_PASSWORD=1234
DB_HOST=localhost
DB_PORT=5436
DB_NAME=search_engine
DB_SSLMODE=disable
EOF

echo "env file created"