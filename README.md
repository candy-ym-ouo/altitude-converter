# Altitude Converter

一个用于户外记录整理的 Go 命令行海拔换算器，支持米（`m`）、英尺（`ft`）、千米（`km`）和海里（`nm`）。可处理单笔终端输入、文本批量输入和 CSV 批量输入，并将结果按指定字段导出为 CSV。

## Requirements and dependencies

- Go 1.22 或更高版本
- 仅使用 Go 标准库，无需安装第三方依赖

## Build and run

```sh
go build -o altitude-converter ./cmd/altitude-converter
./altitude-converter -value 8848.86 -from m -to ft
```

也可直接运行：

```sh
go run ./cmd/altitude-converter -value 10 -from km -to nm
```

## Batch input

文本文件每行一个记录，格式为 `数值 单位`；空行和 `#` 开头的注释会忽略：

```text
8848.86 m
29029 ft
10 km
```

```sh
go run ./cmd/altitude-converter -input hikes.txt -to ft
```

CSV 必须有表头，数值列使用 `value` 或 `altitude`，单位列使用 `unit` 或 `from`：

```csv
value,unit,name
8848.86,m,Everest
3776,ft,Fuji
```

```sh
go run ./cmd/altitude-converter -input hikes.csv -to m -output converted.csv -fields value,from,result,to
```

`-fields` 可选字段为 `value`、`from`、`to`、`result`，以逗号分隔。单位也接受常见英文全称，例如 `meters`、`feet`、`kilometers` 和 `nmi`。

程序会拒绝缺失字段、未知单位、非数字、NaN、无穷大、结构错误的文本行和不完整 CSV 行，并以非零退出码结束。

## Test

```sh
go test ./...
```
