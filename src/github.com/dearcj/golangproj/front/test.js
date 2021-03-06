import {Network} from "./NetworkUtil.js"
import config from "./config.js" //this config is generated by server each time it starts


let allPlayers = {};
let allFishes = {};


let onMessage = (data, networkUtil) => {
  console.log("got message", data);

   let objById = (list,  id) => {
      for (let x of list) {
          if (x.networkObject.ID == id) {
              return x
          }
      }
  }

  if (data.conData) {
      if (data.conData.conMsg == config.ConnectionMessage.CON_SUCCESS_LOGIN) {
          console.log("logged in")
          //successfull login
      }

      if (data.conData.conMsg == config.ConnectionMessage.ERR_NO_SUCH_CREDENTIALS) {
          //unsuccessfull login
      }

      if (data.conData.conMsg == config.ConnectionMessage.RECONNECT_AVAILABLE) {
          //for reconnect issues
      }
  }


  if (data.actions) {
      for (let x of data.actions) {
          for (let a in config.Actions) {
              if (config.Actions[a] == x.Type)
              console.log("get action", a)
          }

          switch (x.Type) {
              case config.Actions.StartScene:
                  console.log("Scene started");
                //here we handle new scene
                //all start objects here  networkUtil.objectsToSync which is = [data.fishes + data.players]

                  //all scene parameters will be here
                  console.log(data.locationData);

                  //after scene started your player id is = networkUtil
                  console.log("PlayerId is = ", networkUtil.playerId);

              break;
              case config.Actions.Appear: //OBJECT APPEARED (FISH OR PLAYER)
                  let isItFish = objById(data.fishes, x.TargetID);
                  let isItPlayer = objById(data.players, x.TargetID);

                  if (isItFish) {
                      allFishes[x.TargetID] = isItFish;
                      console.log("CurveTime", isItFish.CurveTime / config.Second);
                  }

                  if (isItPlayer) {
                      isItPlayer[x.TargetID] = isItPlayer
                  }

                  //x.TargetID is object ID in
                  //networkUtil.objectsToSync which is = [data.fishes + data.players]
              break;
              case config.Actions.MoneyChange:  //OBJECT REMOVED (FISH OR PLAYER)
                    console.log("player ", x.TargetID, " got money: ", x.Value);

                  break;
              case config.Actions.Remove:  //OBJECT REMOVED (FISH OR PLAYER)
                  allFishes[x.TargetID] = null;
                  allPlayers[x.TargetID] = null;
                  //x.TargetID is object ID in
                  //networkUtil.objectsToSync which is = [data.fishes + data.players]
              break;
              case config.Actions.HealthChange:  //Fish took damage => hp changed
                  console.log("fish hp was: ", allFishes[x.TargetID].Hp);
                  allFishes[x.TargetID].Hp += x.Value;
              break;
              case config.Actions.AngleChange:  //Player changed angle
                let angle1 = x.Value;
                let angle2 = x.Value2;
                console.log("Player" + x.TargetID, " changed angles: ", angle1, " ", angle2);
              break
          }
      }


  }
};

let onOpen = (event, networkUtil) => {
  console.log("websocket opened");
};

let network = new Network(onMessage, onOpen);
network.init();


let update = () =>{
  network.process();
  RAF();
};


let RAF = () => {
  window.requestAnimationFrame(update);
};

RAF();

/////////////////////////////
//mouse move handler
/////////////////////////////

document.addEventListener('mousemove', e => {
    network.sendPlayerCommand(config.Commands.CMD_ANGLE_CHANGE, [e.clientX, e.clientY]);
});

//////////////////////////////
//shooting fishes
//////////////////////////////

setInterval(()=>{
    let ids = Object.keys(allFishes);//ids
    let randomId = ids[Math.floor(Math.random()*ids.length)];


    network.sendPlayerCommand(config.Commands.CMD_SHOOT, [randomId]);
},500);

setTimeout(()=>{

//    network.sendPlayerCommand(config.Commands.CMD_CHANGE_GUN, [4]);
}, 5000);