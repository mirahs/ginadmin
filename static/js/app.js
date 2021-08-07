// 确认操作执行(输入 yes 才继续, 避免误操作, 删除重要数据的时候调用)
function msgYn(msg) {
    msg = msg ? msg : '提示: 你确认要操作吗?';
    const i = window.prompt(msg, '请在这里输入 yes 确认操作');
    return i === 'yes';
}

// http 请求
function request(jquery, url, data = {}, callback, type = 'get') {
    jquery.ajax({
        url: url,
        data: data,
        dataType: 'text',
        type: type,

        beforeSend: function () {
            layer.load(2, {
                shade:[0.5, "#333"]
            });
        },
        success: function (res) {
            callback(res, null);
        },
        error: function (xhr) {
            callback(null, xhr);
        },
        complete: function () {
            layer.closeAll("loading");
        },
    });
}

// layui 分页渲染
function laypageRender(laypage, count, curr, limit, query) {
    laypage.render({
        elem: 'page' //这里的 page 是 ID，不用加 # 号
        ,count: count //数据总数
        ,curr: curr //当前页
        ,limit: limit //每页显示的条数
        ,layout: ['count', 'prev', 'page', 'next', 'limit', 'refresh', 'skip']
        ,jump: function (obj, first) {
            if (!first) {
                location.href='?page=' + obj.curr + '&limit=' + obj.limit + query;
            }
        }
    });
}
