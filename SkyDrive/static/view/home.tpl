<html>

<head>
    <script src="http://libs.baidu.com/jquery/2.0.0/jquery.min.js"></script>
    <!-- 最新版本的 Bootstrap 核心 CSS 文件 -->
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@3.3.7/dist/css/bootstrap.min.css"
          integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">

    <!-- 可选的 Bootstrap 主题文件（一般不用引入） -->
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@3.3.7/dist/css/bootstrap-theme.min.css"
          integrity="sha384-rHyoN1iRsVXV4nD0JutlnGaslCJuC7uwjduW9SVrLvRYooPp2bWYgmgJQIXwl/Sp" crossorigin="anonymous">

    <!-- 最新的 Bootstrap 核心 JavaScript 文件 -->
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@3.3.7/dist/js/bootstrap.min.js"
            integrity="sha384-Tc5IQib027qvyjSMfHjOMaLkfuWVxZxUPnCJA7l2mCWNIpG9mGCD8wGNIcPD7Txa" crossorigin="anonymous">
    </script>

    <script src="/static/js/auth.js"></script>
</head>

<body style="width:100%;height:100%">
<div style="width:100%;height:100%;margin:0 0 10px 0;text-align: center;">
    <div style="font-size:20px;font-weight:bold;
                margin:0;background: rgb(195, 228, 250);height:32px;">
        文件云盘首页
    </div>
    <table style="height:100%;width:100%;text-align: left;border-width: 2px; border-color: lightslategrey;">
        <tbody>
        <tr style="margin-bottom: 20px;">
            <td style="width:20%;height: 100%;background: lightsteelblue;">
                <div style="text-align: top;height:20%;margin: 10px 0 0 10px;">
                    <img style="width:80px;height:80px;" src="/static/img/avatar.jpeg"></img><br>
                    用户名: <p id="username" style="color: seagreen">{{.}}</p>
                    注册时间: <p id="regtime" style="color: seagreen"></p>
                </div>
                <div style="height: 80%;"></div>
            </td>
            <td style="width: 3px;height:100%;">
                <div style="width:100%;height: 100%;background:rgb(202, 157, 248);"></div>
            </td>
            <td style="text-align: top;">
                <div>文件列表
                    <button class="btn btn-success" onclick="toUploadFile()" style="float: right;margin-right: 30px;">
                        上传文件
                    </button>
                    <div style="width:100%;height: 1px;background:rgb(202, 157, 248);margin-top: 15px;"></div>
                </div>
                <div style="height:95%;" style="width:100%;">
                    <table id="filetbl" style="margin-left:3%;width:96%;">
                        <thead style="height:50px;border:1px;">
                        <tr style="height:50px;border:1px;">
                            <th>文件hash</th>
                            <th>文件名</th>
                            <th>文件大小</th>
                            <th>上传时间</th>
                            <th>最近更新</th>
                        </tr>
                        </thead>
                    </table>
                </div>
            </td>
        </tr>
        </tbody>
    </table>
</div>
</body>

{{/*<script lang="javascript">*/}}
{{/*    window.onload = function () {*/}}
{{/*        var username = document.getElementById('username');*/}}
{{/*        $.ajax({*/}}
{{/*            url: "/user/info?" + queryParams(),*/}}
{{/*            type: "POST",*/}}
{{/*            error: function (err) {*/}}
{{/*                alert(err);*/}}
{{/*            },*/}}
{{/*            success: function (body) {*/}}
{{/*                var resp = JSON.parse(body);*/}}
{{/*                if (resp.code == 10005) {*/}}
{{/*                    window.location.href = "/static/view/signin.html";*/}}
{{/*                }*/}}
{{/*                document.getElementById("username").innerHTML = resp.data.UserName;*/}}
{{/*                document.getElementById("regtime").innerHTML = resp.data.SignupAt;*/}}
{{/*                updateFileList();*/}}
{{/*            }*/}}
{{/*        });*/}}
{{/*    }*/}}

{{/*    function updateFileList() {*/}}
{{/*        $.ajax({*/}}
{{/*            url: "/file/query?" + queryParams(),*/}}
{{/*            type: "POST",*/}}
{{/*            data: {*/}}
{{/*                limit: 15*/}}
{{/*            },*/}}
{{/*            error: function (err) {*/}}
{{/*                alert(err);*/}}
{{/*            },*/}}
{{/*            success: function (body) {*/}}
{{/*                if (!body) {*/}}
{{/*                    return;*/}}
{{/*                }*/}}
{{/*                var data = JSON.parse(body);*/}}
{{/*                if (!data || data.length <= 0) {*/}}
{{/*                    return;*/}}
{{/*                }*/}}

{{/*                for (var i = 0; i < data.length; i++) {*/}}
{{/*                    var x = document.getElementById('filetbl').insertRow();*/}}
{{/*                    var cell = x.insertCell();*/}}
{{/*                    cell.innerHTML = data[i].FileHash.substr(0, 20) + "...";*/}}

{{/*                    cell = x.insertCell();*/}}
{{/*                    cell.innerHTML = data[i].FileName;*/}}

{{/*                    cell = x.insertCell();*/}}
{{/*                    cell.innerHTML = data[i].FileSize;*/}}

{{/*                    cell = x.insertCell();*/}}
{{/*                    cell.innerHTML = data[i].UploadAt;*/}}

{{/*                    cell = x.insertCell();*/}}
{{/*                    cell.innerHTML = data[i].LastUpdated;*/}}
{{/*                }*/}}
{{/*            }*/}}
{{/*        });*/}}
{{/*    }*/}}

{{/*    function toUploadFile() {*/}}
{{/*        window.location.href = '/file/upload?' + queryParams();*/}}
{{/*    }*/}}

{{/*</script>*/}}

</html>