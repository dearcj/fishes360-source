//import * as $protobuf from "./protobufjs/minimal.js";
let $protobuf = window.protobuf;

// Common aliases
const $Reader = $protobuf.Reader, $Writer = $protobuf.Writer, $util = $protobuf.util;

// Exported root namespace
const $root = $protobuf.roots["default"] || ($protobuf.roots["default"] = {});

export const CustomObject = $root.CustomObject = (() => {

    /**
     * Properties of a CustomObject.
     * @exports ICustomObject
     * @interface ICustomObject
     * @property {INetworkObject|null} [networkObject] CustomObject networkObject
     * @property {number|null} [param1] CustomObject param1
     * @property {number|null} [param2] CustomObject param2
     * @property {number|null} [param3] CustomObject param3
     */

    /**
     * Constructs a new CustomObject.
     * @exports CustomObject
     * @classdesc Represents a CustomObject.
     * @implements ICustomObject
     * @constructor
     * @param {ICustomObject=} [properties] Properties to set
     */
    function CustomObject(properties) {
        if (properties)
            for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                if (properties[keys[i]] != null)
                    this[keys[i]] = properties[keys[i]];
    }

    /**
     * CustomObject networkObject.
     * @member {INetworkObject|null|undefined} networkObject
     * @memberof CustomObject
     * @instance
     */
    CustomObject.prototype.networkObject = null;

    /**
     * CustomObject param1.
     * @member {number} param1
     * @memberof CustomObject
     * @instance
     */
    CustomObject.prototype.param1 = 0;

    /**
     * CustomObject param2.
     * @member {number} param2
     * @memberof CustomObject
     * @instance
     */
    CustomObject.prototype.param2 = 0;

    /**
     * CustomObject param3.
     * @member {number} param3
     * @memberof CustomObject
     * @instance
     */
    CustomObject.prototype.param3 = 0;

    /**
     * Creates a new CustomObject instance using the specified properties.
     * @function create
     * @memberof CustomObject
     * @static
     * @param {ICustomObject=} [properties] Properties to set
     * @returns {CustomObject} CustomObject instance
     */
    CustomObject.create = function create(properties) {
        return new CustomObject(properties);
    };

    /**
     * Encodes the specified CustomObject message. Does not implicitly {@link CustomObject.verify|verify} messages.
     * @function encode
     * @memberof CustomObject
     * @static
     * @param {ICustomObject} message CustomObject message or plain object to encode
     * @param {$protobuf.Writer} [writer] Writer to encode to
     * @returns {$protobuf.Writer} Writer
     */
    CustomObject.encode = function encode(message, writer) {
        if (!writer)
            writer = $Writer.create();
        if (message.networkObject != null && message.hasOwnProperty("networkObject"))
            $root.NetworkObject.encode(message.networkObject, writer.uint32(/* id 1, wireType 2 =*/10).fork()).ldelim();
        if (message.param1 != null && message.hasOwnProperty("param1"))
            writer.uint32(/* id 2, wireType 0 =*/16).int32(message.param1);
        if (message.param2 != null && message.hasOwnProperty("param2"))
            writer.uint32(/* id 3, wireType 0 =*/24).int32(message.param2);
        if (message.param3 != null && message.hasOwnProperty("param3"))
            writer.uint32(/* id 4, wireType 0 =*/32).int32(message.param3);
        return writer;
    };

    /**
     * Encodes the specified CustomObject message, length delimited. Does not implicitly {@link CustomObject.verify|verify} messages.
     * @function encodeDelimited
     * @memberof CustomObject
     * @static
     * @param {ICustomObject} message CustomObject message or plain object to encode
     * @param {$protobuf.Writer} [writer] Writer to encode to
     * @returns {$protobuf.Writer} Writer
     */
    CustomObject.encodeDelimited = function encodeDelimited(message, writer) {
        return this.encode(message, writer).ldelim();
    };

    /**
     * Decodes a CustomObject message from the specified reader or buffer.
     * @function decode
     * @memberof CustomObject
     * @static
     * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
     * @param {number} [length] Message length if known beforehand
     * @returns {CustomObject} CustomObject
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    CustomObject.decode = function decode(reader, length) {
        if (!(reader instanceof $Reader))
            reader = $Reader.create(reader);
        let end = length === undefined ? reader.len : reader.pos + length, message = new $root.CustomObject();
        while (reader.pos < end) {
            let tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.networkObject = $root.NetworkObject.decode(reader, reader.uint32());
                    break;
                case 2:
                    message.param1 = reader.int32();
                    break;
                case 3:
                    message.param2 = reader.int32();
                    break;
                case 4:
                    message.param3 = reader.int32();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    };

    /**
     * Decodes a CustomObject message from the specified reader or buffer, length delimited.
     * @function decodeDelimited
     * @memberof CustomObject
     * @static
     * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
     * @returns {CustomObject} CustomObject
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    CustomObject.decodeDelimited = function decodeDelimited(reader) {
        if (!(reader instanceof $Reader))
            reader = new $Reader(reader);
        return this.decode(reader, reader.uint32());
    };

    /**
     * Verifies a CustomObject message.
     * @function verify
     * @memberof CustomObject
     * @static
     * @param {Object.<string,*>} message Plain object to verify
     * @returns {string|null} `null` if valid, otherwise the reason why it is not
     */
    CustomObject.verify = function verify(message) {
        if (typeof message !== "object" || message === null)
            return "object expected";
        if (message.networkObject != null && message.hasOwnProperty("networkObject")) {
            let error = $root.NetworkObject.verify(message.networkObject);
            if (error)
                return "networkObject." + error;
        }
        if (message.param1 != null && message.hasOwnProperty("param1"))
            if (!$util.isInteger(message.param1))
                return "param1: integer expected";
        if (message.param2 != null && message.hasOwnProperty("param2"))
            if (!$util.isInteger(message.param2))
                return "param2: integer expected";
        if (message.param3 != null && message.hasOwnProperty("param3"))
            if (!$util.isInteger(message.param3))
                return "param3: integer expected";
        return null;
    };

    /**
     * Creates a CustomObject message from a plain object. Also converts values to their respective internal types.
     * @function fromObject
     * @memberof CustomObject
     * @static
     * @param {Object.<string,*>} object Plain object
     * @returns {CustomObject} CustomObject
     */
    CustomObject.fromObject = function fromObject(object) {
        if (object instanceof $root.CustomObject)
            return object;
        let message = new $root.CustomObject();
        if (object.networkObject != null) {
            if (typeof object.networkObject !== "object")
                throw TypeError(".CustomObject.networkObject: object expected");
            message.networkObject = $root.NetworkObject.fromObject(object.networkObject);
        }
        if (object.param1 != null)
            message.param1 = object.param1 | 0;
        if (object.param2 != null)
            message.param2 = object.param2 | 0;
        if (object.param3 != null)
            message.param3 = object.param3 | 0;
        return message;
    };

    /**
     * Creates a plain object from a CustomObject message. Also converts values to other types if specified.
     * @function toObject
     * @memberof CustomObject
     * @static
     * @param {CustomObject} message CustomObject
     * @param {$protobuf.IConversionOptions} [options] Conversion options
     * @returns {Object.<string,*>} Plain object
     */
    CustomObject.toObject = function toObject(message, options) {
        if (!options)
            options = {};
        let object = {};
        if (options.defaults) {
            object.networkObject = null;
            object.param1 = 0;
            object.param2 = 0;
            object.param3 = 0;
        }
        if (message.networkObject != null && message.hasOwnProperty("networkObject"))
            object.networkObject = $root.NetworkObject.toObject(message.networkObject, options);
        if (message.param1 != null && message.hasOwnProperty("param1"))
            object.param1 = message.param1;
        if (message.param2 != null && message.hasOwnProperty("param2"))
            object.param2 = message.param2;
        if (message.param3 != null && message.hasOwnProperty("param3"))
            object.param3 = message.param3;
        return object;
    };

    /**
     * Converts this CustomObject to JSON.
     * @function toJSON
     * @memberof CustomObject
     * @instance
     * @returns {Object.<string,*>} JSON object
     */
    CustomObject.prototype.toJSON = function toJSON() {
        return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
    };

    return CustomObject;
})();

