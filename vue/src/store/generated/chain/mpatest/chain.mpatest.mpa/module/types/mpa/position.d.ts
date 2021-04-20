import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "chain.mpatest.mpa";
export interface Position {
    creator: string;
    id: number;
    collateral: number;
    supply: number;
}
export declare const Position: {
    encode(message: Position, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Position;
    fromJSON(object: any): Position;
    toJSON(message: Position): unknown;
    fromPartial(object: DeepPartial<Position>): Position;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
