  cat <<EOF > deployments/.env
DB_USERNAME=postgres
DB_PASSWORD=1234
DB_HOST=postgres
DB_PORT=5432
DB_NAME=search_engine_db
DB_SSLMODE=disable
EOF

echo "env file created"