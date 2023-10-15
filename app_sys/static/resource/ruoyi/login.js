$(function () {
    validateRule();
    captcha();
    $(".i-checks").iCheck({checkboxClass: "icheckbox_square-green-demo"});
    $('.imgcode').click(function () {
        captcha();
    });
});

$.validator.setDefaults({
    submitHandler: function () {
        login();
    }
});

function captcha() {
    var url = ctx + "/captchaImage?s=" + Math.random();
    $.ajax({
        type: "get",
        url: url,
        success: function (r) {
            if (r.code == "0") {
                $("#imgcode").attr("src", r.data);
                $("#idkey").val(r.idkey);
            }
        }
    });
}

function login() {
    $.modal.loading($("#btnSubmit").data("loading"));
    var username = $("input[name='username']").val().trim();
    var password = $("input[name='password']").val().trim();
    var validateCode = $("input[name='validateCode']").val();
    var idkey = $("#idkey").val();
    var rememberMe = $("input[name='rememberme']").is(':checked');
    $.ajax({
        type: "post",
        url: ctx + "/login",
        data: {
            "username": username,
            "password": password,
            "validateCode": validateCode,
            "idkey": idkey,
            "rememberMe": rememberMe
        },
        success: function (r) {
            if (r.code == 0) {
                window.location = ctx+'/index';
            } else {
                if(r.data>=5){ //显示验证码
                  $("#vcode").show()
                }
                $.modal.closeLoading();
                $('.imgcode').click();
                $.modal.msg(r.msg);
            }
        }
    });
}

function validateRule() {
    var icon = "<i class='fa fa-times-circle'></i> ";
    $("#signupForm").validate({
        rules: {
            username: {
                required: true
            },
            password: {
                required: true
            }
        },
        messages: {
            username: {
                required: icon + "请输入您的用户名",
            },
            password: {
                required: icon + "请输入您的密码",
            }
        }
    })
}
