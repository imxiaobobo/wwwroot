<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>尚硅谷会员注册页面</title>
    <link type="text/css" rel="stylesheet" href="/static/css/style.css">
    <style type="text/css">
        .login_form {
            height: 420px;
            margin-top: 25px;
        }

    </style>
</head>
<body>
<div id="login_header">
    <img class="logo_img" alt="" src="/static/img/logo.gif">
</div>

<div class="login_banner">

    <div id="l_content">
        <span class="login_word">欢迎注册</span>
    </div>

    <div id="content">
        <div class="login_form">
            <div class="login_box">
                <div class="tit">
                    <h1>注册尚硅谷会员</h1>
                    <span class="errorMsg"></span>
                </div>
                <br>
                <div class="form">
                    <form action="/regist" method="post">
                        <label>用户名称：</label>
                        <input class="itxt" type="text" placeholder="请输入用户名" autocomplete="off" tabindex="1"
                               name="username" id="username"/>
                        <br/>
                        <br/>
                        <label>用户密码：</label>
                        <input class="itxt" type="password" placeholder="请输入密码" autocomplete="off" tabindex="1"
                               name="password" id="password"/>
                        <br/>
                        <br/>
                        <label>确认密码：</label>
                        <input class="itxt" type="password" placeholder="确认密码" autocomplete="off" tabindex="1"
                               name="repwd" id="repwd"/>
                        <br/>
                        <br/>

                        <input type="submit" value="注册" id="sub_btn"/>
                    </form>
                </div>

            </div>
        </div>
    </div>
</div>
<div id="bottom">
			<span>
				尚硅谷书城.Copyright &copy;2015
			</span>
</div>

<script src="https://cdn.bootcss.com/jquery/3.4.1/jquery.min.js">
</script>
<script>
    $("#username").change(function () {
        let username = $(this).val();
        $.ajax({
            type: "POST",
            url: "/ajaxUserName",
            data: {username: username},
            success: function (data) {
                if (data == "用户名已存在"){
                    $(".errorMsg").html(data);
                }else{
                    $(".errorMsg").html("");
                }
            },
        })
    })
</script>
</body>
</html>