export const Player = $root.Player = (() => {

    /**
     * Properties of a Player.
     * @exports IPlayer
     * @interface IPlayer
     * @property {number|null} [startPosition] Player startPosition
     * @property {string|null} [name] Player name
     * @property {INetworkObject|null} [networkObject] Player networkObject
     */

    /**
     * Constructs a new Player.
     * @exports Player
     * @classdesc Represents a Player.
     * @implements IPlayer
     * @constructor
     * @param {IPlayer=} [properties] Properties to set
     */
    function Player(properties) {
        if (properties)
            for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                if (properties[keys[i]] != null)
                    this[keys[i]] = properties[keys[i]];
    }

    /**
     * Player startPosition.
     * @member {number} startPosition
     * @memberof Player
     * @instance
     */
    Player.prototype.startPosition = 0;

    /**
     * Player name.
     * @member {string} name
     * @memberof Player
     * @instance
     */
    Player.prototype.name = "";

    /**
     * Player networkObject.
     * @member {INetworkObject|null|undefined} networkObject
     * @memberof Player
     * @instance
     */
    Player.prototype.networkObject = null;

    /**
     * Creates a new Player instance using the specified properties.
     * @function create
     * @memberof Player
     * @static
     * @param {IPlayer=} [properties] Properties to set
     * @returns {Player} Player instance
     */
    Player.create = function create(properties) {
        return new Player(properties);
    };

    /**
     * Encodes the specified Player message. Does not implicitly {@link Player.verify|verify} messages.
     * @function encode
     * @memberof Player
     * @static
     * @param {IPlayer} message Player message or plain object to encode
     * @param {$protobuf.Writer} [writer] Writer to encode to
     * @returns {$protobuf.Writer} Writer
     */
    Player.encode = function encode(message, writer) {
        if (!writer)
            writer = $Writer.create();
        if (message.name != null && message.hasOwnProperty("name"))
            writer.uint32(/* id 1, wireType 2 =*/10).string(message.name);
        if (message.networkObject != null && message.hasOwnProperty("networkObject"))
            $root.NetworkObject.encode(message.networkObject, writer.uint32(/* id 2, wireType 2 =*/18).fork()).ldelim();
        if (message.startPosition != null && message.hasOwnProperty("startPosition"))
            writer.uint32(/* id 3, wireType 0 =*/24).uint32(message.startPosition);
        return writer;
    };

    /**
     * Encodes the specified Player message, length delimited. Does not implicitly {@link Player.verify|verify} messages.
     * @function encodeDelimited
     * @memberof Player
     * @static
     * @param {IPlayer} message Player message or plain object to encode
     * @param {$protobuf.Writer} [writer] Writer to encode to
     * @returns {$protobuf.Writer} Writer
     */
    Player.encodeDelimited = function encodeDelimited(message, writer) {
        return this.encode(message, writer).ldelim();
    };

    /**
     * Decodes a Player message from the specified reader or buffer.
     * @function decode
     * @memberof Player
     * @static
     * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
     * @param {number} [length] Message length if known beforehand
     * @returns {Player} Player
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    Player.decode = function decode(reader, length) {
        if (!(reader instanceof $Reader))
            reader = $Reader.create(reader);
        let end = length === undefined ? reader.len : reader.pos + length, message = new $root.Player();
        while (reader.pos < end) {
            let tag = reader.uint32();
            switch (tag >>> 3) {
                case 3:
                    message.startPosition = reader.uint32();
                    break;
                case 1:
                    message.name = reader.string();
                    break;
                case 2:
                    message.networkObject = $root.NetworkObject.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    };

    /**
     * Decodes a Player message from the specified reader or buffer, length delimited.
     * @function decodeDelimited
     * @memberof Player
     * @static
     * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
     * @returns {Player} Player
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    Player.decodeDelimited = function decodeDelimited(reader) {
        if (!(reader instanceof $Reader))
            reader = new $Reader(reader);
        return this.decode(reader, reader.uint32());
    };

    /**
     * Verifies a Player message.
     * @function verify
     * @memberof Player
     * @static
     * @param {Object.<string,*>} message Plain object to verify
     * @returns {string|null} `null` if valid, otherwise the reason why it is not
     */
    Player.verify = function verify(message) {
        if (typeof message !== "object" || message === null)
            return "object expected";
        if (message.startPosition != null && message.hasOwnProperty("startPosition"))
            if (!$util.isInteger(message.startPosition))
                return "startPosition: integer expected";
        if (message.name != null && message.hasOwnProperty("name"))
            if (!$util.isString(message.name))
                return "name: string expected";
        if (message.networkObject != null && message.hasOwnProperty("networkObject")) {
            let error = $root.NetworkObject.verify(message.networkObject);
            if (error)
                return "networkObject." + error;
        }
        return null;
    };

    /**
     * Creates a Player message from a plain object. Also converts values to their respective internal types.
     * @function fromObject
     * @memberof Player
     * @static
     * @param {Object.<string,*>} object Plain object
     * @returns {Player} Player
     */
    Player.fromObject = function fromObject(object) {
        if (object instanceof $root.Player)
            return object;
        let message = new $root.Player();
        if (object.startPosition != null)
            message.startPosition = object.startPosition >>> 0;
        if (object.name != null)
            message.name = String(object.name);
        if (object.networkObject != null) {
            if (typeof object.networkObject !== "object")
                throw TypeError(".Player.networkObject: object expected");
            message.networkObject = $root.NetworkObject.fromObject(object.networkObject);
        }
        return message;
    };

    /**
     * Creates a plain object from a Player message. Also converts values to other types if specified.
     * @function toObject
     * @memberof Player
     * @static
     * @param {Player} message Player
     * @param {$protobuf.IConversionOptions} [options] Conversion options
     * @returns {Object.<string,*>} Plain object
     */
    Player.toObject = function toObject(message, options) {
        if (!options)
            options = {};
        let object = {};
        if (options.defaults) {
            object.name = "";
            object.networkObject = null;
            object.startPosition = 0;
        }
        if (message.name != null && message.hasOwnProperty("name"))
            object.name = message.name;
        if (message.networkObject != null && message.hasOwnProperty("networkObject"))
            object.networkObject = $root.NetworkObject.toObject(message.networkObject, options);
        if (message.startPosition != null && message.hasOwnProperty("startPosition"))
            object.startPosition = message.startPosition;
        return object;
    };

    /**
     * Converts this Player to JSON.
     * @function toJSON
     * @memberof Player
     * @instance
     * @returns {Object.<string,*>} JSON object
     */
    Player.prototype.toJSON = function toJSON() {
        return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
    };

    return Player;
})();

export const Fish = $root.Fish = (() => {

    /**
     * Properties of a Fish.
     * @exports IFish
     * @interface IFish
     * @property {number|null} [Hp] Fish Hp
     * @property {number|null} [Maxhp] Fish Maxhp
     * @property {number|null} [CurveInx] Fish CurveInx
     * @property {number|null} [FishType] Fish FishType
     * @property {number|Long|null} [StartTime] Fish StartTime
     * @property {number|Long|null} [CurveTime] Fish CurveTime
     * @property {boolean|null} [IsBoss] Fish IsBoss
     * @property {INetworkObject|null} [networkObject] Fish networkObject
     */

    /**
     * Constructs a new Fish.
     * @exports Fish
     * @classdesc Represents a Fish.
     * @implements IFish
     * @constructor
     * @param {IFish=} [properties] Properties to set
     */
    function Fish(properties) {
        if (properties)
            for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                if (properties[keys[i]] != null)
                    this[keys[i]] = properties[keys[i]];
    }

    /**
     * Fish Hp.
     * @member {number} Hp
     * @memberof Fish
     * @instance
     */
    Fish.prototype.Hp = 0;

    /**
     * Fish Maxhp.
     * @member {number} Maxhp
     * @memberof Fish
     * @instance
     */
    Fish.prototype.Maxhp = 0;

    /**
     * Fish CurveInx.
     * @member {number} CurveInx
     * @memberof Fish
     * @instance
     */
    Fish.prototype.CurveInx = 0;

    /**
     * Fish FishType.
     * @member {number} FishType
     * @memberof Fish
     * @instance
     */
    Fish.prototype.FishType = 0;

    /**
     * Fish StartTime.
     * @member {number|Long} StartTime
     * @memberof Fish
     * @instance
     */
    Fish.prototype.StartTime = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

    /**
     * Fish CurveTime.
     * @member {number|Long} CurveTime
     * @memberof Fish
     * @instance
     */
    Fish.prototype.CurveTime = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

    /**
     * Fish IsBoss.
     * @member {boolean} IsBoss
     * @memberof Fish
     * @instance
     */
    Fish.prototype.IsBoss = false;

    /**
     * Fish networkObject.
     * @member {INetworkObject|null|undefined} networkObject
     * @memberof Fish
     * @instance
     */
    Fish.prototype.networkObject = null;

    /**
     * Creates a new Fish instance using the specified properties.
     * @function create
     * @memberof Fish
     * @static
     * @param {IFish=} [properties] Properties to set
     * @returns {Fish} Fish instance
     */
    Fish.create = function create(properties) {
        return new Fish(properties);
    };

    /**
     * Encodes the specified Fish message. Does not implicitly {@link Fish.verify|verify} messages.
     * @function encode
     * @memberof Fish
     * @static
     * @param {IFish} message Fish message or plain object to encode
     * @param {$protobuf.Writer} [writer] Writer to encode to
     * @returns {$protobuf.Writer} Writer
     */
    Fish.encode = function encode(message, writer) {
        if (!writer)
            writer = $Writer.create();
        if (message.Hp != null && message.hasOwnProperty("Hp"))
            writer.uint32(/* id 1, wireType 0 =*/8).uint32(message.Hp);
        if (message.Maxhp != null && message.hasOwnProperty("Maxhp"))
            writer.uint32(/* id 2, wireType 0 =*/16).uint32(message.Maxhp);
        if (message.FishType != null && message.hasOwnProperty("FishType"))
            writer.uint32(/* id 3, wireType 0 =*/24).uint32(message.FishType);
        if (message.networkObject != null && message.hasOwnProperty("networkObject"))
            $root.NetworkObject.encode(message.networkObject, writer.uint32(/* id 4, wireType 2 =*/34).fork()).ldelim();
        if (message.CurveTime != null && message.hasOwnProperty("CurveTime"))
            writer.uint32(/* id 5, wireType 0 =*/40).uint64(message.CurveTime);
        if (message.StartTime != null && message.hasOwnProperty("StartTime"))
            writer.uint32(/* id 6, wireType 0 =*/48).uint64(message.StartTime);
        if (message.CurveInx != null && message.hasOwnProperty("CurveInx"))
            writer.uint32(/* id 7, wireType 0 =*/56).uint32(message.CurveInx);
        if (message.IsBoss != null && message.hasOwnProperty("IsBoss"))
            writer.uint32(/* id 8, wireType 0 =*/64).bool(message.IsBoss);
        return writer;
    };

    /**
     * Encodes the specified Fish message, length delimited. Does not implicitly {@link Fish.verify|verify} messages.
     * @function encodeDelimited
     * @memberof Fish
     * @static
     * @param {IFish} message Fish message or plain object to encode
     * @param {$protobuf.Writer} [writer] Writer to encode to
     * @returns {$protobuf.Writer} Writer
     */
    Fish.encodeDelimited = function encodeDelimited(message, writer) {
        return this.encode(message, writer).ldelim();
    };

    /**
     * Decodes a Fish message from the specified reader or buffer.
     * @function decode
     * @memberof Fish
     * @static
     * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
     * @param {number} [length] Message length if known beforehand
     * @returns {Fish} Fish
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    Fish.decode = function decode(reader, length) {
        if (!(reader instanceof $Reader))
            reader = $Reader.create(reader);
        let end = length === undefined ? reader.len : reader.pos + length, message = new $root.Fish();
        while (reader.pos < end) {
            let tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.Hp = reader.uint32();
                    break;
                case 2:
                    message.Maxhp = reader.uint32();
                    break;
                case 7:
                    message.CurveInx = reader.uint32();
                    break;
                case 3:
                    message.FishType = reader.uint32();
                    break;
                case 6:
                    message.StartTime = reader.uint64();
                    break;
                case 5:
                    message.CurveTime = reader.uint64();
                    break;
                case 8:
                    message.IsBoss = reader.bool();
                    break;
                case 4:
                    message.networkObject = $root.NetworkObject.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    };

    /**
     * Decodes a Fish message from the specified reader or buffer, length delimited.
     * @function decodeDelimited
     * @memberof Fish
     * @static
     * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
     * @returns {Fish} Fish
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    Fish.decodeDelimited = function decodeDelimited(reader) {
        if (!(reader instanceof $Reader))
            reader = new $Reader(reader);
        return this.decode(reader, reader.uint32());
    };

    /**
     * Verifies a Fish message.
     * @function verify
     * @memberof Fish
     * @static
     * @param {Object.<string,*>} message Plain object to verify
     * @returns {string|null} `null` if valid, otherwise the reason why it is not
     */
    Fish.verify = function verify(message) {
        if (typeof message !== "object" || message === null)
            return "object expected";
        if (message.Hp != null && message.hasOwnProperty("Hp"))
            if (!$util.isInteger(message.Hp))
                return "Hp: integer expected";
        if (message.Maxhp != null && message.hasOwnProperty("Maxhp"))
            if (!$util.isInteger(message.Maxhp))
                return "Maxhp: integer expected";
        if (message.CurveInx != null && message.hasOwnProperty("CurveInx"))
            if (!$util.isInteger(message.CurveInx))
                return "CurveInx: integer expected";
        if (message.FishType != null && message.hasOwnProperty("FishType"))
            if (!$util.isInteger(message.FishType))
                return "FishType: integer expected";
        if (message.StartTime != null && message.hasOwnProperty("StartTime"))
            if (!$util.isInteger(message.StartTime) && !(message.StartTime && $util.isInteger(message.StartTime.low) && $util.isInteger(message.StartTime.high)))
                return "StartTime: integer|Long expected";
        if (message.CurveTime != null && message.hasOwnProperty("CurveTime"))
            if (!$util.isInteger(message.CurveTime) && !(message.CurveTime && $util.isInteger(message.CurveTime.low) && $util.isInteger(message.CurveTime.high)))
                return "CurveTime: integer|Long expected";
        if (message.IsBoss != null && message.hasOwnProperty("IsBoss"))
            if (typeof message.IsBoss !== "boolean")
                return "IsBoss: boolean expected";
        if (message.networkObject != null && message.hasOwnProperty("networkObject")) {
            let error = $root.NetworkObject.verify(message.networkObject);
            if (error)
                return "networkObject." + error;
        }
        return null;
    };

    /**
     * Creates a Fish message from a plain object. Also converts values to their respective internal types.
     * @function fromObject
     * @memberof Fish
     * @static
     * @param {Object.<string,*>} object Plain object
     * @returns {Fish} Fish
     */
    Fish.fromObject = function fromObject(object) {
        if (object instanceof $root.Fish)
            return object;
        let message = new $root.Fish();
        if (object.Hp != null)
            message.Hp = object.Hp >>> 0;
        if (object.Maxhp != null)
            message.Maxhp = object.Maxhp >>> 0;
        if (object.CurveInx != null)
            message.CurveInx = object.CurveInx >>> 0;
        if (object.FishType != null)
            message.FishType = object.FishType >>> 0;
        if (object.StartTime != null)
            if ($util.Long)
                (message.StartTime = $util.Long.fromValue(object.StartTime)).unsigned = true;
            else if (typeof object.StartTime === "string")
                message.StartTime = parseInt(object.StartTime, 10);
            else if (typeof object.StartTime === "number")
                message.StartTime = object.StartTime;
            else if (typeof object.StartTime === "object")
                message.StartTime = new $util.LongBits(object.StartTime.low >>> 0, object.StartTime.high >>> 0).toNumber(true);
        if (object.CurveTime != null)
            if ($util.Long)
                (message.CurveTime = $util.Long.fromValue(object.CurveTime)).unsigned = true;
            else if (typeof object.CurveTime === "string")
                message.CurveTime = parseInt(object.CurveTime, 10);
            else if (typeof object.CurveTime === "number")
                message.CurveTime = object.CurveTime;
            else if (typeof object.CurveTime === "object")
                message.CurveTime = new $util.LongBits(object.CurveTime.low >>> 0, object.CurveTime.high >>> 0).toNumber(true);
        if (object.IsBoss != null)
            message.IsBoss = Boolean(object.IsBoss);
        if (object.networkObject != null) {
            if (typeof object.networkObject !== "object")
                throw TypeError(".Fish.networkObject: object expected");
            message.networkObject = $root.NetworkObject.fromObject(object.networkObject);
        }
        return message;
    };

    /**
     * Creates a plain object from a Fish message. Also converts values to other types if specified.
     * @function toObject
     * @memberof Fish
     * @static
     * @param {Fish} message Fish
     * @param {$protobuf.IConversionOptions} [options] Conversion options
     * @returns {Object.<string,*>} Plain object
     */
    Fish.toObject = function toObject(message, options) {
        if (!options)
            options = {};
        let object = {};
        if (options.defaults) {
            object.Hp = 0;
            object.Maxhp = 0;
            object.FishType = 0;
            object.networkObject = null;
            if ($util.Long) {
                let long = new $util.Long(0, 0, true);
                object.CurveTime = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
            } else
                object.CurveTime = options.longs === String ? "0" : 0;
            if ($util.Long) {
                let long = new $util.Long(0, 0, true);
                object.StartTime = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
            } else
                object.StartTime = options.longs === String ? "0" : 0;
            object.CurveInx = 0;
            object.IsBoss = false;
        }
        if (message.Hp != null && message.hasOwnProperty("Hp"))
            object.Hp = message.Hp;
        if (message.Maxhp != null && message.hasOwnProperty("Maxhp"))
            object.Maxhp = message.Maxhp;
        if (message.FishType != null && message.hasOwnProperty("FishType"))
            object.FishType = message.FishType;
        if (message.networkObject != null && message.hasOwnProperty("networkObject"))
            object.networkObject = $root.NetworkObject.toObject(message.networkObject, options);
        if (message.CurveTime != null && message.hasOwnProperty("CurveTime"))
            if (typeof message.CurveTime === "number")
                object.CurveTime = options.longs === String ? String(message.CurveTime) : message.CurveTime;
            else
                object.CurveTime = options.longs === String ? $util.Long.prototype.toString.call(message.CurveTime) : options.longs === Number ? new $util.LongBits(message.CurveTime.low >>> 0, message.CurveTime.high >>> 0).toNumber(true) : message.CurveTime;
        if (message.StartTime != null && message.hasOwnProperty("StartTime"))
            if (typeof message.StartTime === "number")
                object.StartTime = options.longs === String ? String(message.StartTime) : message.StartTime;
            else
                object.StartTime = options.longs === String ? $util.Long.prototype.toString.call(message.StartTime) : options.longs === Number ? new $util.LongBits(message.StartTime.low >>> 0, message.StartTime.high >>> 0).toNumber(true) : message.StartTime;
        if (message.CurveInx != null && message.hasOwnProperty("CurveInx"))
            object.CurveInx = message.CurveInx;
        if (message.IsBoss != null && message.hasOwnProperty("IsBoss"))
            object.IsBoss = message.IsBoss;
        return object;
    };

    /**
     * Converts this Fish to JSON.
     * @function toJSON
     * @memberof Fish
     * @instance
     * @returns {Object.<string,*>} JSON object
     */
    Fish.prototype.toJSON = function toJSON() {
        return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
    };

    return Fish;
})();

