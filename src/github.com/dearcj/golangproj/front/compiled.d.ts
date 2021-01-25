import * as $protobuf from "protobufjs";
/** Properties of a CustomObject. */
export interface ICustomObject {

    /** CustomObject networkObject */
    networkObject?: (INetworkObject|null);

    /** CustomObject param1 */
    param1?: (number|null);

    /** CustomObject param2 */
    param2?: (number|null);

    /** CustomObject param3 */
    param3?: (number|null);
}

/** Represents a CustomObject. */
export class CustomObject implements ICustomObject {

    /**
     * Constructs a new CustomObject.
     * @param [properties] Properties to set
     */
    constructor(properties?: ICustomObject);

    /** CustomObject networkObject. */
    public networkObject?: (INetworkObject|null);

    /** CustomObject param1. */
    public param1: number;

    /** CustomObject param2. */
    public param2: number;

    /** CustomObject param3. */
    public param3: number;

    /**
     * Creates a new CustomObject instance using the specified properties.
     * @param [properties] Properties to set
     * @returns CustomObject instance
     */
    public static create(properties?: ICustomObject): CustomObject;

    /**
     * Encodes the specified CustomObject message. Does not implicitly {@link CustomObject.verify|verify} messages.
     * @param message CustomObject message or plain object to encode
     * @param [writer] Writer to encode to
     * @returns Writer
     */
    public static encode(message: ICustomObject, writer?: $protobuf.Writer): $protobuf.Writer;

    /**
     * Encodes the specified CustomObject message, length delimited. Does not implicitly {@link CustomObject.verify|verify} messages.
     * @param message CustomObject message or plain object to encode
     * @param [writer] Writer to encode to
     * @returns Writer
     */
    public static encodeDelimited(message: ICustomObject, writer?: $protobuf.Writer): $protobuf.Writer;

    /**
     * Decodes a CustomObject message from the specified reader or buffer.
     * @param reader Reader or buffer to decode from
     * @param [length] Message length if known beforehand
     * @returns CustomObject
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): CustomObject;

    /**
     * Decodes a CustomObject message from the specified reader or buffer, length delimited.
     * @param reader Reader or buffer to decode from
     * @returns CustomObject
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): CustomObject;

    /**
     * Verifies a CustomObject message.
     * @param message Plain object to verify
     * @returns `null` if valid, otherwise the reason why it is not
     */
    public static verify(message: [ 'object' ].<string, any>): (string|null);

    /**
     * Creates a CustomObject message from a plain object. Also converts values to their respective internal types.
     * @param object Plain object
     * @returns CustomObject
     */
    public static fromObject(object: [ 'object' ].<string, any>): CustomObject;

    /**
     * Creates a plain object from a CustomObject message. Also converts values to other types if specified.
     * @param message CustomObject
     * @param [options] Conversion options
     * @returns Plain object
     */
    public static toObject(message: CustomObject, options?: $protobuf.IConversionOptions): [ 'object' ].<string, any>;

    /**
     * Converts this CustomObject to JSON.
     * @returns JSON object
     */
    public toJSON(): [ 'object' ].<string, any>;
}

/** Properties of a Player. */
export interface IPlayer {

    /** Player startPosition */
    startPosition?: (number|null);

    /** Player name */
    name?: (string|null);

    /** Player networkObject */
    networkObject?: (INetworkObject|null);
}

/** Represents a Player. */
export class Player implements IPlayer {

    /**
     * Constructs a new Player.
     * @param [properties] Properties to set
     */
    constructor(properties?: IPlayer);

    /** Player startPosition. */
    public startPosition: number;

    /** Player name. */
    public name: string;

    /** Player networkObject. */
    public networkObject?: (INetworkObject|null);

    /**
     * Creates a new Player instance using the specified properties.
     * @param [properties] Properties to set
     * @returns Player instance
     */
    public static create(properties?: IPlayer): Player;

