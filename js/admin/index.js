$(function(){
    //隐藏左侧导航
    $('.hideMenu').click(function(){
        var width = $('.layui-table').width();
        var $admin = $(".layui-layout-admin");
        $admin.toggleClass("showMenu");
        if ($admin.hasClass('showMenu')) {
            $('.hideMenu i').removeClass('layui-icon-shrink-right').addClass('layui-icon-spread-left');
            width += 200;
            $('.layui-table').animate({width: width + 'px'});
        } else {
            $('.hideMenu i').removeClass('layui-icon-spread-left').addClass('layui-icon-shrink-right');
            width -= 200;
            $('.layui-table').animate({width: width + 'px'});
        }
    });

    $('.refresh').click(function(){
        location.reload();
    });
});