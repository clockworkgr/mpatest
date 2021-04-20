import { Reader, Writer } from 'protobufjs/minimal';
import { Position } from '../mpa/position';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
export declare const protobufPackage = "chain.mpatest.mpa";
/** this line is used by starport scaffolding # 3 */
export interface QueryGetPositionRequest {
    id: number;
}
export interface QueryGetPositionResponse {
    Position: Position | undefined;
}
export interface QueryAllPositionRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllPositionResponse {
    Position: Position[];
    pagination: PageResponse | undefined;
}
export declare const QueryGetPositionRequest: {
    encode(message: QueryGetPositionRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetPositionRequest;
    fromJSON(object: any): QueryGetPositionRequest;
    toJSON(message: QueryGetPositionRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetPositionRequest>): QueryGetPositionRequest;
};
export declare const QueryGetPositionResponse: {
    encode(message: QueryGetPositionResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetPositionResponse;
    fromJSON(object: any): QueryGetPositionResponse;
    toJSON(message: QueryGetPositionResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetPositionResponse>): QueryGetPositionResponse;
};
export declare const QueryAllPositionRequest: {
    encode(message: QueryAllPositionRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllPositionRequest;
    fromJSON(object: any): QueryAllPositionRequest;
    toJSON(message: QueryAllPositionRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllPositionRequest>): QueryAllPositionRequest;
};
export declare const QueryAllPositionResponse: {
    encode(message: QueryAllPositionResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllPositionResponse;
    fromJSON(object: any): QueryAllPositionResponse;
    toJSON(message: QueryAllPositionResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllPositionResponse>): QueryAllPositionResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** this line is used by starport scaffolding # 2 */
    Position(request: QueryGetPositionRequest): Promise<QueryGetPositionResponse>;
    PositionAll(request: QueryAllPositionRequest): Promise<QueryAllPositionResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Position(request: QueryGetPositionRequest): Promise<QueryGetPositionResponse>;
    PositionAll(request: QueryAllPositionRequest): Promise<QueryAllPositionResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
