recover()在defer之外调用无效

recover只能在defer语句中的函数内有效。如果直接调用recover()，它不会捕获panic，只是单纯地返回nil，不会有任何效果。