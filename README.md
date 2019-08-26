# vending-machine-cli

console program that imitate the Vending Machine

### how to play
clone this project/repository
```bash
git clone https://github.com/jojoarianto/vending-machine-cli.git
```
run build to download and build binari
```bash
make build
```

run app
```bash
./vending-machine
```

### 

### project structure

```
    cmd/
        app.go
        router.go
    constant
        errorcode.go
    models/
        coin.go
        item.go
        storage.go
    service/
        vending.go
        vending_test.go
    utils/
        coin.go
        coin_test.go
        msgbuilder.go
        msgbuilder_test.go
    main.go
    go.mod
    Makefile
```

### development mode

to run dev mode run
```bash
make run
```

to run test 
```bash
make test
```

see another make commmand on makefile
