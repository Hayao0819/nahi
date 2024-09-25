# exutils

os/execに関するユーティリティ

- `CommandExists`はコマンドがパスの中に存在するか確認します
- `CommandWithStdio`は標準入出力をそれぞれ`/dev/stdin`, `/dev/stdout`, `/dev/stderr`に設定した`exec.Command`を作成します
- `EvalSh`は文字列をシェルコマンドとして解釈して実行します
- `EvalString`はシェルスクリプトによって解釈可能な文字列として解析します