export const NetworkObject = $root.NetworkObject = (() => {

    /**
     * Properties of a NetworkObject.
     * @exports INetworkObject
     * @interface INetworkObject
     * @property {number|null} [ID] NetworkObject ID
     * @property {number|null} [Type] NetworkObject Type
     */

    /**
     * Constructs a new NetworkObject.
     * @exports NetworkObject
     * @classdesc Represents a NetworkObject.
     * @implements INetworkObject
     * @constructor
     * @param {INetworkObject=} [properties] Properties to set
     */
    function NetworkObject(properties) {
        if (properties)
            for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                if (properties[keys[i]] != null)
                    this[keys[i]] = properties[keys[i]];
    }

    /**
     * NetworkObject ID.
     * @member {number} ID
     * @memberof NetworkObject
     * @instance
     */
    NetworkObject.prototype.ID = 0;

    /**
     * NetworkObject Type.
     * @member {number} Type
     * @memberof NetworkObject
     * @instance
     */
    NetworkObject.prototype.Type = 0;

    /**
     * Creates a new NetworkObject instance using the specified properties.
     * @function create
     * @memberof NetworkObject
     * @static
     * @param {INetworkObject=} [properties] Properties to set
     * @returns {NetworkObject} NetworkObject instance
     */
    NetworkObject.create = function create(properties) {
        return new NetworkObject(properties);
    };

    /**
     * Encodes the specified NetworkObject message. Does not implicitly {@link NetworkObject.verify|verify} messages.
     * @function encode
     * @memberof NetworkObject
     * @static
     * @param {INetworkObject} message NetworkObject message or plain object to encode
     * @param {$protobuf.Writer} [writer] Writer to encode to
     * @returns {$protobuf.Writer} Writer
     */
    NetworkObject.encode = function encode(message, writer) {
        if (!writer)
            writer = $Writer.create();
        if (message.ID != null && message.hasOwnProperty("ID"))
            writer.uint32(/* id 1, wireType 0 =*/8).uint32(message.ID);
        if (message.Type != null && message.hasOwnProperty("Type"))
            writer.uint32(/* id 2, wireType 0 =*/16).uint32(message.Type);
        return writer;
    };

    /**
     * Encodes the specified NetworkObject message, length delimited. Does not implicitly {@link NetworkObject.verify|verify} messages.
     * @function encodeDelimited
     * @memberof NetworkObject
     * @static
     * @param {INetworkObject} message NetworkObject message or plain object to encode
     * @param {$protobuf.Writer} [writer] Writer to encode to
     * @returns {$protobuf.Writer} Writer
     */
    NetworkObject.encodeDelimited = function encodeDelimited(message, writer) {
        return this.encode(message, writer).ldelim();
    };

    /**
     * Decodes a NetworkObject message from the specified reader or buffer.
     * @function decode
     * @memberof NetworkObject
     * @static
     * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
     * @param {number} [length] Message length if known beforehand
     * @returns {NetworkObject} NetworkObject
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    NetworkObject.decode = function decode(reader, length) {
        if (!(reader instanceof $Reader))
            reader = $Reader.create(reader);
        let end = length === undefined ? reader.len : reader.pos + length, message = new $root.NetworkObject();
        while (reader.pos < end) {
            let tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.ID = reader.uint32();
                    break;
                case 2:
                    message.Type = reader.uint32();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    };

    /**
     * Decodes a NetworkObject message from the specified reader or buffer, length delimited.
     * @function decodeDelimited
     * @memberof NetworkObject
     * @static
     * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
     * @returns {NetworkObject} NetworkObject
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    NetworkObject.decodeDelimited = function decodeDelimited(reader) {
        if (!(reader instanceof $Reader))
            reader = new $Reader(reader);
        return this.decode(reader, reader.uint32());
    };

    /**
     * Verifies a NetworkObject message.
     * @function verify
     * @memberof NetworkObject
     * @static
     * @param {Object.<string,*>} message Plain object to verify
     * @returns {string|null} `null` if valid, otherwise the reason why it is not
     */
    NetworkObject.verify = function verify(message) {
        if (typeof message !== "object" || message === null)
            return "object expected";
        if (message.ID != null && message.hasOwnProperty("ID"))
            if (!$util.isInteger(message.ID))
                return "ID: integer expected";
        if (message.Type != null && message.hasOwnProperty("Type"))
            if (!$util.isInteger(message.Type))
                return "Type: integer expected";
        return null;
    };

    /**
     * Creates a NetworkObject message from a plain object. Also converts values to their respective internal types.
     * @function fromObject
     * @memberof NetworkObject
     * @static
     * @param {Object.<string,*>} object Plain object
     * @returns {NetworkObject} NetworkObject
     */
    NetworkObject.fromObject = function fromObject(object) {
        if (object instanceof $root.NetworkObject)
            return object;
        let message = new $root.NetworkObject();
        if (object.ID != null)
            message.ID = object.ID >>> 0;
        if (object.Type != null)
            message.Type = object.Type >>> 0;
        return message;
    };

    /**
     * Creates a plain object from a NetworkObject message. Also converts values to other types if specified.
     * @function toObject
     * @memberof NetworkObject
     * @static
     * @param {NetworkObject} message NetworkObject
     * @param {$protobuf.IConversionOptions} [options] Conversion options
     * @returns {Object.<string,*>} Plain object
     */
    NetworkObject.toObject = function toObject(message, options) {
        if (!options)
            options = {};
        let object = {};
        if (options.defaults) {
            object.ID = 0;
            object.Type = 0;
        }
        if (message.ID != null && message.hasOwnProperty("ID"))
            object.ID = message.ID;
        if (message.Type != null && message.hasOwnProperty("Type"))
            object.Type = message.Type;
        return object;
    };

    /**
     * Converts this NetworkObject to JSON.
     * @function toJSON
     * @memberof NetworkObject
     * @instance
     * @returns {Object.<string,*>} JSON object
     */
    NetworkObject.prototype.toJSON = function toJSON() {
        return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
    };

    return NetworkObject;
})();

