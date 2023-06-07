# Guarded Suspension Pattern
保護掛起模式, thread 執行時條件不符，就停下等待，等到條件符合時再開始開始執行

* 問題 :
    * 設計線上即時顯示留言板時，會接受大量的留言，並將其顯示出來，我們需確保留言不出現 race condition，也需確保沒留言時不執行顯示

* 解決方式 :
    * m.messages 讀取時，新增執行條件，即「m.messages 為空的時候等待而不讀取，等到有訊息實再讀取顯示」，此時可以透過 channel 來實現

* channel 特性：
    * 在讀寫時不會產生 race condtion

    * 在讀取時如果沒有元素，會等待到有值再執行

    * 如果使用 unbuffered channel，e.g. make(chan string)，寫入後如果沒有讀取，會在寫入處等待
    
    * 如果使用 buffered channel，e.g. make(chan string, 100)，寫入後如果沒有讀取，可以繼續寫入，直到到達 buffer 數才會等待讀取