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
// 适合 多层iframe嵌套时使用
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
// 适合 多层iframe嵌套时使用
function lv_add_center(title, url,width,height,callback){

    top.layer.open({
        type: 2,
        area: [ width, height],
        fix: false,
        //不固定
        shade: 0.3,
        title: title,
        content: url,
        btn: ['确定', '关闭'],
        // 弹层外区域关闭
        shadeClose: true,
        yes: function(index, layero) {
            let iframeWin = layero.find('iframe')[0];
            iframeWin.contentWindow.submitHandler(index, layero,callback);
        },
        cancel: function(index) {
            return true;
        },
        success: function () {
            $(':focus').blur();
        }
    });
}

document.addEventListener('keydown', function(event) {
    if (event.key === "Enter") {
        // 执行操作
        $.table.search()
    }
});