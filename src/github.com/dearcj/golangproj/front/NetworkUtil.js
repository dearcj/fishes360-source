import {ClientData, ServerData}  from "./compiled.js"


export class Network {


    /*
    Should be called during update (requestAnimationFrame)
     */
    process() {
        if (this.needSend) {
            this.send();
        }
    }

    valOr0(v) {
        return v ? v : 0
    }

    sendPlayerCommand(cmd, params) {
        let command = {
            CommandId: cmd,
            Params: params
        };

        this.playerCommands.push(command);
        this.needSend = true;
    }

    reset() {
        let ids = [];

        for (let x of this.networkObjects) {
            ids.push(x.networkObject.ID);
        }
        this.removeNetworkObjects(ids);

        this.prevMessage = null;
        this.needSend = false;
        this.networkObjects = [];
        this.playerCommands = [];
    }

    send() {
        if (!this.ws) return;

        let message = ClientData.create({
            commands: this.playerCommands
        });
        let buffer = ClientData.encode(message).finish();

        if (this.ws.readyState === 1) {
            if (buffer.length > this.maxMsgSentLength)
                this.maxMsgSentLength = buffer.length;
            this.ws.send(buffer);
        }

        this.playerCommands = [];
        this.needSend = false;
    }

    getObj(ID) {
        if (this.ids[ID]) return this.ids[ID];

        return this.objByID(ID);
    }

    processData(serverData) {
        this.objectsToSync = [];
        if (serverData)
            this.objectsToSync = this.objectsToSync.concat(serverData.customObjects, serverData.players, serverData.enemies);

        //console.log('Server packet with enemy count = ' + serverData.enemies.length);

        if (serverData.conData) {
            this.playerId = this.valOr0(serverData.conData.playerID);
        }



    }

    removeNetworkObjects(Ids) {
        let len = Ids.length;
        for (let j = 0; j < len; ++j) {
            let obj = this.getObj(Ids[j]);
            if (obj) {
                console.log("Removing network object", Ids[j]);
                this.ids[Ids[j]] = null;
                this.networkObjects.splice(this.networkObjects.indexOf(obj), 1);
            }
        }
    }

    mutateObjStates(objectsState, accountGeneral, serverData) {
        for (let newData of this.objectsToSync) {
            let no = newData.networkObject ? newData.networkObject : newData;
            let obj = this.getObj(no.ID);
            if (obj && obj.gameObject) {
                (obj.gameObject).mutateGameObject(newData, accountGeneral, serverData);
            }
        }
    }

    resetConData() {
        if (this.ws) {
            this.objectsToSync = null;
            this.networkObjects = [];
            this.playerCommands = [];
            this.ws.close();
            this.ws = null;
        }
    }

    startWS(changeHost) {
        this.resetConData();

        let log = "";
        let pass = "";
        let anonym = true;
        let host = window.location.hostname;
        if (host === "") host = "localhost";
        if (changeHost) host = changeHost;

        let protocol = location.protocol === "http:" ? "ws://" : "wss://";

        let params = (new URL(document.location)).searchParams;
        let token = params.get('token');
        let puuid = params.get('puuid');


        let port = location.port;//nodeConfig.Server.SSL ? nodeConfig.Server.SSLPort : nodeConfig.Server.Port;
        this.ws = new WebSocket(protocol + host + port + "/ws/?log=" + log + "&pass=" + pass + "&anonym=" + anonym + "&http_key=" + "vkMfSCmAsf" + "&token=" + token + "&puuid=" + puuid );
        this.ws.binaryType = 'arraybuffer';

        this.ws.onopen = (event) => {
            this.totalTicks = 0;
            this.totalDelta = 0;
            console.log("Send Text WS was opened.");
            if (this.onOpen) {
                this.onOpen(event, this)
            }
        };

        setInterval(() => {
            this.needSend = true;
            this.send();
        }, this.PONG_INTERVAL);

        this.ws.onmessage = (event) => {
            let message = ServerData.decode(new Uint8Array(event.data));

            if (this.prevMessageTime) {
                let delta = (new Date()).getTime() - this.prevMessageTime.getTime();
                this.totalTicks++;
                this.totalDelta += delta;
                if (delta > this.maxPing) this.maxPing = delta;
                if (event.data.byteLength > this.maxMsgLength) {
                    this.maxMsgLength = event.data.byteLength;
                }
            }
            this.prevMessageTime = new Date();
            this.processData(message);

            if (this.onMessage)
                this.onMessage(message, this);
        };

        this.ws.onerror = (event) => {
            console.log("Send Text fired an error");
        };

        this.ws.onclose = (event) => {
            console.log("WebSocket instance closed.");
        };
    }

    constructor(onMessage, onOpen) {
        this.maxPing = 0;
        this.objectsToSync = null;

        this.ws = null;// WebSocket;
        this.ids = {};
        this.networkObjects = [];
        this.playerCommands = [];
        this.playerId = -1;
        this.maxMsgSentLength = 0;
        this.maxMsgLength = 0;
        this.totalTicks = 0;
        this.totalDelta = 0;
        this.prevMessageTime = null;
        this.needSend = false;
        this.PONG_INTERVAL = 15000;

        this.onMessage = onMessage;
        this.onOpen = onOpen;
    }

    init() {
        this.startWS();
    }

}