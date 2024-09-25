# Cobra Utils

[spf13/cobra](https://github.com/spf13/cobra)に関する小さい関数をまとめてあります。

- `reflect`を用いて`cobra.Command`のプライベートなフィールドにアクセスします。
- コマンドのテスト可用性と可読性を両立するサブコマンド構成を実現するための関数を提供します。
- `PersistentPreRunE`を再帰的に実行するユーティリティを提供します。