    /**
     * Encodes the specified Player message. Does not implicitly {@link Player.verify|verify} messages.
     * @param message Player message or plain object to encode
     * @param [writer] Writer to encode to
     * @returns Writer
     */
    public static encode(message: IPlayer, writer?: $protobuf.Writer): $protobuf.Writer;

    /**
     * Encodes the specified Player message, length delimited. Does not implicitly {@link Player.verify|verify} messages.
     * @param message Player message or plain object to encode
     * @param [writer] Writer to encode to
     * @returns Writer
     */
    public static encodeDelimited(message: IPlayer, writer?: $protobuf.Writer): $protobuf.Writer;

    /**
     * Decodes a Player message from the specified reader or buffer.
     * @param reader Reader or buffer to decode from
     * @param [length] Message length if known beforehand
     * @returns Player
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): Player;

    /**
     * Decodes a Player message from the specified reader or buffer, length delimited.
     * @param reader Reader or buffer to decode from
     * @returns Player
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): Player;

    /**
     * Verifies a Player message.
     * @param message Plain object to verify
     * @returns `null` if valid, otherwise the reason why it is not
     */
    public static verify(message: [ 'object' ].<string, any>): (string|null);

    /**
     * Creates a Player message from a plain object. Also converts values to their respective internal types.
     * @param object Plain object
     * @returns Player
     */
    public static fromObject(object: [ 'object' ].<string, any>): Player;

    /**
     * Creates a plain object from a Player message. Also converts values to other types if specified.
     * @param message Player
     * @param [options] Conversion options
     * @returns Plain object
     */
    public static toObject(message: Player, options?: $protobuf.IConversionOptions): [ 'object' ].<string, any>;

    /**
     * Converts this Player to JSON.
     * @returns JSON object
     */
    public toJSON(): [ 'object' ].<string, any>;
}

/** Properties of a Fish. */
export interface IFish {

    /** Fish Hp */
    Hp?: (number|null);

    /** Fish Maxhp */
    Maxhp?: (number|null);

    /** Fish CurveInx */
    CurveInx?: (number|null);

    /** Fish FishType */
    FishType?: (number|null);

    /** Fish StartTime */
    StartTime?: (number|Long|null);

    /** Fish CurveTime */
    CurveTime?: (number|Long|null);

    /** Fish IsBoss */
    IsBoss?: (boolean|null);

    /** Fish networkObject */
    networkObject?: (INetworkObject|null);
}

/** Represents a Fish. */
export class Fish implements IFish {

    /**
     * Constructs a new Fish.
     * @param [properties] Properties to set
     */
    constructor(properties?: IFish);

    /** Fish Hp. */
    public Hp: number;

    /** Fish Maxhp. */
    public Maxhp: number;

    /** Fish CurveInx. */
    public CurveInx: number;

    /** Fish FishType. */
    public FishType: number;

    /** Fish StartTime. */
    public StartTime: (number|Long);

    /** Fish CurveTime. */
    public CurveTime: (number|Long);

    /** Fish IsBoss. */
    public IsBoss: boolean;

    /** Fish networkObject. */
    public networkObject?: (INetworkObject|null);

    /**
     * Creates a new Fish instance using the specified properties.
     * @param [properties] Properties to set
     * @returns Fish instance
     */
    public static create(properties?: IFish): Fish;

    /**
     * Encodes the specified Fish message. Does not implicitly {@link Fish.verify|verify} messages.
     * @param message Fish message or plain object to encode
     * @param [writer] Writer to encode to
     * @returns Writer
     */
    public static encode(message: IFish, writer?: $protobuf.Writer): $protobuf.Writer;

    /**
     * Encodes the specified Fish message, length delimited. Does not implicitly {@link Fish.verify|verify} messages.
     * @param message Fish message or plain object to encode
     * @param [writer] Writer to encode to
     * @returns Writer
     */
    public static encodeDelimited(message: IFish, writer?: $protobuf.Writer): $protobuf.Writer;