/**
 * ActionType enum.
 * @exports ActionType
 * @enum {string}
 * @property {number} ANY_VALUE=0 ANY_VALUE value
 */
$root.ActionType = (function() {
    const valuesById = {}, values = Object.create(valuesById);
    values[valuesById[0] = "ANY_VALUE"] = 0;
    return values;
})();

export const Action = $root.Action = (() => {

    /**
     * Properties of an Action.
     * @exports IAction
     * @interface IAction
     * @property {ActionType|null} [Type] Action Type
     * @property {number|null} [Value] Action Value
     * @property {number|null} [Value2] Action Value2
     * @property {number|null} [TargetID] Action TargetID
     */

    /**
     * Constructs a new Action.
     * @exports Action
     * @classdesc Represents an Action.
     * @implements IAction
     * @constructor
     * @param {IAction=} [properties] Properties to set
     */
    function Action(properties) {
        if (properties)
            for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                if (properties[keys[i]] != null)
                    this[keys[i]] = properties[keys[i]];
    }

    /**
     * Action Type.
     * @member {ActionType} Type
     * @memberof Action
     * @instance
     */
    Action.prototype.Type = 0;

    /**
     * Action Value.
     * @member {number} Value
     * @memberof Action
     * @instance
     */
    Action.prototype.Value = 0;

    /**
     * Action Value2.
     * @member {number} Value2
     * @memberof Action
     * @instance
     */
    Action.prototype.Value2 = 0;

    /**
     * Action TargetID.
     * @member {number} TargetID
     * @memberof Action
     * @instance
     */
    Action.prototype.TargetID = 0;

    /**
     * Creates a new Action instance using the specified properties.
     * @function create
     * @memberof Action
     * @static
     * @param {IAction=} [properties] Properties to set
     * @returns {Action} Action instance
     */
    Action.create = function create(properties) {
        return new Action(properties);
    };

    /**
     * Encodes the specified Action message. Does not implicitly {@link Action.verify|verify} messages.
     * @function encode
     * @memberof Action
     * @static
     * @param {IAction} message Action message or plain object to encode
     * @param {$protobuf.Writer} [writer] Writer to encode to
     * @returns {$protobuf.Writer} Writer
     */
    Action.encode = function encode(message, writer) {
        if (!writer)
            writer = $Writer.create();
        if (message.Type != null && message.hasOwnProperty("Type"))
            writer.uint32(/* id 1, wireType 0 =*/8).int32(message.Type);
        if (message.Value != null && message.hasOwnProperty("Value"))
            writer.uint32(/* id 2, wireType 5 =*/21).float(message.Value);
        if (message.TargetID != null && message.hasOwnProperty("TargetID"))
            writer.uint32(/* id 3, wireType 0 =*/24).uint32(message.TargetID);
        if (message.Value2 != null && message.hasOwnProperty("Value2"))
            writer.uint32(/* id 4, wireType 5 =*/37).float(message.Value2);
        return writer;
    };

    /**
     * Encodes the specified Action message, length delimited. Does not implicitly {@link Action.verify|verify} messages.
     * @function encodeDelimited
     * @memberof Action
     * @static
     * @param {IAction} message Action message or plain object to encode
     * @param {$protobuf.Writer} [writer] Writer to encode to
     * @returns {$protobuf.Writer} Writer
     */
    Action.encodeDelimited = function encodeDelimited(message, writer) {
        return this.encode(message, writer).ldelim();
    };

    /**
     * Decodes an Action message from the specified reader or buffer.
     * @function decode
     * @memberof Action
     * @static
     * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
     * @param {number} [length] Message length if known beforehand
     * @returns {Action} Action
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    Action.decode = function decode(reader, length) {
        if (!(reader instanceof $Reader))
            reader = $Reader.create(reader);
        let end = length === undefined ? reader.len : reader.pos + length, message = new $root.Action();
        while (reader.pos < end) {
            let tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.Type = reader.int32();
                    break;
                case 2:
                    message.Value = reader.float();
                    break;
                case 4:
                    message.Value2 = reader.float();
                    break;
                case 3:
                    message.TargetID = reader.uint32();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    };

    /**
     * Decodes an Action message from the specified reader or buffer, length delimited.
     * @function decodeDelimited
     * @memberof Action
     * @static
     * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
     * @returns {Action} Action
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    Action.decodeDelimited = function decodeDelimited(reader) {
        if (!(reader instanceof $Reader))
            reader = new $Reader(reader);
        return this.decode(reader, reader.uint32());
    };

    /**
     * Verifies an Action message.
     * @function verify
     * @memberof Action
     * @static
     * @param {Object.<string,*>} message Plain object to verify
     * @returns {string|null} `null` if valid, otherwise the reason why it is not
     */
    Action.verify = function verify(message) {
        if (typeof message !== "object" || message === null)
            return "object expected";
        if (message.Type != null && message.hasOwnProperty("Type"))
            switch (message.Type) {
                default:
                    return "Type: enum value expected";
                case 0:
                    break;
            }
        if (message.Value != null && message.hasOwnProperty("Value"))
            if (typeof message.Value !== "number")
                return "Value: number expected";
        if (message.Value2 != null && message.hasOwnProperty("Value2"))
            if (typeof message.Value2 !== "number")
                return "Value2: number expected";
        if (message.TargetID != null && message.hasOwnProperty("TargetID"))
            if (!$util.isInteger(message.TargetID))
                return "TargetID: integer expected";
        return null;
    };

    /**
     * Creates an Action message from a plain object. Also converts values to their respective internal types.
     * @function fromObject
     * @memberof Action
     * @static
     * @param {Object.<string,*>} object Plain object
     * @returns {Action} Action
     */
    Action.fromObject = function fromObject(object) {
        if (object instanceof $root.Action)
            return object;
        let message = new $root.Action();
        switch (object.Type) {
            case "ANY_VALUE":
            case 0:
                message.Type = 0;
                break;
        }
        if (object.Value != null)
            message.Value = Number(object.Value);
        if (object.Value2 != null)
            message.Value2 = Number(object.Value2);
        if (object.TargetID != null)
            message.TargetID = object.TargetID >>> 0;
        return message;
    };

    /**
     * Creates a plain object from an Action message. Also converts values to other types if specified.
     * @function toObject
     * @memberof Action
     * @static
     * @param {Action} message Action
     * @param {$protobuf.IConversionOptions} [options] Conversion options
     * @returns {Object.<string,*>} Plain object
     */
    Action.toObject = function toObject(message, options) {
        if (!options)
            options = {};
        let object = {};
        if (options.defaults) {
            object.Type = options.enums === String ? "ANY_VALUE" : 0;
            object.Value = 0;
            object.TargetID = 0;
            object.Value2 = 0;
        }
        if (message.Type != null && message.hasOwnProperty("Type"))
            object.Type = options.enums === String ? $root.ActionType[message.Type] : message.Type;
        if (message.Value != null && message.hasOwnProperty("Value"))
            object.Value = options.json && !isFinite(message.Value) ? String(message.Value) : message.Value;
        if (message.TargetID != null && message.hasOwnProperty("TargetID"))
            object.TargetID = message.TargetID;
        if (message.Value2 != null && message.hasOwnProperty("Value2"))
            object.Value2 = options.json && !isFinite(message.Value2) ? String(message.Value2) : message.Value2;
        return object;
    };

    /**
     * Converts this Action to JSON.
     * @function toJSON
     * @memberof Action
     * @instance
     * @returns {Object.<string,*>} JSON object
     */
    Action.prototype.toJSON = function toJSON() {
        return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
    };

    return Action;
})();

