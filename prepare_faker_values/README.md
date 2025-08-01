# Support memtier_benchmark use prepared faker values

## Guide

### Step1 Prepare fake values

```bash
cd prepare_fake_values
go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct
go build
./prepare_faker_values -count 1000000 -value-length 100
```

### Step2 Run Redis / Valkey / Engula server

```bash
nohup taskset -c 0-3 ./valkey-server --save "" --appendonly no > output.log 2>&1 &
```

### Step3 Build memtier_benchmark and run

```bash
taskset -c 4-7 ./memtier_benchmark -s localhost -p 6379 -d 256 --faker-text-data --test-time=120 --clients=1 --threads=1 --pipeline=1 --show-config --hide-histogram --key-pattern=S:S --key-minimum=1 --key-maximum=10000000
```