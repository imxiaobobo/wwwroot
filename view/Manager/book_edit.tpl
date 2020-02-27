<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>编辑图书</title>
<link type="text/css" rel="stylesheet" href="/static/css/style.css" >
<style type="text/css">
	h1 {
		text-align: center;
		margin-top: 200px;
	}
	
	h1 a {
		color:red;
	}
	
	input {
		text-align: center;
	}
</style>
</head>
<body>
		<div id="header">
			<img class="logo_img" alt="" src="/static/img/logo.gif" >
			<span class="wel_word">编辑图书</span>
			<div>
				<a href="/manager">图书管理</a>
				<a href="/">返回商城</a>
			</div>
		</div>
		
		<div id="main">
			<form action="/bookEdit" method="POST">
				<input type="hidden" name="bookid" value="{{.ID}}" />
				<table>
					<tr>
						<td>名称</td>
						<td>价格</td>
						<td>作者</td>
						<td>销量</td>
						<td>库存</td>
						<td colspan="2">操作</td>
					</tr>		
					<tr>
						<td><input name="title" type="text" value="{{.Title}}"/></td>
						<td><input name="price" type="text" value="{{.Price}}"/></td>
						<td><input name="author" type="text" value="{{.Author}}"/></td>
						<td><input name="sales" type="text" value="{{.Sales}}"/></td>
						<td><input name="stock" type="text" value="{{.Stock}}"/></td>
						<td><input type="submit" value="提交"/></td>
					</tr>		
				</table>
			</form>
			
	
		</div>
		
		<div id="bottom">
			<span>
				尚硅谷书城.Copyright &copy;2015
			</span>
		</div>
</body>
</html>