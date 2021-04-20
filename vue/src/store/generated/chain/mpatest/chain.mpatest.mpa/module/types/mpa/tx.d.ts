import { Reader, Writer } from 'protobufjs/minimal';
export declare const protobufPackage = "chain.mpatest.mpa";
/** this line is used by starport scaffolding # proto/tx/message */
export interface MsgCreatePosition {
    creator: string;
    collateral: number;
    supply: number;
}
export interface MsgCreatePositionResponse {
    id: number;
}
export interface MsgUpdatePosition {
    creator: string;
    id: number;
    collateral: number;
    supply: number;
}
export interface MsgUpdatePositionResponse {
}
export interface MsgDeletePosition {
    creator: string;
    id: number;
}
export interface MsgDeletePositionResponse {
}
export declare const MsgCreatePosition: {
    encode(message: MsgCreatePosition, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreatePosition;
    fromJSON(object: any): MsgCreatePosition;
    toJSON(message: MsgCreatePosition): unknown;
    fromPartial(object: DeepPartial<MsgCreatePosition>): MsgCreatePosition;
};
export declare const MsgCreatePositionResponse: {
    encode(message: MsgCreatePositionResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreatePositionResponse;
    fromJSON(object: any): MsgCreatePositionResponse;
    toJSON(message: MsgCreatePositionResponse): unknown;
    fromPartial(object: DeepPartial<MsgCreatePositionResponse>): MsgCreatePositionResponse;
};
export declare const MsgUpdatePosition: {
    encode(message: MsgUpdatePosition, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdatePosition;
    fromJSON(object: any): MsgUpdatePosition;
    toJSON(message: MsgUpdatePosition): unknown;
    fromPartial(object: DeepPartial<MsgUpdatePosition>): MsgUpdatePosition;
};
export declare const MsgUpdatePositionResponse: {
    encode(_: MsgUpdatePositionResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdatePositionResponse;
    fromJSON(_: any): MsgUpdatePositionResponse;
    toJSON(_: MsgUpdatePositionResponse): unknown;
    fromPartial(_: DeepPartial<MsgUpdatePositionResponse>): MsgUpdatePositionResponse;
};
export declare const MsgDeletePosition: {
    encode(message: MsgDeletePosition, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeletePosition;
    fromJSON(object: any): MsgDeletePosition;
    toJSON(message: MsgDeletePosition): unknown;
    fromPartial(object: DeepPartial<MsgDeletePosition>): MsgDeletePosition;
};
export declare const MsgDeletePositionResponse: {
    encode(_: MsgDeletePositionResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeletePositionResponse;
    fromJSON(_: any): MsgDeletePositionResponse;
    toJSON(_: MsgDeletePositionResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeletePositionResponse>): MsgDeletePositionResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    /** this line is used by starport scaffolding # proto/tx/rpc */
    CreatePosition(request: MsgCreatePosition): Promise<MsgCreatePositionResponse>;
    UpdatePosition(request: MsgUpdatePosition): Promise<MsgUpdatePositionResponse>;
    DeletePosition(request: MsgDeletePosition): Promise<MsgDeletePositionResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    CreatePosition(request: MsgCreatePosition): Promise<MsgCreatePositionResponse>;
    UpdatePosition(request: MsgUpdatePosition): Promise<MsgUpdatePositionResponse>;
    DeletePosition(request: MsgDeletePosition): Promise<MsgDeletePositionResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
