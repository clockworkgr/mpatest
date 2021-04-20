/* eslint-disable */
import { Reader, util, configure, Writer } from 'protobufjs/minimal'
import * as Long from 'long'

export const protobufPackage = 'chain.mpatest.mpa'

/** this line is used by starport scaffolding # proto/tx/message */
export interface MsgCreatePosition {
  creator: string
  collateral: number
  supply: number
}

export interface MsgCreatePositionResponse {
  id: number
}

export interface MsgUpdatePosition {
  creator: string
  id: number
  collateral: number
  supply: number
}

export interface MsgUpdatePositionResponse {}

export interface MsgDeletePosition {
  creator: string
  id: number
}

export interface MsgDeletePositionResponse {}

const baseMsgCreatePosition: object = { creator: '', collateral: 0, supply: 0 }

export const MsgCreatePosition = {
  encode(message: MsgCreatePosition, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.collateral !== 0) {
      writer.uint32(16).int32(message.collateral)
    }
    if (message.supply !== 0) {
      writer.uint32(24).int32(message.supply)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreatePosition {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgCreatePosition } as MsgCreatePosition
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.collateral = reader.int32()
          break
        case 3:
          message.supply = reader.int32()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgCreatePosition {
    const message = { ...baseMsgCreatePosition } as MsgCreatePosition
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.collateral !== undefined && object.collateral !== null) {
      message.collateral = Number(object.collateral)
    } else {
      message.collateral = 0
    }
    if (object.supply !== undefined && object.supply !== null) {
      message.supply = Number(object.supply)
    } else {
      message.supply = 0
    }
    return message
  },

  toJSON(message: MsgCreatePosition): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.collateral !== undefined && (obj.collateral = message.collateral)
    message.supply !== undefined && (obj.supply = message.supply)
    return obj
  },

  fromPartial(object: DeepPartial<MsgCreatePosition>): MsgCreatePosition {
    const message = { ...baseMsgCreatePosition } as MsgCreatePosition
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.collateral !== undefined && object.collateral !== null) {
      message.collateral = object.collateral
    } else {
      message.collateral = 0
    }
    if (object.supply !== undefined && object.supply !== null) {
      message.supply = object.supply
    } else {
      message.supply = 0
    }
    return message
  }
}

const baseMsgCreatePositionResponse: object = { id: 0 }

export const MsgCreatePositionResponse = {
  encode(
    message: MsgCreatePositionResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id)
    }
    return writer
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreatePositionResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = {
      ...baseMsgCreatePositionResponse
    } as MsgCreatePositionResponse
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

  fromJSON(object: any): MsgCreatePositionResponse {
    const message = {
      ...baseMsgCreatePositionResponse
    } as MsgCreatePositionResponse
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id)
    } else {
      message.id = 0
    }
    return message
  },

  toJSON(message: MsgCreatePositionResponse): unknown {
    const obj: any = {}
    message.id !== undefined && (obj.id = message.id)
    return obj
  },

  fromPartial(
    object: DeepPartial<MsgCreatePositionResponse>
  ): MsgCreatePositionResponse {
    const message = {
      ...baseMsgCreatePositionResponse
    } as MsgCreatePositionResponse
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id
    } else {
      message.id = 0
    }
    return message
  }
}

const baseMsgUpdatePosition: object = {
  creator: '',
  id: 0,
  collateral: 0,
  supply: 0
}

export const MsgUpdatePosition = {
  encode(message: MsgUpdatePosition, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id)
    }
    if (message.collateral !== 0) {
      writer.uint32(24).int32(message.collateral)
    }
    if (message.supply !== 0) {
      writer.uint32(32).int32(message.supply)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdatePosition {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgUpdatePosition } as MsgUpdatePosition
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.id = longToNumber(reader.uint64() as Long)
          break
        case 3:
          message.collateral = reader.int32()
          break
        case 4:
          message.supply = reader.int32()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgUpdatePosition {
    const message = { ...baseMsgUpdatePosition } as MsgUpdatePosition
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id)
    } else {
      message.id = 0
    }
    if (object.collateral !== undefined && object.collateral !== null) {
      message.collateral = Number(object.collateral)
    } else {
      message.collateral = 0
    }
    if (object.supply !== undefined && object.supply !== null) {
      message.supply = Number(object.supply)
    } else {
      message.supply = 0
    }
    return message
  },

  toJSON(message: MsgUpdatePosition): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.id !== undefined && (obj.id = message.id)
    message.collateral !== undefined && (obj.collateral = message.collateral)
    message.supply !== undefined && (obj.supply = message.supply)
    return obj
  },

  fromPartial(object: DeepPartial<MsgUpdatePosition>): MsgUpdatePosition {
    const message = { ...baseMsgUpdatePosition } as MsgUpdatePosition
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id
    } else {
      message.id = 0
    }
    if (object.collateral !== undefined && object.collateral !== null) {
      message.collateral = object.collateral
    } else {
      message.collateral = 0
    }
    if (object.supply !== undefined && object.supply !== null) {
      message.supply = object.supply
    } else {
      message.supply = 0
    }
    return message
  }
}