export const ConnectionData = $root.ConnectionData = (() => {

    /**
     * Properties of a ConnectionData.
     * @exports IConnectionData
     * @interface IConnectionData
     * @property {number|null} [roomID] ConnectionData roomID
     * @property {number|null} [playerID] ConnectionData playerID
     * @property {number|null} [conMsg] ConnectionData conMsg
     */

    /**
     * Constructs a new ConnectionData.
     * @exports ConnectionData
     * @classdesc Represents a ConnectionData.
     * @implements IConnectionData
     * @constructor
     * @param {IConnectionData=} [properties] Properties to set
     */
    function ConnectionData(properties) {
        if (properties)
            for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                if (properties[keys[i]] != null)
                    this[keys[i]] = properties[keys[i]];
    }

    /**
     * ConnectionData roomID.
     * @member {number} roomID
     * @memberof ConnectionData
     * @instance
     */
    ConnectionData.prototype.roomID = 0;

    /**
     * ConnectionData playerID.
     * @member {number} playerID
     * @memberof ConnectionData
     * @instance
     */
    ConnectionData.prototype.playerID = 0;

    /**
     * ConnectionData conMsg.
     * @member {number} conMsg
     * @memberof ConnectionData
     * @instance
     */
    ConnectionData.prototype.conMsg = 0;

    /**
     * Creates a new ConnectionData instance using the specified properties.
     * @function create
     * @memberof ConnectionData
     * @static
     * @param {IConnectionData=} [properties] Properties to set
     * @returns {ConnectionData} ConnectionData instance
     */
    ConnectionData.create = function create(properties) {
        return new ConnectionData(properties);
    };

    /**
     * Encodes the specified ConnectionData message. Does not implicitly {@link ConnectionData.verify|verify} messages.
     * @function encode
     * @memberof ConnectionData
     * @static
     * @param {IConnectionData} message ConnectionData message or plain object to encode
     * @param {$protobuf.Writer} [writer] Writer to encode to
     * @returns {$protobuf.Writer} Writer
     */
    ConnectionData.encode = function encode(message, writer) {
        if (!writer)
            writer = $Writer.create();
        if (message.roomID != null && message.hasOwnProperty("roomID"))
            writer.uint32(/* id 1, wireType 0 =*/8).uint32(message.roomID);
        if (message.playerID != null && message.hasOwnProperty("playerID"))
            writer.uint32(/* id 2, wireType 0 =*/16).uint32(message.playerID);
        if (message.conMsg != null && message.hasOwnProperty("conMsg"))
            writer.uint32(/* id 3, wireType 0 =*/24).uint32(message.conMsg);
        return writer;
    };

    /**
     * Encodes the specified ConnectionData message, length delimited. Does not implicitly {@link ConnectionData.verify|verify} messages.
     * @function encodeDelimited
     * @memberof ConnectionData
     * @static
     * @param {IConnectionData} message ConnectionData message or plain object to encode
     * @param {$protobuf.Writer} [writer] Writer to encode to
     * @returns {$protobuf.Writer} Writer
     */
    ConnectionData.encodeDelimited = function encodeDelimited(message, writer) {
        return this.encode(message, writer).ldelim();
    };

    /**
     * Decodes a ConnectionData message from the specified reader or buffer.
     * @function decode
     * @memberof ConnectionData
     * @static
     * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
     * @param {number} [length] Message length if known beforehand
     * @returns {ConnectionData} ConnectionData
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    ConnectionData.decode = function decode(reader, length) {
        if (!(reader instanceof $Reader))
            reader = $Reader.create(reader);
        let end = length === undefined ? reader.len : reader.pos + length, message = new $root.ConnectionData();
        while (reader.pos < end) {
            let tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.roomID = reader.uint32();
                    break;
                case 2:
                    message.playerID = reader.uint32();
                    break;
                case 3:
                    message.conMsg = reader.uint32();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    };

    /**
     * Decodes a ConnectionData message from the specified reader or buffer, length delimited.
     * @function decodeDelimited
     * @memberof ConnectionData
     * @static
     * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
     * @returns {ConnectionData} ConnectionData
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    ConnectionData.decodeDelimited = function decodeDelimited(reader) {
        if (!(reader instanceof $Reader))
            reader = new $Reader(reader);
        return this.decode(reader, reader.uint32());
    };

    /**
     * Verifies a ConnectionData message.
     * @function verify
     * @memberof ConnectionData
     * @static
     * @param {Object.<string,*>} message Plain object to verify
     * @returns {string|null} `null` if valid, otherwise the reason why it is not
     */
    ConnectionData.verify = function verify(message) {
        if (typeof message !== "object" || message === null)
            return "object expected";
        if (message.roomID != null && message.hasOwnProperty("roomID"))
            if (!$util.isInteger(message.roomID))
                return "roomID: integer expected";
        if (message.playerID != null && message.hasOwnProperty("playerID"))
            if (!$util.isInteger(message.playerID))
                return "playerID: integer expected";
        if (message.conMsg != null && message.hasOwnProperty("conMsg"))
            if (!$util.isInteger(message.conMsg))
                return "conMsg: integer expected";
        return null;
    };

    /**
     * Creates a ConnectionData message from a plain object. Also converts values to their respective internal types.
     * @function fromObject
     * @memberof ConnectionData
     * @static
     * @param {Object.<string,*>} object Plain object
     * @returns {ConnectionData} ConnectionData
     */
    ConnectionData.fromObject = function fromObject(object) {
        if (object instanceof $root.ConnectionData)
            return object;
        let message = new $root.ConnectionData();
        if (object.roomID != null)
            message.roomID = object.roomID >>> 0;
        if (object.playerID != null)
            message.playerID = object.playerID >>> 0;
        if (object.conMsg != null)
            message.conMsg = object.conMsg >>> 0;
        return message;
    };

    /**
     * Creates a plain object from a ConnectionData message. Also converts values to other types if specified.
     * @function toObject
     * @memberof ConnectionData
     * @static
     * @param {ConnectionData} message ConnectionData
     * @param {$protobuf.IConversionOptions} [options] Conversion options
     * @returns {Object.<string,*>} Plain object
     */
    ConnectionData.toObject = function toObject(message, options) {
        if (!options)
            options = {};
        let object = {};
        if (options.defaults) {
            object.roomID = 0;
            object.playerID = 0;
            object.conMsg = 0;
        }
        if (message.roomID != null && message.hasOwnProperty("roomID"))
            object.roomID = message.roomID;
        if (message.playerID != null && message.hasOwnProperty("playerID"))
            object.playerID = message.playerID;
        if (message.conMsg != null && message.hasOwnProperty("conMsg"))
            object.conMsg = message.conMsg;
        return object;
    };

    /**
     * Converts this ConnectionData to JSON.
     * @function toJSON
     * @memberof ConnectionData
     * @instance
     * @returns {Object.<string,*>} JSON object
     */
    ConnectionData.prototype.toJSON = function toJSON() {
        return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
    };

    return ConnectionData;
})();

export const LocationData = $root.LocationData = (() => {

    /**
     * Properties of a LocationData.
     * @exports ILocationData
     * @interface ILocationData
     * @property {string|null} [locationName] LocationData locationName
     * @property {number|null} [roomId] LocationData roomId
     */

    /**
     * Constructs a new LocationData.
     * @exports LocationData
     * @classdesc Represents a LocationData.
     * @implements ILocationData
     * @constructor
     * @param {ILocationData=} [properties] Properties to set
     */
    function LocationData(properties) {
        if (properties)
            for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                if (properties[keys[i]] != null)
                    this[keys[i]] = properties[keys[i]];
    }

    /**
     * LocationData locationName.
     * @member {string} locationName
     * @memberof LocationData
     * @instance
     */
    LocationData.prototype.locationName = "";

    /**
     * LocationData roomId.
     * @member {number} roomId
     * @memberof LocationData
     * @instance
     */
    LocationData.prototype.roomId = 0;

    /**
     * Creates a new LocationData instance using the specified properties.
     * @function create
     * @memberof LocationData
     * @static
     * @param {ILocationData=} [properties] Properties to set
     * @returns {LocationData} LocationData instance
     */
    LocationData.create = function create(properties) {
        return new LocationData(properties);
    };

    /**
     * Encodes the specified LocationData message. Does not implicitly {@link LocationData.verify|verify} messages.
     * @function encode
     * @memberof LocationData
     * @static
     * @param {ILocationData} message LocationData message or plain object to encode
     * @param {$protobuf.Writer} [writer] Writer to encode to
     * @returns {$protobuf.Writer} Writer
     */
    LocationData.encode = function encode(message, writer) {
        if (!writer)
            writer = $Writer.create();
        if (message.locationName != null && message.hasOwnProperty("locationName"))
            writer.uint32(/* id 2, wireType 2 =*/18).string(message.locationName);
        if (message.roomId != null && message.hasOwnProperty("roomId"))
            writer.uint32(/* id 3, wireType 0 =*/24).uint32(message.roomId);
        return writer;
    };

    /**
     * Encodes the specified LocationData message, length delimited. Does not implicitly {@link LocationData.verify|verify} messages.
     * @function encodeDelimited
     * @memberof LocationData
     * @static
     * @param {ILocationData} message LocationData message or plain object to encode
     * @param {$protobuf.Writer} [writer] Writer to encode to
     * @returns {$protobuf.Writer} Writer
     */
    LocationData.encodeDelimited = function encodeDelimited(message, writer) {
        return this.encode(message, writer).ldelim();
    };

    /**
     * Decodes a LocationData message from the specified reader or buffer.
     * @function decode
     * @memberof LocationData
     * @static
     * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
     * @param {number} [length] Message length if known beforehand
     * @returns {LocationData} LocationData
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    LocationData.decode = function decode(reader, length) {
        if (!(reader instanceof $Reader))
            reader = $Reader.create(reader);
        let end = length === undefined ? reader.len : reader.pos + length, message = new $root.LocationData();
        while (reader.pos < end) {
            let tag = reader.uint32();
            switch (tag >>> 3) {
                case 2:
                    message.locationName = reader.string();
                    break;
                case 3:
                    message.roomId = reader.uint32();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    };

    /**
     * Decodes a LocationData message from the specified reader or buffer, length delimited.
     * @function decodeDelimited
     * @memberof LocationData
     * @static
     * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
     * @returns {LocationData} LocationData
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    LocationData.decodeDelimited = function decodeDelimited(reader) {
        if (!(reader instanceof $Reader))
            reader = new $Reader(reader);
        return this.decode(reader, reader.uint32());
    };

    /**
     * Verifies a LocationData message.
     * @function verify
     * @memberof LocationData
     * @static
     * @param {Object.<string,*>} message Plain object to verify
     * @returns {string|null} `null` if valid, otherwise the reason why it is not
     */
    LocationData.verify = function verify(message) {
        if (typeof message !== "object" || message === null)
            return "object expected";
        if (message.locationName != null && message.hasOwnProperty("locationName"))
            if (!$util.isString(message.locationName))
                return "locationName: string expected";
        if (message.roomId != null && message.hasOwnProperty("roomId"))
            if (!$util.isInteger(message.roomId))
                return "roomId: integer expected";
        return null;
    };

    /**
     * Creates a LocationData message from a plain object. Also converts values to their respective internal types.
     * @function fromObject
     * @memberof LocationData
     * @static
     * @param {Object.<string,*>} object Plain object
     * @returns {LocationData} LocationData
     */
    LocationData.fromObject = function fromObject(object) {
        if (object instanceof $root.LocationData)
            return object;
        let message = new $root.LocationData();
        if (object.locationName != null)
            message.locationName = String(object.locationName);
        if (object.roomId != null)
            message.roomId = object.roomId >>> 0;
        return message;
    };

    /**
     * Creates a plain object from a LocationData message. Also converts values to other types if specified.
     * @function toObject
     * @memberof LocationData
     * @static
     * @param {LocationData} message LocationData
     * @param {$protobuf.IConversionOptions} [options] Conversion options
     * @returns {Object.<string,*>} Plain object
     */
    LocationData.toObject = function toObject(message, options) {
        if (!options)
            options = {};
        let object = {};
        if (options.defaults) {
            object.locationName = "";
            object.roomId = 0;
        }
        if (message.locationName != null && message.hasOwnProperty("locationName"))
            object.locationName = message.locationName;
        if (message.roomId != null && message.hasOwnProperty("roomId"))
            object.roomId = message.roomId;
        return object;
    };

    /**
     * Converts this LocationData to JSON.
     * @function toJSON
     * @memberof LocationData
     * @instance
     * @returns {Object.<string,*>} JSON object
     */
    LocationData.prototype.toJSON = function toJSON() {
        return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
    };

    return LocationData;
})();

