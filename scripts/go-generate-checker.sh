make go_generate
clean=$(git status | grep "nothing to commit" || true)
if [ -z "$clean" ]; then
  git status
  echo "生成コードが一致しません。make go_generateを実行してください。"
  exit 1
fi
