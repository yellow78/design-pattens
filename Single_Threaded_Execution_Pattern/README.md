# Single Threaded Execution Pattern
單線程執行模式

* 問題 :
    * like.Add()並沒有 goroutine safe(執行緒安全)，造成了讀寫like.Count產生了 race condition，即四個 goroutine 會同時讀寫 count，導致有 goroutine 在寫入後，其他 goroutine 卻沒有讀到最新資料的狀況

* 解決方式 :
    * like.Add()裡新增 lock 來確保 function只會被一個goroutine讀寫，當一個 goroutine 取得 like.Add()的 lock，即有權限執行此 function，除非此 goroutine unlock，讓其他 goroutine 有機會取得 lock，否則其他 goroutine 只能等待

