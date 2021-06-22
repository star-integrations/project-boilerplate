actual=`wc -c back/web/docs/swagger.json | awk '{print $1}'`
make swag
expect=`wc -c back/web/docs/swagger.json | awk '{print $1}'`
if [ $actual -ne $expect ]
then
  echo "生成コードが一致しません。make swagを実行してください。"
  exit 1
fi
