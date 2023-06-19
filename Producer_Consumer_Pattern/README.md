# Producer Consumer Pattern
生產消費模式

* 多個 Producer 程序存到一個 queue 中，再由其他 Consumer 程序讀取 queue 執行，這樣的話可以使 Producer 與 Consumer 程序間沒有直接關係，他們只依賴 queue，即可解耦

* 問題 :
    * 多個使用者叫車，不同的司機接單會收到此使用者的資訊

* 解決方式 :
    * 設計一個 job channel 搜集使用者的叫車單與資訊, 不斷監聽 job channel 是否有新的叫車，如果有的話就執行載客服務
