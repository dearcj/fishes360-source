/**
 * Created by MSI on 24.01.2017.
 */
'use strict';
var Parser = require('../front/cocos/assets/scripts/Parser.js');
var W3CWebSocket = require('websocket').w3cwebsocket;
var http = require('http');
var WebSocketClient = require('websocket').client;
var clients = [];

var  host = "5.200.55.232:5050";
//var  host = "10.10.0.2:5050";
//const host = "localhost:5050";


setTimeout(function () {
}, 99999999);

var download = function(url){
    http.request(url, function(response) {

        response.on('data', function(chunk) {
       });
    }).end();
};

var packets = 0;

setInterval(function send(){
    console.log(packets + ' packets per second');
    packets = 0
}, 1000);

function createClient(room) {

    //if (Math.random() < 0.3)
    //download('http://' + host + '/front/cocos/assets/tiles2.png');

    var client = new W3CWebSocket('ws://' + host + '/ws/' + room, null, 'http://localhost/');
    client.clientID = 0;

    clients.push(client);
    console.log('added client to room: ' + room + '; total: ' + clients.length);

    client.onerror = function(e) {
        console.log('Connection Error');
    };

    client.onopen = function() {
        console.log('WebSocket Client Connected');

    };

    client.onclose = function() {
        console.log('echo-protocol Client Closed');
    };



    setInterval(function changeDir() {
        if (!client.me) return;
        var v = 300;
        var angle = Math.random()* Math.PI*2;
        client.VX = Math.cos(angle) * v;
        client.VY = Math.sin(angle) * v;

        client.BVX = Math.cos(angle + Math.PI / 2) * v * 2;
        client.BVY = Math.sin(angle + Math.PI / 2) * v * 2;

        client.bullet = {
            ID: -1,
            CID: client.clientID,
            POS: client.me.POS,
            V: [client.BVX, client.BVY],
            TYPE: 4,
            OWNER: client.me.OWNER
        };
        client.clientID++;
    }, 1000);

        client.onmessage = function(e) {
        if (typeof e.data === 'string') {
            if (!client.playerId) {
                client.playerId = e.data;
            } else {
                client.data = Parser.parse(e.data);
                client.me = null;
                for (var i = 0 ;i < client.data.objectsToSync.length; ++i) {
                    if (client.data.objectsToSync[i].ID == client.playerId) {
                        client.me = client.data.objectsToSync[i];
                    }
                }

                if (client.me && client.VX && client.VY) {
                    client.me.V[0] = client.VX;
                    client.me.V[1] = client.VY;
                }
            }

            //console.log("Received: '" + e.data + "'");
        }
    };
}


var CLIENTS_COUNT = 150;
var NUM_ROOMS = 100;
for (var i = 0; i < CLIENTS_COUNT; ++i) {
    var room = Math.floor(Math.random() * (NUM_ROOMS));
    setTimeout(
        createClient.bind(this, room)
    , 200*i)
}

console.time("interval");
setInterval(function send(){
    console.timeEnd("interval");
    console.time("interval");
    var cl = clients.length;
    for (var i = 0; i < cl; ++i) {
        var client = clients[i];
        if (client.me) {
            var str = Parser.serialize(client.me);

            if (client.bullet) {
                str = str + ';'+ Parser.serialize(client.bullet);
                client.bullet = null;
            }
            if (client.readyState == 1) {
                client.send(str);
                packets++;
            }
        }
    }


}, 40);


process.stdin.resume();
process.on('SIGINT', function () {
    console.log('aborted all clients');

    for (var i = 0; i < clients.length; ++i)
        clients[i].close();

    setTimeout(function() {
        process.exit (0);
    })
});

