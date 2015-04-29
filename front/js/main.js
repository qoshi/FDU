zt = {
    host : "http://127.0.0:8080/",
    map : null,
    whichMap : "",
    mapContainer : {
        width : 0,
        height : 0,
    },
    worldMap : {
    tooltip : {
        trigger: 'item',
        formatter: '{b}'
    },
    series : [
        {
            name: '世界地图',
            type: 'map',
            mapType: 'world',
            roam: true,
            selectedMode : 'single',
            itemStyle:{
                normal:{label:{show:false}},
                emphasis:{label:{show:true}}
            },
            data:[],
        }
    ]
    },
    chinaMap : {
        tooltip : {
            trigger: 'item',
            formatter: '{b}'
        },
        series : [
            {
                name: '中国',
                type: 'map',
                mapType: 'china',
                selectedMode : 'multiple',
                roam: true,
                itemStyle:{
                    normal:{label:{show:true}},
                    emphasis:{label:{show:true}}
                },
                data:[
                    {name:'广东',selected:true}
                ]
            }
        ]
    }
}
$(document).ready(function() {
    $('#fullpage').fullpage();
    init();
    bindEvent();
});


function init() {
    zt.chart = echarts.init(document.getElementById('map'));
    zt.chart.setOption(zt.chinaMap);
    // changeToMap("chinaMap");
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
    var url = zt.host +"/nP";
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
    var url = zt.host="/nS";
    var data = {
        //TODO
        name :"",
        location : 123,
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
    zt.chart.setOption(zt[str]);
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
}