export const AccountGeneral = $root.AccountGeneral = (() => {

    /**
     * Properties of an AccountGeneral.
     * @exports IAccountGeneral
     * @interface IAccountGeneral
     * @property {number|null} [Money] AccountGeneral Money
     * @property {string|null} [Username] AccountGeneral Username
     */

    /**
     * Constructs a new AccountGeneral.
     * @exports AccountGeneral
     * @classdesc Represents an AccountGeneral.
     * @implements IAccountGeneral
     * @constructor
     * @param {IAccountGeneral=} [properties] Properties to set
     */
    function AccountGeneral(properties) {
        if (properties)
            for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                if (properties[keys[i]] != null)
                    this[keys[i]] = properties[keys[i]];
    }

    /**
     * AccountGeneral Money.
     * @member {number} Money
     * @memberof AccountGeneral
     * @instance
     */
    AccountGeneral.prototype.Money = 0;

    /**
     * AccountGeneral Username.
     * @member {string} Username
     * @memberof AccountGeneral
     * @instance
     */
    AccountGeneral.prototype.Username = "";

    /**
     * Creates a new AccountGeneral instance using the specified properties.
     * @function create
     * @memberof AccountGeneral
     * @static
     * @param {IAccountGeneral=} [properties] Properties to set
     * @returns {AccountGeneral} AccountGeneral instance
     */
    AccountGeneral.create = function create(properties) {
        return new AccountGeneral(properties);
    };

    /**
     * Encodes the specified AccountGeneral message. Does not implicitly {@link AccountGeneral.verify|verify} messages.
     * @function encode
     * @memberof AccountGeneral
     * @static
     * @param {IAccountGeneral} message AccountGeneral message or plain object to encode
     * @param {$protobuf.Writer} [writer] Writer to encode to
     * @returns {$protobuf.Writer} Writer
     */
    AccountGeneral.encode = function encode(message, writer) {
        if (!writer)
            writer = $Writer.create();
        if (message.Money != null && message.hasOwnProperty("Money"))
            writer.uint32(/* id 1, wireType 5 =*/13).float(message.Money);
        if (message.Username != null && message.hasOwnProperty("Username"))
            writer.uint32(/* id 2, wireType 2 =*/18).string(message.Username);
        return writer;
    };

    /**
     * Encodes the specified AccountGeneral message, length delimited. Does not implicitly {@link AccountGeneral.verify|verify} messages.
     * @function encodeDelimited
     * @memberof AccountGeneral
     * @static
     * @param {IAccountGeneral} message AccountGeneral message or plain object to encode
     * @param {$protobuf.Writer} [writer] Writer to encode to
     * @returns {$protobuf.Writer} Writer
     */
    AccountGeneral.encodeDelimited = function encodeDelimited(message, writer) {
        return this.encode(message, writer).ldelim();
    };

    /**
     * Decodes an AccountGeneral message from the specified reader or buffer.
     * @function decode
     * @memberof AccountGeneral
     * @static
     * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
     * @param {number} [length] Message length if known beforehand
     * @returns {AccountGeneral} AccountGeneral
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    AccountGeneral.decode = function decode(reader, length) {
        if (!(reader instanceof $Reader))
            reader = $Reader.create(reader);
        let end = length === undefined ? reader.len : reader.pos + length, message = new $root.AccountGeneral();
        while (reader.pos < end) {
            let tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.Money = reader.float();
                    break;
                case 2:
                    message.Username = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    };

    /**
     * Decodes an AccountGeneral message from the specified reader or buffer, length delimited.
     * @function decodeDelimited
     * @memberof AccountGeneral
     * @static
     * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
     * @returns {AccountGeneral} AccountGeneral
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    AccountGeneral.decodeDelimited = function decodeDelimited(reader) {
        if (!(reader instanceof $Reader))
            reader = new $Reader(reader);
        return this.decode(reader, reader.uint32());
    };

    /**
     * Verifies an AccountGeneral message.
     * @function verify
     * @memberof AccountGeneral
     * @static
     * @param {Object.<string,*>} message Plain object to verify
     * @returns {string|null} `null` if valid, otherwise the reason why it is not
     */
    AccountGeneral.verify = function verify(message) {
        if (typeof message !== "object" || message === null)
            return "object expected";
        if (message.Money != null && message.hasOwnProperty("Money"))
            if (typeof message.Money !== "number")
                return "Money: number expected";
        if (message.Username != null && message.hasOwnProperty("Username"))
            if (!$util.isString(message.Username))
                return "Username: string expected";
        return null;
    };

    /**
     * Creates an AccountGeneral message from a plain object. Also converts values to their respective internal types.
     * @function fromObject
     * @memberof AccountGeneral
     * @static
     * @param {Object.<string,*>} object Plain object
     * @returns {AccountGeneral} AccountGeneral
     */
    AccountGeneral.fromObject = function fromObject(object) {
        if (object instanceof $root.AccountGeneral)
            return object;
        let message = new $root.AccountGeneral();
        if (object.Money != null)
            message.Money = Number(object.Money);
        if (object.Username != null)
            message.Username = String(object.Username);
        return message;
    };

    /**
     * Creates a plain object from an AccountGeneral message. Also converts values to other types if specified.
     * @function toObject
     * @memberof AccountGeneral
     * @static
     * @param {AccountGeneral} message AccountGeneral
     * @param {$protobuf.IConversionOptions} [options] Conversion options
     * @returns {Object.<string,*>} Plain object
     */
    AccountGeneral.toObject = function toObject(message, options) {
        if (!options)
            options = {};
        let object = {};
        if (options.defaults) {
            object.Money = 0;
            object.Username = "";
        }
        if (message.Money != null && message.hasOwnProperty("Money"))
            object.Money = options.json && !isFinite(message.Money) ? String(message.Money) : message.Money;
        if (message.Username != null && message.hasOwnProperty("Username"))
            object.Username = message.Username;
        return object;
    };

    /**
     * Converts this AccountGeneral to JSON.
     * @function toJSON
     * @memberof AccountGeneral
     * @instance
     * @returns {Object.<string,*>} JSON object
     */
    AccountGeneral.prototype.toJSON = function toJSON() {
        return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
    };

    return AccountGeneral;
})();