    /**
     * Decodes a Fish message from the specified reader or buffer.
     * @param reader Reader or buffer to decode from
     * @param [length] Message length if known beforehand
     * @returns Fish
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): Fish;

    /**
     * Decodes a Fish message from the specified reader or buffer, length delimited.
     * @param reader Reader or buffer to decode from
     * @returns Fish
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): Fish;

    /**
     * Verifies a Fish message.
     * @param message Plain object to verify
     * @returns `null` if valid, otherwise the reason why it is not
     */
    public static verify(message: [ 'object' ].<string, any>): (string|null);

    /**
     * Creates a Fish message from a plain object. Also converts values to their respective internal types.
     * @param object Plain object
     * @returns Fish
     */
    public static fromObject(object: [ 'object' ].<string, any>): Fish;

    /**
     * Creates a plain object from a Fish message. Also converts values to other types if specified.
     * @param message Fish
     * @param [options] Conversion options
     * @returns Plain object
     */
    public static toObject(message: Fish, options?: $protobuf.IConversionOptions): [ 'object' ].<string, any>;

    /**
     * Converts this Fish to JSON.
     * @returns JSON object
     */
    public toJSON(): [ 'object' ].<string, any>;
}

/** Properties of a NetworkObject. */
export interface INetworkObject {

    /** NetworkObject ID */
    ID?: (number|null);

    /** NetworkObject Type */
    Type?: (number|null);
}

/** Represents a NetworkObject. */
export class NetworkObject implements INetworkObject {

    /**
     * Constructs a new NetworkObject.
     * @param [properties] Properties to set
     */
    constructor(properties?: INetworkObject);

    /** NetworkObject ID. */
    public ID: number;

    /** NetworkObject Type. */
    public Type: number;

    /**
     * Creates a new NetworkObject instance using the specified properties.
     * @param [properties] Properties to set
     * @returns NetworkObject instance
     */
    public static create(properties?: INetworkObject): NetworkObject;

    /**
     * Encodes the specified NetworkObject message. Does not implicitly {@link NetworkObject.verify|verify} messages.
     * @param message NetworkObject message or plain object to encode
     * @param [writer] Writer to encode to
     * @returns Writer
     */
    public static encode(message: INetworkObject, writer?: $protobuf.Writer): $protobuf.Writer;

    /**
     * Encodes the specified NetworkObject message, length delimited. Does not implicitly {@link NetworkObject.verify|verify} messages.
     * @param message NetworkObject message or plain object to encode
     * @param [writer] Writer to encode to
     * @returns Writer
     */
    public static encodeDelimited(message: INetworkObject, writer?: $protobuf.Writer): $protobuf.Writer;

    /**
     * Decodes a NetworkObject message from the specified reader or buffer.
     * @param reader Reader or buffer to decode from
     * @param [length] Message length if known beforehand
     * @returns NetworkObject
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): NetworkObject;

    /**
     * Decodes a NetworkObject message from the specified reader or buffer, length delimited.
     * @param reader Reader or buffer to decode from
     * @returns NetworkObject
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): NetworkObject;

    /**
     * Verifies a NetworkObject message.
     * @param message Plain object to verify
     * @returns `null` if valid, otherwise the reason why it is not
     */
    public static verify(message: [ 'object' ].<string, any>): (string|null);

    /**
     * Creates a NetworkObject message from a plain object. Also converts values to their respective internal types.
     * @param object Plain object
     * @returns NetworkObject
     */
    public static fromObject(object: [ 'object' ].<string, any>): NetworkObject;

    /**
     * Creates a plain object from a NetworkObject message. Also converts values to other types if specified.
     * @param message NetworkObject
     * @param [options] Conversion options
     * @returns Plain object
     */
    public static toObject(message: NetworkObject, options?: $protobuf.IConversionOptions): [ 'object' ].<string, any>;

    /**
     * Converts this NetworkObject to JSON.
     * @returns JSON object
     */
    public toJSON(): [ 'object' ].<string, any>;
}

/** ActionType enum. */
export enum ActionType {
    ANY_VALUE = 0
}

