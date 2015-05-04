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
    getWorld()
    getChina()
    // zt.chart.setOption(zt.infoChart)
    test();
}

function getChina() {
    var url = zt.host +"/gC?callback=?";
    $.getJSON(url,function(result){
        if ( result.success == false ) {
            return
        }
        setMapData("chinaMap",result.data)
    });
}

function getWorld() {
    var url = zt.host +"/gW?callback=?";
    $.getJSON(url,function(result){
        if ( result.success == false ) {
            return
        }
        zt.worldMap.series[0].data = result.data;
    });
}

function getTop() {
    var url = zt.host+"/gT?callback=?";
    $.getJSON(url,function(result){
        if ( result.success == false )  {
            return
        }
        console.log(result);
        setChartData(result.data);
    });
}

function getSiginNum() {
    var url = zt.host +"/getSiginCount?callback=?";
    $.getJSON(url,function(result){
        if ( result.success == false ) {
            return
        }
        $("userCount").html(result.data);
    });
}


function newPost() {
    var url = zt.host +"/nM?callback=?";
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
    var url = zt.host+"/nS?callback=?";
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
        newSignin();
        newPost();
    });
}


function test() {
    var name = "test";
    var location = "中国";
    var message = "这里是测试留言测试留言测试留言";
    for (var i = 0; i < 10; i++) {
        appendMessage(name,location,i);
        appendSignin(name,location,i);
    }
    // setMapData("chinaMap",testdataChina);

    var url = zt.host+"/socket.io"
    // console.log(url)
    var s = io.connect(url)
    s.on('open',function(){
        console.log("haha");
        s.emit('signin');
    })
    s.on('data',function(data){
        console.log(data);
    })
    s.on('disconnect',function(){
        console.log("oops");
    })
    getTop();
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


function setMapData(pointer,data) {
    zt[pointer].series[0].data = data;
    zt.map.setOption(zt[pointer]);
}

function setChartData(data) {
    zt.infoChart.xAxis[0].data = data.Axis;
    zt.infoChart.series[0].data = data.Count;
    zt.chart.setOption(zt.infoChart)
}
