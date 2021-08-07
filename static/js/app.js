function msg_yn(msg) {
    msg = msg ? msg : '提示：删除将无法恢复！你确认要删除吗？';
    const i = window.prompt(msg, '请在这里输入 yes 确认操作');
    return i === 'yes';
}

function request(jquery, url, data = {}, callback, type = 'get') {
    data['is_layui'] = 1;
    jquery.ajax({
        url: url,
        data: data,
        dataType: 'text',
        type: type,

        beforeSend: function() {
            layer.load(2, {
                shade:[0.5, "#333"]
            });
        },
        success: function(res) {
            callback(res, null);
        },
        error: function(xhr) {
            callback(null, xhr);
        },
        complete: function() {
            layer.closeAll("loading");
        },
    });
}


function laypage_render(laypage, count, curr, limit, query) {
    laypage.render({
        elem: 'page' //这里的 page 是 ID，不用加 # 号
        ,count: count //数据总数
        ,curr: curr //当前页
        ,limit: limit //每页显示的条数
        ,layout: ['count', 'prev', 'page', 'next', 'limit', 'refresh', 'skip']
        ,jump: function (obj, first){
            // obj包含了当前分页的所有参数，比如：
            // console.log(obj.curr);  //得到当前页，以便向服务端请求对应页的数据。
            // console.log(obj.limit); //得到每页显示的条数
            // 首次不执行,使用原始的curr,后面需要自己通过回传来更新
            if (!first) {
                location.href='?page=' + obj.curr + '&limit=' + obj.limit + query;
            }
        }
    });
}
