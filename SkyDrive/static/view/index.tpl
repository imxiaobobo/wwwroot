<html>

<head>
    <!-- bootstrap 4.x is supported. You can also use the bootstrap css 3.3.x versions -->
    <link rel="stylesheet" href="/static/css/bootstrap.min.css">
    <link href="/static/css/fileinput.min.css" media="all" rel="stylesheet" type="text/css"/>
    <!-- if using RTL (Right-To-Left) orientation, load the RTL CSS file after fileinput.css by uncommenting below -->
    <!-- link href="https://cdnjs.cloudflare.com/ajax/libs/bootstrap-fileinput/4.4.9/css/fileinput-rtl.min.css" media="all" rel="stylesheet" type="text/css" /-->
    <script src="/static/js/jquery-3.2.1.min.js"></script>
    <!-- piexif.min.js is needed for auto orienting image files OR when restoring exif data in resized images and when you
          wish to resize images before upload. This must be loaded before fileinput.min.js -->
    <script src="/static/js/piexif.min.js" type="text/javascript"></script>
    <!-- sortable.min.js is only needed if you wish to sort / rearrange files in initial preview.
          This must be loaded before fileinput.min.js -->
    <script src="/static/js/sortable.min.js" type="text/javascript"></script>
    <!-- purify.min.js is only needed if you wish to purify HTML content in your preview for
          HTML files. This must be loaded before fileinput.min.js -->
    <script src="/static/js/purify.min.js" type="text/javascript"></script>
    <!-- popper.min.js below is needed if you use bootstrap 4.x. You can also use the bootstrap js
         3.3.x versions without popper.min.js. -->
    <script src="/static/js/popper.min.js"></script>
    <!-- bootstrap.min.js below is needed if you wish to zoom and preview file content in a detail modal
          dialog. bootstrap 4.x is supported. You can also use the bootstrap js 3.3.x versions. -->
    <script src="/static/js/bootstrap.min.js" type="text/javascript">
    </script>
    <!-- the main fileinput plugin file -->
    <script src="/static/js/fileinput.min.js"></script>
    <!-- optionally if you need a theme like font awesome theme you can include it as mentioned below -->
    <script src="/static/js/theme.js"></script>
    <script src="/static/js/auth.js"></script>
</head>

<body style="width:100%;height:100%;text-align:center;">
<div style="width:100%;height:100%;margin:0 0 10px 0;text-align: center;">
    <div style="font-size:20px;font-weight:bold;color:#ddd;
      margin:0;padding-top:3px;background:#383e4b;height:40px;">
        文件上传
    </div>
    <div style="width:60%;height:30%;text-align:center;margin: 50px auto;">
        <form id="upForm" action="/file/upload" method="post" enctype="multipart/form-data">
            <!-- <input id="file" name="file" type="file" class="file" data-msg-placeholder="选择文件"> -->
            <input id="file" type="file" name="file"/>
            <button id="上传" type="submit">upload</button>
        </form>
    </div>
</div>
</body>

<script lang="javascript">
    // $(document).ready(function () {
    //   var upEntry = localStorage.getItem("uploadEntry");
    //   if (upEntry != "") {
    //     alert("http://" + upEntry + "/file/upload");
    //     document.getElementById("upForm").action = "http://" + upEntry + "/file/upload";
    //   }
    // });

    function onUpload() {
        var upUrl = "/file/upload?" + queryParams();
        var upEntry = localStorage.getItem("uploadEntry");
        if (upEntry != "") {
            if (upEntry.indexOf("http:") >= 0) {
                upUrl = upEntry + "/file/upload?" + queryParams();
            } else {
                upUrl = "http://" + upEntry + "/file/upload?" + queryParams();
            }
        }
        $.ajax({
            url: upUrl,
            type: 'POST',
            cache: false,
            data: new FormData($('#upForm')[0]),
            processData: false,
            contentType: false,
            error: function (err) {
                alert('请求报错信息: ' + JSON.stringify(err));
            },
            success: function (res) {
                alert('请求返回结果: ' + JSON.stringify(res));
                if (res.code == 0) {
                    window.location.href = "/static/view/home.html";
                }
            }
        });
    }
</script>

</html>