/** Properties of an Action. */
export interface IAction {

    /** Action Type */
    Type?: (ActionType|null);

    /** Action Value */
    Value?: (number|null);

    /** Action Value2 */
    Value2?: (number|null);

    /** Action TargetID */
    TargetID?: (number|null);
}

/** Represents an Action. */
export class Action implements IAction {

    /**
     * Constructs a new Action.
     * @param [properties] Properties to set
     */
    constructor(properties?: IAction);

    /** Action Type. */
    public Type: ActionType;

    /** Action Value. */
    public Value: number;

    /** Action Value2. */
    public Value2: number;

    /** Action TargetID. */
    public TargetID: number;

    /**
     * Creates a new Action instance using the specified properties.
     * @param [properties] Properties to set
     * @returns Action instance
     */
    public static create(properties?: IAction): Action;

    /**
     * Encodes the specified Action message. Does not implicitly {@link Action.verify|verify} messages.
     * @param message Action message or plain object to encode
     * @param [writer] Writer to encode to
     * @returns Writer
     */
    public static encode(message: IAction, writer?: $protobuf.Writer): $protobuf.Writer;

    /**
     * Encodes the specified Action message, length delimited. Does not implicitly {@link Action.verify|verify} messages.
     * @param message Action message or plain object to encode
     * @param [writer] Writer to encode to
     * @returns Writer
     */
    public static encodeDelimited(message: IAction, writer?: $protobuf.Writer): $protobuf.Writer;

    /**
     * Decodes an Action message from the specified reader or buffer.
     * @param reader Reader or buffer to decode from
     * @param [length] Message length if known beforehand
     * @returns Action
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): Action;

    /**
     * Decodes an Action message from the specified reader or buffer, length delimited.
     * @param reader Reader or buffer to decode from
     * @returns Action
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): Action;

    /**
     * Verifies an Action message.
     * @param message Plain object to verify
     * @returns `null` if valid, otherwise the reason why it is not
     */
    public static verify(message: [ 'object' ].<string, any>): (string|null);

    /**
     * Creates an Action message from a plain object. Also converts values to their respective internal types.
     * @param object Plain object
     * @returns Action
     */
    public static fromObject(object: [ 'object' ].<string, any>): Action;

    /**
     * Creates a plain object from an Action message. Also converts values to other types if specified.
     * @param message Action
     * @param [options] Conversion options
     * @returns Plain object
     */
    public static toObject(message: Action, options?: $protobuf.IConversionOptions): [ 'object' ].<string, any>;

    /**
     * Converts this Action to JSON.
     * @returns JSON object
     */
    public toJSON(): [ 'object' ].<string, any>;
}

/** Properties of a ConnectionData. */
export interface IConnectionData {

    /** ConnectionData roomID */
    roomID?: (number|null);

    /** ConnectionData playerID */
    playerID?: (number|null);

    /** ConnectionData conMsg */
    conMsg?: (number|null);
}

/** Represents a ConnectionData. */
export class ConnectionData implements IConnectionData {

    /**
     * Constructs a new ConnectionData.
     * @param [properties] Properties to set
     */
    constructor(properties?: IConnectionData);

    /** ConnectionData roomID. */
    public roomID: number;

    /** ConnectionData playerID. */
    public playerID: number;

    /** ConnectionData conMsg. */
    public conMsg: number;

    /**
     * Creates a new ConnectionData instance using the specified properties.
     * @param [properties] Properties to set
     * @returns ConnectionData instance
     */
    public static create(properties?: IConnectionData): ConnectionData;

    /**
     * Encodes the specified ConnectionData message. Does not implicitly {@link ConnectionData.verify|verify} messages.
     * @param message ConnectionData message or plain object to encode
     * @param [writer] Writer to encode to
     * @returns Writer
     */
    public static encode(message: IConnectionData, writer?: $protobuf.Writer): $protobuf.Writer;

