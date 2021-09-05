
input="$1"
#echo "package main\n\nimport (\n        \"database/sql\"\n )\n\n func sendSQL() int {\n         var id int\n         var db *sql.DB\n         db.QueryRow(\`\n $query\n \`).Scan(&id)\n        return id\n }"|sqlformat|tail +11|head -n -2|sed 's/`).Scan(&id)//g'
#echo "package main\n\nimport (\n        \"database/sql\"\n )\n\n func sendSQL() int {\n         var id int\n         var db *sql.DB\n         db.QueryRow(\`\n $query\n a\`).Scan(&id)\n        return id\n }"|sqlformat |tail +11|head -n -2|sed 's/a`).Scan(&id)//g'|sed '$s/).Scan(&id)//g'|awk 'NF' |sed '$s/a`,//g'

echo "$input"|./sqlformatter/bin/sql-formatter
	

#echo "package main\n\nimport (\n        \"database/sql\"\n )\n\n func sendSQL() int {\n         var id int\n         var db *sql.DB\n         db.QueryRow(\`\n $query\n \`).Scan(&id)\n        return id\n }"|sqlformat|tail +11|head -n -2|sed 's/`).Scan(&id)//g'

#echo "package main\n\nimport (\n        \"database/sql\"\n )\n\n func sendSQL() int {\n         var id int\n         var db *sql.DB\n         db.QueryRow(\`\n $query\n a\`).Scan(&id)\n        return id\n }"|sqlformat |tail +11|head -n -2|sed 's/a`).Scan(&id)//g'|sed '$s/).Scan(&id)//g'|awk 'NF' |sed '$s/a`,//g'


#echo "$input"|./sqlformatter/bin/sql-formatter

