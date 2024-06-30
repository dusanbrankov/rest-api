#!/usr/bin/env bash

# INFO:
# This script creates a new MySQL database and user

tmpfile="$(mktemp -t tmp.XXXXXXXXXX)"
cleanup() { rm -f "$tmpfile"; }
trap cleanup EXIT

err() {
	echo "error: $*" >&2
	exit 1
}

# Ensure the script is run as root
if (( EUID != 0 )); then
	err "this script must be run as root"
fi

db_name="${1:-restapi}"
db_pass="$(tr -dc 'A-Za-z0-9!@#%^&*()' </dev/urandom | head -c 32)"

sql_script="init-db.sql"
if [ ! -f "$sql_script" ]; then
	err "$sql_script does not exist"
fi

sed "s/DB_NAME/$db_name/ ; s/DB_PASSWORD/$db_pass/" "$sql_script" > "$tmpfile"

echo "creating database and user..."
if [ ! -f "$HOME/.my.cnf" ]; then
	err "$HOME/.my.cnf does not exist"
fi
mysql < "$tmpfile"

echo "adding database credentials to .env file..."
echo -e "DB_NAME=$db_name\nDB_USER=${db_name}_user\nDB_PASSWORD=$db_pass" >> .env
chown "$SUDO_USER": .env
chmod 600 .env

echo "done"