    /**
     * Encodes the specified ConnectionData message, length delimited. Does not implicitly {@link ConnectionData.verify|verify} messages.
     * @param message ConnectionData message or plain object to encode
     * @param [writer] Writer to encode to
     * @returns Writer
     */
    public static encodeDelimited(message: IConnectionData, writer?: $protobuf.Writer): $protobuf.Writer;

    /**
     * Decodes a ConnectionData message from the specified reader or buffer.
     * @param reader Reader or buffer to decode from
     * @param [length] Message length if known beforehand
     * @returns ConnectionData
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): ConnectionData;

    /**
     * Decodes a ConnectionData message from the specified reader or buffer, length delimited.
     * @param reader Reader or buffer to decode from
     * @returns ConnectionData
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): ConnectionData;

    /**
     * Verifies a ConnectionData message.
     * @param message Plain object to verify
     * @returns `null` if valid, otherwise the reason why it is not
     */
    public static verify(message: [ 'object' ].<string, any>): (string|null);

    /**
     * Creates a ConnectionData message from a plain object. Also converts values to their respective internal types.
     * @param object Plain object
     * @returns ConnectionData
     */
    public static fromObject(object: [ 'object' ].<string, any>): ConnectionData;

    /**
     * Creates a plain object from a ConnectionData message. Also converts values to other types if specified.
     * @param message ConnectionData
     * @param [options] Conversion options
     * @returns Plain object
     */
    public static toObject(message: ConnectionData, options?: $protobuf.IConversionOptions): [ 'object' ].<string, any>;

    /**
     * Converts this ConnectionData to JSON.
     * @returns JSON object
     */
    public toJSON(): [ 'object' ].<string, any>;
}

/** Properties of a LocationData. */
export interface ILocationData {

    /** LocationData locationName */
    locationName?: (string|null);

    /** LocationData roomId */
    roomId?: (number|null);
}

/** Represents a LocationData. */
export class LocationData implements ILocationData {

    /**
     * Constructs a new LocationData.
     * @param [properties] Properties to set
     */
    constructor(properties?: ILocationData);

    /** LocationData locationName. */
    public locationName: string;

    /** LocationData roomId. */
    public roomId: number;

    /**
     * Creates a new LocationData instance using the specified properties.
     * @param [properties] Properties to set
     * @returns LocationData instance
     */
    public static create(properties?: ILocationData): LocationData;

    /**
     * Encodes the specified LocationData message. Does not implicitly {@link LocationData.verify|verify} messages.
     * @param message LocationData message or plain object to encode
     * @param [writer] Writer to encode to
     * @returns Writer
     */
    public static encode(message: ILocationData, writer?: $protobuf.Writer): $protobuf.Writer;

    /**
     * Encodes the specified LocationData message, length delimited. Does not implicitly {@link LocationData.verify|verify} messages.
     * @param message LocationData message or plain object to encode
     * @param [writer] Writer to encode to
     * @returns Writer
     */
    public static encodeDelimited(message: ILocationData, writer?: $protobuf.Writer): $protobuf.Writer;

    /**
     * Decodes a LocationData message from the specified reader or buffer.
     * @param reader Reader or buffer to decode from
     * @param [length] Message length if known beforehand
     * @returns LocationData
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): LocationData;

    /**
     * Decodes a LocationData message from the specified reader or buffer, length delimited.
     * @param reader Reader or buffer to decode from
     * @returns LocationData
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): LocationData;

    /**
     * Verifies a LocationData message.
     * @param message Plain object to verify
     * @returns `null` if valid, otherwise the reason why it is not
     */
    public static verify(message: [ 'object' ].<string, any>): (string|null);

    /**
     * Creates a LocationData message from a plain object. Also converts values to their respective internal types.
     * @param object Plain object
     * @returns LocationData
     */
    public static fromObject(object: [ 'object' ].<string, any>): LocationData;

    /**
     * Creates a plain object from a LocationData message. Also converts values to other types if specified.
     * @param message LocationData
     * @param [options] Conversion options
     * @returns Plain object
     */
    public static toObject(message: LocationData, options?: $protobuf.IConversionOptions): [ 'object' ].<string, any>;

