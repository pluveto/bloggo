export_filename=`pwd`/install/schema.sql
sqlite3 tmp.db .schema > $export_filename
echo "exported to" $export_filename