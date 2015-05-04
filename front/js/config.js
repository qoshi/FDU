ccolor = ["#CB140D", "#EF5450", "#F48A8A", "#F2B5B4", "#F3D4D4", "#BCBCBC", "#D8D8D8"]
testdataChina = [
    {name: '北京',value: Math.random()},
    {name: '天津',value: Math.random()},
    {name: '上海',value: Math.random()},
    {name: '重庆',value: Math.random()},
    {name: '河北',value: Math.random()},
    {name: '河南',value: Math.random()},
    {name: '云南',value: Math.random()},
    {name: '辽宁',value: Math.random()},
    {name: '黑龙江',value: Math.random()},
    {name: '湖南',value: Math.random()},
    {name: '安徽',value: Math.random()},
    {name: '山东',value: Math.random()},
    {name: '新疆',value: Math.random()},
    {name: '江苏',value: Math.random()},
    {name: '浙江',value: Math.random()},
    {name: '江西',value: Math.random()},
    {name: '湖北',value: Math.random()},
    {name: '广西',value: Math.random()},
    {name: '甘肃',value: Math.random()},
    {name: '山西',value: Math.random()},
    {name: '内蒙古',value: Math.random()},
    {name: '陕西',value: Math.random()},
    {name: '吉林',value: Math.random()},
    {name: '福建',value: Math.random()},
    {name: '贵州',value: Math.random()},
    {name: '广东',value: Math.random()},
    {name: '青海',value: Math.random()},
    {name: '西藏',value: Math.random()},
    {name: '四川',value: Math.random()},
    {name: '宁夏',value: Math.random()},
    {name: '海南',value: Math.random()},
    {name: '台湾',value: Math.random()},
    {name: '香港',value: Math.random()},
    {name: '澳门',value: Math.random()}
];
testdataWorld = [];

zt = {
    host : "http://127.0.0.1:9999",
    map : null,
    chart : null,
    whichMap : "",
    color : ccolor,
    temp:{
        signin:null,
        message:null,
    },
    mapContainer : {
        width : 0,
        height : 0,
    },
    worldMap : {
        tooltip : {
            trigger: 'item',
            formatter: '{b}'
        },
        dataRange: {
            x : 'right',
            min: 0,
            max: 1,
            color:ccolor,
            text:[],          
        },
        series : [
            {
                name: 'world',
                type: 'map',
                mapType: 'world',
                roam: 'move',
                itemStyle:{
                    normal:{label:{show:false}},
                    emphasis:{label:{show:true}}
                },
                data:null,//testdataChina,
            }
        ]
    },
    chinaMap : {
        tooltip : {
            trigger: 'item',
            formatter: '{b}'
        },
        dataRange: {
            x : 'right',
            min: 0,
            max: 1,
            color: ccolor,
            text:[],          
        },
        series : [
            {
                name: 'china',
                type: 'map',
                mapType: 'china',
                roam: "move",
                itemStyle:{
                    normal:{label:{show:true}},
                    emphasis:{label:{show:true}}
                },
                data:null,
            }
        ]
    },
    infoChart : {
        gird : [{
            borderWidth : 0, 
        }],
        xAxis : [
            {
                type : 'category',
                data : ['中国','美国','英国','其他'],
                splitLine:{
                    　　　　show:false
                    　　},
                axisLabel:{
                    textStyle : {
                        color:'white',
                        fontSize : 18,
                    }
                }
            }
        ],
        yAxis : [
            {
                show:false,
                splitLine:{
                    　　　　show:false
                    　　}
            }
        ],
        series : [
            {
                itemStyle: {
                    normal: {
                        label: {
                            show: true,
                            position: 'top',
                        }
                    }
                },
                name:'蒸发量',
                type:'bar',
                data:[2.0, 4.9, 7.0,10],
            }
        ]
    }
}

zt.temp.signin = '<div class="tempS row"><div class="nameS col-md-3">{name}</div><div class="locationS col-md-6">在 {location} 签到</div><div class="col-md-3">{time}</div></div>'

zt.temp.message = '<div class="temM row"><div class="nameM col-md-3">{name}</div><div class="locationM col-md-3">{location}</div><div class="message col-md-6">{message}</div></div>'