    /**
     * Converts this LocationData to JSON.
     * @returns JSON object
     */
    public toJSON(): [ 'object' ].<string, any>;
}

/** Properties of an AccountGeneral. */
export interface IAccountGeneral {

    /** AccountGeneral Money */
    Money?: (number|null);

    /** AccountGeneral Username */
    Username?: (string|null);
}

/** Represents an AccountGeneral. */
export class AccountGeneral implements IAccountGeneral {

    /**
     * Constructs a new AccountGeneral.
     * @param [properties] Properties to set
     */
    constructor(properties?: IAccountGeneral);

    /** AccountGeneral Money. */
    public Money: number;

    /** AccountGeneral Username. */
    public Username: string;

    /**
     * Creates a new AccountGeneral instance using the specified properties.
     * @param [properties] Properties to set
     * @returns AccountGeneral instance
     */
    public static create(properties?: IAccountGeneral): AccountGeneral;

    /**
     * Encodes the specified AccountGeneral message. Does not implicitly {@link AccountGeneral.verify|verify} messages.
     * @param message AccountGeneral message or plain object to encode
     * @param [writer] Writer to encode to
     * @returns Writer
     */
    public static encode(message: IAccountGeneral, writer?: $protobuf.Writer): $protobuf.Writer;

    /**
     * Encodes the specified AccountGeneral message, length delimited. Does not implicitly {@link AccountGeneral.verify|verify} messages.
     * @param message AccountGeneral message or plain object to encode
     * @param [writer] Writer to encode to
     * @returns Writer
     */
    public static encodeDelimited(message: IAccountGeneral, writer?: $protobuf.Writer): $protobuf.Writer;

    /**
     * Decodes an AccountGeneral message from the specified reader or buffer.
     * @param reader Reader or buffer to decode from
     * @param [length] Message length if known beforehand
     * @returns AccountGeneral
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): AccountGeneral;

    /**
     * Decodes an AccountGeneral message from the specified reader or buffer, length delimited.
     * @param reader Reader or buffer to decode from
     * @returns AccountGeneral
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): AccountGeneral;

    /**
     * Verifies an AccountGeneral message.
     * @param message Plain object to verify
     * @returns `null` if valid, otherwise the reason why it is not
     */
    public static verify(message: [ 'object' ].<string, any>): (string|null);

    /**
     * Creates an AccountGeneral message from a plain object. Also converts values to their respective internal types.
     * @param object Plain object
     * @returns AccountGeneral
     */
    public static fromObject(object: [ 'object' ].<string, any>): AccountGeneral;

    /**
     * Creates a plain object from an AccountGeneral message. Also converts values to other types if specified.
     * @param message AccountGeneral
     * @param [options] Conversion options
     * @returns Plain object
     */
    public static toObject(message: AccountGeneral, options?: $protobuf.IConversionOptions): [ 'object' ].<string, any>;

    /**
     * Converts this AccountGeneral to JSON.
     * @returns JSON object
     */
    public toJSON(): [ 'object' ].<string, any>;
}

/** Properties of a ServerData. */
export interface IServerData {

    /** ServerData conData */
    conData?: (IConnectionData|null);

    /** ServerData accountGeneral */
    accountGeneral?: (IAccountGeneral|null);

    /** ServerData locationData */
    locationData?: (ILocationData|null);

    /** ServerData curves */
    curves?: (string[]|null);

    /** ServerData customObjects */
    customObjects?: (ICustomObject[]|null);

    /** ServerData fishes */
    fishes?: (IFish[]|null);

    /** ServerData players */
    players?: (IPlayer[]|null);

    /** ServerData actions */
    actions?: (IAction[]|null);
}

/** Represents a ServerData. */
export class ServerData implements IServerData {

    /**
     * Constructs a new ServerData.
     * @param [properties] Properties to set
     */
    constructor(properties?: IServerData);