export const ServerData = $root.ServerData = (() => {

    /**
     * Properties of a ServerData.
     * @exports IServerData
     * @interface IServerData
     * @property {IConnectionData|null} [conData] ServerData conData
     * @property {IAccountGeneral|null} [accountGeneral] ServerData accountGeneral
     * @property {ILocationData|null} [locationData] ServerData locationData
     * @property {Array.<string>|null} [curves] ServerData curves
     * @property {Array.<ICustomObject>|null} [customObjects] ServerData customObjects
     * @property {Array.<IFish>|null} [fishes] ServerData fishes
     * @property {Array.<IPlayer>|null} [players] ServerData players
     * @property {Array.<IAction>|null} [actions] ServerData actions
     */

    /**
     * Constructs a new ServerData.
     * @exports ServerData
     * @classdesc Represents a ServerData.
     * @implements IServerData
     * @constructor
     * @param {IServerData=} [properties] Properties to set
     */
    function ServerData(properties) {
        this.curves = [];
        this.customObjects = [];
        this.fishes = [];
        this.players = [];
        this.actions = [];
        if (properties)
            for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                if (properties[keys[i]] != null)
                    this[keys[i]] = properties[keys[i]];
    }

    /**
     * ServerData conData.
     * @member {IConnectionData|null|undefined} conData
     * @memberof ServerData
     * @instance
     */
    ServerData.prototype.conData = null;

    /**
     * ServerData accountGeneral.
     * @member {IAccountGeneral|null|undefined} accountGeneral
     * @memberof ServerData
     * @instance
     */
    ServerData.prototype.accountGeneral = null;

    /**
     * ServerData locationData.
     * @member {ILocationData|null|undefined} locationData
     * @memberof ServerData
     * @instance
     */
    ServerData.prototype.locationData = null;

    /**
     * ServerData curves.
     * @member {Array.<string>} curves
     * @memberof ServerData
     * @instance
     */
    ServerData.prototype.curves = $util.emptyArray;

    /**
     * ServerData customObjects.
     * @member {Array.<ICustomObject>} customObjects
     * @memberof ServerData
     * @instance
     */
    ServerData.prototype.customObjects = $util.emptyArray;

    /**
     * ServerData fishes.
     * @member {Array.<IFish>} fishes
     * @memberof ServerData
     * @instance
     */
    ServerData.prototype.fishes = $util.emptyArray;

    /**
     * ServerData players.
     * @member {Array.<IPlayer>} players
     * @memberof ServerData
     * @instance
     */
    ServerData.prototype.players = $util.emptyArray;

    /**
     * ServerData actions.
     * @member {Array.<IAction>} actions
     * @memberof ServerData
     * @instance
     */
    ServerData.prototype.actions = $util.emptyArray;

    /**
     * Creates a new ServerData instance using the specified properties.
     * @function create
     * @memberof ServerData
     * @static
     * @param {IServerData=} [properties] Properties to set
     * @returns {ServerData} ServerData instance
     */
    ServerData.create = function create(properties) {
        return new ServerData(properties);
    };

    /**
     * Encodes the specified ServerData message. Does not implicitly {@link ServerData.verify|verify} messages.
     * @function encode
     * @memberof ServerData
     * @static
     * @param {IServerData} message ServerData message or plain object to encode
     * @param {$protobuf.Writer} [writer] Writer to encode to
     * @returns {$protobuf.Writer} Writer
     */
    ServerData.encode = function encode(message, writer) {
        if (!writer)
            writer = $Writer.create();
        if (message.conData != null && message.hasOwnProperty("conData"))
            $root.ConnectionData.encode(message.conData, writer.uint32(/* id 2, wireType 2 =*/18).fork()).ldelim();
        if (message.accountGeneral != null && message.hasOwnProperty("accountGeneral"))
            $root.AccountGeneral.encode(message.accountGeneral, writer.uint32(/* id 3, wireType 2 =*/26).fork()).ldelim();
        if (message.locationData != null && message.hasOwnProperty("locationData"))
            $root.LocationData.encode(message.locationData, writer.uint32(/* id 4, wireType 2 =*/34).fork()).ldelim();
        if (message.curves != null && message.curves.length)
            for (let i = 0; i < message.curves.length; ++i)
                writer.uint32(/* id 5, wireType 2 =*/42).string(message.curves[i]);
        if (message.customObjects != null && message.customObjects.length)
            for (let i = 0; i < message.customObjects.length; ++i)
                $root.CustomObject.encode(message.customObjects[i], writer.uint32(/* id 6, wireType 2 =*/50).fork()).ldelim();
        if (message.fishes != null && message.fishes.length)
            for (let i = 0; i < message.fishes.length; ++i)
                $root.Fish.encode(message.fishes[i], writer.uint32(/* id 7, wireType 2 =*/58).fork()).ldelim();
        if (message.players != null && message.players.length)
            for (let i = 0; i < message.players.length; ++i)
                $root.Player.encode(message.players[i], writer.uint32(/* id 8, wireType 2 =*/66).fork()).ldelim();
        if (message.actions != null && message.actions.length)
            for (let i = 0; i < message.actions.length; ++i)
                $root.Action.encode(message.actions[i], writer.uint32(/* id 9, wireType 2 =*/74).fork()).ldelim();
        return writer;
    };

    /**
     * Encodes the specified ServerData message, length delimited. Does not implicitly {@link ServerData.verify|verify} messages.
     * @function encodeDelimited
     * @memberof ServerData
     * @static
     * @param {IServerData} message ServerData message or plain object to encode
     * @param {$protobuf.Writer} [writer] Writer to encode to
     * @returns {$protobuf.Writer} Writer
     */
    ServerData.encodeDelimited = function encodeDelimited(message, writer) {
        return this.encode(message, writer).ldelim();
    };

    /**
     * Decodes a ServerData message from the specified reader or buffer.
     * @function decode
     * @memberof ServerData
     * @static
     * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
     * @param {number} [length] Message length if known beforehand
     * @returns {ServerData} ServerData
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    ServerData.decode = function decode(reader, length) {
        if (!(reader instanceof $Reader))
            reader = $Reader.create(reader);
        let end = length === undefined ? reader.len : reader.pos + length, message = new $root.ServerData();
        while (reader.pos < end) {
            let tag = reader.uint32();
            switch (tag >>> 3) {
                case 2:
                    message.conData = $root.ConnectionData.decode(reader, reader.uint32());
                    break;
                case 3:
                    message.accountGeneral = $root.AccountGeneral.decode(reader, reader.uint32());
                    break;
                case 4:
                    message.locationData = $root.LocationData.decode(reader, reader.uint32());
                    break;
                case 5:
                    if (!(message.curves && message.curves.length))
                        message.curves = [];
                    message.curves.push(reader.string());
                    break;
                case 6:
                    if (!(message.customObjects && message.customObjects.length))
                        message.customObjects = [];
                    message.customObjects.push($root.CustomObject.decode(reader, reader.uint32()));
                    break;
                case 7:
                    if (!(message.fishes && message.fishes.length))
                        message.fishes = [];
                    message.fishes.push($root.Fish.decode(reader, reader.uint32()));
                    break;
                case 8:
                    if (!(message.players && message.players.length))
                        message.players = [];
                    message.players.push($root.Player.decode(reader, reader.uint32()));
                    break;
                case 9:
                    if (!(message.actions && message.actions.length))
                        message.actions = [];
                    message.actions.push($root.Action.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    };

    /**
     * Decodes a ServerData message from the specified reader or buffer, length delimited.
     * @function decodeDelimited
     * @memberof ServerData
     * @static
     * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
     * @returns {ServerData} ServerData
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    ServerData.decodeDelimited = function decodeDelimited(reader) {
        if (!(reader instanceof $Reader))
            reader = new $Reader(reader);
        return this.decode(reader, reader.uint32());
    };

    /**
     * Verifies a ServerData message.
     * @function verify
     * @memberof ServerData
     * @static
     * @param {Object.<string,*>} message Plain object to verify
     * @returns {string|null} `null` if valid, otherwise the reason why it is not
     */
    ServerData.verify = function verify(message) {
        if (typeof message !== "object" || message === null)
            return "object expected";
        if (message.conData != null && message.hasOwnProperty("conData")) {
            let error = $root.ConnectionData.verify(message.conData);
            if (error)
                return "conData." + error;
        }
        if (message.accountGeneral != null && message.hasOwnProperty("accountGeneral")) {
            let error = $root.AccountGeneral.verify(message.accountGeneral);
            if (error)
                return "accountGeneral." + error;
        }
        if (message.locationData != null && message.hasOwnProperty("locationData")) {
            let error = $root.LocationData.verify(message.locationData);
            if (error)
                return "locationData." + error;
        }
        if (message.curves != null && message.hasOwnProperty("curves")) {
            if (!Array.isArray(message.curves))
                return "curves: array expected";
            for (let i = 0; i < message.curves.length; ++i)
                if (!$util.isString(message.curves[i]))
                    return "curves: string[] expected";
        }
        if (message.customObjects != null && message.hasOwnProperty("customObjects")) {
            if (!Array.isArray(message.customObjects))
                return "customObjects: array expected";
            for (let i = 0; i < message.customObjects.length; ++i) {
                let error = $root.CustomObject.verify(message.customObjects[i]);
                if (error)
                    return "customObjects." + error;
            }
        }
        if (message.fishes != null && message.hasOwnProperty("fishes")) {
            if (!Array.isArray(message.fishes))
                return "fishes: array expected";
            for (let i = 0; i < message.fishes.length; ++i) {
                let error = $root.Fish.verify(message.fishes[i]);
                if (error)
                    return "fishes." + error;
            }
        }
        if (message.players != null && message.hasOwnProperty("players")) {
            if (!Array.isArray(message.players))
                return "players: array expected";
            for (let i = 0; i < message.players.length; ++i) {
                let error = $root.Player.verify(message.players[i]);
                if (error)
                    return "players." + error;
            }
        }
        if (message.actions != null && message.hasOwnProperty("actions")) {
            if (!Array.isArray(message.actions))
                return "actions: array expected";
            for (let i = 0; i < message.actions.length; ++i) {
                let error = $root.Action.verify(message.actions[i]);
                if (error)
                    return "actions." + error;
            }
        }
        return null;
    };

    /**
     * Creates a ServerData message from a plain object. Also converts values to their respective internal types.
     * @function fromObject
     * @memberof ServerData
     * @static
     * @param {Object.<string,*>} object Plain object
     * @returns {ServerData} ServerData
     */
    ServerData.fromObject = function fromObject(object) {
        if (object instanceof $root.ServerData)
            return object;
        let message = new $root.ServerData();
        if (object.conData != null) {
            if (typeof object.conData !== "object")
                throw TypeError(".ServerData.conData: object expected");
            message.conData = $root.ConnectionData.fromObject(object.conData);
        }
        if (object.accountGeneral != null) {
            if (typeof object.accountGeneral !== "object")
                throw TypeError(".ServerData.accountGeneral: object expected");
            message.accountGeneral = $root.AccountGeneral.fromObject(object.accountGeneral);
        }
        if (object.locationData != null) {
            if (typeof object.locationData !== "object")
                throw TypeError(".ServerData.locationData: object expected");
            message.locationData = $root.LocationData.fromObject(object.locationData);
        }
        if (object.curves) {
            if (!Array.isArray(object.curves))
                throw TypeError(".ServerData.curves: array expected");
            message.curves = [];
            for (let i = 0; i < object.curves.length; ++i)
                message.curves[i] = String(object.curves[i]);
        }
        if (object.customObjects) {
            if (!Array.isArray(object.customObjects))
                throw TypeError(".ServerData.customObjects: array expected");
            message.customObjects = [];
            for (let i = 0; i < object.customObjects.length; ++i) {
                if (typeof object.customObjects[i] !== "object")
                    throw TypeError(".ServerData.customObjects: object expected");
                message.customObjects[i] = $root.CustomObject.fromObject(object.customObjects[i]);
            }
        }
        if (object.fishes) {
            if (!Array.isArray(object.fishes))
                throw TypeError(".ServerData.fishes: array expected");
            message.fishes = [];
            for (let i = 0; i < object.fishes.length; ++i) {
                if (typeof object.fishes[i] !== "object")
                    throw TypeError(".ServerData.fishes: object expected");
                message.fishes[i] = $root.Fish.fromObject(object.fishes[i]);
            }
        }
        if (object.players) {
            if (!Array.isArray(object.players))
                throw TypeError(".ServerData.players: array expected");
            message.players = [];
            for (let i = 0; i < object.players.length; ++i) {
                if (typeof object.players[i] !== "object")
                    throw TypeError(".ServerData.players: object expected");
                message.players[i] = $root.Player.fromObject(object.players[i]);
            }
        }
        if (object.actions) {
            if (!Array.isArray(object.actions))
                throw TypeError(".ServerData.actions: array expected");
            message.actions = [];
            for (let i = 0; i < object.actions.length; ++i) {
                if (typeof object.actions[i] !== "object")
                    throw TypeError(".ServerData.actions: object expected");
                message.actions[i] = $root.Action.fromObject(object.actions[i]);
            }
        }
        return message;
    };

    /**
     * Creates a plain object from a ServerData message. Also converts values to other types if specified.
     * @function toObject
     * @memberof ServerData
     * @static
     * @param {ServerData} message ServerData
     * @param {$protobuf.IConversionOptions} [options] Conversion options
     * @returns {Object.<string,*>} Plain object
     */
    ServerData.toObject = function toObject(message, options) {
        if (!options)
            options = {};
        let object = {};
        if (options.arrays || options.defaults) {
            object.curves = [];
            object.customObjects = [];
            object.fishes = [];
            object.players = [];
            object.actions = [];
        }
        if (options.defaults) {
            object.conData = null;
            object.accountGeneral = null;
            object.locationData = null;
        }
        if (message.conData != null && message.hasOwnProperty("conData"))
            object.conData = $root.ConnectionData.toObject(message.conData, options);
        if (message.accountGeneral != null && message.hasOwnProperty("accountGeneral"))
            object.accountGeneral = $root.AccountGeneral.toObject(message.accountGeneral, options);
        if (message.locationData != null && message.hasOwnProperty("locationData"))
            object.locationData = $root.LocationData.toObject(message.locationData, options);
        if (message.curves && message.curves.length) {
            object.curves = [];
            for (let j = 0; j < message.curves.length; ++j)
                object.curves[j] = message.curves[j];
        }
        if (message.customObjects && message.customObjects.length) {
            object.customObjects = [];
            for (let j = 0; j < message.customObjects.length; ++j)
                object.customObjects[j] = $root.CustomObject.toObject(message.customObjects[j], options);
        }
        if (message.fishes && message.fishes.length) {
            object.fishes = [];
            for (let j = 0; j < message.fishes.length; ++j)
                object.fishes[j] = $root.Fish.toObject(message.fishes[j], options);
        }
        if (message.players && message.players.length) {
            object.players = [];
            for (let j = 0; j < message.players.length; ++j)
                object.players[j] = $root.Player.toObject(message.players[j], options);
        }
        if (message.actions && message.actions.length) {
            object.actions = [];
            for (let j = 0; j < message.actions.length; ++j)
                object.actions[j] = $root.Action.toObject(message.actions[j], options);
        }
        return object;
    };

    /**
     * Converts this ServerData to JSON.
     * @function toJSON
     * @memberof ServerData
     * @instance
     * @returns {Object.<string,*>} JSON object
     */
    ServerData.prototype.toJSON = function toJSON() {
        return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
    };

    return ServerData;
})();

