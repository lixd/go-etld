# effective Top-Level Domain Parser

effective top level domain parser. parse the subdomain, domain, and eTLD from a host string. used go embed,so require go 1.16+

> 一个简单的有效顶级域名解析器，从 host 中解析子域名和有效顶级域名。

主要根据 [publicsuffix](https://publicsuffix.org/list/effective_tld_names.dat) 提供的 eTLD 列表进行解析。



## Glossary

| 缩写  | 全称                        | 含义                    |
| ----- | ------------------------------ | ----------------------- |
| TLD   | Top-Level Domain               | 顶级域名                |
| gTLD  | Generic top-level domain       | 通用顶级域名            |
| nTLD  | National Top-Level Domain      | 国家顶级域名            |
| ccTLD | Country Code Top Level Domain  | 国家顶级域名,nTLD 的别名 |
| iTLD  | International Top Level Domain | 国际顶级域名,比较少见   |
| eTLD | effective Top-Level-Domain | 有效顶级域名 |


## Usage

```go
info := etld.Parse("www.github.com")
// info.Subdomain = "www"
// info.Domain = "github"
// info.Suffix = "com"
```

> 编译前，建议通过 make renew 更新 eTLD 列表。
