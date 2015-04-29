$(document).ready(function() {
    $('#fullpage').fullpage();
    init();
    bindEvent();
});


function init() {
    $("#map").css({
        "width":"100%",
        "height":parseInt(document.body.clientHeight)-37-231+"px"
    })
    zt.chart = echarts.init(document.getElementById('chart'));
    zt.map = echarts.init(document.getElementById('map'));
    zt.map.on("click",function(e){
        console.log(e);
    })
    changeToMap("chinaMap");
    zt.chart.setOption(zt.infoChart)
    test();
}


function getSiginNum() {
    var url = zt.host +"/getSiginCount";
    $.get(url,function(result){
        if ( result.success == false ) {
            return
        }
        $("userCount").html(result.data);
    });
}

function newPost() {
    var url = zt.host +"/nM";
    console.log(url);
    var data = {
        name :$("#msgName").val(),
        country :$("#msgLocation").val(),
        context :$("#msgContext").val()
    }
    $.post(url,data,function(result){
        if ( result.success == false ) {
            alert("留言失败");
        } else {
            alert("留言成功");
        }
    });
}

function newSignin() {
    var url = zt.host+"/nS";
    console.log(url);
    var data = {
        //TODO
        name :$("#msgName").val(),
        location : 123,
        context :$("#msgContext").val()
    }
    $.post(url,data,function(result){
        if ( result.success == false ) {
            alert("留言成功");
        } else {
            alert("签到成功");
        }
    });
}

function changeToMap(str) {
    if (zt.whichMap === str) {
        return true;
    }
    console.log(zt[str]);
    zt.map.setOption(zt[str]);
    zt.whichMap = str;
    return false;
}

function bindEvent() {
    $("#worldBut").on("click",function(){
        if ( changeToMap("worldMap") === true ) {
            return false;
        }
        $("#worldBut").addClass("btnMA");
        $("#chinaBut").removeClass("btnMA");
        return false;
    });
    $("#chinaBut").on("click",function(){
        if ( changeToMap("chinaMap") === true ) {
            return false;
        }
        $("#chinaBut").addClass("btnMA");
        $("#worldBut").removeClass("btnMA");
        return false;
    });
    $("#sendMessage").on("click",function(){
        console.log("hahaha");
        newSignin();
        newPost();
    });
}


function test() {
    var name = "sb";
    var location = "中国";
    var message = "今天好运气啊，老狼请吃鸡啊，你打电话我不接，你打他有啥用啊？？你打电话我不接";
    for (var i = 0; i < 10; i++) {
        appendMessage(name,location,i);
        appendSignin(name,location,i);
    }
}

function appendMessage(name,location,message) {
    var str = zt.temp.message;
    var $main = $("#message")
    str = str.replace('{name}',name);
    str = str.replace('{location}',location);
    str = str.replace('{message}',message);
    $main.append($(str));

}

function appendSignin(name,location,time) {
    var str = zt.temp.signin;
    var $main = $("#signinContainer");
    str = str.replace('{name}',name);
    str = str.replace('{location}',location);
    str = str.replace('{time}',time);
    $main.prepend($(str));
}


