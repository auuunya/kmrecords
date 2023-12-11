package kmrecords

var (
	HTBORDER      = 18 //在没有大小边界的窗口边框中。
	HTBOTTOM      = 15 //在可调整大小的窗口的下水平边框中（用户可以单击鼠标以垂直调整窗口大小）。
	HTBOTTOMLEFT  = 16 //在可调整大小的窗口的边框左下角（用户可以单击鼠标以对角线调整窗口大小）。
	HTBOTTOMRIGHT = 17 //在可调整大小的窗口的边框右下角（用户可以单击鼠标以对角线调整窗口大小）。
	HTCAPTION     = 2  //在标题栏中。
	HTCLIENT      = 1  //在客户端区中。
	HTCLOSE       = 20 //在关闭按钮中。
	HTERROR       = -2 //在屏幕背景上或窗口之间的分割线上（与 HTNOWHERE 相同，只是 DefWindowProc 函数会生成系统蜂鸣音以指示错误）。
	HTGROWBOX     = 4  //在大小框中（与 HTSIZE 相同）。
	HTHELP        = 21 //在帮助按钮中。
	HTHSCROLL     = 6  //在水平滚动条中。
	HTLEFT        = 10 //在可调整大小的窗口的左边框中（用户可以单击鼠标以水平调整窗口大小）。
	HTMENU        = 5  //在菜单中。
	HTMAXBUTTON   = 9  //在最大化按钮中。
	HTMINBUTTON   = 8  //在最小化按钮中。
	HTNOWHERE     = 0  //在屏幕背景上，或在窗口之间的分隔线上。
	HTREDUCE      = 8  //在最小化按钮中。
	HTRIGHT       = 11 //在可调整大小的窗口的右左边框中（用户可以单击鼠标以水平调整窗口大小）。
	HTSIZE        = 4  //在大小框中（与 HTGROWBOX 相同）。
	HTSYSMENU     = 3  //在窗口菜单或子窗口的关闭按钮中。
	HTTOP         = 12 //在窗口的上水平边框中。
	HTTOPLEFT     = 13 //在窗口边框的左上角。
	HTTOPRIGHT    = 14 //在窗口边框的右上角。
	HTTRANSPARENT = -1 //在同一线程当前由另一个窗口覆盖的窗口中（消息将发送到同一线程中的基础窗口，直到其中一个窗口返回不是 HTTRANSPARENT 的代码）。
	HTVSCROLL     = 7  //在垂直滚动条中。
	HTZOOM        = 9  //在最大化按钮中。
)
