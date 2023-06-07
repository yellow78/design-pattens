# Read Write Lock Pattern
讀寫鎖定模式。多讀單寫

* 問題 :
    * lock 分為 read 與 write 兩種，讓 lock 效能更佳，read 行為不會改變資料，所以 read lock 時可以再同時 read，達到「多讀」，而 write 行為會改變資料，所以 write lock 不能再同時 read 與 write 時，達到「單寫」
    
    * read, write 中在取得 lock 的時候，只要碰到以下兩種「衝突 conflict」狀態，就會等待此衝突結束，再取得 lock：

        * read 與 write 的 lock 同時存在
        * write 與 write 的 lock 同時存在

    * 而在以下情形，會直接取得 lock，不會等待此狀態結束：

        * read 與 read 的 lock 同時存在

* 解決方式 :
    1. 在寫入處Add()與讀取處Show()中加入sync.Mutex{}的 lock，使讀寫不同時進行
    * 此方式讀寫都會lock, 但讀取不會改變資料可以不用lock, 運行時間較差

    2. sync.Mutex{}換成sync.RWMutex{}讀寫鎖，Show()部分的 lock 換成 read-lock
    * 改為讀寫鎖, 系統只要注意改變資料時其他線程完成讀取即可
    

