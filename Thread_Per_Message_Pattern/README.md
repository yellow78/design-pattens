# Thread Per Message Pattern
消息線程模式, 每個任務互不相關當task沒有執行順序或資料保護等需求, 可將task同時執行

* 問題 :
    * 設計一個推播新聞系統，會將新的新聞逐個推播出去，提出效能提升方案

* 解決方式 :
    * 使用goroutine多執行序併發處理任務

* channel 特性：
    * 在讀寫時不會產生 race condtion

    * 在讀取時如果沒有元素，會等待到有值再執行

    * 如果使用 unbuffered channel，e.g. make(chan string)，寫入後如果沒有讀取，會在寫入處等待
    
    * 如果使用 buffered channel，e.g. make(chan string, 100)，寫入後如果沒有讀取，可以繼續寫入，直到到達 buffer 數才會等待讀取