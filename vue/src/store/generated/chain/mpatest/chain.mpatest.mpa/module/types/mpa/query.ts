/* eslint-disable */
import { Reader, util, configure, Writer } from 'protobufjs/minimal'
import * as Long from 'long'
import { Position } from '../mpa/position'
import {
  PageRequest,
  PageResponse
} from '../cosmos/base/query/v1beta1/pagination'

export const protobufPackage = 'chain.mpatest.mpa'

/** this line is used by starport scaffolding # 3 */
export interface QueryGetPositionRequest {
  id: number
}

export interface QueryGetPositionResponse {
  Position: Position | undefined
}

export interface QueryAllPositionRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllPositionResponse {
  Position: Position[]
  pagination: PageResponse | undefined
}

const baseQueryGetPositionRequest: object = { id: 0 }

export const QueryGetPositionRequest = {
  encode(
    message: QueryGetPositionRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetPositionRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = {
      ...baseQueryGetPositionRequest
    } as QueryGetPositionRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long)
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetPositionRequest {
    const message = {
      ...baseQueryGetPositionRequest
    } as QueryGetPositionRequest
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id)
    } else {
      message.id = 0
    }
    return message
  },

  toJSON(message: QueryGetPositionRequest): unknown {
    const obj: any = {}
    message.id !== undefined && (obj.id = message.id)
    return obj
  },

  fromPartial(
    object: DeepPartial<QueryGetPositionRequest>
  ): QueryGetPositionRequest {
    const message = {
      ...baseQueryGetPositionRequest
    } as QueryGetPositionRequest
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id
    } else {
      message.id = 0
    }
    return message
  }
}

const baseQueryGetPositionResponse: object = {}

export const QueryGetPositionResponse = {
  encode(
    message: QueryGetPositionResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.Position !== undefined) {
      Position.encode(message.Position, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetPositionResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = {
      ...baseQueryGetPositionResponse
    } as QueryGetPositionResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.Position = Position.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetPositionResponse {
    const message = {
      ...baseQueryGetPositionResponse
    } as QueryGetPositionResponse
    if (object.Position !== undefined && object.Position !== null) {
      message.Position = Position.fromJSON(object.Position)
    } else {
      message.Position = undefined
    }
    return message
  },

  toJSON(message: QueryGetPositionResponse): unknown {
    const obj: any = {}
    message.Position !== undefined &&
      (obj.Position = message.Position
        ? Position.toJSON(message.Position)
        : undefined)
    return obj
  },

  fromPartial(
    object: DeepPartial<QueryGetPositionResponse>
  ): QueryGetPositionResponse {
    const message = {
      ...baseQueryGetPositionResponse
    } as QueryGetPositionResponse
    if (object.Position !== undefined && object.Position !== null) {
      message.Position = Position.fromPartial(object.Position)
    } else {
      message.Position = undefined
    }
    return message
  }
}

const baseQueryAllPositionRequest: object = {}

export const QueryAllPositionRequest = {
  encode(
    message: QueryAllPositionRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllPositionRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = {
      ...baseQueryAllPositionRequest
    } as QueryAllPositionRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllPositionRequest {
    const message = {
      ...baseQueryAllPositionRequest
    } as QueryAllPositionRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllPositionRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined)
    return obj
  },

  fromPartial(
    object: DeepPartial<QueryAllPositionRequest>
  ): QueryAllPositionRequest {
    const message = {
      ...baseQueryAllPositionRequest
    } as QueryAllPositionRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllPositionResponse: object = {}

export const QueryAllPositionResponse = {
  encode(
    message: QueryAllPositionResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.Position) {
      Position.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllPositionResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = {
      ...baseQueryAllPositionResponse
    } as QueryAllPositionResponse
    message.Position = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.Position.push(Position.decode(reader, reader.uint32()))
          break
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllPositionResponse {
    const message = {
      ...baseQueryAllPositionResponse
    } as QueryAllPositionResponse
    message.Position = []
    if (object.Position !== undefined && object.Position !== null) {
      for (const e of object.Position) {
        message.Position.push(Position.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllPositionResponse): unknown {
    const obj: any = {}
    if (message.Position) {
      obj.Position = message.Position.map((e) =>
        e ? Position.toJSON(e) : undefined
      )
    } else {
      obj.Position = []
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined)
    return obj
  },

  fromPartial(
    object: DeepPartial<QueryAllPositionResponse>
  ): QueryAllPositionResponse {
    const message = {
      ...baseQueryAllPositionResponse
    } as QueryAllPositionResponse
    message.Position = []
    if (object.Position !== undefined && object.Position !== null) {
      for (const e of object.Position) {
        message.Position.push(Position.fromPartial(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

/** Query defines the gRPC querier service. */
export interface Query {
  /** this line is used by starport scaffolding # 2 */
  Position(request: QueryGetPositionRequest): Promise<QueryGetPositionResponse>
  PositionAll(
    request: QueryAllPositionRequest
  ): Promise<QueryAllPositionResponse>
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  Position(
    request: QueryGetPositionRequest
  ): Promise<QueryGetPositionResponse> {
    const data = QueryGetPositionRequest.encode(request).finish()
    const promise = this.rpc.request(
      'chain.mpatest.mpa.Query',
      'Position',
      data
    )
    return promise.then((data) =>
      QueryGetPositionResponse.decode(new Reader(data))
    )
  }

  PositionAll(
    request: QueryAllPositionRequest
  ): Promise<QueryAllPositionResponse> {
    const data = QueryAllPositionRequest.encode(request).finish()
    const promise = this.rpc.request(
      'chain.mpatest.mpa.Query',
      'PositionAll',
      data
    )
    return promise.then((data) =>
      QueryAllPositionResponse.decode(new Reader(data))
    )
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>
}

declare var self: any | undefined
declare var window: any | undefined
var globalThis: any = (() => {
  if (typeof globalThis !== 'undefined') return globalThis
  if (typeof self !== 'undefined') return self
  if (typeof window !== 'undefined') return window
  if (typeof global !== 'undefined') return global
  throw 'Unable to locate global object'
})()

type Builtin = Date | Function | Uint8Array | string | number | undefined
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error('Value is larger than Number.MAX_SAFE_INTEGER')
  }
  return long.toNumber()
}

if (util.Long !== Long) {
  util.Long = Long as any
  configure()
}
