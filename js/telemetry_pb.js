// source: telemetry.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

goog.exportSymbol('proto.Telemetry', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.Telemetry = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.Telemetry, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.Telemetry.displayName = 'proto.Telemetry';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.Telemetry.prototype.toObject = function(opt_includeInstance) {
  return proto.Telemetry.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.Telemetry} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.Telemetry.toObject = function(includeInstance, msg) {
  var f, obj = {
    load1: jspb.Message.getFloatingPointFieldWithDefault(msg, 1, 0.0),
    load2: jspb.Message.getFloatingPointFieldWithDefault(msg, 2, 0.0),
    load3: jspb.Message.getFloatingPointFieldWithDefault(msg, 3, 0.0),
    memoryused: jspb.Message.getFieldWithDefault(msg, 4, 0),
    memoryfree: jspb.Message.getFieldWithDefault(msg, 5, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.Telemetry}
 */
proto.Telemetry.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.Telemetry;
  return proto.Telemetry.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.Telemetry} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.Telemetry}
 */
proto.Telemetry.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readFloat());
      msg.setLoad1(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readFloat());
      msg.setLoad2(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readFloat());
      msg.setLoad3(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setMemoryused(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setMemoryfree(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.Telemetry.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.Telemetry.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.Telemetry} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.Telemetry.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getLoad1();
  if (f !== 0.0) {
    writer.writeFloat(
      1,
      f
    );
  }
  f = message.getLoad2();
  if (f !== 0.0) {
    writer.writeFloat(
      2,
      f
    );
  }
  f = message.getLoad3();
  if (f !== 0.0) {
    writer.writeFloat(
      3,
      f
    );
  }
  f = message.getMemoryused();
  if (f !== 0) {
    writer.writeInt64(
      4,
      f
    );
  }
  f = message.getMemoryfree();
  if (f !== 0) {
    writer.writeInt64(
      5,
      f
    );
  }
};


/**
 * optional float load1 = 1;
 * @return {number}
 */
proto.Telemetry.prototype.getLoad1 = function() {
  return /** @type {number} */ (jspb.Message.getFloatingPointFieldWithDefault(this, 1, 0.0));
};


/**
 * @param {number} value
 * @return {!proto.Telemetry} returns this
 */
proto.Telemetry.prototype.setLoad1 = function(value) {
  return jspb.Message.setProto3FloatField(this, 1, value);
};


/**
 * optional float load2 = 2;
 * @return {number}
 */
proto.Telemetry.prototype.getLoad2 = function() {
  return /** @type {number} */ (jspb.Message.getFloatingPointFieldWithDefault(this, 2, 0.0));
};


/**
 * @param {number} value
 * @return {!proto.Telemetry} returns this
 */
proto.Telemetry.prototype.setLoad2 = function(value) {
  return jspb.Message.setProto3FloatField(this, 2, value);
};


/**
 * optional float load3 = 3;
 * @return {number}
 */
proto.Telemetry.prototype.getLoad3 = function() {
  return /** @type {number} */ (jspb.Message.getFloatingPointFieldWithDefault(this, 3, 0.0));
};


/**
 * @param {number} value
 * @return {!proto.Telemetry} returns this
 */
proto.Telemetry.prototype.setLoad3 = function(value) {
  return jspb.Message.setProto3FloatField(this, 3, value);
};


/**
 * optional int64 memoryUsed = 4;
 * @return {number}
 */
proto.Telemetry.prototype.getMemoryused = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/**
 * @param {number} value
 * @return {!proto.Telemetry} returns this
 */
proto.Telemetry.prototype.setMemoryused = function(value) {
  return jspb.Message.setProto3IntField(this, 4, value);
};


/**
 * optional int64 memoryFree = 5;
 * @return {number}
 */
proto.Telemetry.prototype.getMemoryfree = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/**
 * @param {number} value
 * @return {!proto.Telemetry} returns this
 */
proto.Telemetry.prototype.setMemoryfree = function(value) {
  return jspb.Message.setProto3IntField(this, 5, value);
};


goog.object.extend(exports, proto);