const baseMsgUpdatePositionResponse: object = {}

export const MsgUpdatePositionResponse = {
  encode(
    _: MsgUpdatePositionResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUpdatePositionResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = {
      ...baseMsgUpdatePositionResponse
    } as MsgUpdatePositionResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(_: any): MsgUpdatePositionResponse {
    const message = {
      ...baseMsgUpdatePositionResponse
    } as MsgUpdatePositionResponse
    return message
  },

  toJSON(_: MsgUpdatePositionResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(
    _: DeepPartial<MsgUpdatePositionResponse>
  ): MsgUpdatePositionResponse {
    const message = {
      ...baseMsgUpdatePositionResponse
    } as MsgUpdatePositionResponse
    return message
  }
}

const baseMsgDeletePosition: object = { creator: '', id: 0 }

export const MsgDeletePosition = {
  encode(message: MsgDeletePosition, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeletePosition {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgDeletePosition } as MsgDeletePosition
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.id = longToNumber(reader.uint64() as Long)
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgDeletePosition {
    const message = { ...baseMsgDeletePosition } as MsgDeletePosition
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id)
    } else {
      message.id = 0
    }
    return message
  },

  toJSON(message: MsgDeletePosition): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.id !== undefined && (obj.id = message.id)
    return obj
  },

  fromPartial(object: DeepPartial<MsgDeletePosition>): MsgDeletePosition {
    const message = { ...baseMsgDeletePosition } as MsgDeletePosition
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id
    } else {
      message.id = 0
    }
    return message
  }
}

const baseMsgDeletePositionResponse: object = {}

export const MsgDeletePositionResponse = {
  encode(
    _: MsgDeletePositionResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeletePositionResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = {
      ...baseMsgDeletePositionResponse
    } as MsgDeletePositionResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(_: any): MsgDeletePositionResponse {
    const message = {
      ...baseMsgDeletePositionResponse
    } as MsgDeletePositionResponse
    return message
  },

  toJSON(_: MsgDeletePositionResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(
    _: DeepPartial<MsgDeletePositionResponse>
  ): MsgDeletePositionResponse {
    const message = {
      ...baseMsgDeletePositionResponse
    } as MsgDeletePositionResponse
    return message
  }
}

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  CreatePosition(request: MsgCreatePosition): Promise<MsgCreatePositionResponse>
  UpdatePosition(request: MsgUpdatePosition): Promise<MsgUpdatePositionResponse>
  DeletePosition(request: MsgDeletePosition): Promise<MsgDeletePositionResponse>
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  CreatePosition(
    request: MsgCreatePosition
  ): Promise<MsgCreatePositionResponse> {
    const data = MsgCreatePosition.encode(request).finish()
    const promise = this.rpc.request(
      'chain.mpatest.mpa.Msg',
      'CreatePosition',
      data
    )
    return promise.then((data) =>
      MsgCreatePositionResponse.decode(new Reader(data))
    )
  }

  UpdatePosition(
    request: MsgUpdatePosition
  ): Promise<MsgUpdatePositionResponse> {
    const data = MsgUpdatePosition.encode(request).finish()
    const promise = this.rpc.request(
      'chain.mpatest.mpa.Msg',
      'UpdatePosition',
      data
    )
    return promise.then((data) =>
      MsgUpdatePositionResponse.decode(new Reader(data))
    )
  }

  DeletePosition(
    request: MsgDeletePosition
  ): Promise<MsgDeletePositionResponse> {
    const data = MsgDeletePosition.encode(request).finish()
    const promise = this.rpc.request(
      'chain.mpatest.mpa.Msg',
      'DeletePosition',
      data
    )
    return promise.then((data) =>
      MsgDeletePositionResponse.decode(new Reader(data))
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
