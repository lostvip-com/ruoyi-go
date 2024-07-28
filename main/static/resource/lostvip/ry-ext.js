function lv_detail_right(title, url,width,height){
    if (height==undefined || height<=0 || height=="") {
        height = '100%'
    }
    top.layer.open({
        type: 2,
        offset: 'r',
        anim: 'slideLeft',
        move: false,
        title: title,
        shade: 0.3,
        shadeClose: true,
        area: [ width, height],
        content: url
    });
}

function lv_edit_right(title, url,width,height,callback){
    if (height==undefined || height<=0 || height=="") {
        height = '100%'
    }
    top.layer.open({
        type: 2,
        offset: 'r',
        anim: 'slideLeft',
        move: false,
        title: title,
        shade: 0.3,
        shadeClose: true,
        area: [ width, height],
        btn: ['确定', '关闭'],
        content: url,
        yes: function(index, layero) {
            var iframeWin = layero.find('iframe')[0];
            iframeWin.contentWindow.submitHandler(index, layero,callback);
        },
        cancel: function(index) {
            return true;
        },
        success: function () { //加载成功
            $(':focus').blur();
        }
    });
}

