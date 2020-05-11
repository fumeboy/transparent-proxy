transparent-proxy

一個展示了 「透明代理」原理並且做了簡單「負載均衡」的 示例項目。

一 # 透明代理

1）通過 iptables 將 example/http-client 發送的請求轉發到 proxy

2）proxy 獲取該請求的目的地址

3）proxy 根據真實的後端地址，創建與 example/http-server 的連接，然後就可以通過 「client-proxy」、「proxy-server」兩個連接實現數據的轉發。


二 # 負載均衡

簡單起見，並沒有引入真實的服務註冊機，而是用一個臨時的寫死的字典（director/registry/nodes）代替

1）example/http-client 發送的請求到 1.0.0.1:20080。
  「1.0.0.1」 是一個 「假IP」或者說「虛擬IP（VIP）」，它滿足這樣的格式：“1.0.0.${id}”

2）proxy 獲取該請求的目的地址，然後從服務註冊機獲取該 「假IP」 對應的 「真IP」。這一步是服務發現

3）同上面的第三步
