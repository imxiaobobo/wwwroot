<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>图书管理</title>
    <link type="text/css" rel="stylesheet" href="/static/css/style.css">
</head>
<body>

<div id="header">
    <img class="logo_img" alt="" src="/static/img/logo.gif">
    <span class="wel_word">图书管理系统</span>
    <div>
        <a href="javascript:;">图书管理</a>
        <a href="/">返回商城</a>
    </div>
</div>

<div id="main">
    <table>
        <tr>
            <td>名称</td>
            <td>价格</td>
            <td>作者</td>
            <td>销量</td>
            <td>库存</td>
            <td colspan="2">操作</td>
        </tr>
        {{range .Book}}
            <tr>
                <td>{{.Title}}</td>
                <td>{{.Price}}</td>
                <td>{{.Author}}</td>
                <td>{{.Sales}}</td>
                <td>{{.Stock}}</td>
                <td><a href="/bookEdit?bookid={{.ID}}">修改</a></td>
                <td><a id="{{.Title}}" class="bookDel" href="/bookDel?bookid={{.ID}}">删除</a></td>
            </tr>
        {{end}}
        <tr>
            <td></td>
            <td></td>
            <td></td>
            <td></td>
            <td></td>
            <td></td>
            <td><a href="/addBook">添加图书</a></td>
        </tr>
    </table>

    <div id="page_nav">
        {{if .IsHasPrev}}
            <a href="#">首页</a>
            <a href="/bookManager?pageNo={{.GetPrev}}">上一页</a>
        {{end}}
        当前是第{{.PageNo}}页,共{{.Count}}页,共{{.Total}}条记录
        {{if .IsHasNext}}
            <a href="/bookManager?pageNo={{.GetNext}}">下一页</a>
            <a href="#">末页</a>
        {{end}}
    </div>
</div>

<div id="bottom">
		<span>
			尚硅谷书城.Copyright &copy;2015
		</span>
</div>

<script src="static/script/jquery-1.7.2.js">
</script>
<script>
    $(".bookDel").click(function () {
        var title = $(this).attr("id")//获取书名
        return confirm("确定要删除「" + title + "」么?")
    })
</script>
</body>
</html>