    /** ServerData conData. */
    public conData?: (IConnectionData|null);

    /** ServerData accountGeneral. */
    public accountGeneral?: (IAccountGeneral|null);

    /** ServerData locationData. */
    public locationData?: (ILocationData|null);

    /** ServerData curves. */
    public curves: [ 'Array' ].<string>;

    /** ServerData customObjects. */
    public customObjects: [ 'Array' ].<ICustomObject>;

    /** ServerData fishes. */
    public fishes: [ 'Array' ].<IFish>;

    /** ServerData players. */
    public players: [ 'Array' ].<IPlayer>;

    /** ServerData actions. */
    public actions: [ 'Array' ].<IAction>;

    /**
     * Creates a new ServerData instance using the specified properties.
     * @param [properties] Properties to set
     * @returns ServerData instance
     */
    public static create(properties?: IServerData): ServerData;

    /**
     * Encodes the specified ServerData message. Does not implicitly {@link ServerData.verify|verify} messages.
     * @param message ServerData message or plain object to encode
     * @param [writer] Writer to encode to
     * @returns Writer
     */
    public static encode(message: IServerData, writer?: $protobuf.Writer): $protobuf.Writer;

    /**
     * Encodes the specified ServerData message, length delimited. Does not implicitly {@link ServerData.verify|verify} messages.
     * @param message ServerData message or plain object to encode
     * @param [writer] Writer to encode to
     * @returns Writer
     */
    public static encodeDelimited(message: IServerData, writer?: $protobuf.Writer): $protobuf.Writer;

    /**
     * Decodes a ServerData message from the specified reader or buffer.
     * @param reader Reader or buffer to decode from
     * @param [length] Message length if known beforehand
     * @returns ServerData
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): ServerData;

    /**
     * Decodes a ServerData message from the specified reader or buffer, length delimited.
     * @param reader Reader or buffer to decode from
     * @returns ServerData
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): ServerData;

    /**
     * Verifies a ServerData message.
     * @param message Plain object to verify
     * @returns `null` if valid, otherwise the reason why it is not
     */
    public static verify(message: [ 'object' ].<string, any>): (string|null);

    /**
     * Creates a ServerData message from a plain object. Also converts values to their respective internal types.
     * @param object Plain object
     * @returns ServerData
     */
    public static fromObject(object: [ 'object' ].<string, any>): ServerData;

    /**
     * Creates a plain object from a ServerData message. Also converts values to other types if specified.
     * @param message ServerData
     * @param [options] Conversion options
     * @returns Plain object
     */
    public static toObject(message: ServerData, options?: $protobuf.IConversionOptions): [ 'object' ].<string, any>;

    /**
     * Converts this ServerData to JSON.
     * @returns JSON object
     */
    public toJSON(): [ 'object' ].<string, any>;
}

/** Properties of a Command. */
export interface ICommand {

    /** Command CommandId */
    CommandId?: (number|null);

    /** Command Params */
    Params?: (number[]|null);
}

/** Represents a Command. */
export class Command implements ICommand {

    /**
     * Constructs a new Command.
     * @param [properties] Properties to set
     */
    constructor(properties?: ICommand);

    /** Command CommandId. */
    public CommandId: number;

    /** Command Params. */
    public Params: [ 'Array' ].<number>;

    /**
     * Creates a new Command instance using the specified properties.
     * @param [properties] Properties to set
     * @returns Command instance
     */
    public static create(properties?: ICommand): Command;

    /**
     * Encodes the specified Command message. Does not implicitly {@link Command.verify|verify} messages.
     * @param message Command message or plain object to encode
     * @param [writer] Writer to encode to
     * @returns Writer
     */
    public static encode(message: ICommand, writer?: $protobuf.Writer): $protobuf.Writer;

    /**
     * Encodes the specified Command message, length delimited. Does not implicitly {@link Command.verify|verify} messages.
     * @param message Command message or plain object to encode
     * @param [writer] Writer to encode to
     * @returns Writer
     */
    public static encodeDelimited(message: ICommand, writer?: $protobuf.Writer): $protobuf.Writer;

