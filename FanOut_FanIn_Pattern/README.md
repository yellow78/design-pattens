# FanOut FanIn Pattern
扇入扇出模式

* FanIn(扇入)
    * 收集資料
    * 將多個輸入通道(channels)中的數據聚合到一個輸出通道中

* FanIn(扇出)
    * 分發任務
    * 將一個任務分發給多個併發處理(worker)

* 問題 :
    * 設計一個資訊看板系統，需要到不同server拿取資料，資料沒有順序性, 加快效能

* 解決方式 :
    * 實作Producer()、Task()、Consumer()來分別分發任務、執行任務、收集資料
