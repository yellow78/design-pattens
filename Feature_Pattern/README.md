# Feature Pattern
特徵模式, 異步io每個任務處理完畢取得回傳之結果, 可接續處理相對應之邏輯

* 問題 :
    * 設計一個推播新聞系統，會將新的新聞逐個推播出去，並記錄執行時間

* 解決方式 :
    * goroutine執行後, 實作中間物件取得併發結果, 在此使用channel實作
