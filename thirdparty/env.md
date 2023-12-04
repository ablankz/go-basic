## 環境変数呼び出しに関するサードパティライブラリ

- github.com/joho/godotenv
    ``` .env
    ENV=debug
    ```
    ``` main.go
    package main

    import (
        "os"
        "fmt"
        "log"
        dotenv "github.com/joho/godotenv"
    )

    func main() {
        if err := dotenv.Load(); err != nil {
            log.Fatalln(err)
        }
        env := os.Getenv("ENV")
        fmt.Println(env)
    }
    ```
- gopkg.in/ini.v1
    ``` config.ini
    [dev]
    ENV=debug
    ```
    ``` main.go
    package main

    import (
            "fmt"
            "log"
            "gopkg.in/ini.v1"
    )

    // ConfigList is ...
    type ConfigList struct {
            ENV        string
    }

    func main() {
            config, err := ini.Load("config.ini")
            if err != nil {
                    log.Fatalf("config.iniの読み込み失敗：%v", err)
            }
            cfg = ConfigList{
                    ENV:        config.Section("dev").Key("ENV").MustString(""),
            }
            fmt.Println(cfg.ENV)
    }
    ```
- github.com/caarlos0/env
    ``` bash
    $ echo $ENV
    debug
    ```
    ``` main.go
    package main

    import (
        "log"
        "fmt"
        "github.com/caarlos0/env"
    )

    type ServerConfig struct {
        Env     string `env:"ENV,required"`
    }

    func main() {
        var config ServerConfig
        if err := env.Parse(&config); err != nil {
            log.Fatalln(err)
        }
        fmt.Println(config.Env)
    }
    ```