    /**
     * Decodes a Command message from the specified reader or buffer.
     * @param reader Reader or buffer to decode from
     * @param [length] Message length if known beforehand
     * @returns Command
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): Command;

    /**
     * Decodes a Command message from the specified reader or buffer, length delimited.
     * @param reader Reader or buffer to decode from
     * @returns Command
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): Command;

    /**
     * Verifies a Command message.
     * @param message Plain object to verify
     * @returns `null` if valid, otherwise the reason why it is not
     */
    public static verify(message: [ 'object' ].<string, any>): (string|null);

    /**
     * Creates a Command message from a plain object. Also converts values to their respective internal types.
     * @param object Plain object
     * @returns Command
     */
    public static fromObject(object: [ 'object' ].<string, any>): Command;

    /**
     * Creates a plain object from a Command message. Also converts values to other types if specified.
     * @param message Command
     * @param [options] Conversion options
     * @returns Plain object
     */
    public static toObject(message: Command, options?: $protobuf.IConversionOptions): [ 'object' ].<string, any>;

    /**
     * Converts this Command to JSON.
     * @returns JSON object
     */
    public toJSON(): [ 'object' ].<string, any>;
}

/** Properties of a ClientData. */
export interface IClientData {

    /** ClientData roomId */
    roomId?: (number|null);

    /** ClientData commands */
    commands?: (ICommand[]|null);
}

/** Represents a ClientData. */
export class ClientData implements IClientData {

    /**
     * Constructs a new ClientData.
     * @param [properties] Properties to set
     */
    constructor(properties?: IClientData);

    /** ClientData roomId. */
    public roomId: number;

    /** ClientData commands. */
    public commands: [ 'Array' ].<ICommand>;

    /**
     * Creates a new ClientData instance using the specified properties.
     * @param [properties] Properties to set
     * @returns ClientData instance
     */
    public static create(properties?: IClientData): ClientData;

    /**
     * Encodes the specified ClientData message. Does not implicitly {@link ClientData.verify|verify} messages.
     * @param message ClientData message or plain object to encode
     * @param [writer] Writer to encode to
     * @returns Writer
     */
    public static encode(message: IClientData, writer?: $protobuf.Writer): $protobuf.Writer;

    /**
     * Encodes the specified ClientData message, length delimited. Does not implicitly {@link ClientData.verify|verify} messages.
     * @param message ClientData message or plain object to encode
     * @param [writer] Writer to encode to
     * @returns Writer
     */
    public static encodeDelimited(message: IClientData, writer?: $protobuf.Writer): $protobuf.Writer;

    /**
     * Decodes a ClientData message from the specified reader or buffer.
     * @param reader Reader or buffer to decode from
     * @param [length] Message length if known beforehand
     * @returns ClientData
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): ClientData;

    /**
     * Decodes a ClientData message from the specified reader or buffer, length delimited.
     * @param reader Reader or buffer to decode from
     * @returns ClientData
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): ClientData;

    /**
     * Verifies a ClientData message.
     * @param message Plain object to verify
     * @returns `null` if valid, otherwise the reason why it is not
     */
    public static verify(message: [ 'object' ].<string, any>): (string|null);

    /**
     * Creates a ClientData message from a plain object. Also converts values to their respective internal types.
     * @param object Plain object
     * @returns ClientData
     */
    public static fromObject(object: [ 'object' ].<string, any>): ClientData;

    /**
     * Creates a plain object from a ClientData message. Also converts values to other types if specified.
     * @param message ClientData
     * @param [options] Conversion options
     * @returns Plain object
     */
    public static toObject(message: ClientData, options?: $protobuf.IConversionOptions): [ 'object' ].<string, any>;

    /**
     * Converts this ClientData to JSON.
     * @returns JSON object
     */
    public toJSON(): [ 'object' ].<string, any>;
}
