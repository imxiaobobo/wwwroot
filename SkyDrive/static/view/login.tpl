<html>

<head>
    <script src="http://libs.baidu.com/jquery/2.0.0/jquery.min.js"></script>
    <!-- 最新版本的 Bootstrap 核心 CSS 文件 -->
    <link rel="stylesheet"
          href="https://cdn.jsdelivr.net/npm/bootstrap@3.3.7/dist/css/bootstrap.min.css"
          integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">

    <!-- 可选的 Bootstrap 主题文件（一般不用引入） -->
    <link rel="stylesheet"
          href="https://cdn.jsdelivr.net/npm/bootstrap@3.3.7/dist/css/bootstrap-theme.min.css"
          integrity="sha384-rHyoN1iRsVXV4nD0JutlnGaslCJuC7uwjduW9SVrLvRYooPp2bWYgmgJQIXwl/Sp" crossorigin="anonymous">

    <!-- 最新的 Bootstrap 核心 JavaScript 文件 -->
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@3.3.7/dist/js/bootstrap.min.js"
            integrity="sha384-Tc5IQib027qvyjSMfHjOMaLkfuWVxZxUPnCJA7l2mCWNIpG9mGCD8wGNIcPD7Txa"
            crossorigin="anonymous"></script>
    <script lang="javascript">
    </script>
</head>

<body>
<form>
    <div style="width:500px;margin:10px auto;text-align: center;">
        <div style="font-size:28px;font-weight:bold;margin:0px auto;">用户登录</div>
        <br/>
        <table style="width:100%;text-align: left;">
            <tbody>
            <tr style="margin-bottom: 20px;">
                <td>
                    <span class="p">*</span>
                    <label for="username" class="l"> 用户名:</label>
                </td>
                <td>
                    <input id="username" type="text"
                           style="height:30px;width:250px;padding-right:50px;">
                </td>
            </tr>
            <tr>
                <td><br></td>
                <td></td>
            </tr>
            <tr style="margin-bottom: 20px;">
                <td>
                    <span class="p">*</span>
                    <label for="password" class="l"> 密码:</label>
                </td>
                <td>
                    <input id="password" type="text"
                           style="height:30px;width:250px;padding-right:50px;">
                </td>
            </tr>
            <tr>
                <td><br></td>
                <td></td>
            </tr>
            <tr>
                <td>
                </td>
                <td>
                    <input class="btn btn-success" type="button"
                           style="margin:0 auto;width:250px;"
                           value="登录"
                           onclick='onSignin()'/>
                </td>
            </tr>
            </tbody>
        </table>
    </div>
</form>
</body>

<script lang="javascript">
    function onSignin() {
        var username = document.getElementById('username');
        var password = document.getElementById('password');
        $.ajax({
            url: "/user/login",
            type: "POST",
            data: {"username": username.value, "password": password.value},
            error: function (err) {
                alert(err);
            },
            success: function (body) {
                var data = JSON.parse(body)
                console.log(data)
                if (data.code == 404) {
                    alert(data.msg);
                    return;
                }else{
                    // localStorage.setItem("token", data.data.Token)
                    // localStorage.setItem("username", data.data.Username)
                    window.location.href = data.data.Location+"?username="+data.data.Username+"&token="+data.data.Token;
                }

            }
        });
    }
</script>
</html>