export const Command = $root.Command = (() => {

    /**
     * Properties of a Command.
     * @exports ICommand
     * @interface ICommand
     * @property {number|null} [CommandId] Command CommandId
     * @property {Array.<number>|null} [Params] Command Params
     */

    /**
     * Constructs a new Command.
     * @exports Command
     * @classdesc Represents a Command.
     * @implements ICommand
     * @constructor
     * @param {ICommand=} [properties] Properties to set
     */
    function Command(properties) {
        this.Params = [];
        if (properties)
            for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                if (properties[keys[i]] != null)
                    this[keys[i]] = properties[keys[i]];
    }

    /**
     * Command CommandId.
     * @member {number} CommandId
     * @memberof Command
     * @instance
     */
    Command.prototype.CommandId = 0;

    /**
     * Command Params.
     * @member {Array.<number>} Params
     * @memberof Command
     * @instance
     */
    Command.prototype.Params = $util.emptyArray;

    /**
     * Creates a new Command instance using the specified properties.
     * @function create
     * @memberof Command
     * @static
     * @param {ICommand=} [properties] Properties to set
     * @returns {Command} Command instance
     */
    Command.create = function create(properties) {
        return new Command(properties);
    };

    /**
     * Encodes the specified Command message. Does not implicitly {@link Command.verify|verify} messages.
     * @function encode
     * @memberof Command
     * @static
     * @param {ICommand} message Command message or plain object to encode
     * @param {$protobuf.Writer} [writer] Writer to encode to
     * @returns {$protobuf.Writer} Writer
     */
    Command.encode = function encode(message, writer) {
        if (!writer)
            writer = $Writer.create();
        if (message.CommandId != null && message.hasOwnProperty("CommandId"))
            writer.uint32(/* id 1, wireType 0 =*/8).int32(message.CommandId);
        if (message.Params != null && message.Params.length) {
            writer.uint32(/* id 2, wireType 2 =*/18).fork();
            for (let i = 0; i < message.Params.length; ++i)
                writer.int32(message.Params[i]);
            writer.ldelim();
        }
        return writer;
    };

    /**
     * Encodes the specified Command message, length delimited. Does not implicitly {@link Command.verify|verify} messages.
     * @function encodeDelimited
     * @memberof Command
     * @static
     * @param {ICommand} message Command message or plain object to encode
     * @param {$protobuf.Writer} [writer] Writer to encode to
     * @returns {$protobuf.Writer} Writer
     */
    Command.encodeDelimited = function encodeDelimited(message, writer) {
        return this.encode(message, writer).ldelim();
    };

    /**
     * Decodes a Command message from the specified reader or buffer.
     * @function decode
     * @memberof Command
     * @static
     * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
     * @param {number} [length] Message length if known beforehand
     * @returns {Command} Command
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    Command.decode = function decode(reader, length) {
        if (!(reader instanceof $Reader))
            reader = $Reader.create(reader);
        let end = length === undefined ? reader.len : reader.pos + length, message = new $root.Command();
        while (reader.pos < end) {
            let tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.CommandId = reader.int32();
                    break;
                case 2:
                    if (!(message.Params && message.Params.length))
                        message.Params = [];
                    if ((tag & 7) === 2) {
                        let end2 = reader.uint32() + reader.pos;
                        while (reader.pos < end2)
                            message.Params.push(reader.int32());
                    } else
                        message.Params.push(reader.int32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    };

    /**
     * Decodes a Command message from the specified reader or buffer, length delimited.
     * @function decodeDelimited
     * @memberof Command
     * @static
     * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
     * @returns {Command} Command
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    Command.decodeDelimited = function decodeDelimited(reader) {
        if (!(reader instanceof $Reader))
            reader = new $Reader(reader);
        return this.decode(reader, reader.uint32());
    };

    /**
     * Verifies a Command message.
     * @function verify
     * @memberof Command
     * @static
     * @param {Object.<string,*>} message Plain object to verify
     * @returns {string|null} `null` if valid, otherwise the reason why it is not
     */
    Command.verify = function verify(message) {
        if (typeof message !== "object" || message === null)
            return "object expected";
        if (message.CommandId != null && message.hasOwnProperty("CommandId"))
            if (!$util.isInteger(message.CommandId))
                return "CommandId: integer expected";
        if (message.Params != null && message.hasOwnProperty("Params")) {
            if (!Array.isArray(message.Params))
                return "Params: array expected";
            for (let i = 0; i < message.Params.length; ++i)
                if (!$util.isInteger(message.Params[i]))
                    return "Params: integer[] expected";
        }
        return null;
    };

    /**
     * Creates a Command message from a plain object. Also converts values to their respective internal types.
     * @function fromObject
     * @memberof Command
     * @static
     * @param {Object.<string,*>} object Plain object
     * @returns {Command} Command
     */
    Command.fromObject = function fromObject(object) {
        if (object instanceof $root.Command)
            return object;
        let message = new $root.Command();
        if (object.CommandId != null)
            message.CommandId = object.CommandId | 0;
        if (object.Params) {
            if (!Array.isArray(object.Params))
                throw TypeError(".Command.Params: array expected");
            message.Params = [];
            for (let i = 0; i < object.Params.length; ++i)
                message.Params[i] = object.Params[i] | 0;
        }
        return message;
    };

    /**
     * Creates a plain object from a Command message. Also converts values to other types if specified.
     * @function toObject
     * @memberof Command
     * @static
     * @param {Command} message Command
     * @param {$protobuf.IConversionOptions} [options] Conversion options
     * @returns {Object.<string,*>} Plain object
     */
    Command.toObject = function toObject(message, options) {
        if (!options)
            options = {};
        let object = {};
        if (options.arrays || options.defaults)
            object.Params = [];
        if (options.defaults)
            object.CommandId = 0;
        if (message.CommandId != null && message.hasOwnProperty("CommandId"))
            object.CommandId = message.CommandId;
        if (message.Params && message.Params.length) {
            object.Params = [];
            for (let j = 0; j < message.Params.length; ++j)
                object.Params[j] = message.Params[j];
        }
        return object;
    };

    /**
     * Converts this Command to JSON.
     * @function toJSON
     * @memberof Command
     * @instance
     * @returns {Object.<string,*>} JSON object
     */
    Command.prototype.toJSON = function toJSON() {
        return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
    };

    return Command;
})();

export const ClientData = $root.ClientData = (() => {

    /**
     * Properties of a ClientData.
     * @exports IClientData
     * @interface IClientData
     * @property {number|null} [roomId] ClientData roomId
     * @property {Array.<ICommand>|null} [commands] ClientData commands
     */

    /**
     * Constructs a new ClientData.
     * @exports ClientData
     * @classdesc Represents a ClientData.
     * @implements IClientData
     * @constructor
     * @param {IClientData=} [properties] Properties to set
     */
    function ClientData(properties) {
        this.commands = [];
        if (properties)
            for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                if (properties[keys[i]] != null)
                    this[keys[i]] = properties[keys[i]];
    }

    /**
     * ClientData roomId.
     * @member {number} roomId
     * @memberof ClientData
     * @instance
     */
    ClientData.prototype.roomId = 0;

    /**
     * ClientData commands.
     * @member {Array.<ICommand>} commands
     * @memberof ClientData
     * @instance
     */
    ClientData.prototype.commands = $util.emptyArray;

    /**
     * Creates a new ClientData instance using the specified properties.
     * @function create
     * @memberof ClientData
     * @static
     * @param {IClientData=} [properties] Properties to set
     * @returns {ClientData} ClientData instance
     */
    ClientData.create = function create(properties) {
        return new ClientData(properties);
    };

    /**
     * Encodes the specified ClientData message. Does not implicitly {@link ClientData.verify|verify} messages.
     * @function encode
     * @memberof ClientData
     * @static
     * @param {IClientData} message ClientData message or plain object to encode
     * @param {$protobuf.Writer} [writer] Writer to encode to
     * @returns {$protobuf.Writer} Writer
     */
    ClientData.encode = function encode(message, writer) {
        if (!writer)
            writer = $Writer.create();
        if (message.roomId != null && message.hasOwnProperty("roomId"))
            writer.uint32(/* id 2, wireType 0 =*/16).int32(message.roomId);
        if (message.commands != null && message.commands.length)
            for (let i = 0; i < message.commands.length; ++i)
                $root.Command.encode(message.commands[i], writer.uint32(/* id 3, wireType 2 =*/26).fork()).ldelim();
        return writer;
    };

    /**
     * Encodes the specified ClientData message, length delimited. Does not implicitly {@link ClientData.verify|verify} messages.
     * @function encodeDelimited
     * @memberof ClientData
     * @static
     * @param {IClientData} message ClientData message or plain object to encode
     * @param {$protobuf.Writer} [writer] Writer to encode to
     * @returns {$protobuf.Writer} Writer
     */
    ClientData.encodeDelimited = function encodeDelimited(message, writer) {
        return this.encode(message, writer).ldelim();
    };

    /**
     * Decodes a ClientData message from the specified reader or buffer.
     * @function decode
     * @memberof ClientData
     * @static
     * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
     * @param {number} [length] Message length if known beforehand
     * @returns {ClientData} ClientData
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    ClientData.decode = function decode(reader, length) {
        if (!(reader instanceof $Reader))
            reader = $Reader.create(reader);
        let end = length === undefined ? reader.len : reader.pos + length, message = new $root.ClientData();
        while (reader.pos < end) {
            let tag = reader.uint32();
            switch (tag >>> 3) {
                case 2:
                    message.roomId = reader.int32();
                    break;
                case 3:
                    if (!(message.commands && message.commands.length))
                        message.commands = [];
                    message.commands.push($root.Command.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    };

    /**
     * Decodes a ClientData message from the specified reader or buffer, length delimited.
     * @function decodeDelimited
     * @memberof ClientData
     * @static
     * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
     * @returns {ClientData} ClientData
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    ClientData.decodeDelimited = function decodeDelimited(reader) {
        if (!(reader instanceof $Reader))
            reader = new $Reader(reader);
        return this.decode(reader, reader.uint32());
    };

    /**
     * Verifies a ClientData message.
     * @function verify
     * @memberof ClientData
     * @static
     * @param {Object.<string,*>} message Plain object to verify
     * @returns {string|null} `null` if valid, otherwise the reason why it is not
     */
    ClientData.verify = function verify(message) {
        if (typeof message !== "object" || message === null)
            return "object expected";
        if (message.roomId != null && message.hasOwnProperty("roomId"))
            if (!$util.isInteger(message.roomId))
                return "roomId: integer expected";
        if (message.commands != null && message.hasOwnProperty("commands")) {
            if (!Array.isArray(message.commands))
                return "commands: array expected";
            for (let i = 0; i < message.commands.length; ++i) {
                let error = $root.Command.verify(message.commands[i]);
                if (error)
                    return "commands." + error;
            }
        }
        return null;
    };

    /**
     * Creates a ClientData message from a plain object. Also converts values to their respective internal types.
     * @function fromObject
     * @memberof ClientData
     * @static
     * @param {Object.<string,*>} object Plain object
     * @returns {ClientData} ClientData
     */
    ClientData.fromObject = function fromObject(object) {
        if (object instanceof $root.ClientData)
            return object;
        let message = new $root.ClientData();
        if (object.roomId != null)
            message.roomId = object.roomId | 0;
        if (object.commands) {
            if (!Array.isArray(object.commands))
                throw TypeError(".ClientData.commands: array expected");
            message.commands = [];
            for (let i = 0; i < object.commands.length; ++i) {
                if (typeof object.commands[i] !== "object")
                    throw TypeError(".ClientData.commands: object expected");
                message.commands[i] = $root.Command.fromObject(object.commands[i]);
            }
        }
        return message;
    };

    /**
     * Creates a plain object from a ClientData message. Also converts values to other types if specified.
     * @function toObject
     * @memberof ClientData
     * @static
     * @param {ClientData} message ClientData
     * @param {$protobuf.IConversionOptions} [options] Conversion options
     * @returns {Object.<string,*>} Plain object
     */
    ClientData.toObject = function toObject(message, options) {
        if (!options)
            options = {};
        let object = {};
        if (options.arrays || options.defaults)
            object.commands = [];
        if (options.defaults)
            object.roomId = 0;
        if (message.roomId != null && message.hasOwnProperty("roomId"))
            object.roomId = message.roomId;
        if (message.commands && message.commands.length) {
            object.commands = [];
            for (let j = 0; j < message.commands.length; ++j)
                object.commands[j] = $root.Command.toObject(message.commands[j], options);
        }
        return object;
    };

    /**
     * Converts this ClientData to JSON.
     * @function toJSON
     * @memberof ClientData
     * @instance
     * @returns {Object.<string,*>} JSON object
     */
    ClientData.prototype.toJSON = function toJSON() {
        return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
    };

    return ClientData;
})();

export { $root as default };
