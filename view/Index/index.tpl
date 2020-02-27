<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>书城首页</title>
    <link type="text/css" rel="stylesheet" href="/static/css/style.css">
</head>
<body>

<div id="header">
    <img class="logo_img" alt="" src="/static/img/logo.gif">
    <span class="wel_word">网上书城</span>
    <div>
        <a href="/login">登录</a> |
        <a href="/regist">注册</a> &nbsp;&nbsp;
        <a href="/pages/cart/cart.html">购物车</a>
        <a href="/manager">后台管理</a>
    </div>
</div>

<div id="main">
    <div id="book">
        <div class="book_cond">
            <form action="/" method="get">
            价格：<input type="text" name="min"> 元 - <input type="text" name="max"> 元
            <button>查询</button>
            </form>
        </div>
        <div style="text-align: center">
            <span>您的购物车中有3件商品</span>
            <div>
                您刚刚将<span style="color: red">时间简史</span>加入到了购物车中
            </div>
        </div>
        {{range .Book}}
        <div class="b_list">
            <div class="img_div">
                <img class="book_img" alt="" src="/static/img/default.jpg"/>
            </div>
            <div class="book_info">
                <div class="book_name">
                    <span class="sp1">书名:</span>
                    <span class="sp2">{{.Title}}</span>
                </div>
                <div class="book_author">
                    <span class="sp1">作者:</span>
                    <span class="sp2">{{.Price}}</span>
                </div>
                <div class="book_price">
                    <span class="sp1">价格:</span>
                    <span class="sp2">￥{{.Author}}</span>
                </div>
                <div class="book_sales">
                    <span class="sp1">销量:</span>
                    <span class="sp2">{{.Sales}}</span>
                </div>
                <div class="book_amount">
                    <span class="sp1">库存:</span>
                    <span class="sp2">{{.Stock}}</span>
                </div>
                <div class="book_add">
                    <button>加入购物车</button>
                </div>
            </div>
        </div>
        {{end}}
    </div>

    <div id="page_nav">
        {{if .IsHasPrev}}
            <a href="#">首页</a>
            <a href="/?pageNo={{.GetPrev}}">上一页</a>
        {{end}}
        当前是第{{.PageNo}}页,共{{.Count}}页,共{{.Total}}条记录
        {{if .IsHasNext}}
            <a href="/?pageNo={{.GetNext}}">下一页</a>
            <a href="#">末页</a>
        {{end}}
    </div>

</div>

<div id="bottom">
		<span>
			尚硅谷书城.Copyright &copy;2015
		</span>
</div>
</body>
</html>